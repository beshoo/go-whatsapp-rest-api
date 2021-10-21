package whatsapp

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

//
func (wh *WaHandler) ProcessQue() {
	workerCount := 4
	chQueue := make(chan func() error)
	var wgErr599 sync.WaitGroup

	// Signal management
	go func() {
		for {
			select {
			case signal := <-wh.Signal:
				log.Infof("Process Que Recieved Signal %v", signal)


				if wh.State != signal {
					wh.State = signal
				}


                                if wh.State == STOP {
                                        return
                                }



			}
		}
	}()

	go func() {
		sleepTime := time.Millisecond * time.Duration(1000*60/(*rate))
		log.Infof("Sleep time is %v for rate %d", sleepTime, *rate)

		for work := range wh.Queue {
		CHECKSTATE:
		 log.Infof("Is %v", wh.State)
			switch wh.State {
			case STOP:
				close(chQueue)
				log.Info("Stop sending go routine")
				for bID, quitChannel := range wh.ChBrodcastQuit {
					log.Info("Stop bID,%v", bID)
					quitChannel <- struct{}{}
				}
				return
			case PAUSE:
				time.Sleep(time.Millisecond * 5000)
				log.Info("Try RESUME")
				wh.SendSignal(RESUME)
				goto CHECKSTATE

			case RECONNECT:
				wgErr599.Wait()
				log.Info("Try Disconnect")
				wh.Wac.Disconnect()
				log.Info(" Disconnect")
			TRYRECONNECT:
				time.Sleep(time.Millisecond * 5000)
				log.Info("Try Reconect")
				err := ReconnectForHandler(wh)
				if err != nil {
					log.Infof("Reconnect failed %v", err)
					if wh.State == RECONNECT {
						// try only if state is reconnect
						goto TRYRECONNECT
					}
				}
				log.Info("Restart Sending")
				// not through channel
				wh.State = RESUME
				chQueue <- work
			case RESUME:
				time.Sleep(sleepTime)
				chQueue <- work
			}
		}
	}()

	// Workers Sending mssagess
	for i := 0; i < workerCount; i++ {
		// this has to annoymous to access local state
		go func(id int) {
			for work := range chQueue {
				Retry(100, time.Millisecond*100, func() error {
					wgErr599.Add(1)
					err := work()
					wgErr599.Done()
					if err != nil {
						errorMessage := err.Error()
						if errorMessage == "message sending responded with %!d(float64=599)" {
							log.Errorf("Send in 599 %d", id)
							wh.SendSignal(RECONNECT)
							go func() {
								// trying till state changes/ stopped
								sleepTime := time.Millisecond * time.Duration(1000*60/(*rate))
								for {
									time.Sleep(sleepTime)
									select {
									case wh.Queue <- work:
										return
									default:
										switch wh.State {
										case STOP:
											return
										}
									}
								}
							}()
							return nil
						}
						if errorMessage != "sending message timed out" {
							return err
						}
					}
					return nil
				})

			}
		}(i)
	}

}
