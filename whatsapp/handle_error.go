package whatsapp

import (
        "time"
//      "fmt"
        wa "github.com/Rhymen/go-whatsapp"
        log "github.com/sirupsen/logrus"
)

//
func (wh *WaHandler) HandleError(err error) {

        if e, ok := err.(*wa.ErrConnectionClosed); ok {
                if wh.Wac == nil {
                        log.Errorf("ErrConnectionClosed Wac is nill")
                        return
                }
                if wh.Wac.Info == nil {
                        log.Errorf("ErrConnectionClosed Wac Info is nill")
                        return
                }
                log.Errorf("%v ErrConnectionClosed %v", wh.Wac.Info.Wid, e)

                json_disconnect := int(wh.ReloginCounter[wh.Wac.Info.Wid].ReloginCount)
				
                if json_disconnect == 0 {
				
				timerConnectionClosed := time.NewTicker(10 * time.Second)
				log.Errorf("Unknone Deconnecting......%v" , wh.Wac.Info.Wid)
						
                        for range timerConnectionClosed.C {

                                        log.Infof("1-Reconnecting... ")
                                        err := ReconnectForHandler(wh)
                                        if err == nil {
                                                        log.Infof("%v Reconnected ... sessionId %v", wh.Wac.Info.Wid, wh.SessionID)
                                                        go wh.NotifyConnectivity(true)
                                                        wh.Signal <- RESUME
                                                        timerConnectionClosed.Stop()
                                        } else {
                                                log.Errorf("1-Reconnecting Error %v", err)
                                 	
						if err.Error() == "admin login responded with 401"  {
                                                        // Delete if 401
                                                        log.Infof("Admin login responded with 401")

                                                        wh.logout()
							timerConnectionClosed.Stop()
                                                        go wh.NotifyLogout(wh.Wac.Info.Wid)
							return
                                                }

                                        }
                        }

                }else{
				log.Info("Deconnecting..from JSON")
				wh.reloginCounter(wh.Wac.Info.Wid , 0)
		     }

        }

        if e, ok := err.(*wa.ErrConnectionFailed); ok {

                log.Errorf("%v ErrConnectionFailed %v", *wh.SessionID, e )
                go wh.NotifyConnectivity(false)

                if wh.Wac.Info != nil {

                wh.Signal <- PAUSE

                }

                timer := time.NewTicker(30 * time.Second)
                log.Errorf("handle_error.go pass PAUSE")
 
               for range timer.C {

                        log.Infof("Reconnecting... ")
                        err := ReconnectForHandler(wh)
                        if err == nil {
                                log.Infof("%v Reconnected ... sessionId %v", wh.Wac.Info.Wid, wh.SessionID)
                                go wh.NotifyConnectivity(true)
                                wh.Signal <- RESUME
                                timer.Stop()
                        }
                        if err != nil {

						log.Errorf("handle_error: Reconnecting Error %v", err)
                                                if err.Error() == "admin login responded with 401"  {
                                                        // Delete if 401
                                                        wh.logout()
                                                        timer.Stop()
                                                        go wh.NotifyLogout(wh.Wac.Info.Wid)
                                                        return
                                                }

                        }
                }
        }
}

