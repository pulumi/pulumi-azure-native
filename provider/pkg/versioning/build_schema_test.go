package versioning

import (
	"path"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

// To update snapshots run:
// UPDATE_SNAPS=true go test -v ./provider/pkg/versioning -run TestNetwork_2020_02_01

func TestNetwork_2020_02_01(t *testing.T) {
	buildSchemaArgs := BuildSchemaArgs{
		Specs: ReadSpecsArgs{
			NamespaceFilter: "Network",
			VersionsFilter:  "2020-02-01",
		},
		RootDir: path.Join("..", "..", ".."),
		Version: "2.0.0",
	}
	buildSchemaResult, err := BuildSchema(buildSchemaArgs)
	if err != nil {
		panic(err)
	}
	snaps.MatchSnapshot(t, buildSchemaResult.PackageSpec)
	snaps.MatchSnapshot(t, buildSchemaResult.Metadata)
}

func TestCache_2023_04_01(t *testing.T) {
	buildSchemaArgs := BuildSchemaArgs{
		Specs: ReadSpecsArgs{
			NamespaceFilter: "Cache",
			VersionsFilter:  "2023-04-01",
		},
		RootDir: path.Join("..", "..", ".."),
		Version: "2.0.0",
	}
	buildSchemaResult, err := BuildSchema(buildSchemaArgs)
	if err != nil {
		panic(err)
	}
	snaps.MatchSnapshot(t, buildSchemaResult.PackageSpec)
	snaps.MatchSnapshot(t, buildSchemaResult.Metadata)
}

func TestCognitiveServices_2023_05_01(t *testing.T) {
	buildSchemaArgs := BuildSchemaArgs{
		Specs: ReadSpecsArgs{
			NamespaceFilter: "CognitiveServices",
			VersionsFilter:  "2023-05-01",
		},
		RootDir: path.Join("..", "..", ".."),
		Version: "2.0.0",
	}
	buildSchemaResult, err := BuildSchema(buildSchemaArgs)
	if err != nil {
		panic(err)
	}
	snaps.MatchSnapshot(t, buildSchemaResult.PackageSpec)
	snaps.MatchSnapshot(t, buildSchemaResult.Metadata)
}
