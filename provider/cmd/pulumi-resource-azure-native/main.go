// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	_ "embed"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/version"

	"github.com/edsrzf/mmap-go"
)

var providerName = "azure-native"

//go:embed metadata-compact.json
var azureApiResources string

var schemaBytes []byte

func unsafeStringToBytes(data string) []byte {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&data))

	var bytes []byte
	bytes = unsafe.Slice((*byte)(unsafe.Pointer(hdr.Data)), hdr.Len)

	return bytes
}

func loadSchemaFile() ([]byte, error) {
	// Memoize in case we're called multiple times, as we are going to leak this memory that we mmap.
	if schemaBytes != nil {
		return schemaBytes, nil
	}

	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(filepath.Dir(exe), "schema-azure-native.json")

	schemaFile, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	schemaBytes, err = mmap.Map(schemaFile, mmap.RDONLY, 0)
	if err != nil {
		schemaFile.Close()
		return nil, err
	}

	return schemaBytes, nil
}

func schemaLoader() ([]byte, string, error) {
	schemaBytes, err := loadSchemaFile()
	if err != nil {
		return nil, "", err
	}

	schemaString := *(*string)(unsafe.Pointer(&schemaBytes))

	return schemaBytes, schemaString, nil
}

func main() {
	provider.Serve(providerName, version.Version, schemaLoader, unsafeStringToBytes(azureApiResources))
}
