package whatsapp

import (
	"fmt"
	"time"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
	wa "github.com/Rhymen/go-whatsapp"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

//
type HistoryHandler struct {
	Wac       *wa.Conn
	SessionID *string
	items     []*models.ChatItem
}

func (h *HistoryHandler) messageInfo(info wa.MessageInfo) models.MessageInfo {
	remoteJID := NumberFromWidJid(info.RemoteJid)
	senderID := NumberFromWidJid(h.Wac.Info.Wid)
	timeStamp := time.Unix(int64(info.Timestamp), 0)

	from := remoteJID
	to := senderID
	if info.FromMe {
		from = senderID
		to = remoteJID
	}

	return models.MessageInfo{
		Owner: &models.MessageInfoOwner{
			Number:    &senderID,
			SessionID: h.SessionID,
		},
		FromMe:    &info.FromMe,
		MessageID: &info.Id,
		PushName:  info.Source.GetPushName(),
		From:      &from,
		To:        &to,
		Status:    int64(info.Status),
		Timestamp: strfmt.DateTime(timeStamp),
	}
}

func (h *HistoryHandler) context(context wa.ContextInfo) models.MessageContext {
	return models.MessageContext{
		IsForwarded:     &context.IsForwarded,
		Participant:     context.Participant,
		QuotedMessageID: context.QuotedMessageID,
	}
}

//
func GetChatHistory(wac *wa.Conn, sessionID *string, number string, chunk int, beforeMessageID *string, fromMe *bool) []*models.ChatItem {
	handler := &HistoryHandler{
		Wac:       wac,
		SessionID: sessionID,
	}
	jid := WidJidFromNumber(number)
	if beforeMessageID != nil {
		log.Info("before messageID %v", *beforeMessageID)
		wac.LoadChatMessages(jid, chunk, *beforeMessageID, *fromMe, false, handler)
	} else {
		wac.LoadChatMessages(jid, chunk, "", true, false, handler)
	}
	return handler.items
}

//
func (h *HistoryHandler) ShouldCallSynchronously() bool {
	return true
}

//
func (h *HistoryHandler) HandleError(err error) {
	log.Infof("Error occured while retrieving chat history: %v", err)
}

//
func (h *HistoryHandler) HandleTextMessage(message wa.TextMessage) {

	messageInfo := h.messageInfo(message.Info)
	messageContext := h.context(message.ContextInfo)

	chatItem := models.ChatItem{
		MessageType:    "TEXT",
		MessageInfo:    &messageInfo,
		MessageContext: &messageContext,
		Text:           message.Text,
	}
	h.items = append(h.items, &chatItem)
}

//
func (h *HistoryHandler) HandleImageMessage(message wa.ImageMessage) {
	messageInfo := h.messageInfo(message.Info)
	messageContext := h.context(message.ContextInfo)

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
	chatItem := models.ChatItem{
		MessageType:    "IMAGE",
		MessageInfo:    &messageInfo,
		MessageContext: &messageContext,
		Caption:        message.Caption,
		ImageThumb:     thumbnailURL,
		Image:          imageURL,
	}
	h.items = append(h.items, &chatItem)
}

//
func (h *HistoryHandler) HandleVideoMessage(message wa.VideoMessage) {
	messageInfo := h.messageInfo(message.Info)
	messageContext := h.context(message.ContextInfo)
	video, err := message.Download()
	if err != nil {
		log.Errorf("Image download failed %v ", err)
		return
	}
	videoURL, err := URLFromBytes(video, GetFileNameFromType(message.Info.Id, message.Type))
	if err != nil {
		log.Errorf("Video url failed %v ", err)
		return
	}

	thumbnailURL, err := URLFromBytes(message.Thumbnail, fmt.Sprintf("%v_thumb.png", message.Info.Id))
	if err != nil {
		log.Errorf("Thumb url failed %v ", err)
		return
	}

	chatItem := models.ChatItem{
		MessageType:    "VIDEO",
		MessageInfo:    &messageInfo,
		MessageContext: &messageContext,
		Caption:        message.Caption,
		VideoThumb:     videoURL,
		Video:          thumbnailURL,
	}
	h.items = append(h.items, &chatItem)
}

//
func (h *HistoryHandler) HandleAudioMessage(message wa.AudioMessage) {
	messageInfo := h.messageInfo(message.Info)
	messageContext := h.context(message.ContextInfo)
	audio, err := message.Download()
	if err != nil {
		log.Errorf("Audio download failed %v ", err)
		return
	}

	audioURL, err := URLFromBytes(audio, GetFileNameFromType(message.Info.Id, message.Type))
	if err != nil {
		log.Errorf("Audio url failed %v ", err)
		return
	}

	chatItem := models.ChatItem{
		MessageType:    "AUDIO",
		MessageInfo:    &messageInfo,
		MessageContext: &messageContext,
		AudioLength:    fmt.Sprintf("%d", message.Length),
		Audio:          audioURL,
	}
	h.items = append(h.items, &chatItem)

}

//
func (h *HistoryHandler) HandleDocumentMessage(message wa.DocumentMessage) {
	messageInfo := h.messageInfo(message.Info)
	messageContext := h.context(message.ContextInfo)
	doc, err := message.Download()
	if err != nil {
		log.Errorf("Doc download failed %v", err)
	}
	docURL, err := URLFromBytes(doc, GetFileNameFromType(message.Info.Id, message.Type))
	if err != nil {
		log.Errorf("Doc url failed %v ", err)
		return
	}

	chatItem := models.ChatItem{
		MessageType:    "DOC",
		MessageInfo:    &messageInfo,
		MessageContext: &messageContext,
		DocTitle:       message.Title,
		PageCount:      int64(message.PageCount),
		Doc:            docURL,
	}

	h.items = append(h.items, &chatItem)
}
