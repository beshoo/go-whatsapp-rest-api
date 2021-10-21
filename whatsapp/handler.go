package whatsapp

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
	"strconv"
	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/storage"
	wa "github.com/Rhymen/go-whatsapp"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

var fromMe = flag.Bool("fromMe", false, "Should recieve self messages")
var nodePort = flag.String("nodePort","8090","Pass node port")
func (wh *WaHandler) numberOnline(number string, isOnline bool, lastSeen string) {
	wh.rwm.Lock()
	wh.OnlineNumbers[number] = OnlineStatus{
		IsOnline: isOnline,
		LastSeen: lastSeen,
	}
	wh.rwm.Unlock()
}

//

func (wh *WaHandler) reloginCounter(number string , count int) {
	log.Infof("Relogin Counter %v %v", number, count)
        wh.rwm.Lock()
        wh.ReloginCounter[number] = ReloginStatus{
                ReloginCount: count,
        }
        wh.rwm.Unlock()
}


//
func (wh *WaHandler) IsOnline(number string) OnlineStatus {
	wh.rwm.Lock()
	defer wh.rwm.Unlock()
	online, ok := wh.OnlineNumbers[number]
	if !ok {
		return OnlineStatus{
			false,
			online.LastSeen,
		}
	}
	return online
}

func (wh *WaHandler) logout() {
	storage.DeleteSession(*wh.SessionID)
	if wh.State != STOP {
		wh.Signal <- STOP
	}
	delete(Connections, *(wh.SessionID))
	wh.Wac.RemoveHandlers()
	wh.Wac.Disconnect()
}

//
func (wh *WaHandler) messageInfo(info wa.MessageInfo) models.MessageInfo {
	remoteJID := NumberFromWidJid(info.RemoteJid)
	senderID := NumberFromWidJid(wh.Wac.Info.Wid)
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
			SessionID: wh.SessionID,
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

//
/*
func (wh *WaHandler) context(context wa.ContextInfo) models.MessageContext {
	return models.MessageContext{
		IsForwarded:     &context.IsForwarded,
		Participant:     context.Participant,
		QuotedMessageID: context.QuotedMessageID,
	}
}
*/

func (wh *WaHandler) context(context wa.ContextInfo) models.MessageContext {
  p_url := ""
      var p_price int64;
      p_currency_code := ""
      p_id := ""
      p_title :=""
      p_description:= ""
      p_retailer_id:= ""
      p_type := false
      
    if context.QuotedMessage != nil {
      QuotedMessage :=  context.QuotedMessage;
      if QuotedMessage.ProductMessage != nil {
              p_type = true
              ProductMessage := QuotedMessage.ProductMessage.Product;

	      if ProductMessage.Url != nil{
                p_url = *ProductMessage.Url
              }

              if ProductMessage.ProductId != nil{ 
              p_id = *ProductMessage.ProductId
              }
             
	      if ProductMessage.CurrencyCode != nil{ 
              p_currency_code = *ProductMessage.CurrencyCode
              }
              
              if ProductMessage.Title != nil {
              p_title = *ProductMessage.Title
              }
              
              if ProductMessage.Description != nil {
              p_description = *ProductMessage.Description
              }
              
              if ProductMessage.RetailerId != nil {
                p_retailer_id = *ProductMessage.RetailerId
              }

              if ProductMessage.PriceAmount1000 != nil {
              p_price = *ProductMessage.PriceAmount1000
              }

          }
      }
    
	return models.MessageContext{
		IsForwarded:     &context.IsForwarded,
		Participant:     context.Participant,
		QuotedMessageID: context.QuotedMessageID,
        ProductMessageURL: p_url,
        ProductMessageID: p_id,
        ProductMessageTitle: p_title,
        ProductMessageDescription : p_description,
        ProductMessageRetailerID:p_retailer_id,
        ProductMessageType : p_type,
        ProductMessageCurrencyCode:p_currency_code,
        ProductMessagePrice: strconv.Itoa(int(p_price)),
        
	}
}

func requestWith(url, contentType string, body io.Reader) error {
	client := http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Connection", "close")
	req.Close = true
	if err != nil {
		log.Errorf("Request Creation failed %v", err)
		return err
	}
	req.Header.Add("Content-Type", contentType)
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	// request is successful or url is not there so mark it as read
	if res.StatusCode == 200 || res.StatusCode == 404 {
		//log.Errorf("Request Status  %v", res.StatusCode)
		return nil
	}
	return fmt.Errorf("Request failed with status %d", res.StatusCode)
}

//
func (wh *WaHandler) NotifyLogout(number string) {
	if wh.Hook == nil {
		log.Warnf("No hook specified for %v", number)
		return
	}

	sessionID := strfmt.UUID4(*wh.SessionID)
	number = NumberFromWidJid(number)
	now := strfmt.DateTime(time.Now())

	logoutMessage := models.NotifyLogout{
		SessionID: &sessionID,
		Number:    &number,
		Timestamp: &now,
	}
	url := fmt.Sprintf("%v/notify/logout", *wh.Hook)
	messageBytes, err := json.Marshal(logoutMessage)
	if err != nil {
		log.Errorf("Logout hook Request Json Marshal failed with error %v", err)
		return
	}
	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(url, "application/json", bytes.NewReader(messageBytes))
	})
	if err != nil {
		log.Errorf("Logout hook Request Failed with error %v", err)
	}
}

//

func (wh *WaHandler) NotifyDisconnect(number string) {
        if wh.Hook == nil {
                log.Warnf("No hook specified for %v", number)
                return
        }

        sessionID := strfmt.UUID4(*wh.SessionID)
        number = NumberFromWidJid(number)
        now := strfmt.DateTime(time.Now())

        logoutMessage := models.NotifyLogout{
                SessionID: &sessionID,
                Number:    &number,
                Timestamp: &now,
        }
        url := fmt.Sprintf("%v/notify/disconnect", *wh.Hook)
        messageBytes, err := json.Marshal(logoutMessage)
        if err != nil {
                log.Errorf("Logout hook Request Json Marshal failed with error %v", err)
                return
        }
        err = Retry(retrySend, time.Millisecond*500, func() error {
                return requestWith(url, "application/json", bytes.NewReader(messageBytes))
        })
        if err != nil {
                log.Errorf("Logout hook Request Failed with error %v", err)
        }
}


//
func (wh *WaHandler) NotifyConnectivity(isConnected bool) {

        if wh.Wac == nil {
                return
        }


        if wh.Wac.Info == nil {
                return
        }

	if wh.Hook == nil {
		log.Warnf("No hook specified for %v", wh.Wac.Info.Wid)
		return
	}
	sessionID := strfmt.UUID4(*wh.SessionID)
	number := NumberFromWidJid(wh.Wac.Info.Wid)

	notifyMessage := models.NotifyConnectivity{
		SessionID:   &sessionID,
		Number:      &number,
		IsConnected: &isConnected,
	}

	url := fmt.Sprintf("%v/notify/connectivity", *wh.Hook)
	messageBytes, err := json.Marshal(notifyMessage)
	if err != nil {
		log.Errorf("Logout hook Request Json Marshal failed with error %v", err)
		return
	}
	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(url, "application/json", bytes.NewReader(messageBytes))
	})
	if err != nil {
		log.Errorf("Logout hook Request Failed with error %v", err)
	}
}



func (wh *WaHandler) NotifyReceive(number string,to string) {
        if wh.Hook == nil {
                log.Warnf("No hook specified for %v", number)
                return
        }

        sessionID := strfmt.UUID4(*wh.SessionID)
        number = NumberFromWidJid(number)
	to_number := NumberFromWidJid(to)
        now := strfmt.DateTime(time.Now())

        logoutMessage := models.NotifyReceive{
                SessionID: &sessionID,
                Number:    &number,
		To:    &to_number,
                Timestamp: &now,
        }
        url := fmt.Sprintf("%v/notify/receive", *wh.Hook)
        messageBytes, err := json.Marshal(logoutMessage)
        if err != nil {
                log.Errorf("notify/receive hook Request Json Marshal failed with error %v", err)
                return
        }
        err = Retry(retrySend, time.Millisecond*500, func() error {
                return requestWith(url, "application/json", bytes.NewReader(messageBytes))
        })
        if err != nil {
                log.Errorf("Logout hook Request Failed with error %v", err)
        }
}



func (wh *WaHandler) NotifyScanQr() {

        sessionID := strfmt.UUID4(*wh.SessionID)
        now := strfmt.DateTime(time.Now())

        logoutMessage := models.NotifyReceive{
                SessionID: &sessionID,
                Timestamp: &now,
        }

        url := fmt.Sprintf("%v:%v", "https://mujeebk.com" , *nodePort)
        messageBytes, err := json.Marshal(logoutMessage)
        if err != nil {
                log.Errorf("notify/receive hook Request Json Marshal failed with error %v", err)
                return
        }
        err = Retry(retrySend, time.Millisecond*500, func() error {
                return requestWith(url, "application/json", bytes.NewReader(messageBytes))
        })
        if err != nil {
                log.Errorf("Logout hook Request Failed with error %v", err)
        }
}

