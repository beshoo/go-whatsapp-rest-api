package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"runtime"
	wapi "bitbucket.org/rockyOO7/wa-api/api"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/number"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/profile"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/send"

	storage "bitbucket.org/rockyOO7/wa-api/storage"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	wa "bitbucket.org/rockyOO7/wa-api/whatsapp"
	"github.com/go-openapi/loads"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")
var dbType = flag.String("db", "", "Db Type")
var dbURL = flag.String("dburl", "", "Db url")

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Error(err)
		return
	}
	api := operations.NewWaAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	flag.Parse()
	server.Port = *portFlag
	router(api)
	server.SetHandler(handlerCors(swaggerUI(api.Serve(nil))))

	if dbURL != nil {
		// storage.Init(storage.MYSQL, "root:passwd@/wapidb")
		storage.Init(storage.MYSQL, *dbURL)
		defer storage.Close()
		sessions, err := storage.GetAllSessions()
		if err != nil {
			panic(fmt.Sprintf("Get sessions failed %v", err))
		}
		wa.StartWithSession(sessions)
	}

	if err := server.Serve(); err != nil {
		log.Errorf("Cannot start server %v", err)
	}
}


func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}


func swaggerUI(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		PrintMemUsage()
		log.Infof("Url %v", r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/api") {
		handler.ServeHTTP(w, r)

		} else {
			http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
		}

	})

}

func handlerCors(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
}

func router(api *operations.WaAPI) {
	api.ProfileScanQrHandler = profile.ScanQrHandlerFunc(wapi.ScanQr)
	api.ProfileProfileHandler = profile.ProfileHandlerFunc(wapi.GetProfile)
	api.ProfileSetHookHandler = profile.SetHookHandlerFunc(wapi.SetHook)
	api.ProfileGetContactsHandler = profile.GetContactsHandlerFunc(wapi.GetContacts)
	api.ProfileIsConnectedHandler = profile.IsConnectedHandlerFunc(wapi.IsConnected)
	api.ProfileConnectHandler = profile.ConnectHandlerFunc(wapi.Connect)
	api.ProfileDisconnectHandler = profile.DisconnectHandlerFunc(wapi.Disconnect)
	api.ProfileLogoutHandler = profile.LogoutHandlerFunc(wapi.Logout)
	api.SendSendVcardHandler = send.SendVcardHandlerFunc(wapi.SendVcard)
	api.SendSendLinkHandler = send.SendLinkHandlerFunc(wapi.SendLink)
	api.SendSendReadHandler = send.SendReadHandlerFunc(wapi.SendReadAck)
	api.SendSendTextHandler = send.SendTextHandlerFunc(wapi.SendText)
	api.SendSendLocationHandler = send.SendLocationHandlerFunc(wapi.SendLocation)
	api.SendSendImageHandler = send.SendImageHandlerFunc(wapi.SendImage)
	api.SendSendVideoHandler = send.SendVideoHandlerFunc(wapi.SendVideo)
	api.SendSendAudioHandler = send.SendAudioHandlerFunc(wapi.SendAudio)
	api.SendSendDocHandler = send.SendDocHandlerFunc(wapi.SendDoc)
	api.SendSendAudioRecordHandler = send.SendAudioRecordHandlerFunc(wapi.SendAudioRecord)
	api.NumberHasWhatsAppHandler = number.HasWhatsAppHandlerFunc(wapi.HasWhatApp)
	api.NumberIsOnlineHandler = number.IsOnlineHandlerFunc(wapi.IsOnline)
	api.NumberGetChatsHandler = number.GetChatsHandlerFunc(wapi.GetChats)
	api.NumberGetAvatarHandler = number.GetAvatarHandlerFunc(wapi.GetAvatar)

}
