// Copyright 2025, Pulumi Corporation.  All rights reserved.

package util

import (
	"bytes"
	"compress/gzip"
	"io"
)

func GzipBytes(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	if _, err := gzipWriter.Write(input); err != nil {
		return nil, err
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GunzipBytes(input []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
