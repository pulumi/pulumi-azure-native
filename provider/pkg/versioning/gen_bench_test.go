package versioning

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

// BenchmarkGen benchmarks the generation of the Pulumi schema.
// Run by executing `go test -benchmem -run=^$ -tags all -bench ^BenchmarkGen github.com/pulumi/pulumi-azure-native/provider/pkg/versioning -memprofile=gen.mem.pprof` in the provider folder.
// View results with `go tool pprof -http=: gen.mem.pprof`
func BenchmarkGen(b *testing.B) {
	b.ReportAllocs()

	rootDir := path.Join("..", "..", "..")

	b.ResetTimer()

	versionSources, err := ReadVersionSources(rootDir)
	if err != nil {
		b.Fatal(err)
	}

	specs, err := openapi.ReadAzureProviders(path.Join(rootDir, "azure-rest-api-specs"), "*", "")
	if err != nil {
		b.Fatal(err)
	}

	versionMetadata, err := calculateVersionMetadata(V2, versionSources, specs)
	if err != nil {
		b.Fatal(err)
	}

	specs = openapi.ApplyProvidersTransformations(specs, versionMetadata.V2Lock, nil, versionMetadata.Deprecated, map[string][]string{})

	gen.PulumiSchema(specs)
}
