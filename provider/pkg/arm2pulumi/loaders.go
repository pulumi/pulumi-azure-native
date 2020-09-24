package arm2pulumi

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

var pkgSpec *schema.PackageSpec
var metadata *provider.AzureAPIMetadata

// loadMetadata loads the serialized/compressed metadata generated during
// schema generation from metadata.go
func loadMetadata() (*provider.AzureAPIMetadata, error) {
	if metadata == nil {
		var err error
		metadata, err = provider.LoadMetadata(azureApiResources)
		if err != nil {
			return nil, fmt.Errorf("loading metadata: %w", err)
		}
	}
	return metadata, nil
}

// loadSchema loads the serialized/compressed schema generated during
// generation from schema.go
func loadSchema() (*schema.PackageSpec, error) {
	if pkgSpec == nil {
		uncompressed, err := gzip.NewReader(bytes.NewReader(pulumiSchema))
		if err != nil {
			return nil, fmt.Errorf("loading schema: %w", err)
		}

		if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
			return nil, fmt.Errorf("deserializing schema: %w", err)
		}
		if err = uncompressed.Close(); err != nil {
			return nil, fmt.Errorf("closing uncompress stream for schema: %w", err)
		}
		// embed version because go codegen is particularly sensitive to this.
		if pkgSpec.Version == "" {
			pkgSpec.Version = version.Version
		}
	}
	return pkgSpec, nil
}
