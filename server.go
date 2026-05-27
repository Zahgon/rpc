// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"net/http"
)

// ----------------------------------------------------------------------------
// Codec
// ----------------------------------------------------------------------------

// Codec creates a CodecRequest to process each request.
type Codec interface {
	NewRequest(*http.Request) CodecRequest
}

// CodecRequest decodes a request and encodes a response using a specific
// serialization scheme.
type CodecRequest interface {
	// Reads request and returns the RPC method name.
	Method() (string, error)
	// Reads request filling the RPC method args.
	ReadRequest(interface{}) error
	// Writes response using the RPC method reply. The error parameter is
	// the error returned by the method call, if any.
	WriteResponse(http.ResponseWriter, interface{}, error) error
}

// ----------------------------------------------------------------------------
// Server
// ----------------------------------------------------------------------------

// NewServer returns a new RPC server.
func NewServer() *Server { _ = "STUB: not implemented"; return nil }

// RequestInfo contains all the information we pass to before/after functions
type RequestInfo struct {
	Method     string
	Error      error
	Request    *http.Request
	StatusCode int
}

// Server serves registered RPC services using registered codecs.
type Server struct {
	codecs        map[string]Codec
	services      *serviceMap
	interceptFunc func(i *RequestInfo) *http.Request
	beforeFunc    func(i *RequestInfo)
	afterFunc     func(i *RequestInfo)
}

// RegisterCodec adds a new codec to the server.
//
// Codecs are defined to process a given serialization scheme, e.g., JSON or
// XML. A codec is chosen based on the "Content-Type" header from the request,
// excluding the charset definition.
func (s *Server) RegisterCodec(codec Codec, contentType string) { _ = "STUB: not implemented"; return }

// RegisterService adds a new service to the server.
//
// The name parameter is optional: if empty it will be inferred from
// the receiver type name.
//
// Methods from the receiver will be extracted if these rules are satisfied:
//
//   - The receiver is exported (begins with an upper case letter) or local
//     (defined in the package registering the service).
//   - The method name is exported.
//   - The method has three arguments: *http.Request, *args, *reply.
//   - All three arguments are pointers.
//   - The second and third arguments are exported or local.
//   - The method has return type error.
//
// All other methods are ignored.
func (s *Server) RegisterService(receiver interface{}, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterTCPService adds a new TCP service to the server.
// No HTTP request struct will be passed to the service methods.
//
// The name parameter is optional: if empty it will be inferred from
// the receiver type name.
//
// Methods from the receiver will be extracted if these rules are satisfied:
//
//   - The receiver is exported (begins with an upper case letter) or local
//     (defined in the package registering the service).
//   - The method name is exported.
//   - The method has two arguments: *args, *reply.
//   - Both arguments are pointers.
//   - Both arguments are exported or local.
//   - The method has return type error.
//
// All other methods are ignored.
func (s *Server) RegisterTCPService(receiver interface{}, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// HasMethod returns true if the given method is registered.
//
// The method uses a dotted notation as in "Service.Method".
func (s *Server) HasMethod(method string) bool { _ = "STUB: not implemented"; return false }

// RegisterInterceptFunc registers the specified function as the function
// that will be called before every request. The function is allowed to intercept
// the request e.g. add values to the context.
//
// Note: Only one function can be registered, subsequent calls to this
// method will overwrite all the previous functions.
func (s *Server) RegisterInterceptFunc(f func(i *RequestInfo) *http.Request) {
	_ = "STUB: not implemented"
	return

	// RegisterBeforeFunc registers the specified function as the function
	// that will be called before every request.
	//
	// Note: Only one function can be registered, subsequent calls to this
	// method will overwrite all the previous functions.
}

func (s *Server) RegisterBeforeFunc(f func(i *RequestInfo)) {
	_ = "STUB: not implemented"

	// RegisterAfterFunc registers the specified function as the function
	// that will be called after every request
	//
	// Note: Only one function can be registered, subsequent calls to this
	// method will overwrite all the previous functions.
	return
}

func (s *Server) RegisterAfterFunc(f func(i *RequestInfo)) {
	_ = "STUB: not implemented"

	// ServeHTTP
	return
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// If Content-Type is not set and only one codec has been registered,
// then default to that codec.

// Create a new codec request.

// Get service method to be called.

// Decode the args.

// Call the registered Intercept Function

// Call the registered Before Function

// Call the service method.

// omit the HTTP request if the service method doesn't accept it

// Cast the result to error if needed.

// Prevents Internet Explorer from MIME-sniffing a response away
// from the declared content-type

// Encode the response.

// Call the registered After Function

func (s *Server) writeError(w http.ResponseWriter, status int, msg string) {
	_ = "STUB: not implemented"
	return
}
