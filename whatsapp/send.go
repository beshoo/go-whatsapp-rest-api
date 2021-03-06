package whatsapp

import (
        "bytes"
        "image/jpeg"
        "io"
"io/ioutil"
"fmt"
"strconv"
        "net/http"
        "time"
        "bitbucket.org/rockyOO7/wa-api/gen/models"
        "bitbucket.org/rockyOO7/wa-api/storage"
        wa "github.com/Rhymen/go-whatsapp"
        "github.com/disintegration/imaging"
	"github.com/Rhymen/go-whatsapp/binary/proto"
        log "github.com/sirupsen/logrus"
)

type work func(item *models.NumberReplyIds) error

func numberChannel(numberReplyIds []*models.NumberReplyIds) (channel <-chan models.NumberReplyIds) {
        numChannel := make(chan models.NumberReplyIds)
        go func() {
                for _, item := range numberReplyIds {
                        numChannel <- *item
                }
                close(numChannel)
        }()
        return numChannel
}

func sender(wh *WaHandler, broadcastID *string, numberReplyIds []*models.NumberReplyIds, quitChannel chan struct{}, f work) {

        go func() {
                //var wg sync.WaitGroup
                log.Info(len(numberReplyIds))
                //wg.Add(len(numberReplyIds))
                for _, item := range numberReplyIds {
                item:= item
                        select {
                        case wh.Queue <- func() error {
                                err := f(item)
                                if err != nil {
                                        return err
                                }
                                //wg.Done()
                                return nil
                        }:
                        case <-quitChannel:
                                err := storage.UpdateBroadcastStatus(broadcastID, storage.STOPPED)
                                if err != nil {
                                        log.Errorf("On Quit Broadcast Status update failed  %v", err)
                                }
                                return
                        }
                }
                log.Info("Out of loop")
                // wg.Wait()
                err := storage.UpdateBroadcastStatus(broadcastID, storage.COMPLETED)
                if err != nil {
                        log.Errorf("Broadcast Status update failed %v", err)
                }
        }()
}

//
func SendText(wh *WaHandler, data *models.TextMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.TEXT)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        if err != nil {
                return nil, err
        }

        work := func(item *models.NumberReplyIds) error {
                err := sendText(wh.Wac, *data.Text, *item.Number, item.ReplyTo, broadcastID)
                if err != nil {
                        return err
                }
                return nil
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}


func sendText(wac *wa.Conn, text, number string, reply *models.MessageItem, broadcastID *string) error {

        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.TextMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id: ID,
                },
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Text: text,
        }

        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID: ID,
                BroadcastID: broadcastID,
                From: NumberFromWidJid(wac.Info.Wid),
                To: number,
                Status: storage.SENT,
                Direction: storage.OUTGOING,
                MediaType: storage.TEXT,
                SentAt: &now,
        })
        if err != nil {
                return err
        }


        if wac == nil{
                log.Error("WAC is null Sendtext  failed")
                return nil
        }

/*        _, err = wac.Presence(JID, wa.PresenceComposing)
        if err != nil {
                log.Infof("Send Text Presence %v", err)
                return err
        }

        time.Sleep(100 * time.Millisecond)
*/
        log.Infof("Send Text Start Sending")
        _, err = wac.Send(msg)
        if err != nil {

                return err
        }

        return nil
}


func SendLocation(wh *WaHandler, data *models.LocationMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.TEXT)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        if err != nil {
                return nil, err
        }

        work := func(item *models.NumberReplyIds) error {
	lat, _ := strconv.ParseFloat(*data.Lat , 64)
	lng, _ := strconv.ParseFloat(*data.Lng, 64)
                err := sendLocation(wh.Wac, lat, lng ,*item.Number, item.ReplyTo, broadcastID)
                if err != nil {
                        return err
                }
                return nil
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}



func contextFromReplyID(wac *wa.Conn, wid, to string, reply *models.MessageItem) wa.ContextInfo {
        if reply == nil {
                return wa.ContextInfo{}
        }

        fromMe := *reply.FromMe
        qMessage := GetMessageForID(wac, to, *reply.ID, fromMe)
        participant := to
        if fromMe {
                participant = NumberFromWidJid(wid)
        }

        log.Info(participant)
        return wa.ContextInfo{
                QuotedMessageID: *reply.ID,
                Participant:     WidJidFromNumber(participant),
                QuotedMessage:   qMessage,
        }
}



//
func SendImage(wh *WaHandler, data *models.ImageMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }
        image, err := DownloadMediaFromURL(data.Image.String())
        if err != nil {
                log.Errorf("Error DownloadMediaFromURL %v", err)
                return nil, err
        }

        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.IMAGE)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        var imageBuff bytes.Buffer
        thumbReader := io.TeeReader(image, &imageBuff)
        thumb, err := GetImageThumbnail(thumbReader)
//      imageReader := imageBuff.Bytes()

        if err != nil {
                log.Errorf("Error GetImageThumbnail %v", err)
                return nil, err
        }

        work := func(item *models.NumberReplyIds) error {
                return sendImage(wh.Wac, thumb, image, data.Caption, *item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}




func sendImage(wac *wa.Conn, thumb []byte, image io.Reader, caption, number string, reply *models.MessageItem, broadcastID *string) error {
        img := image.(*bytes.Reader)
        img.Seek(0,io.SeekStart)
        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.ImageMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id:        ID,
                },
                Type:        "image/jpeg",
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Content:     img,
                Caption:     caption,
                Thumbnail:   thumb,
        }
        //log.Errorf("%v image: %v",number, image)

        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID:  ID,
                BroadcastID: broadcastID,
                From:        NumberFromWidJid(wac.Info.Wid),
                To:          number,
                Status:      storage.SENT,
                Direction:   storage.OUTGOING,
                MediaType:   storage.IMAGE,
                SentAt:      &now,
        })
        if err != nil {
                log.Errorf("Error CreateMessage %v", err)
                return err
        }
//      wac.Presence(JID, wa.PresenceComposing)
//      time.Sleep(100 * time.Millisecond)
        _, err = wac.Send(msg)
        if err != nil {
                log.Errorf("Error sendImage %v to %v ID:%v JID:%v", err , number,ID,JID)
                return err
        }

        return nil

}

//
func SendVideo(wh *WaHandler, data *models.VideoMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }

        video, err := DownloadMediaFromURL(data.Video.String())
        if err != nil {
                return nil, err
        }
        var thumbnail []byte
        if data.VideoThumbnail.String() != "" {
                r, err := DownloadMediaFromURL(data.VideoThumbnail.String())
                if err == nil {
                        buf := new(bytes.Buffer)
                        buf.ReadFrom(r)
                        thumbnail = buf.Bytes()
                } else {
                        log.Errorf("Io read failed for thumbnail %v", err)
                }
        }
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.VIDEO)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }

        work := func(item *models.NumberReplyIds) error {
                return sendVideo(wh.Wac, video, thumbnail, data.Caption, *item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}

func sendVideo(wac *wa.Conn, video io.Reader, thumbnail []byte, caption, number string, reply *models.MessageItem, broadcastID *string) error {

        // buf := &bytes.Buffer{}
        // nRead, err := io.Copy(buf, video)
        // log.Infof("Check Size and throw error %v", nRead)
        //fmt.Printf("%v", video)

        vid := video.(*bytes.Reader)
        vid.Seek(0, io.SeekStart)
        ID := MessageID()
        JID := WidJidFromNumber(number)

        msg := wa.VideoMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id:        ID,
                },
                Type:        "video/mp4",
                Thumbnail:   thumbnail,
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Content:     vid,
                Caption:     caption,
        }
        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID:  ID,
                BroadcastID: broadcastID,
                From:        NumberFromWidJid(wac.Info.Wid),
                To:          number,
                Status:      storage.SENT,
                Direction:   storage.OUTGOING,
                MediaType:   storage.VIDEO,
                SentAt:      &now,
        })

        if err != nil {
                return err
        }

//      wac.Presence(JID, wa.PresenceComposing)
//      time.Sleep(100 * time.Millisecond)
        _, err = wac.Send(msg)
        if err != nil {
                log.Infof("Error %v", err)
//              log.Info(err)
                return err
        }
        return nil

}

//
func SendAudio(wh *WaHandler, data *models.AudioMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }
        audio, err := DownloadMediaFromURL(data.Audio.String())
        if err != nil {
                return nil, err
        }
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.AUDIO)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }

        work := func(item *models.NumberReplyIds) error {
                return sendAudio(wh.Wac, audio, *item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}

func sendAudio(wac *wa.Conn, audio io.Reader, number string, reply *models.MessageItem, broadcastID *string) error {

        audioReader := audio.(*bytes.Reader)
        audioReader.Seek(0, io.SeekStart)


        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.AudioMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id:        ID,
                },
                Type:        "audio/mpeg",
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Content:     audioReader,
        }

        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID:  ID,
                BroadcastID: broadcastID,
                From:        NumberFromWidJid(wac.Info.Wid),
                To:          number,
                Status:      storage.SENT,
                Direction:   storage.OUTGOING,
                MediaType:   storage.AUDIO,
                SentAt:      &now,
        })
        if err != nil {
                return err
        }
//        wac.Presence(JID, wa.PresenceComposing)
//        time.Sleep(100 * time.Millisecond)
        _, err = wac.Send(msg)
        if err != nil {
                log.Info(err)
                return err
        }
        return nil
}

//
func SendDoc(wh *WaHandler, data *models.DocMessage, docType string) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }
        doc, err := DownloadMediaFromURL(data.Doc.String())
        if err != nil {
                return nil, err
        }
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.AUDIO)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        work := func(item *models.NumberReplyIds) error {
                return sendDoc(wh.Wac, doc, docType, *data.Title, *item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)

        return broadcastID, nil
}

func getDocType(dtype string) string {
        docType := ""
        switch dtype {
        case "DOC":
                docType = "application/msword"
        case "DOCX":
                docType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        case "CSV":
                docType = "text/csv"
        case "XLS":
                docType = "application/vnd.ms-excel"
        case "XLSX":
                docType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        case "PDF":
                docType = "application/pdf"
        case "PPT":
                docType = "application/vnd.ms-powerpoint"
        case "PPTX":
                docType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
        case "GZ":
                docType = "application/gzip"
        case "ZIP":
                docType = "application/zip"
        case "7z":
                docType = "application/x-7z-compressed"
        case "TEXT":
                docType = "text/plain"
        }
        return docType
}

func sendDoc(wac *wa.Conn, doc io.Reader, docType string, title, number string, reply *models.MessageItem, broadcastID *string) error {

        docReader := doc.(*bytes.Reader)
        docReader.Seek(0, io.SeekStart)

        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.DocumentMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id:        ID,
                },
                Type:        getDocType(docType),
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Content:     docReader,
                Title:       title,
        }
        log.Infof("Doc %v", msg)

        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID:  ID,
                BroadcastID: broadcastID,
                From:        NumberFromWidJid(wac.Info.Wid),
                To:          number,
                Status:      storage.SENT,
                Direction:   storage.OUTGOING,
                MediaType:   storage.AUDIO,
                SentAt:      &now,
        })
        if err != nil {
                return err
        }
//        wac.Presence(JID, wa.PresenceComposing)
//        time.Sleep(100 * time.Millisecond)
        _, err = wac.Send(msg)
        if err != nil {
                log.Info(err)
                return err
        }

        return nil
}

//
func GetImageThumbnail(image io.Reader) ([]byte, error) {

        img, err := imaging.Decode(image)
        if err != nil {
                return nil, err
        }

        //b := img.Bounds()
        //imgWidth := b.Max.X
        //imgHeight := b.Max.Y

        thumbWidth := 140
        thumbHeight := 140
/*
        if imgWidth > imgHeight {
                thumbHeight = 56
        } else {
                thumbWidth = 56
        }
*/
        thumb := imaging.Thumbnail(img, thumbWidth, thumbHeight, imaging.CatmullRom)

        buf := new(bytes.Buffer)
        err = jpeg.Encode(buf, thumb, nil)
        if err != nil {
                return nil, err
        }

        return buf.Bytes(), nil
}



//
func DownloadMediaFromURL(url string) (io.Reader , error) {

        client := http.DefaultClient
        resp, err := client.Get(url)
        if err != nil {
                return nil, err
        }
        respByte, err := ioutil.ReadAll(resp.Body)
//      respByte := io.LimitReader(resp.Body, 1e6 * 100)

        if err != nil {
                log.Errorf("Download file problem  %v", err)
                return nil, err
        }

        r := bytes.NewReader(respByte)

        return r, nil
}


func StreamToByte(stream io.Reader) []byte {
        buf := new(bytes.Buffer)
        buf.ReadFrom(stream)
        return buf.Bytes()
}


func sendLocation(wac *wa.Conn, lat float64,lng float64, number string, reply *models.MessageItem, broadcastID *string) error {
//log.Infof("Send Text Start to %v", number)
        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.LiveLocationMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id: ID,
                },
				
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                AccuracyInMeters: 1,
		DegreesLongitude: lng,
		DegreesLatitude: lat,
		SequenceNumber: 1,
        }
	//log.Infof("Send Text Mysql insert")
        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID: ID,
                BroadcastID: broadcastID,
                From: NumberFromWidJid(wac.Info.Wid),
                To: number,
                Status: storage.SENT,
                Direction: storage.OUTGOING,
                MediaType: storage.TEXT,
                SentAt: &now,
        })
        if err != nil {
                return err
        }


        if wac == nil{
                log.Error("WAC is null GPS  failed")
                return nil
        }

        _, err = wac.Send(msg)
        if err != nil {
                log.Infof("Send GPS %v", err)
                return err
        }

        return nil
}


func SendAudioRecord(wh *WaHandler, data *models.AudioMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }
log.Infof("sendAudioRecord %v" , data.Audio.String())

        audio, err := DownloadMediaFromURL(data.Audio.String())
        if err != nil {
                return nil, err
        }
        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.AUDIO)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        work := func(item *models.NumberReplyIds) error {
                return sendAudioRecord(wh.Wac, audio, *item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}

func sendAudioRecord(wac *wa.Conn, audio io.Reader, number string, reply *models.MessageItem, broadcastID *string) error {
log.Info("sendAudioRecord")
        audioReader := audio.(*bytes.Reader)
        audioReader.Seek(0, io.SeekStart)
        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.AudioMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id: ID,
                },
                Type: "audio/ogg; codecs=opus",
                ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
                Content: audioReader,
                Ptt: true,
        }
        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID: ID,
                BroadcastID: broadcastID,
                From: NumberFromWidJid(wac.Info.Wid),
                To: number,
                Status: storage.SENT,
                Direction: storage.OUTGOING,
                MediaType: storage.AUDIO,
                SentAt: &now,
        })
        if err != nil {
                return err
        }
        _, err = wac.Send(msg)
        if err != nil {
                log.Info(err)
                return err
        }
        return nil
}


func createVCard(display, number string) string {
card := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nN:;%v;;;\nFN:%v\nTEL;type=CELL;waid=%v:+%v\nEND:VCARD", display, display, number, number)
return card
}

func SendVcard(wh *WaHandler, data *models.ContactMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }

        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.VCARD)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        work := func(item *models.NumberReplyIds) error {
	vcard := createVCard(*data.DisplayName,*data.Number)
                return sendVcard(wh.Wac, vcard, *data.DisplayName ,*item.Number, item.ReplyTo, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}

func sendVcard(wac *wa.Conn, vcard string, displayname string, number string, reply *models.MessageItem, broadcastID *string) error {
log.Info("sendVcard")

        ID := MessageID()
        JID := WidJidFromNumber(number)
        msg := wa.ContactMessage{
                Info: wa.MessageInfo{
                        RemoteJid: JID,
                        Timestamp: uint64(time.Now().Unix()),
                        Id: ID,
                },
		DisplayName: displayname,
                Vcard: vcard,
		ContextInfo: contextFromReplyID(wac, wac.Info.Wid, number, reply),
        }
        now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID: ID,
                BroadcastID: broadcastID,
                From: NumberFromWidJid(wac.Info.Wid),
                To: number,
                Status: storage.SENT,
                Direction: storage.OUTGOING,
                MediaType: storage.VCARD,
                SentAt: &now,
        })
        if err != nil {
                return err
        }
        _, err = wac.Send(msg)
        if err != nil {
                log.Info(err)
                return err
        }
        return nil
}

func SendLink(wh *WaHandler, data *models.LinkMessage) (*string, error) {
        err := ValidateNumber(data.NumberReplyIds)
        quitChannel := make(chan struct{})
        if err != nil {
                return nil, err
        }

        var broadcastID *string
        if len(data.NumberReplyIds) > 1 {
                broadcastID, err = storage.CreateBroadcast(storage.LINK)
                wh.ChBrodcastQuit[*broadcastID] = quitChannel
        }
        work := func(item *models.NumberReplyIds) error {
    
                return sendLink(wh.Wac, *data.Title ,*data.Text,*data.URL,*data.Imageurl ,*data.Description, *data.Messagetype  ,*item.Number, broadcastID)
        }
        sender(wh, broadcastID, data.NumberReplyIds, quitChannel, work)
        return broadcastID, nil
}

func sendLink(wac *wa.Conn, title string, text string, url string, imageurl string,  description string, messagetype string, number string ,broadcastID *string) error {
	if text != ""{
	text += "\n" + url
	}else{
	text = url
	}

	image, errimg := DownloadMediaFromURL(imageurl)
        if errimg != nil {
                log.Errorf("Error DownloadMediaFromURL %v", errimg)
                return errimg
        }

        var imageBuff bytes.Buffer
        thumbReader := io.TeeReader(image, &imageBuff)
        thumb, errthump := GetImageThumbnail(thumbReader)

        if errthump != nil {
                        log.Errorf("Error GetImageThumbnail %v", errthump)
                        return  errthump
        }

        ID := MessageID()
        JID := WidJidFromNumber(number)
        ts := uint64(time.Now().Unix())
        status := proto.WebMessageInfo_PENDING
        fromMe := true
	ptype := proto.ExtendedTextMessage_NONE.Enum()

	if messagetype == "video" {
	ptype = proto.ExtendedTextMessage_VIDEO.Enum()
	}

        revocation := &proto.WebMessageInfo{
                Key: &proto.MessageKey{
                        FromMe: &fromMe,
                        Id: &ID,
                        RemoteJid: &JID,
                },
                MessageTimestamp: &ts,
                Message: &proto.Message{
                ExtendedTextMessage: &proto.ExtendedTextMessage{
                        MatchedText: &url,
                        CanonicalUrl: &url,
                        Description: &description,
                        Title: &title,
                        Text: &text,
                        JpegThumbnail: thumb,
			PreviewType :  ptype,
                },
                },
                Status: &status,
        }

	now := time.Now()
        err := storage.CreateMessage(storage.WAMessage{
                WhatsappID: ID,
                BroadcastID: broadcastID,
                From: NumberFromWidJid(wac.Info.Wid),
                To: number,
                Status: storage.SENT,
                Direction: storage.OUTGOING,
                MediaType: storage.LINK,
                SentAt: &now,
        })

        if err != nil {
                return err
        }
	log.Infof("#%v",revocation)
        _, err = wac.Send(revocation)
        if err != nil {
                log.Info(err)
                return err
        }
        return nil
}
