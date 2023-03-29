package gen

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

func BenchmarkGen(b *testing.B) {
	b.ReportAllocs()

	specs, err := openapi.ReadAzureProviders("../../../azure-rest-api-specs/", "Network")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PulumiSchema(specs)
	}
}
