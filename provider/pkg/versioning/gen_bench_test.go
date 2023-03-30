package versioning

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

func BenchmarkGen(b *testing.B) {
	b.ReportAllocs()

	rootDir := path.Join("..", "..", "..")

	versionSources, err := ReadVersionSources(rootDir)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	// Test each provider
	for k := range versionSources.v2Spec {
		specs, err := openapi.ReadAzureProviders(path.Join(rootDir, "azure-rest-api-specs"), k)
		if err != nil {
			b.Fatal(err)
		}

		versionMetadata, err := calculateVersionMetadata(versionSources, specs)
		if err != nil {
			b.Fatal(err)
		}

		specs = openapi.ApplyProvidersTransformations(specs, versionMetadata.V2Lock, nil, versionMetadata.Deprecated, map[string][]string{})

		// for i := 0; i < b.N; i++ {
		gen.PulumiSchema(specs)
		// }
	}
}
