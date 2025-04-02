// Copyright 2025, Pulumi Corporation.  All rights reserved.

package util

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// ZipBytes zips the given input into a single-file archive.
func ZipBytes(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	f, err := zipWriter.Create("schema.json")
	if err != nil {
		return nil, err
	}

	if _, err := f.Write(input); err != nil {
		return nil, err
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnzipBytes expects an archive with a single file and returns the unzipped contents of that file.
func UnzipBytes(input []byte) (string, error) {
	start := time.Now()
	defer func() {
		logging.V(9).Infof("unzip took %v", time.Since(start))
	}()

	r, err := zip.NewReader(bytes.NewReader(input), int64(len(input)))
	if err != nil {
		return "", err
	}

	if len(r.File) != 1 {
		return "", fmt.Errorf("expected exactly one file in the zip, got %d", len(r.File))
	}

	var result []byte
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()

		result, err = io.ReadAll(rc)
		if err != nil {
			return "", err
		}
	}
	return string(result), nil
}
