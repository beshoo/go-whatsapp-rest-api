package whatsapp

import (
	"encoding/json"
	"fmt"
//	"time"
	"bitbucket.org/rockyOO7/wa-api/storage"
	log "github.com/sirupsen/logrus"
)

//
func (wh *WaHandler) HandleJsonMessage(message string) {
	var msg []interface{}
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Errorf("json parsing error %v", err)
		return
	}
//	 log.Info(message)
	if msg[0] == "Stream" {
		if msg[1] == "asleep" {
			go wh.NotifyConnectivity(false)
			wh.SendSignal(PAUSE)

		}
	}
	if msg[0] == "Conn" {
		go wh.NotifyConnectivity(true)
		wh.SendSignal(RESUME)
	}
	if msg[0] == "Presence" {
		updates := msg[1].(map[string]interface{})
		number := NumberFromWidJid(fmt.Sprintf("%v", updates["id"]))
		isOnline := false
		lastSeen := fmt.Sprintf("%v", updates["t"])
		available := fmt.Sprintf("%v", updates["type"])
		if available == "available" {
			isOnline = true
		}
		wh.numberOnline(number, isOnline, lastSeen)

	}

        if msg[0] == "Cmd" {
	jsonres := msg[1].(map[string]interface{})
                if jsonres["type"] == "disconnect" {
			wh.reloginCounter(wh.Wac.Info.Wid , 1)
                       // go wh.NotifyConnectivity(false)
                       // wh.SendSignal(PAUSE)
			if value, ok := jsonres["kind"]; ok {
				if value == "replaced" {

	                                log.Info("Kind replace")					
					go wh.NotifyDisconnect(wh.Wac.Info.Wid)
					}else{

					log.Infof("Json connection not handeld var %v",value)
					
					}

				} else {
                                        log.Info("Json Logout")
                                        wh.logout()
                                        go wh.NotifyLogout(wh.Wac.Info.Wid)

				}
                }
        }



	if msg[0] == "Msg" {
		acknowledgements := msg[1].(map[string]interface{})
		cmd := acknowledgements["cmd"]
		// log.Info(cmd)
		if cmd != "ack" {
			return
		}
		ack := acknowledgements["ack"].(float64)
		whatsappID := acknowledgements["id"]
		time := acknowledgements["t"].(float64)
		timeInt := uint64(time)
		from := acknowledgements["from"]
		to := acknowledgements["to"]
		// log.Info(ack)
		if ack == 2.0 {
			storage.UpdateMessageReceivedAt(whatsappID.(string), timeInt)

			go wh.NotifyReceive(from.(string),to.(string))

		} else if ack == 3.0 {
			storage.UpdateMessageReadAt(whatsappID.(string), timeInt)
		}
	}
}
