package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"strings"
	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/storage"
	wa "github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
)
//var StartTime = uint64(time.Now().Unix())
//
func (wh *WaHandler) HandleImageMessage(message wa.ImageMessage) {

if  message.Info.Timestamp < StartTime {
                return
        }

if strings.Contains(message.Info.RemoteJid, "@g.us") {
       return
}

	if message.Info.FromMe {
		return
	}

if uint64(time.Now().Unix()) > message.Info.Timestamp {
receiveTimeDef := uint64(time.Now().Unix()) - message.Info.Timestamp

if  receiveTimeDef  > allowTimeDef  {
                log.Infof("Old MessageID %v %v ",wh.Wac.Info.Wid ,message.Info.Id)
                return
        }
}
	from := NumberFromWidJid(message.Info.RemoteJid)
	to := NumberFromWidJid(wh.Wac.Info.Wid)
	if from == "status" || to == "status" {
		return
	}
	if storage.MessageExist(message.Info.Id, storage.INCOMING) {
		return
	}
	now := time.Now()
	storage.CreateMessage(storage.WAMessage{
		WhatsappID: message.Info.Id,
		From:       NumberFromWidJid(message.Info.RemoteJid),
		To:         NumberFromWidJid(wh.Wac.Info.Wid),
		Status:     storage.INCOM,
		Direction:  storage.INCOMING,
		MediaType:  storage.IMAGE,
		IncomingAt: &now,
	})

	if wh.Hook == nil {
		log.Warnf("No hook specified for %v", wh.Wac.Info.Wid)
		return
	}

	messageInfo := wh.messageInfo(message.Info)
	messageContext := wh.context(message.ContextInfo)
	image, err := message.Download()
	if err != nil {
		log.Errorf("Image download failed %v ", err)
		return
	}

	imageURL, err := URLFromBytes(image, GetFileNameFromType(message.Info.Id, message.Type))
	if err != nil {
		log.Errorf("Image url failed %v ", err)
		return
	}

	thumbnailURL, err := URLFromBytes(message.Thumbnail, fmt.Sprintf("%v_thumb.png", message.Info.Id))
	if err != nil {
		log.Errorf("Thumb url failed %v ", err)
		return
	}
	imageMessage := models.MessageImage{
		ContextInfo: &messageContext,
		MessageInfo: &messageInfo,
		Caption:     message.Caption,
		Image:       &imageURL,
		Thumbnail:   &thumbnailURL,
	}
	messageBytes, err := json.Marshal(imageMessage)
	if err != nil {
		log.Errorf("Image Request Json Marshal failed with error %v", err)
		return
	}

	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(fmt.Sprintf("%v/message/image", *wh.Hook), "application/json", bytes.NewReader(messageBytes))
	})

	if err != nil {
		log.Errorf("Image Request Failed with error %v", err)
	}

}
