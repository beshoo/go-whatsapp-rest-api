package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//
type Database string

//
const (
	MYSQL     Database = "mysql"
	POSTGRESS Database = "postgres"
	SQLITE3   Database = "sqlite3"
	MSSQL     Database = "mssql"
)

//
var db *gorm.DB

//
func Init(database Database, url string) error {
	if database == MYSQL {
		url = fmt.Sprintf("%v?parseTime=True", url)
		log.Info(url)
	}
	dbType := string(database)
	gormDB, err := gorm.Open(dbType, url)
	if err != nil {
		log.Infof("Database Init failed for %v db url %v", database, url)
		return err
	}
	db = gormDB
	db.AutoMigrate(&WASession{}, &WABroadcast{}, &WAMessage{})
	return nil
}

//
func Close() {
	log.Info("Closing Database")
	if db == nil {
		log.Warn("No database connected")
		return
	}
	err := db.Close()
	if err != nil {
		log.Errorf("Error closing %v", err)
	}
}

//
func CreateSession(session WASession) error {
	if db == nil {
		log.Warnf("No database connected %v", db)
		return nil
	}
	return db.Create(&session).Error
}

//

func UpdateSession(session WASession) {
db.LogMode(true)
	if db == nil {
		log.Warn("No database connected")
		return

	}
log.Infof("SessionId",session.ID)
db.Model(&session).Update(session)
db.LogMode(false)

}

func UpdateSessionDuplicated(session WASession) {
db.LogMode(true)
        if db == nil {
                log.Warn("No database connected")
                return

        }
log.Infof("SessionId",session.ID)

err := db.Table("wa_sessions").Where("wid = ?", session.Wid).Updates(&session).Error
        if err != nil {
                log.Errorf("Update Error %v", err)
        }

db.LogMode(false)

}



//
func DeleteSession(sessionID string) {
	if db == nil {
		log.Warnf("No database connected %v", db)
		return
	}
	log.Infof("Delete Session ID %v ", sessionID)
	err := db.Unscoped().Delete(&WASession{
		ID: sessionID,
	}).Error
	if err != nil {
		log.Errorf("Delete failed: %v", err)
	}

}





//
func UpdateSessionHook(sessionID string, hook string) error {
	if db == nil {
		log.Warn("No database connected")
		return nil
	}
	log.Infof("SessionId %v %v", sessionID, hook)
	session := &WASession{ID: sessionID}
	return db.Model(&session).Update("hook_url", hook).Error
}

//
func GetSessionForID(sessionID string) (*WASession, error) {
//db.LogMode(true)
	session := &WASession{
		ID: sessionID,
	}
	if db == nil {
		log.Warn("No database connected")
		return nil, errors.New("No database connected")
	}
	err := db.Where(session).Find(session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

//
func GetAllSessions() ([]WASession, error) {
	var sessions []WASession
	if db == nil {
		log.Warn("No database connected")
		return sessions, nil
	}
	err := db.Find(&sessions).Error
	if err != nil {
		return sessions, err
	}
	return sessions, nil
}

//
func CreateBroadcast(mediaType MediaType) (*string, error) {
	if db == nil {
		log.Warn("No database connected")
		return nil, nil
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	bID := id.String()
	b := WABroadcast{
		ID:        bID,
		MediaType: mediaType,
		Status:    ONGOING,
	}
	err = db.Create(&b).Error
	return &bID, err
}

//
func UpdateBroadcastStatus(broadcastID *string, status BroadcastStatus) error {
	if db == nil {
		log.Warn("No database connected")
		return nil
	}
	if broadcastID == nil {
		return nil
	}
	broadcast := &WABroadcast{
		ID: *broadcastID,
	}
	return db.Model(broadcast).Update("status", string(status)).Error
}

//
func CreateMessage(message WAMessage) error {
	if db == nil {
		log.Warn("No database connected")
		return nil
	}
	return db.Create(&message).Error
}

//
func MessageExist(messageID string, direction Direction) bool {
	if db == nil {
		log.Warn("No database connected")
		return false
	}
	message := &WAMessage{
		WhatsappID: messageID,
		Direction:  direction,
	}
	return !db.Where(message).First(message).RecordNotFound()
}

//
func UpdateMessageReceivedAt(messageID string, utime uint64) error {
	if db == nil {
		log.Warn("No database connected")
		return nil
	}
//db.LogMode(true)

	message := &WAMessage{
		WhatsappID: messageID,
		Direction:  OUTGOING,
	}
	t := time.Unix(int64(utime), 0)
	return db.Model(&WAMessage{}).Where(message).Updates(
		WAMessage{
			RecievedAt: &t,
			Status:     RECIEVED,
		}).Error

}

//
func UpdateMessageReadAt(messageID string, utime uint64) error {

	if db == nil {
		log.Warn("No database connected")
		return nil
	}
//db.LogMode(true)

	message := &WAMessage{
		WhatsappID: messageID,
		Direction:  OUTGOING,
	}
	t := time.Unix(int64(utime), 0)
	return db.Model(&WAMessage{}).Where(message).Updates(WAMessage{
		ReadAt: &t,
		Status: READ,
	}).Error
}

//
func UpdateMessageHookedAt(messageID string, time time.Time) {
	if db == nil {
		log.Warn("No database connected")
		return
	}
}
