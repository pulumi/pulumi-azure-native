// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	_ "embed"
	"reflect"
	"unsafe"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
)

var providerName = "azure-native"

//go:embed metadata-compact.json
var azureApiResources string

//go:embed schema-full.json.zip
var pulumiSchema string

//go:embed schema-default-versions.json.zip
var defaultSchema string

func unsafeStringToBytes(data string) []byte {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&data))

	var bytes []byte
	bytes = unsafe.Slice((*byte)(unsafe.Pointer(hdr.Data)), hdr.Len)

	return bytes
}

func main() {
	provider.Serve(
		providerName,
		version.Version,
		unsafeStringToBytes(pulumiSchema),
		unsafeStringToBytes(defaultSchema),
		unsafeStringToBytes(azureApiResources),
	)
}
