package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/solo-io/squash/pkg/restapi/operations/debugconfig"
	"github.com/solo-io/squash/pkg/restapi/operations/debugsessions"
)

// NewSquashAPI creates a new Squash instance
func NewSquashAPI(spec *loads.Document) *SquashAPI {
	return &SquashAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		DebugconfigAddDebugConfigHandler: debugconfig.AddDebugConfigHandlerFunc(func(params debugconfig.AddDebugConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigAddDebugConfig has not yet been implemented")
		}),
		DebugconfigDeleteDebugConfigHandler: debugconfig.DeleteDebugConfigHandlerFunc(func(params debugconfig.DeleteDebugConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigDeleteDebugConfig has not yet been implemented")
		}),
		DebugconfigGetDebugConfigHandler: debugconfig.GetDebugConfigHandlerFunc(func(params debugconfig.GetDebugConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigGetDebugConfig has not yet been implemented")
		}),
		DebugconfigGetDebugConfigsHandler: debugconfig.GetDebugConfigsHandlerFunc(func(params debugconfig.GetDebugConfigsParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigGetDebugConfigs has not yet been implemented")
		}),
		DebugconfigPopContainerToDebugHandler: debugconfig.PopContainerToDebugHandlerFunc(func(params debugconfig.PopContainerToDebugParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigPopContainerToDebug has not yet been implemented")
		}),
		DebugsessionsPopDebugSessionHandler: debugsessions.PopDebugSessionHandlerFunc(func(params debugsessions.PopDebugSessionParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugsessionsPopDebugSession has not yet been implemented")
		}),
		DebugsessionsPutDebugSessionHandler: debugsessions.PutDebugSessionHandlerFunc(func(params debugsessions.PutDebugSessionParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugsessionsPutDebugSession has not yet been implemented")
		}),
		DebugconfigUpdateDebugConfigHandler: debugconfig.UpdateDebugConfigHandlerFunc(func(params debugconfig.UpdateDebugConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation DebugconfigUpdateDebugConfig has not yet been implemented")
		}),
	}
}

/*SquashAPI Squash debugget api sample. */
type SquashAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// DebugconfigAddDebugConfigHandler sets the operation handler for the add debug config operation
	DebugconfigAddDebugConfigHandler debugconfig.AddDebugConfigHandler
	// DebugconfigDeleteDebugConfigHandler sets the operation handler for the delete debug config operation
	DebugconfigDeleteDebugConfigHandler debugconfig.DeleteDebugConfigHandler
	// DebugconfigGetDebugConfigHandler sets the operation handler for the get debug config operation
	DebugconfigGetDebugConfigHandler debugconfig.GetDebugConfigHandler
	// DebugconfigGetDebugConfigsHandler sets the operation handler for the get debug configs operation
	DebugconfigGetDebugConfigsHandler debugconfig.GetDebugConfigsHandler
	// DebugconfigPopContainerToDebugHandler sets the operation handler for the pop container to debug operation
	DebugconfigPopContainerToDebugHandler debugconfig.PopContainerToDebugHandler
	// DebugsessionsPopDebugSessionHandler sets the operation handler for the pop debug session operation
	DebugsessionsPopDebugSessionHandler debugsessions.PopDebugSessionHandler
	// DebugsessionsPutDebugSessionHandler sets the operation handler for the put debug session operation
	DebugsessionsPutDebugSessionHandler debugsessions.PutDebugSessionHandler
	// DebugconfigUpdateDebugConfigHandler sets the operation handler for the update debug config operation
	DebugconfigUpdateDebugConfigHandler debugconfig.UpdateDebugConfigHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SquashAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SquashAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SquashAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SquashAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SquashAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SquashAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SquashAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SquashAPI
func (o *SquashAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DebugconfigAddDebugConfigHandler == nil {
		unregistered = append(unregistered, "debugconfig.AddDebugConfigHandler")
	}

	if o.DebugconfigDeleteDebugConfigHandler == nil {
		unregistered = append(unregistered, "debugconfig.DeleteDebugConfigHandler")
	}

	if o.DebugconfigGetDebugConfigHandler == nil {
		unregistered = append(unregistered, "debugconfig.GetDebugConfigHandler")
	}

	if o.DebugconfigGetDebugConfigsHandler == nil {
		unregistered = append(unregistered, "debugconfig.GetDebugConfigsHandler")
	}

	if o.DebugconfigPopContainerToDebugHandler == nil {
		unregistered = append(unregistered, "debugconfig.PopContainerToDebugHandler")
	}

	if o.DebugsessionsPopDebugSessionHandler == nil {
		unregistered = append(unregistered, "debugsessions.PopDebugSessionHandler")
	}

	if o.DebugsessionsPutDebugSessionHandler == nil {
		unregistered = append(unregistered, "debugsessions.PutDebugSessionHandler")
	}

	if o.DebugconfigUpdateDebugConfigHandler == nil {
		unregistered = append(unregistered, "debugconfig.UpdateDebugConfigHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SquashAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SquashAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SquashAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SquashAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SquashAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the squash API
func (o *SquashAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SquashAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/debugconfig"] = debugconfig.NewAddDebugConfig(o.context, o.DebugconfigAddDebugConfigHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/debugconfig/{debugConfigId}"] = debugconfig.NewDeleteDebugConfig(o.context, o.DebugconfigDeleteDebugConfigHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/debugconfig/{debugConfigId}"] = debugconfig.NewGetDebugConfig(o.context, o.DebugconfigGetDebugConfigHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/debugconfig"] = debugconfig.NewGetDebugConfigs(o.context, o.DebugconfigGetDebugConfigsHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/debugconfig/platform/containers/{node}/latest"] = debugconfig.NewPopContainerToDebug(o.context, o.DebugconfigPopContainerToDebugHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/debugconfig/{debugConfigId}/session"] = debugsessions.NewPopDebugSession(o.context, o.DebugsessionsPopDebugSessionHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/debugconfig/{debugConfigId}/session"] = debugsessions.NewPutDebugSession(o.context, o.DebugsessionsPutDebugSessionHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/debugconfig/{debugConfigId}"] = debugconfig.NewUpdateDebugConfig(o.context, o.DebugconfigUpdateDebugConfigHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SquashAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *SquashAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}