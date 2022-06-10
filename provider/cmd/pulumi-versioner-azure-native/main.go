package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"os"
	"path"
)

func main() {
	var err error
	calculate := flag.NewFlagSet("calculate", flag.ExitOnError)
	outputPath := calculate.String("o", "", "Output path")
	version := calculate.Int("v", 1, "Version to generate")

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case calculate.Name():
		calculate.Parse(args)
		if *outputPath == "" {
			fmt.Println("Output path must be specified")
			os.Exit(0)
		}
		err = calculateVersion(*version, *outputPath)

	default:
		fmt.Printf("Unknown command %q\n", command)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func calculateVersion(version int, outputPath string) error {
	providers, err := openapi.SpecVersions()
	if err != nil {
		return err
	}
	var providerDefaults openapi.ProviderDefaults
	switch version {
	case 1:
		providerDefaults, err = openapi.CalculateProviderDefaults(providers)
		if err != nil {
			return err
		}
	default:
		return errors.Errorf("Unknown version %v", version)
	}
	defaultVersions := formatDefaultVersions(providerDefaults)
	formatted, err := json.MarshalIndent(defaultVersions, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	return emitFile(outputPath, formatted)
}

type ProviderResourceDefaultVersions = map[string]map[string]string

func formatDefaultVersions(providerDefaults openapi.ProviderDefaults) ProviderResourceDefaultVersions {
	providerResourceDefaultVersions := make(ProviderResourceDefaultVersions)
	for providerName, resources := range providerDefaults {
		resourceVersions := make(map[string]string)
		for k, resource := range resources.Resources {
			resourceVersions[k] = resource.Swagger.Info.Version
		}
		for k, invoke := range resources.Invokes {
			resourceVersions[k] = invoke.Swagger.Info.Version
		}
		providerResourceDefaultVersions[providerName] = resourceVersions
	}
	return providerResourceDefaultVersions
}

func emitFile(outPath string, contents []byte) error {
	if err := tools.EnsureDir(path.Dir(outPath)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(outPath)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}
