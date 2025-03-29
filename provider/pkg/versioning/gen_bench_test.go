package versioning

import (
	"path"
	"testing"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

// BenchmarkGen benchmarks the generation of the Pulumi schema.
// Run by executing `go test -benchmem -run=^$ -tags all -bench ^BenchmarkGen github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versioning -memprofile=gen.mem.pprof` in the provider folder.
// View results with `go tool pprof -http=: gen.mem.pprof`
func BenchmarkGen(b *testing.B) {
	b.ReportAllocs()

	rootDir := path.Join("..", "..", "..")

	b.ResetTimer()

	specs, _, err := openapi.ReadAzureModules(path.Join(rootDir, "azure-rest-api-specs"), "*", "")
	if err != nil {
		b.Fatal(err)
	}

	versionSources, err := ReadVersionSources(rootDir, specs, 2)
	if err != nil {
		b.Fatal(err)
	}

	versionMetadata, err := calculateVersionMetadata(versionSources)
	if err != nil {
		b.Fatal(err)
	}

	specs = openapi.ApplyTransformations(specs, versionMetadata.DefaultVersions, nil, versionSources.RemovedVersions, nil)

	gen.PulumiSchema(rootDir, specs, versionMetadata, semver.MustParse("2.0.0"), false /* onlyExplicitVersions */)
}
