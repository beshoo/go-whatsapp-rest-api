package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
//	"strings"
	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"github.com/go-openapi/strfmt"
	wa "github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
)

//
func (wh *WaHandler) HandleBatteryMessage(message wa.BatteryMessage) {


//	fmt.Printf("%#v", message)
	if wh.Hook == nil {
		log.Warnf("No hook specified for %v", wh.Wac.Info.Wid)
		return
	}

	sessionID := strfmt.UUID4(*wh.SessionID)
        number := NumberFromWidJid(wh.Wac.Info.Wid)
	percent := int64(message.Percentage)
	contactMessage := models.MessageBattery{
		SessionID:   &sessionID,
                Number:      &number,
		Plugged: &message.Plugged,
		Powersave: &message.Powersave,
		Percentage: &percent,
	}

	messageBytes, err := json.Marshal(contactMessage)
	if err != nil {
		log.Errorf("Request Json Marshal failed with error %v", err)
		return
	}

	err = Retry(retrySend, time.Millisecond*500, func() error {
		return requestWith(fmt.Sprintf("%v/power/battery", *wh.Hook), "application/json", bytes.NewReader(messageBytes))
	})

	if err != nil {
		log.Errorf("Contact Request Failed with error %v", err)
	}


}
