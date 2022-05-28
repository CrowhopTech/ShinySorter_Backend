// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
)

//go:generate swagger generate server --target ../../server --name ShinySorter --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.ShinySorterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ShinySorterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.SetImageContentMaxParseMemory = 32 << 20

	if api.CheckHealthHandler == nil {
		api.CheckHealthHandler = operations.CheckHealthHandlerFunc(func(params operations.CheckHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckHealth has not yet been implemented")
		})
	}
	if api.CreateImageHandler == nil {
		api.CreateImageHandler = operations.CreateImageHandlerFunc(func(params operations.CreateImageParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateImage has not yet been implemented")
		})
	}
	if api.CreateQuestionHandler == nil {
		api.CreateQuestionHandler = operations.CreateQuestionHandlerFunc(func(params operations.CreateQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateQuestion has not yet been implemented")
		})
	}
	if api.CreateTagHandler == nil {
		api.CreateTagHandler = operations.CreateTagHandlerFunc(func(params operations.CreateTagParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateTag has not yet been implemented")
		})
	}
	if api.DeleteQuestionHandler == nil {
		api.DeleteQuestionHandler = operations.DeleteQuestionHandlerFunc(func(params operations.DeleteQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteQuestion has not yet been implemented")
		})
	}
	if api.DeleteTagHandler == nil {
		api.DeleteTagHandler = operations.DeleteTagHandlerFunc(func(params operations.DeleteTagParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteTag has not yet been implemented")
		})
	}
	if api.GetImageByIDHandler == nil {
		api.GetImageByIDHandler = operations.GetImageByIDHandlerFunc(func(params operations.GetImageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetImageByID has not yet been implemented")
		})
	}
	if api.GetImageContentHandler == nil {
		api.GetImageContentHandler = operations.GetImageContentHandlerFunc(func(params operations.GetImageContentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetImageContent has not yet been implemented")
		})
	}
	if api.ListImagesHandler == nil {
		api.ListImagesHandler = operations.ListImagesHandlerFunc(func(params operations.ListImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListImages has not yet been implemented")
		})
	}
	if api.ListQuestionsHandler == nil {
		api.ListQuestionsHandler = operations.ListQuestionsHandlerFunc(func(params operations.ListQuestionsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListQuestions has not yet been implemented")
		})
	}
	if api.ListTagsHandler == nil {
		api.ListTagsHandler = operations.ListTagsHandlerFunc(func(params operations.ListTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListTags has not yet been implemented")
		})
	}
	if api.PatchImageByIDHandler == nil {
		api.PatchImageByIDHandler = operations.PatchImageByIDHandlerFunc(func(params operations.PatchImageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PatchImageByID has not yet been implemented")
		})
	}
	if api.PatchQuestionByIDHandler == nil {
		api.PatchQuestionByIDHandler = operations.PatchQuestionByIDHandlerFunc(func(params operations.PatchQuestionByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PatchQuestionByID has not yet been implemented")
		})
	}
	if api.PatchTagByIDHandler == nil {
		api.PatchTagByIDHandler = operations.PatchTagByIDHandlerFunc(func(params operations.PatchTagByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PatchTagByID has not yet been implemented")
		})
	}
	if api.SetImageContentHandler == nil {
		api.SetImageContentHandler = operations.SetImageContentHandlerFunc(func(params operations.SetImageContentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.SetImageContent has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
