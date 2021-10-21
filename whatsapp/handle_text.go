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

var StartTime = uint64(time.Now().Unix())

//
func (wh *WaHandler) HandleTextMessage(message wa.TextMessage) {

        if !*fromMe && message.Info.FromMe {

                return
        }

//log.Infof("Message status %v",message.Info.Status)
if strings.Contains(message.Info.RemoteJid, "@g.us") {
       return
}


if  message.Info.Timestamp  < StartTime  {
		return
	}


if uint64(time.Now().Unix()) > message.Info.Timestamp {

        receiveTimeDef := uint64(time.Now().Unix()) - message.Info.Timestamp

	if  receiveTimeDef  > allowTimeDef  {
		log.Infof("Old MessageID %v %v %v sec %v %v - %v",wh.Wac.Info.Wid ,message.Info.Id, message.Text , receiveTimeDef , uint64(time.Now().Unix()) , message.Info.Timestamp)
                return
        }
}



//fmt.Printf("%#v", message) 
	remoteJID := NumberFromWidJid(message.Info.RemoteJid)
	senderID := NumberFromWidJid(wh.Wac.Info.Wid)

	from := remoteJID
	to := senderID
	direction := storage.INCOMING
	status := storage.INCOM

	if message.Info.FromMe {
		from = senderID
		to = remoteJID
		direction = storage.OUTGOING
		status = storage.SENT
	}
	if from == "status" || to == "status" {
		return
	}

	log.Infof("MessageID %v %v", message.Info.Id, message.Text)
	if storage.MessageExist(message.Info.Id, direction) {
		log.Info("Message already exists")
		return
	}
	if !message.Info.FromMe {
		log.Infof("number  %v", wh.Wac.Info.Wid)
	}

	//wh.Wac.Read(message.Info.RemoteJid, message.Info.Id)
	now := time.Now()
	storage.CreateMessage(storage.WAMessage{
		WhatsappID: message.Info.Id,
		From:       from,
		To:         to,
		Status:     status,
		Direction:  direction,
		MediaType:  storage.TEXT,
		IncomingAt: &now,
	})

	if wh.Hook == nil {
		log.Warnf("No hook specified for %v %v", wh.Wac.Info.Wid,*wh.SessionID)
		return
	}

	messageInfo := wh.messageInfo(message.Info)
	messageContext := wh.context(message.ContextInfo)
	textMessage := models.MessageText{
		ContextInfo: &messageContext,
		MessageInfo: &messageInfo,
		Text:        &message.Text,
	}
	messageBytes, err := json.Marshal(textMessage)
	if err != nil {
		log.Errorf("%v Text Request Json Marshal failed with error %v", wh.Wac.Info.Wid, err)
		return
	}
	//log.Errorf("POST : %v", bytes.NewReader(messageBytes))
	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(fmt.Sprintf("%v/message/text", *wh.Hook), "application/json", bytes.NewReader(messageBytes))
	})

	if err != nil {
		log.Errorf("%v Text Request Failed with error %v", err)
	}

}
