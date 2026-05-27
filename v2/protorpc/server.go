// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protorpc

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/rpc/v2"
)

var null = json.RawMessage([]byte("null"))

// ----------------------------------------------------------------------------
// Request and Response
// ----------------------------------------------------------------------------

// serverRequest represents a ProtoRPC request received by the server.
type serverRequest struct {
	// A String containing the name of the method to be invoked.
	Method string `json:"method"`
	// An Array of objects to pass as arguments to the method.
	Params *json.RawMessage `json:"params"`
	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id *json.RawMessage `json:"id"`
}

// serverResponse represents a ProtoRPC response returned by the server.
type serverResponse struct {
	// The Object that was returned by the invoked method. This must be null
	// in case there was an error invoking the method.
	Result interface{} `json:"result"`
	// An Error object if there was an error invoking the method. It must be
	// null if there was no error.
	Error interface{} `json:"error"`
	// This must be the same id as the request it is responding to.
	Id *json.RawMessage `json:"id"`
}

// ----------------------------------------------------------------------------
// Codec
// ----------------------------------------------------------------------------

// NewCodec returns a new ProtoRPC Codec.
func NewCodec() *Codec {
	_ = "STUB: not implemented"

	// Codec creates a CodecRequest to process each request.
	return nil
}

type Codec struct {
}

// NewRequest returns a CodecRequest.
func (c *Codec) NewRequest(r *http.Request) rpc.CodecRequest {
	_ = "STUB: not implemented"
	return *new(rpc.CodecRequest)
}

// ----------------------------------------------------------------------------
// CodecRequest
// ----------------------------------------------------------------------------

// newCodecRequest returns a new CodecRequest.
func newCodecRequest(r *http.Request) rpc.CodecRequest {
	_ = "STUB: not implemented"
	// Decode the request body and check if RPC method is valid.
	return *new(rpc.CodecRequest)
}

// Copy request body for decoding and access of underlying methods

// Close original body

// Add close method to buffer and pass as request body

// CodecRequest decodes and encodes a single request.
type CodecRequest struct {
	request *serverRequest
	err     error
}

// Method returns the RPC method for the current request.
//
// The method uses a dotted notation as in "Service.Method".
func (c *CodecRequest) Method() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadRequest fills the request object for the RPC method.
func (c *CodecRequest) ReadRequest(args interface{}) error { _ = "STUB: not implemented"; return nil }

// WriteResponse encodes the response and writes it to the ResponseWriter.
func (c *CodecRequest) WriteResponse(w http.ResponseWriter, reply interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *CodecRequest) WriteError(w http.ResponseWriter, status int, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *CodecRequest) writeServerResponse(w http.ResponseWriter, status int, res *serverResponse) {
	_ = "STUB: not implemented"
	return
}
