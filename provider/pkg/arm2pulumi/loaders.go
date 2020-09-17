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

func loadMetadata() (*provider.AzureAPIMetadata, error) {
	metadata, err := provider.LoadMetadata(azureApiResources)
	if err != nil {
		return nil, fmt.Errorf("loading metadata: %w", err)
	}
	return metadata, nil
}

func loadSchema() (*schema.PackageSpec, error) {
	uncompressed, err := gzip.NewReader(bytes.NewReader(pulumiSchema))
	if err != nil {
		return nil, fmt.Errorf("loading schema: %w", err)
	}

	var pkgSpec schema.PackageSpec
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
	return &pkgSpec, nil
}
