// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/files"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/questions"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/tags"
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
	// files.SetFileContentMaxParseMemory = 32 << 20

	if api.CheckHealthHandler == nil {
		api.CheckHealthHandler = operations.CheckHealthHandlerFunc(func(params operations.CheckHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckHealth has not yet been implemented")
		})
	}
	if api.FilesCreateFileHandler == nil {
		api.FilesCreateFileHandler = files.CreateFileHandlerFunc(func(params files.CreateFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.CreateFile has not yet been implemented")
		})
	}
	if api.QuestionsCreateQuestionHandler == nil {
		api.QuestionsCreateQuestionHandler = questions.CreateQuestionHandlerFunc(func(params questions.CreateQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation questions.CreateQuestion has not yet been implemented")
		})
	}
	if api.TagsCreateTagHandler == nil {
		api.TagsCreateTagHandler = tags.CreateTagHandlerFunc(func(params tags.CreateTagParams) middleware.Responder {
			return middleware.NotImplemented("operation tags.CreateTag has not yet been implemented")
		})
	}
	if api.QuestionsDeleteQuestionHandler == nil {
		api.QuestionsDeleteQuestionHandler = questions.DeleteQuestionHandlerFunc(func(params questions.DeleteQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation questions.DeleteQuestion has not yet been implemented")
		})
	}
	if api.TagsDeleteTagHandler == nil {
		api.TagsDeleteTagHandler = tags.DeleteTagHandlerFunc(func(params tags.DeleteTagParams) middleware.Responder {
			return middleware.NotImplemented("operation tags.DeleteTag has not yet been implemented")
		})
	}
	if api.FilesGetFileByIDHandler == nil {
		api.FilesGetFileByIDHandler = files.GetFileByIDHandlerFunc(func(params files.GetFileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation files.GetFileByID has not yet been implemented")
		})
	}
	if api.FilesGetFileContentHandler == nil {
		api.FilesGetFileContentHandler = files.GetFileContentHandlerFunc(func(params files.GetFileContentParams) middleware.Responder {
			return middleware.NotImplemented("operation files.GetFileContent has not yet been implemented")
		})
	}
	if api.FilesListFilesHandler == nil {
		api.FilesListFilesHandler = files.ListFilesHandlerFunc(func(params files.ListFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation files.ListFiles has not yet been implemented")
		})
	}
	if api.QuestionsListQuestionsHandler == nil {
		api.QuestionsListQuestionsHandler = questions.ListQuestionsHandlerFunc(func(params questions.ListQuestionsParams) middleware.Responder {
			return middleware.NotImplemented("operation questions.ListQuestions has not yet been implemented")
		})
	}
	if api.TagsListTagsHandler == nil {
		api.TagsListTagsHandler = tags.ListTagsHandlerFunc(func(params tags.ListTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation tags.ListTags has not yet been implemented")
		})
	}
	if api.FilesPatchFileByIDHandler == nil {
		api.FilesPatchFileByIDHandler = files.PatchFileByIDHandlerFunc(func(params files.PatchFileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation files.PatchFileByID has not yet been implemented")
		})
	}
	if api.QuestionsPatchQuestionByIDHandler == nil {
		api.QuestionsPatchQuestionByIDHandler = questions.PatchQuestionByIDHandlerFunc(func(params questions.PatchQuestionByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation questions.PatchQuestionByID has not yet been implemented")
		})
	}
	if api.TagsPatchTagByIDHandler == nil {
		api.TagsPatchTagByIDHandler = tags.PatchTagByIDHandlerFunc(func(params tags.PatchTagByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation tags.PatchTagByID has not yet been implemented")
		})
	}
	if api.QuestionsReorderQuestionsHandler == nil {
		api.QuestionsReorderQuestionsHandler = questions.ReorderQuestionsHandlerFunc(func(params questions.ReorderQuestionsParams) middleware.Responder {
			return middleware.NotImplemented("operation questions.ReorderQuestions has not yet been implemented")
		})
	}
	if api.FilesSetFileContentHandler == nil {
		api.FilesSetFileContentHandler = files.SetFileContentHandlerFunc(func(params files.SetFileContentParams) middleware.Responder {
			return middleware.NotImplemented("operation files.SetFileContent has not yet been implemented")
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
