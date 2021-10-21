package api

import (
	"encoding/json"
	"fmt"
	"time"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/profile"
	"bitbucket.org/rockyOO7/wa-api/storage"
	wa "bitbucket.org/rockyOO7/wa-api/whatsapp"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

//
func ScanQr(params profile.ScanQrParams) middleware.Responder {
	qr := make(chan string)
	errCh := make(chan error)

	go func() {
		sessionID := params.SessionID.String()
		handler, err := wa.Login(qr, params.ProxyURL, sessionID)
		if err != nil {
			errCh <- err
		}
		wa.Connections[sessionID] = handler
	}()

	select {
	case err := <-errCh:
		errText := err.Error()
		return profile.NewScanQrDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errText,
		})
	case qrText := <-qr:
		return profile.NewScanQrOK().WithPayload(&models.QRCode{
			Base64: qrText,
		})

	}

}

//
func GetProfile(params profile.ProfileParams) middleware.Responder {

	type ThumbURL struct {
		EURL   string `json:"eurl"`
		Tag    string `json:"tag"`
		Status int64  `json:"status"`
	}

	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewProfileDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}


        if handler.Wac.Info == nil {
                errorText = "Invalid INFO wid"
                return profile.NewProfileDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }


	jid := wa.WidJidFromNumber(wa.NumberFromWidJid(handler.Wac.Info.Wid))
	thumbChan, err := handler.Wac.GetProfilePicThumb(jid)
	if err != nil {
		errorText = err.Error()
		return profile.NewProfileDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	thumb := <-thumbChan
	thumbnail := ThumbURL{}
	err = json.Unmarshal([]byte(thumb), &thumbnail)
	if err != nil {
		errorText = err.Error()
		return profile.NewProfileDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	url := ""
	if thumbnail.Status != 404 {
		url = thumbnail.EURL
	}
	phoneNumber := fmt.Sprintf("+%v", wa.NumberFromWidJid(handler.Wac.Info.Wid))
	return profile.NewProfileOK().WithPayload(&models.Profile{
		PhoneNumber: &phoneNumber,
		ProfilePic:  strfmt.URI(url),
	})
}

//
func SetHook(params profile.SetHookParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]

	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewSetHookDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	hookURL := params.HookURL.String()
	err := storage.UpdateSessionHook(sessionID, hookURL)
	if err != nil {
		errorText = "Hook update failed"
		return profile.NewSetHookDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	handler.Hook = &hookURL
	log.Infof("SetHook %v", handler.Hook)
	return profile.NewSetHookOK()
}

//
func GetContacts(params profile.GetContactsParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	//name := ""
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewGetContactsDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	var contacts []*models.ContactItem

	for _, item := range handler.Wac.Store.Contacts {

        //log.Infof("Name %#v", item)
/*
		name := item.Name
		
                if name == "" {
                name = item.Short
                }			
*/

name := "Unknown"

if item.Name != "" {

name = item.Name

}else if item.Notify != "" {

name = item.Notify

}else if item.Short != "" {

 name = item.Short

}

  	        number := wa.NumberFromWidJid(item.Jid)
		contact := models.ContactItem{
			Name:   &name,
			Number: &number,
		}
		contacts = append(contacts, &contact)
	}
	return profile.NewGetContactsOK().WithPayload(contacts)
}

//
func IsConnected(params profile.IsConnectedParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewIsConnectedDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	isConnected, err := handler.Wac.AdminTest()
	if err != nil {
		isConnected = false
		return profile.NewIsConnectedOK().WithPayload(&profile.IsConnectedOKBody{
			IsConnected: &isConnected,
		})
	}
	return profile.NewIsConnectedOK().WithPayload(&profile.IsConnectedOKBody{
		IsConnected: &isConnected,
	})
}

//
func Connect(params profile.ConnectParams) middleware.Responder {

	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewConnectDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}

	session, err := storage.GetSessionForID(sessionID)
	if err != nil {
		errorText = err.Error()
		return profile.NewConnectDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}


	isConnected, err := handler.Wac.AdminTest()
        if err == nil && isConnected {
                isConnected = true
                return profile.NewIsConnectedOK().WithPayload(&profile.IsConnectedOKBody{
                        IsConnected: &isConnected,
                })
        }


	newSession, err := handler.Wac.RestoreWithSession(wa.SessionFromWASession(*session))
	if err != nil {
		errorText = err.Error()
	
	handler.SendSignal(wa.STOP)
        log.Info("After logout Signal")
        time.Sleep(time.Millisecond * 1000)
	handler.Wac.Logout()
		return profile.NewConnectDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	handler.Signal <- wa.RESUME
	storage.UpdateSession(wa.WASessionFromSession(sessionID, session.ProxyURL, session.HookURL, newSession))
	return profile.NewConnectOK()
}

//
func Disconnect(params profile.DisconnectParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]

	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewDisconnectDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}

	handler.SendSignal(wa.PAUSE)
	time.Sleep(time.Millisecond * 1000)
	_, err := handler.Wac.Disconnect()
	if err != nil {
		errorText = err.Error()
		return profile.NewDisconnectDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	return profile.NewDisconnectOK()
}

//
func Logout(params profile.LogoutParams) middleware.Responder {
	errorText := ""
	sessionID := params.SessionID.String()
	handler, ok := wa.Connections[sessionID]
	if !ok {
		errorText = "Invalid Session Id"
		return profile.NewLogoutDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}
	log.Info("Before logout Signal")
	handler.SendSignal(wa.STOP)
	log.Info("After logout Signal")
	time.Sleep(time.Millisecond * 1000)

       if handler.Wac == nil {
           errorText = "Wac is null"
                return profile.NewLogoutDefault(500).WithPayload(&models.Error{
                        Code:    500,
                        Message: &errorText,
                })
        }


	err := handler.Wac.Logout()
	if err != nil {
		errorText = err.Error()
		return profile.NewLogoutDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errorText,
		})
	}

	return profile.NewLogoutOK()
}



