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
func (wh *WaHandler) HandleDocumentMessage(message wa.DocumentMessage) {

if  message.Info.Timestamp  < StartTime  {
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
		MediaType:  storage.DOC,
		IncomingAt: &now,
	})

	if wh.Hook == nil {
		log.Warnf("No hook specified for %v", wh.Wac.Info.Wid)
		return
	}

	messageInfo := wh.messageInfo(message.Info)
	messageContext := wh.context(message.ContextInfo)
	doc, err := message.Download()
	if err != nil {
		log.Errorf("Doc download failed %v", err)
	}
	docURL, err := URLFromBytes(doc, GetFileNameFromType(message.Info.Id, message.Type))
	if err != nil {
		log.Errorf("Doc url failed %v ", err)
		return
	}
	docMessage := models.MessageDoc{
		ContextInfo: &messageContext,
		MessageInfo: &messageInfo,
		PageCount:   fmt.Sprintf("%d", message.PageCount),
		Title:       message.Title,
		Doc:         docURL,
	}

	messageBytes, err := json.Marshal(docMessage)
	if err != nil {
		log.Errorf("Doc Request Json Marshal failed with error %v", err)
		return
	}
	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(fmt.Sprintf("%v/message/doc", *wh.Hook), "application/json", bytes.NewReader(messageBytes))
	})

	if err != nil {
		log.Errorf("Doc Request Failed with error %v", err)
	}
}
