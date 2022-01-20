// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package oapigen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Creates a new task
	// (POST /v1/tasks)
	CreateTask(w http.ResponseWriter, r *http.Request, params CreateTaskParams)
	// Deletes a task by name
	// (DELETE /v1/tasks/{name})
	DeleteTaskByName(w http.ResponseWriter, r *http.Request, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateTaskParams

	// ------------- Optional query parameter "run" -------------
	if paramValue := r.URL.Query().Get("run"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "run", r.URL.Query(), &params.Run)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter run: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteTaskByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteTaskByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTaskByName(w, r, name)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/tasks", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/tasks/{name}", wrapper.DeleteTaskByName)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"GVIIHQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
