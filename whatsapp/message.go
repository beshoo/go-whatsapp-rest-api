package whatsapp

import (
	wa "github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	log "github.com/sirupsen/logrus"
)

//
type MessageHandler struct {
	Wac     *wa.Conn
	message *proto.WebMessageInfo
}

//
func GetMessageForID(wac *wa.Conn, number, messageID string, fromMe bool) *proto.Message {
log.Infof("GetMessageForID")
	handler := &MessageHandler{
		Wac: wac,
	}
	jid := WidJidFromNumber(number)
	wac.LoadChatMessages(jid, 1, messageID, fromMe, false, handler)
	messageID = handler.message.GetKey().GetId()
	fromMe = handler.message.GetKey().GetFromMe()
	wac.LoadChatMessages(jid, 1, messageID, fromMe, true, handler)
	messageID = handler.message.GetKey().GetId()
	return handler.message.GetMessage()
}

//
func (h *MessageHandler) ShouldCallSynchronously() bool {
	return true
}

//
func (h *MessageHandler) HandleError(err error) {
	log.Infof("The Error occured while retrieving message: %v", err)
}

//
func (h *MessageHandler) HandleRawMessage(m *proto.WebMessageInfo) {
	h.message = m
}
