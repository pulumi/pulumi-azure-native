// Copyright 2021, Pulumi Corporation.  All rights reserved.

package main

import (
	_ "embed"
	"reflect"
	"unsafe"

	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/encoding/json"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

//go:embed metadata-compact.json
var azureApiResources string

//go:embed schema-full.json
var pulumiSchema string

func main() {
	var readFrom io.Reader
	var langs string

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [arm-template path] lang\n", os.Args[0])
		fmt.Print("arm-template   Path to arm-template or assumed to be stdin if omitted\n")
		fmt.Print("lang           Comma separated list of languages - e.g. nodejs,python,dotnet,go\n")
		fmt.Println()
		return
	}
	if len(os.Args) == 2 {
		readFrom = os.Stdin
		langs = os.Args[1]
	} else {
		var err error
		filePath := os.Args[1]
		readFrom, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to read from file: %s: %+v", filePath, err)
		}
		langs = os.Args[2]
	}

	pkgSpec, err := loadSchema()
	if err != nil {
		log.Fatalf("Failed to load schema: %+v", err)
	}

	metadata, err := loadMetadata()
	if err != nil {
		log.Fatalf("Failed to load metadata: %+v", err)
	}

	renderer := arm2pulumi.NewRenderer(pkgSpec, metadata)
	log.Print("Starting render\n")
	dir := time.Now()
	body, diagnostics, err := renderer.RenderIRFromReader(readFrom)
	if err != nil {
		log.Fatalf("Failure rendering IR from template: %+v", err)
	}
	dirComplete := time.Since(dir)

	fmt.Printf("%v\n", body)

	languages := strings.Split(langs, ",")
	dpro := time.Now()
	programsMap, _, err := renderer.RenderPrograms(body, languages)
	if err != nil {
		log.Printf("Failure rendering programs: %+v", err)
	}
	dproComplete := time.Since(dpro)

	log.Printf("IR generation: %v ms, program gen %v ms", dirComplete.Milliseconds(), dproComplete.Milliseconds())

	for k, v := range programsMap {
		fmt.Printf("Language: %s\n", k)
		fmt.Println()
		fmt.Printf("%s\n", v)
		fmt.Println()
		fmt.Println()
	}
	for k, diags := range diagnostics {
		fmt.Printf("Diagnostics for %s\n", k)
		for _, diag := range diags {
			fmt.Printf("WARN: [%s] at '%s' - '%s'\n", diag.Severity, diag.SourceToken, diag.Description)
		}
		fmt.Println()
	}
}

// loadMetadata loads the serialized/compressed metadata generated during
// schema generation from metadata.go
func loadMetadata() (*resources.AzureAPIMetadata, error) {
	metadata, err := provider.LoadMetadata([]byte(azureApiResources))
	if err != nil {
		return nil, fmt.Errorf("loading metadata: %w", err)
	}
	return metadata, nil
}

// loadSchema loads the serialized/compressed schema generated during
// generation from schema.go
func loadSchema() (*schema.PackageSpec, error) {
	var pkgSpec schema.PackageSpec
	if err := decodeString(pulumiSchema, &pkgSpec); err != nil {
		return nil, fmt.Errorf("deserializing schema: %w", err)
	}
	// embed version because go codegen is particularly sensitive to this.
	if pkgSpec.Version == "" {
		pkgSpec.Version = version.Version
	}
	return &pkgSpec, nil
}

func decodeString(data string, v interface{}) error {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&data))

	var rawBytes []byte
	rawBytes = unsafe.Slice((*byte)(unsafe.Pointer(hdr.Data)), hdr.Len)

	_, err := json.Parse(rawBytes, &v, json.ZeroCopy)
	return err
}
