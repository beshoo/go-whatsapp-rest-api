package whatsapp

import (
	"flag"
	"net/http"
	"net/url"
	"sync"
//	"errors"
	"time"
//	"fmt"
	"bitbucket.org/rockyOO7/wa-api/storage"
	wa "github.com/Rhymen/go-whatsapp"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

var longName = "Windows"
var shortName = "Chrome"
var clientVersion = "10"
var rate = flag.Int("rate", 200, "Max number of messages per minute")
var timeoutInt = flag.Int("timeout", 20 , "Timeout in seconds")
var retrySend = 2
var allowTimeDef = uint64(60*5)
//var timeout = time.Duration(*timeoutInt) * time.Second

//
type OnlineStatus struct {
	IsOnline bool
	LastSeen string
}

type ReloginStatus struct {
        ReloginCount int
}



//
type WaHandler struct {
	Wac             *wa.Conn
	lastMessageTime uint64
	SessionID       *string
	Hook            *string
	Session         *wa.Session
	State           Signal
	Signal          chan Signal
	Queue           chan func() error
	ChBrodcastQuit  map[string]chan struct{}
	OnlineNumbers   map[string]OnlineStatus

	ReloginCounter  map[string]ReloginStatus
	rwm             sync.Mutex
}


type version struct {
	major int
	minor int
	patch int
}

//
var Connections = make(map[string]*WaHandler)

//
type Signal string

//
const (
	RECONNECT Signal = "RECONNECT"
	STOP      Signal = "STOP"
	RESUME    Signal = "RESUME"
	PAUSE     Signal = "PAUSE"
)

//
func (wh *WaHandler) SendSignal(signal Signal) {
	if wh.State != signal {
		wh.Signal <- signal
	}
}

//
func StartWithSession(sessions []storage.WASession) {
	for _, waSession := range sessions {
		//log.Infof("Get Session: %v", waSession)
		session := SessionFromWASession(waSession)
		go func(waSession storage.WASession) {
                        //var counter int
			attempt := 500
                        counter := 1
			Retry(attempt, time.Millisecond*1000, func() error {
				wac, err := NewConn(waSession.ProxyURL)
				if err != nil {
					log.Errorf("StartWithSession Connect failed%v for wid %v", err, session.Wid)
					return err
				}
				wac.SetClientName(longName, shortName , clientVersion)
				sessionID := waSession.ID
				hook := waSession.HookURL
				wh := WaHandler{
					Wac:             wac,
					lastMessageTime: uint64(time.Now().Unix()),
					SessionID:       &sessionID,
					Hook:            hook,
					State:           RESUME,
					Signal:          make(chan Signal),
					Queue:           make(chan func() error),
					ChBrodcastQuit:  make(map[string]chan struct{}),
					OnlineNumbers:   make(map[string]OnlineStatus),
					ReloginCounter:  make(map[string]ReloginStatus),
				}
				newSession, err := wac.RestoreWithSession(session)
				if err != nil {
					if err.Error() == "admin login responded with 401" {
						// Delete if 401
						log.Infof("Admin login responded with 401")
						go wh.NotifyLogout(waSession.Wid)
						storage.DeleteSession(waSession.ID)
						log.Infof("Delete session for number %v", waSession.Wid)
						return nil
					}
//					log.Errorf("%v :: %v ReconnectForHandler :: Restore failed %v",counter, session.Wid, err)

					if counter == attempt {

						log.Errorf("Session Connect failed  for wid %v",waSession.ID)
						go wh.NotifyLogout(waSession.Wid)
                                                storage.DeleteSession(waSession.ID)
                                                log.Infof("Delete session for number %v", waSession.Wid)
                                                return nil
					}
					counter++
					return err
				}

				wh.Session = &newSession
				wac.AddHandler(&wh)
				storage.UpdateSession(waSession)
				Connections[waSession.ID] = &wh
				wh.reloginCounter(wh.Wac.Info.Wid , 0)
				wh.ProcessQue()
				return nil
			})
		}(waSession)
	}
}

//
func ReconnectForHandler(wh *WaHandler) error {
log.Infof("Handel session %v", *wh.SessionID)

	waSession, err := storage.GetSessionForID(*wh.SessionID)
	if err != nil {
		log.Errorf(" ReconnectForHandler storage")
		wh.logout()		
		return err
	}
        
	session := SessionFromWASession(*waSession)

	wac, err := NewConn(waSession.ProxyURL)
	if err != nil {
		log.Errorf("ReconnectForHandler Connect failed%v ", err)
		return err
	}
	wh.Wac.Disconnect()
	wac.SetClientName(longName, shortName ,clientVersion )
	newSession, err := wac.RestoreWithSession(session)
	if err != nil {
		log.Errorf("ReconnectForHandler Restore failed%v", err)
		wh.Wac.Disconnect()
		return err
	}
	wh.Wac.RemoveHandler(wh)
	wh.Wac = wac
	wh.Session = &newSession
	wac.AddHandler(wh)
	storage.UpdateSession(WASessionFromSession(*wh.SessionID, waSession.ProxyURL, wh.Hook, newSession))
	return nil
}

// 
func checkWac(wac *wa.Conn){

                if wac.Info == nil {
			wac.Disconnect()
//			log.Errorf("Wac Disconnected")
                }
}


//
func NewConn(proxyURLString *string) (*wa.Conn, error) {
	var err error
	var wac *wa.Conn
	var timeout = time.Duration(*timeoutInt) * time.Second
        log.Infof("TimeOutFlag: %v , TimeOut: %v", *timeoutInt,timeout)

	if proxyURLString != nil {
		proxyURL, err := url.Parse(*proxyURLString)
		if err != nil {
			return nil, err
		}
		proxy := http.ProxyURL(proxyURL)
		wac, err = wa.NewConnWithProxy(timeout, proxy)
	} else {
		wac, err = wa.NewConn(timeout)
	}
/*
	v, err := wa.CheckCurrentServerVersion()
	if err != nil {
		return wac,err
	}
	serverVersion := &version{major: v[0], minor: v[1], patch: v[2]}
	fmt.Printf("Server has version %d.%d.%d\n", serverVersion.major, serverVersion.minor, serverVersion.patch)

	wac.SetClientVersion(serverVersion.major, serverVersion.minor, serverVersion.patch)
*/
	if err != nil {
		log.Info("Error in connect %v", err)
		wac.Disconnect()
		return nil, err
	}

	time.AfterFunc(15*time.Second , func() { checkWac(wac) })
	return wac, nil
}

//
func Login(qr chan string, proxyURL *strfmt.URI, sessionID string) (*WaHandler, error) {

	var proxyURLString *string
	if proxyURL != nil {
		temp := proxyURL.String()
		proxyURLString = &temp
	}

	var wac, err = NewConn(proxyURLString)
	if err != nil {
		return nil, err
	}
	wac.SetClientName(longName, shortName, clientVersion)
	wh := WaHandler{
		Wac:             wac,
		lastMessageTime: uint64(time.Now().Unix()),
		SessionID:       &sessionID,
		State:           RESUME,
		Signal:          make(chan Signal),
		Queue:           make(chan func() error),
		ChBrodcastQuit:  make(map[string]chan struct{}),
		OnlineNumbers:   make(map[string]OnlineStatus),
		ReloginCounter:  make(map[string]ReloginStatus),
	}
	wac.AddHandler(&wh)
	session, err := wac.Login(qr)
	if err != nil {
		return nil, err
	}

	err = storage.CreateSession(WASessionFromSession(sessionID, proxyURLString, nil, session))
	if err != nil {
		log.Errorf("Session Storing failed %v", err)
		//return nil, err
		 storage.UpdateSessionDuplicated(WASessionFromSession(sessionID, proxyURLString, nil, session))

	}
//	go wh.NotifyScanQr()
	wh.reloginCounter(wh.Wac.Info.Wid , 0)
	wh.ProcessQue()
	return &wh, nil
}
