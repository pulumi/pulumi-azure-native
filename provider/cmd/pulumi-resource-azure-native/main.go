// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	_ "embed"
	"reflect"
	"unsafe"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/version"
)

var providerName = "azure-native"

//go:embed metadata-compact.json
var azureApiResources string

//go:embed schema-full.json
var pulumiSchema string

func unsafeStringToBytes(data string) []byte {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&data))

	var bytes []byte
	bytes = unsafe.Slice((*byte)(unsafe.Pointer(hdr.Data)), hdr.Len)

	return bytes
}

func main() {
	provider.Serve(providerName, version.Version, unsafeStringToBytes(pulumiSchema), pulumiSchema, unsafeStringToBytes(azureApiResources))
}
