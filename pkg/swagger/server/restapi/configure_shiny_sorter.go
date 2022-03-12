// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
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

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()
	api.VideoH264Producer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoH264 producer has not yet been implemented")
	})
	api.VideoH265Producer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoH265 producer has not yet been implemented")
	})
	api.VideoJPEGProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoJPEG producer has not yet been implemented")
	})
	api.VideoMP4Producer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoMP4 producer has not yet been implemented")
	})
	api.VideoMp4Producer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoMp4 producer has not yet been implemented")
	})
	api.VideoMpeg4GenericProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoMpeg4Generic producer has not yet been implemented")
	})
	api.VideoOggProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoOgg producer has not yet been implemented")
	})
	api.VideoRawProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoRaw producer has not yet been implemented")
	})
	api.VideoWebmProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoWebm producer has not yet been implemented")
	})
	api.VideoWebpProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("videoWebp producer has not yet been implemented")
	})

	if api.CheckHealthHandler == nil {
		api.CheckHealthHandler = operations.CheckHealthHandlerFunc(func(params operations.CheckHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckHealth has not yet been implemented")
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
	if api.GetImagesHandler == nil {
		api.GetImagesHandler = operations.GetImagesHandlerFunc(func(params operations.GetImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetImages has not yet been implemented")
		})
	}
	if api.PatchImageByIDHandler == nil {
		api.PatchImageByIDHandler = operations.PatchImageByIDHandlerFunc(func(params operations.PatchImageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PatchImageByID has not yet been implemented")
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
