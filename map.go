// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"net/http"
	"reflect"
	"sync"
)

var (
	// Precompute the reflect.Type of error and http.Request
	typeOfError   = reflect.TypeOf((*error)(nil)).Elem()
	typeOfRequest = reflect.TypeOf((*http.Request)(nil)).Elem()
)

// ----------------------------------------------------------------------------
// service
// ----------------------------------------------------------------------------

type service struct {
	name     string                    // name of service
	rcvr     reflect.Value             // receiver of methods for the service
	rcvrType reflect.Type              // type of the receiver
	methods  map[string]*serviceMethod // registered methods
	passReq  bool
}

type serviceMethod struct {
	method    reflect.Method // receiver method
	argsType  reflect.Type   // type of the request argument
	replyType reflect.Type   // type of the response argument
}

// ----------------------------------------------------------------------------
// serviceMap
// ----------------------------------------------------------------------------

// serviceMap is a registry for services.
type serviceMap struct {
	mutex    sync.Mutex
	services map[string]*service
}

// register adds a new service using reflection to extract its methods.
func (m *serviceMap) register(rcvr interface{}, name string, passReq bool) error {
	_ = "STUB: not implemented"
	// Setup service.
	return nil
}

// Setup methods.

// offset the parameter indexes by one if the
// service methods accept an HTTP request pointer

// Method must be exported.

// Method needs four ins: receiver, *http.Request, *args, *reply.

// If the service methods accept an HTTP request pointer

// First argument must be a pointer and must be http.Request.

// Next argument must be a pointer and must be exported.

// Next argument must be a pointer and must be exported.

// Method needs one out: error.

// Add to the map.

// get returns a registered service given a method name.
//
// The method name uses a dotted notation as in "Service.Method".
func (m *serviceMap) get(method string) (*service, *serviceMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// isExported returns true of a string is an exported (upper case) name.
func isExported(name string) bool { _ = "STUB: not implemented"; return false }

// isExportedOrBuiltin returns true if a type is exported or a builtin.
func isExportedOrBuiltin(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// PkgPath will be non-empty even for an exported type,
// so we need to check the type name as well.
