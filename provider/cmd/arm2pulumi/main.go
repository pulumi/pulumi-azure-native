package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/codegen/schema"
)

func main() {
	var readFrom io.Reader
	var langs string

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [arm-template path] lang\n", os.Args[0])
		fmt.Printf("arm-template   Path to arm-template or assumed to be stdin if omitted\n")
		fmt.Printf("lang           Comma separated list of languages - e.g. nodejs,python,dotnet,go")
	}
	if len(os.Args) == 1 {
		readFrom = os.Stdin
		langs = os.Args[1]
	} else {
		var err error
		filePath := os.Args[1]
		readFrom, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Failed to read from file: %s: %+v", filePath, err)
		}
		langs = os.Args[2]
	}

	metadata, err := provider.LoadMetadata(azureApiResources)
	if err != nil {
		log.Fatalf("Failed to load metadata: %+v", err)
	}

	uncompressed, err := gzip.NewReader(bytes.NewReader(pulumiSchema))
	if err != nil {
		log.Fatalf("Failed to decompress schema: %+v", err)
	}

	var pkgSpec schema.PackageSpec
	if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
		log.Fatalf("Failure unmarshalling resource map: %+v", err)
	}
	if err = uncompressed.Close(); err != nil {
		log.Fatalf("Failure Closing uncompress stream for resource map: %+v", err)
	}

	_, err = gen.Render(readFrom, metadata)
	if err != nil {
		log.Fatalf("Failure rendering PCL from template: %+v", err)
	}
	fmt.Printf("langs: %s", langs)
	return

}
