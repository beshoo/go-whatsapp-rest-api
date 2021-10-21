package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

//
type WASession struct {
	ID          string
	ClientID    string
	ClientToken string
	ServerToken string
	EncKey      string
	MacKey      string
	Wid         string
	HookURL     *string
	ProxyURL    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}





//
type BroadcastStatus string

//
const (
	COMPLETED BroadcastStatus = "COMPLETED"
	ONGOING   BroadcastStatus = "ONGOING"
	STOPPED   BroadcastStatus = "STOPPED"
)

//
type WABroadcast struct {
	ID        string
	Status    BroadcastStatus
	MediaType MediaType
	Messages  []WAMessage
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//
type MessageStatus string

//
const (
	ERR      MessageStatus = "ERROR"
	SENT     MessageStatus = "SENT"
	RECIEVED MessageStatus = "RECIEVED"
	READ     MessageStatus = "READ"
	INCOM    MessageStatus = "INCOMING"
	HOOKED   MessageStatus = "HOOKED"
)

//
type Direction string

//
const (
	INCOMING Direction = "INCOMING"
	OUTGOING Direction = "OUTGOING"
)

//
type MediaType string

//
const (
	TEXT    MediaType = "TEXT"
	VIDEO   MediaType = "VIDEO"
	IMAGE   MediaType = "IMAGE"
	VCARD   MediaType = "VCARD"
	LINK    MediaType = "LINK"
	DOC     MediaType = "DOC"
	AUDIO   MediaType = "AUDIO"
	CONTACT MediaType = "CONTACT"
	LOC     MediaType = "LOC"
	LIVELOC MediaType = "LIVE_LOC"
)

//
type WAMessage struct {
	gorm.Model
	WhatsappID  string
	BroadcastID *string
	From        string
	To          string
	Status      MessageStatus
	Direction   Direction
	MediaType   MediaType
	SentAt      *time.Time
	RecievedAt  *time.Time
	ReadAt      *time.Time
	IncomingAt  *time.Time
	HookedAt    *time.Time
}
