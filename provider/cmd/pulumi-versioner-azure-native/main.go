package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/versioning"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"gopkg.in/yaml.v3"
)

func main() {
	f := flag.NewFlagSet("", flag.ExitOnError)
	outputDir := f.String("o", "versions", "Output directory")
	versionFile := f.String("version", "", "Version file input e.g. \"v1.json\"")
	err := f.Parse(os.Args[2:])
	args := os.Args[1:2]

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}

	err = handleCommand(args, *outputDir, *versionFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleCommand(args []string, outputDir, versionFile string) error {
	var target string
	switch len(args) {
	case 0:
		target = "all"
	case 1:
		target = args[0]
	default:
		return fmt.Errorf("only one target can be specified")
	}
	switch target {
	case "all":
		return writeAll(outputDir)
	case "spec":
		return spec(outputDir)
	case "active":
		return active(outputDir)
	case "deprecated":
		if versionFile == "" {
			return fmt.Errorf("-version is required")
		}
		return deprecated(outputDir, versionFile)
	case "pending":
		if versionFile == "" {
			return fmt.Errorf("-version is required")
		}
		return pending(outputDir, versionFile)
	case "v1":
		return v1(outputDir)
	case "v2":
		return v2(outputDir)
	default:
		re := regexp.MustCompile(`^(v\d+)-config$`)
		matches := re.FindStringSubmatch(target)
		if len(matches) == 2 {
			return vnextConfig(outputDir, matches[1])
		}
		return fmt.Errorf("unknown target: %q", target)
	}
}

func writeAll(outputDir string) error {
	providers, err := openapi.SpecVersions("*")
	if err != nil {
		return err
	}

	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return err
	}

	specVersions := versioning.FindSpecVersions(providers)
	specResourceVersions := versioning.FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1Config, err := versioning.ReadDefaultConfig(path.Join(outputDir, "v1-config.yaml"))
	if err != nil {
		return err
	}
	// No separate V1 curation yet
	v1Curation := make(versioning.Curations)
	v1, err := versioning.DefaultConfigToCuratedVersion(specVersions, v1Config, v1Curation)
	if err != nil {
		return err
	}

	deprecated := versioning.FindDeprecations(specVersions, v1)
	specAfterRemovals := versioning.RemoveDeprecations(specVersions, deprecated)
	v2Config, err := versioning.ReadDefaultConfig(path.Join(outputDir, "v2-config.yaml"))
	if err != nil {
		return err
	}
	v2Curation, err := versioning.ReadManualCurations(path.Join(outputDir, "v2-curation.yaml"))
	if err != nil {
		return err
	}
	v2, err := versioning.DefaultConfigToCuratedVersion(specAfterRemovals, v2Config, v2Curation)
	if err != nil {
		return err
	}

	return emitFiles(outputDir, map[Filename]Data{
		"spec.json":           specVersions,
		"spec-resources.json": specResourceVersions,
		"v1.json":             v1,
		"deprecated.json":     deprecated,
		"active.json":         activePathVersionsJson,
		"pending.json":        versioning.FindNewerVersions(specVersions, v1),
		"v2.json":             v2,
	})
}

func spec(outputDir string) error {
	providers, err := openapi.SpecVersions("*")
	if err != nil {
		return err
	}

	specVersions := versioning.FindSpecVersions(providers)
	specResourceVersions := versioning.FormatResourceVersions(specVersions)

	return emitFiles(outputDir, map[Filename]Data{
		"spec.json":           specVersions,
		"spec-resources.json": specResourceVersions,
	})
}

func active(outputDir string) error {
	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return err
	}

	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	return emitFiles(outputDir, map[Filename]Data{
		"active.json": activePathVersionsJson,
	})
}

func deprecated(outputDir, versionFile string) error {
	specVersions, err := versioning.ReadSpecVersions(path.Join(outputDir, "spec.json"))
	if err != nil {
		return err
	}

	v1, err := openapi.ReadCuratedVersion(path.Join(outputDir, versionFile))
	if err != nil {
		return err
	}

	deprecated := versioning.FindDeprecations(specVersions, v1)

	return emitFiles(outputDir, map[Filename]Data{
		"deprecated.json": deprecated,
	})
}

func pending(outputDir, versionFile string) error {
	specVersions, err := versioning.ReadSpecVersions(path.Join(outputDir, "spec.json"))
	if err != nil {
		return err
	}

	v1, err := openapi.ReadCuratedVersion(path.Join(outputDir, versionFile))
	if err != nil {
		return err
	}

	pending := versioning.FindNewerVersions(specVersions, v1)

	return emitFiles(outputDir, map[Filename]Data{
		"pending.json": pending,
	})
}

func v1(outputDir string) error {
	specVersions, err := versioning.ReadSpecVersions(path.Join(outputDir, "spec.json"))
	if err != nil {
		return err
	}

	v1Config, err := versioning.ReadDefaultConfig(path.Join(outputDir, "v1-config.yaml"))
	if err != nil {
		return err
	}

	// No separate V1 curation yet
	v1Curation := make(versioning.Curations)
	v1, err := versioning.DefaultConfigToCuratedVersion(specVersions, v1Config, v1Curation)
	if err != nil {
		return err
	}

	return emitFiles(outputDir, map[Filename]Data{
		"v1.json": v1,
	})
}

func v2(outputDir string) error {
	specVersions, err := versioning.ReadSpecVersions(path.Join(outputDir, "spec.json"))
	if err != nil {
		return err
	}

	deprecated, err := openapi.ReadDeprecated()
	if err != nil {
		return err
	}

	v2Config, err := versioning.ReadDefaultConfig(path.Join(outputDir, "v2-config.yaml"))
	if err != nil {
		return err
	}

	curations, err := versioning.ReadManualCurations(path.Join(outputDir, "v2-curation.yaml"))
	if err != nil {
		return err
	}

	specAfterRemovals := versioning.RemoveDeprecations(specVersions, deprecated)
	v2, err := versioning.DefaultConfigToCuratedVersion(specAfterRemovals, v2Config, curations)
	if err != nil {
		return err
	}

	return emitFiles(outputDir, map[Filename]Data{
		"v2.json": v2,
	})
}

func vnextConfig(outputDir, version string) error {
	specVersions, err := versioning.ReadSpecVersions(path.Join(outputDir, "spec.json"))
	if err != nil {
		return err
	}

	deprecated, err := openapi.ReadDeprecated()
	if err != nil {
		return err
	}

	curations, err := versioning.ReadManualCurations(path.Join(outputDir, version+"-curation.yaml"))
	if err != nil {
		return err
	}

	specAfterRemovals := versioning.RemoveDeprecations(specVersions, deprecated)

	vConfig := versioning.BuildDefaultConfig(specAfterRemovals, curations, versioning.DefaultConfig{})
	if err != nil {
		return err
	}

	violations := versioning.ValidateDefaultConfig(vConfig, curations)
	if len(violations) > 0 {
		fmt.Printf("Warning: %d curation violations found:\n", len(violations))
		for _, v := range violations {
			fmt.Printf("  %s: %s\n", v.Provider, v.Detail)
		}
	}

	filename := version + "-config.yaml"

	return emitFiles(outputDir, map[Filename]Data{
		filename: vConfig,
	})
}

type Filename = string
type Data = interface{}

// emitFiles writes serializes and writes multiple files. If if the filename ends in `.yaml` then it will be marshalled
// as YAML instead of JSON.
func emitFiles(outDir string, files map[Filename]Data) error {
	for filename, data := range files {
		outPath := path.Join(outDir, filename)
		err := emitFile(outPath, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func emitFile(outputPath string, data Data) error {
	var formatted []byte
	var err error
	if strings.HasSuffix(outputPath, ".yaml") {
		formatted, err = yaml.Marshal(data)
		if err != nil {
			return errors.Wrap(err, "marshaling YAML")
		}
	} else {
		formatted, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return errors.Wrap(err, "marshaling JSON")
		}
	}

	if err := tools.EnsureDir(path.Dir(outputPath)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(formatted)
	return err
}
