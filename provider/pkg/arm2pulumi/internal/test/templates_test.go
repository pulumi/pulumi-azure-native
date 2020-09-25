package test

import (
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testdataPath = filepath.Join("../testdata", "*/", "*.json")

func loadMetadata(t *testing.T) *provider.AzureAPIMetadata {
	bytes, err := ioutil.ReadFile("../../../../cmd/pulumi-resource-azure-nextgen/metadata.json")
	require.NoError(t, err)
	var metadata provider.AzureAPIMetadata
	require.NoError(t, json.Unmarshal(bytes, &metadata))
	return &metadata
}

func loadSchema(t *testing.T) *schema.PackageSpec {
	bytes, err := ioutil.ReadFile("../../../../cmd/pulumi-resource-azure-nextgen/schema-full.json")
	require.NoError(t, err)
	var pkgSpec schema.PackageSpec
	require.NoError(t, json.Unmarshal(bytes, &pkgSpec))
	return &pkgSpec
}

func TestTemplateCoverage(t *testing.T) {
	matches, err := filepath.Glob(testdataPath)
	require.NoError(t, err)

	renderer := arm2pulumi.NewRenderer(loadSchema(t), loadMetadata(t))
	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			body, _, err := renderer.RenderFileIR(match)
			require.NoError(t, err)
			rendered, _, err := renderer.RenderPrograms(body, []string{"nodejs"})
			require.NoError(t, err)
			f, err := os.Open(fmt.Sprintf("%s.ts", match))
			require.NoError(t, err)
			defer f.Close()
			expected, err := ioutil.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, rendered["nodejs"], string(expected), match)
		})

	}
}
