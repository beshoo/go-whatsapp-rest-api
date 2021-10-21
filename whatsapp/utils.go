package whatsapp

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
	"bitbucket.org/rockyOO7/wa-api/storage"
	wa "github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
)

//
func TimeStringFromUInt64(val uint64) string {
	t := time.Unix(int64(val), 0)
	return t.Format(time.RFC3339)
}

//
func NumberFromWidJid(numberID string) string {
	return numberID[:strings.Index(numberID, "@")]
}

//
func WidJidFromNumber(number string) string {
	return number + "@s.whatsapp.net"
}

//
func WidJidFromNumberC(number string) string {
	return number + "@c.us"
}

//
func ValidateNumber(numberReplyIds []*models.NumberReplyIds) error {
	reg, _ := regexp.Compile("^[0-9]+$")
	for _, item := range numberReplyIds {
		if !reg.MatchString(*item.Number) {
			return fmt.Errorf("Invalid phone number %v", *item.Number)
		}
//                fmt.Printf("Number:%v", *item.Number)

	}
	return nil
}

//
func MessageID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return strings.ToUpper(hex.EncodeToString(b))
}

//
func Retry(attempts int, sleep time.Duration, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = f()
		if err == nil {
			return nil
		}
		jitter := time.Duration(rand.Int63n(int64(sleep)))
		sleep = sleep + jitter/2
		time.Sleep(sleep)
	}
	return fmt.Errorf("Retry failed after %d attempts, last err: %v", attempts, err)
}

//
func GetFileNameFromType(name, fileType string) string {
	if strings.Contains(fileType, ";") {
		fileType = fileType[:strings.Index(fileType, ";")]
	}
	fileType = fileType[strings.Index(fileType, "/")+1:]
	name = fmt.Sprintf("%v.%v", name, fileType)
	log.Infof("Name %v", name)
	return name
}

//
func FromBase64(data string) []byte {
	bytes, _ := base64.StdEncoding.DecodeString(data)
	return bytes
}

//
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//
func WASessionFromSession(sessionID string, proxyURL, hook *string, session wa.Session) storage.WASession {
	return storage.WASession{
		ID:          sessionID,
		HookURL:     hook,
		ClientID:    session.ClientId,
		ClientToken: session.ClientToken,
		ServerToken: session.ServerToken,
		EncKey:      ToBase64(session.EncKey),
		MacKey:      ToBase64(session.MacKey),
		Wid:         session.Wid,
		ProxyURL:    proxyURL,
	}
}

//
func SessionFromWASession(session storage.WASession) wa.Session {
	return wa.Session{
		ClientId:    session.ClientID,
		ClientToken: session.ClientToken,
		ServerToken: session.ServerToken,
		EncKey:      FromBase64(session.EncKey),
		MacKey:      FromBase64(session.MacKey),
		Wid:         session.Wid,
	}
}

//
func URLFromBytes(bytes []byte, fileName string) (string, error) {
	fileName = fmt.Sprintf("static/media/%v", fileName)
	log.Info(fileName)
	mediaPath := filepath.Join(".", "static/media")
	err := os.MkdirAll(mediaPath, os.ModePerm.Perm())
	if err != nil {
		return "", err
	}
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return "", err
	}
	_, err = f.Write(bytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/%v", fileName), nil
}
