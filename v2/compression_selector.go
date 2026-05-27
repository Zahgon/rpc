// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
)

// gzipWriter writes and closes the gzip writer.
type gzipWriter struct {
	w *gzip.Writer
}

func (gw *gzipWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// gzipEncoder implements the gzip compressed http encoder.
type gzipEncoder struct {
}

func (enc *gzipEncoder) Encode(w http.ResponseWriter) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// flateWriter writes and closes the flate writer.
type flateWriter struct {
	w *flate.Writer
}

func (fw *flateWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// flateEncoder implements the flate compressed http encoder.
type flateEncoder struct {
}

func (enc *flateEncoder) Encode(w http.ResponseWriter) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// CompressionSelector generates the compressed http encoder.
type CompressionSelector struct {
}

// Select method selects the correct compression encoder based on http HEADER.
func (*CompressionSelector) Select(r *http.Request) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}
