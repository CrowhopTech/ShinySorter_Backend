// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewShinySorterAPI creates a new ShinySorter instance
func NewShinySorterAPI(spec *loads.Document) *ShinySorterAPI {
	return &ShinySorterAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),
		TxtProducer:  runtime.TextProducer(),

		CheckHealthHandler: CheckHealthHandlerFunc(func(params CheckHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation CheckHealth has not yet been implemented")
		}),
		CreateFileHandler: CreateFileHandlerFunc(func(params CreateFileParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateFile has not yet been implemented")
		}),
		CreateQuestionHandler: CreateQuestionHandlerFunc(func(params CreateQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateQuestion has not yet been implemented")
		}),
		CreateTagHandler: CreateTagHandlerFunc(func(params CreateTagParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateTag has not yet been implemented")
		}),
		DeleteQuestionHandler: DeleteQuestionHandlerFunc(func(params DeleteQuestionParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteQuestion has not yet been implemented")
		}),
		DeleteTagHandler: DeleteTagHandlerFunc(func(params DeleteTagParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteTag has not yet been implemented")
		}),
		GetFileByIDHandler: GetFileByIDHandlerFunc(func(params GetFileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetFileByID has not yet been implemented")
		}),
		GetFileContentHandler: GetFileContentHandlerFunc(func(params GetFileContentParams) middleware.Responder {
			return middleware.NotImplemented("operation GetFileContent has not yet been implemented")
		}),
		ListFilesHandler: ListFilesHandlerFunc(func(params ListFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListFiles has not yet been implemented")
		}),
		ListQuestionsHandler: ListQuestionsHandlerFunc(func(params ListQuestionsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListQuestions has not yet been implemented")
		}),
		ListTagsHandler: ListTagsHandlerFunc(func(params ListTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListTags has not yet been implemented")
		}),
		PatchFileByIDHandler: PatchFileByIDHandlerFunc(func(params PatchFileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatchFileByID has not yet been implemented")
		}),
		PatchQuestionByIDHandler: PatchQuestionByIDHandlerFunc(func(params PatchQuestionByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatchQuestionByID has not yet been implemented")
		}),
		PatchTagByIDHandler: PatchTagByIDHandlerFunc(func(params PatchTagByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatchTagByID has not yet been implemented")
		}),
		SetFileContentHandler: SetFileContentHandlerFunc(func(params SetFileContentParams) middleware.Responder {
			return middleware.NotImplemented("operation SetFileContent has not yet been implemented")
		}),
	}
}

/*ShinySorterAPI Endpoint definitions for the shiny-sorter file sorting project */
type ShinySorterAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - text/plain
	TxtProducer runtime.Producer

	// CheckHealthHandler sets the operation handler for the check health operation
	CheckHealthHandler CheckHealthHandler
	// CreateFileHandler sets the operation handler for the create file operation
	CreateFileHandler CreateFileHandler
	// CreateQuestionHandler sets the operation handler for the create question operation
	CreateQuestionHandler CreateQuestionHandler
	// CreateTagHandler sets the operation handler for the create tag operation
	CreateTagHandler CreateTagHandler
	// DeleteQuestionHandler sets the operation handler for the delete question operation
	DeleteQuestionHandler DeleteQuestionHandler
	// DeleteTagHandler sets the operation handler for the delete tag operation
	DeleteTagHandler DeleteTagHandler
	// GetFileByIDHandler sets the operation handler for the get file by Id operation
	GetFileByIDHandler GetFileByIDHandler
	// GetFileContentHandler sets the operation handler for the get file content operation
	GetFileContentHandler GetFileContentHandler
	// ListFilesHandler sets the operation handler for the list files operation
	ListFilesHandler ListFilesHandler
	// ListQuestionsHandler sets the operation handler for the list questions operation
	ListQuestionsHandler ListQuestionsHandler
	// ListTagsHandler sets the operation handler for the list tags operation
	ListTagsHandler ListTagsHandler
	// PatchFileByIDHandler sets the operation handler for the patch file by Id operation
	PatchFileByIDHandler PatchFileByIDHandler
	// PatchQuestionByIDHandler sets the operation handler for the patch question by ID operation
	PatchQuestionByIDHandler PatchQuestionByIDHandler
	// PatchTagByIDHandler sets the operation handler for the patch tag by ID operation
	PatchTagByIDHandler PatchTagByIDHandler
	// SetFileContentHandler sets the operation handler for the set file content operation
	SetFileContentHandler SetFileContentHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ShinySorterAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ShinySorterAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ShinySorterAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ShinySorterAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ShinySorterAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ShinySorterAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ShinySorterAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ShinySorterAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ShinySorterAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ShinySorterAPI
func (o *ShinySorterAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.CheckHealthHandler == nil {
		unregistered = append(unregistered, "CheckHealthHandler")
	}
	if o.CreateFileHandler == nil {
		unregistered = append(unregistered, "CreateFileHandler")
	}
	if o.CreateQuestionHandler == nil {
		unregistered = append(unregistered, "CreateQuestionHandler")
	}
	if o.CreateTagHandler == nil {
		unregistered = append(unregistered, "CreateTagHandler")
	}
	if o.DeleteQuestionHandler == nil {
		unregistered = append(unregistered, "DeleteQuestionHandler")
	}
	if o.DeleteTagHandler == nil {
		unregistered = append(unregistered, "DeleteTagHandler")
	}
	if o.GetFileByIDHandler == nil {
		unregistered = append(unregistered, "GetFileByIDHandler")
	}
	if o.GetFileContentHandler == nil {
		unregistered = append(unregistered, "GetFileContentHandler")
	}
	if o.ListFilesHandler == nil {
		unregistered = append(unregistered, "ListFilesHandler")
	}
	if o.ListQuestionsHandler == nil {
		unregistered = append(unregistered, "ListQuestionsHandler")
	}
	if o.ListTagsHandler == nil {
		unregistered = append(unregistered, "ListTagsHandler")
	}
	if o.PatchFileByIDHandler == nil {
		unregistered = append(unregistered, "PatchFileByIDHandler")
	}
	if o.PatchQuestionByIDHandler == nil {
		unregistered = append(unregistered, "PatchQuestionByIDHandler")
	}
	if o.PatchTagByIDHandler == nil {
		unregistered = append(unregistered, "PatchTagByIDHandler")
	}
	if o.SetFileContentHandler == nil {
		unregistered = append(unregistered, "SetFileContentHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ShinySorterAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ShinySorterAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *ShinySorterAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ShinySorterAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ShinySorterAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/plain":
			result["text/plain"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ShinySorterAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the shiny sorter API
func (o *ShinySorterAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ShinySorterAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthz"] = NewCheckHealth(o.context, o.CheckHealthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/files/{id}"] = NewCreateFile(o.context, o.CreateFileHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/questions"] = NewCreateQuestion(o.context, o.CreateQuestionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tags"] = NewCreateTag(o.context, o.CreateTagHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/questions/{id}"] = NewDeleteQuestion(o.context, o.DeleteQuestionHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tags/{id}"] = NewDeleteTag(o.context, o.DeleteTagHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/files/{id}"] = NewGetFileByID(o.context, o.GetFileByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/files/contents/{id}"] = NewGetFileContent(o.context, o.GetFileContentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/files"] = NewListFiles(o.context, o.ListFilesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/questions"] = NewListQuestions(o.context, o.ListQuestionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tags"] = NewListTags(o.context, o.ListTagsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/files/{id}"] = NewPatchFileByID(o.context, o.PatchFileByIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/questions/{id}"] = NewPatchQuestionByID(o.context, o.PatchQuestionByIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/tags/{id}"] = NewPatchTagByID(o.context, o.PatchTagByIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/files/contents/{id}"] = NewSetFileContent(o.context, o.SetFileContentHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ShinySorterAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ShinySorterAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ShinySorterAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ShinySorterAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ShinySorterAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
