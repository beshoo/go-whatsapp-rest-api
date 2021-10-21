package api

import (
	"encoding/json"
	"time"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/number"
	wa "bitbucket.org/rockyOO7/wa-api/whatsapp"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

//
func HasWhatApp(params number.HasWhatsAppParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return number.NewHasWhatsAppDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	phoneNumber := params.PhoneNumber
	value, err := handler.Wac.Exist(wa.WidJidFromNumber(phoneNumber))

	if err != nil {
		errorText = err.Error()
		return number.NewHasWhatsAppDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}

	response := <-value
	var resp map[string]interface{}
	if err = json.Unmarshal([]byte(response), &resp); err != nil {
		errorText = err.Error()
		return number.NewHasWhatsAppDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	hasNumber := false

	if int(resp["status"].(float64)) != 200 {
		hasNumber = false
	}
	if int(resp["status"].(float64)) == 200 {
		hasNumber = true
	}
	return number.NewHasWhatsAppOK().WithPayload(&number.HasWhatsAppOKBody{
		HasWhatsApp: &hasNumber,
	})

}

//
func IsOnline(params number.IsOnlineParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	phoneNumber := params.PhoneNumber
	value, err := handler.Wac.SubscribePresence(wa.WidJidFromNumberC(phoneNumber))
	if err != nil {
		errorText = err.Error()
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	response := <-value
	var resp map[string]interface{}
	if err = json.Unmarshal([]byte(response), &resp); err != nil {
		errorText = err.Error()
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	log.Infof("Online Status %v", resp)
	if int(resp["status"].(float64)) != 200 {
		errorText = "Presence failed"
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	time.Sleep(1 * time.Second)
	// online := handler.OnlineNumbers[phoneNumber]
	online := handler.IsOnline(phoneNumber)

	log.Infof("Number %v is %v", phoneNumber, online)
	return number.NewIsOnlineOK().WithPayload(&number.IsOnlineOKBody{
		IsOnline: &online.IsOnline,
		LastSeen: online.LastSeen,
	})
}

//
func GetChats(params number.GetChatsParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	phoneNumber := params.PhoneNumber
	if params.NumberOfMessages > 300 {
		errorText = "Max 300 messages allowed at one time"
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	if params.NumberOfMessages < 2 {
		errorText = "Min 2 messages allowed at one time"
		return number.NewIsOnlineDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	chunk := int(params.NumberOfMessages)
	// chunk := int(params.NumberOfMessages) - 1
	chats := wa.GetChatHistory(handler.Wac, &sessionID, phoneNumber, chunk, params.BeforeMessageID, params.FromMe)
	return number.NewGetChatsOK().WithPayload(chats)
}




func GetAvatar(params number.GetAvatarParams) middleware.Responder {

        type ThumbURL struct {
                EURL   string `json:"eurl"`
                Tag    string `json:"tag"`
                Status int64  `json:"status"`
        }

        errorText := ""
        sessionID := params.SessionID.String()
	number_url := params.PhoneNumber
        handler, ok := wa.Connections[sessionID]
        if !ok {
                errorText = "Invalid Session Id"
                return number.NewGetAvatarDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }

        if handler.Wac == nil {
                errorText = "Invalid Wac"
                return number.NewGetAvatarDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }


        jid := wa.WidJidFromNumber(number_url)
        thumbChan, err := handler.Wac.GetProfilePicThumb(jid)
        if err != nil {
                errorText = err.Error()
                return number.NewGetAvatarDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }
        thumb := <-thumbChan
        thumbnail := ThumbURL{}
        err = json.Unmarshal([]byte(thumb), &thumbnail)
        if err != nil {
                errorText = err.Error()
                return number.NewGetAvatarDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }
        url := ""
        if thumbnail.Status != 404 {
                url = thumbnail.EURL
        }
      
		return number.NewGetAvatarOK().WithPayload(&models.Profile{
                PhoneNumber: &number_url,
                ProfilePic:  strfmt.URI(url),
        })

}

