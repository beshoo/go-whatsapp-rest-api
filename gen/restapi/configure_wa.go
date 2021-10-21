// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/hooks"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/number"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/profile"
	"bitbucket.org/rockyOO7/wa-api/gen/restapi/operations/send"
)

//go:generate swagger generate server --target ../../gen --name Wa --spec ../../static/spec.yml --exclude-main

func configureFlags(api *operations.WaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.HooksPostMessageAudioHandler == nil {
		api.HooksPostMessageAudioHandler = hooks.PostMessageAudioHandlerFunc(func(params hooks.PostMessageAudioParams) middleware.Responder {
			return middleware.NotImplemented("operation hooks.PostMessageAudio has not yet been implemented")
		})
	}
	if api.HooksPostMessageDocHandler == nil {
		api.HooksPostMessageDocHandler = hooks.PostMessageDocHandlerFunc(func(params hooks.PostMessageDocParams) middleware.Responder {
			return middleware.NotImplemented("operation hooks.PostMessageDoc has not yet been implemented")
		})
	}
	if api.HooksPostMessageImageHandler == nil {
		api.HooksPostMessageImageHandler = hooks.PostMessageImageHandlerFunc(func(params hooks.PostMessageImageParams) middleware.Responder {
			return middleware.NotImplemented("operation hooks.PostMessageImage has not yet been implemented")
		})
	}
	if api.HooksPostMessageTextHandler == nil {
		api.HooksPostMessageTextHandler = hooks.PostMessageTextHandlerFunc(func(params hooks.PostMessageTextParams) middleware.Responder {
			return middleware.NotImplemented("operation hooks.PostMessageText has not yet been implemented")
		})
	}
	if api.HooksPostMessageVideoHandler == nil {
		api.HooksPostMessageVideoHandler = hooks.PostMessageVideoHandlerFunc(func(params hooks.PostMessageVideoParams) middleware.Responder {
			return middleware.NotImplemented("operation hooks.PostMessageVideo has not yet been implemented")
		})
	}
	if api.ProfileConnectHandler == nil {
		api.ProfileConnectHandler = profile.ConnectHandlerFunc(func(params profile.ConnectParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.Connect has not yet been implemented")
		})
	}
	if api.ProfileDisconnectHandler == nil {
		api.ProfileDisconnectHandler = profile.DisconnectHandlerFunc(func(params profile.DisconnectParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.Disconnect has not yet been implemented")
		})
	}
	if api.NumberGetChatsHandler == nil {
		api.NumberGetChatsHandler = number.GetChatsHandlerFunc(func(params number.GetChatsParams) middleware.Responder {
			return middleware.NotImplemented("operation number.GetChats has not yet been implemented")
		})
	}
	if api.ProfileGetContactsHandler == nil {
		api.ProfileGetContactsHandler = profile.GetContactsHandlerFunc(func(params profile.GetContactsParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetContacts has not yet been implemented")
		})
	}
	if api.NumberHasWhatsAppHandler == nil {
		api.NumberHasWhatsAppHandler = number.HasWhatsAppHandlerFunc(func(params number.HasWhatsAppParams) middleware.Responder {
			return middleware.NotImplemented("operation number.HasWhatsApp has not yet been implemented")
		})
	}
	if api.ProfileIsConnectedHandler == nil {
		api.ProfileIsConnectedHandler = profile.IsConnectedHandlerFunc(func(params profile.IsConnectedParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.IsConnected has not yet been implemented")
		})
	}
	if api.NumberIsOnlineHandler == nil {
		api.NumberIsOnlineHandler = number.IsOnlineHandlerFunc(func(params number.IsOnlineParams) middleware.Responder {
			return middleware.NotImplemented("operation number.IsOnline has not yet been implemented")
		})
	}
	if api.ProfileLogoutHandler == nil {
		api.ProfileLogoutHandler = profile.LogoutHandlerFunc(func(params profile.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.Logout has not yet been implemented")
		})
	}
	if api.ProfileProfileHandler == nil {
		api.ProfileProfileHandler = profile.ProfileHandlerFunc(func(params profile.ProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.Profile has not yet been implemented")
		})
	}
	if api.ProfileScanQrHandler == nil {
		api.ProfileScanQrHandler = profile.ScanQrHandlerFunc(func(params profile.ScanQrParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ScanQr has not yet been implemented")
		})
	}
	if api.SendSendAudioHandler == nil {
		api.SendSendAudioHandler = send.SendAudioHandlerFunc(func(params send.SendAudioParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendAudio has not yet been implemented")
		})
	}
	if api.SendSendDocHandler == nil {
		api.SendSendDocHandler = send.SendDocHandlerFunc(func(params send.SendDocParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendDoc has not yet been implemented")
		})
	}
	if api.SendSendImageHandler == nil {
		api.SendSendImageHandler = send.SendImageHandlerFunc(func(params send.SendImageParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendImage has not yet been implemented")
		})
	}
	if api.SendSendReadHandler == nil {
		api.SendSendReadHandler = send.SendReadHandlerFunc(func(params send.SendReadParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendRead has not yet been implemented")
		})
	}
	if api.SendSendTextHandler == nil {
		api.SendSendTextHandler = send.SendTextHandlerFunc(func(params send.SendTextParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendText has not yet been implemented")
		})
	}
	if api.SendSendVideoHandler == nil {
		api.SendSendVideoHandler = send.SendVideoHandlerFunc(func(params send.SendVideoParams) middleware.Responder {
			return middleware.NotImplemented("operation send.SendVideo has not yet been implemented")
		})
	}
	if api.ProfileSetHookHandler == nil {
		api.ProfileSetHookHandler = profile.SetHookHandlerFunc(func(params profile.SetHookParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.SetHook has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
