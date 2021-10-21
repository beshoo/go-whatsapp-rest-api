package api

import (
	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/send"
	wa "bitbucket.org/rockyOO7/wa-api/whatsapp"
	"github.com/go-openapi/runtime/middleware"
)

//
func SendReadAck(params send.SendReadParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendReadDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}

        if handler.Wac == nil {
	   errorText = "Wac is null"
                return send.NewSendReadDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }


	remoteJID := wa.WidJidFromNumber(*params.Data.Number)
	_, err := handler.Wac.Read(remoteJID, *params.Data.MessageID)
	if err != nil {
		errorText = err.Error()
		return send.NewSendReadDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	return send.NewSendReadOK()
}

//
func SendText(params send.SendTextParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendTextDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	ID, err := wa.SendText(handler, params.Data)
	if err != nil {
		errorText = err.Error()
		return send.NewSendTextDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	broadcastID := ""
	if ID != nil {
		broadcastID = *ID
	}
	return send.NewSendTextOK().WithPayload(&models.BroadcastStatus{
		BroadcastID: broadcastID,
		Status:      models.BroadcastStatusStatusProcessing,
	})

}

//
func SendImage(params send.SendImageParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendImageDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	ID, err := wa.SendImage(handler, params.Data)
	if err != nil {
		errorText = err.Error()
		return send.NewSendImageDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	broadcastID := ""
	if ID != nil {
		broadcastID = *ID
	}
	return send.NewSendImageOK().WithPayload(&models.BroadcastStatus{
		BroadcastID: broadcastID,
		Status:      models.BroadcastStatusStatusProcessing,
	})
}

//
func SendVideo(params send.SendVideoParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendVideoDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	ID, err := wa.SendVideo(handler, params.Data)
	if err != nil {
		errorText = err.Error()
		return send.NewSendVideoDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	broadcastID := ""
	if ID != nil {
		broadcastID = *ID
	}
	return send.NewSendVideoOK().WithPayload(&models.BroadcastStatus{
		BroadcastID: broadcastID,
		Status:      models.BroadcastStatusStatusProcessing,
	})
}

//
func SendAudio(params send.SendAudioParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendAudioDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	ID, err := wa.SendAudio(handler, params.Data)

	if err != nil {
		errorText = err.Error()
		return send.NewSendAudioDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	broadcastID := ""
	if ID != nil {
		broadcastID = *ID
	}
	return send.NewSendAudioOK().WithPayload(&models.BroadcastStatus{
		BroadcastID: broadcastID,
		Status:      models.BroadcastStatusStatusProcessing,
	})
}

//
func SendDoc(params send.SendDocParams) middleware.Responder {
	errorText := ""
	sessionID := params.Data.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return send.NewSendDocDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	ID, err := wa.SendDoc(handler, params.Data, *params.Data.DocType)
	if err != nil {
		errorText = err.Error()
		return send.NewSendDocDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	broadcastID := ""
	if ID != nil {
		broadcastID = *ID
	}
	return send.NewSendDocOK().WithPayload(&models.BroadcastStatus{
		BroadcastID: broadcastID,
		Status:      models.BroadcastStatusStatusProcessing,
	})
}


func SendAudioRecord(params send.SendAudioRecordParams) middleware.Responder {
        errorText := ""
        sessionID := params.Data.SessionID.String()
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return send.NewSendAudioRecordDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        ID, err := wa.SendAudioRecord(handler, params.Data)
        if err != nil {
                errorText = err.Error()
                return send.NewSendAudioRecordDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        broadcastID := ""
        if ID != nil {
                broadcastID = *ID
        }
        return send.NewSendAudioRecordOK().WithPayload(&models.BroadcastStatus{
                BroadcastID: broadcastID,
                Status: models.BroadcastStatusStatusProcessing,
        })
}



func SendLocation(params send.SendLocationParams) middleware.Responder {
        errorText := ""
        sessionID := params.Data.SessionID.String()
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return send.NewSendLocationDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }
        ID, err := wa.SendLocation(handler, params.Data)
        if err != nil {
                errorText = err.Error()
                return send.NewSendLocationDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }
        broadcastID := ""
        if ID != nil {
                broadcastID = *ID
        }
        return send.NewSendLocationOK().WithPayload(&models.BroadcastStatus{
                BroadcastID: broadcastID,
                Status:      models.BroadcastStatusStatusProcessing,
        })

}



func SendVcard(params send.SendVcardParams) middleware.Responder {
        errorText := ""
        sessionID := params.Data.SessionID.String()
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return send.NewSendVcardDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        ID, err := wa.SendVcard(handler, params.Data)
        if err != nil {
                errorText = err.Error()
                return send.NewSendVcardDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        broadcastID := ""
        if ID != nil {
                broadcastID = *ID
        }
        return send.NewSendVcardOK().WithPayload(&models.BroadcastStatus{
                BroadcastID: broadcastID,
                Status: models.BroadcastStatusStatusProcessing,
        })
}



func SendLink(params send.SendLinkParams) middleware.Responder {
        errorText := ""
        sessionID := params.Data.SessionID.String()
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return send.NewSendLinkDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        ID, err := wa.SendLink(handler, params.Data)
        if err != nil {
                errorText = err.Error()
                return send.NewSendLinkDefault(500).WithPayload(&models.Error{
                        Code: 500,
                        Message: &errorText,
                })
        }
        broadcastID := ""
        if ID != nil {
                broadcastID = *ID
        }
        return send.NewSendLinkOK().WithPayload(&models.BroadcastStatus{
                BroadcastID: broadcastID,
                Status: models.BroadcastStatusStatusProcessing,
        })
}



/*
func SendTyping(params send.SendReadParams) middleware.Responder {
        errorText := ""
        sessionID := params.Data.SessionID.String()
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return send.NewSendReadDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }

        if handler.Wac == nil {
           errorText = "Wac is null"
                return send.NewSendReadDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }


        remoteJID := wa.WidJidFromNumber(*params.Data.Number)
        _, err := handler.Wac.wa.Presence(remoteJID, handler.PresenceComposing)
        if err != nil {
                errorText = err.Error()
                return send.NewSendReadDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }
        return send.NewSendReadOK()
}
*/
