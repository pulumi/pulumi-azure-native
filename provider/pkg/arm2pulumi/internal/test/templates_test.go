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
			body, diags, err := renderer.RenderFileIR(match)
			require.NoError(t, err)
			for k, v := range diags {
				t.Logf("Diagnostics for %s", k)
				for _, diag := range v {
					t.Logf("[%s] '%s' @%s - %s", diag.Severity, diag.SourceToken, diag.SourceElement, diag.Description)
				}
			}
			fmt.Printf("%s\n%s\n", match, body)
			var langs []string
			var expected []string
			extensions := map[string]string{
				"nodejs": ".ts",
				"dotnet": ".cs",
				"python": ".py",
				"go":     ".go",
			}
			sorted := []string{"nodejs", "dotnet", "python", "go"}
			for _, lang := range sorted {
				extension := extensions[lang]
				_, err := os.Stat(fmt.Sprintf("%s%s", match, extension))
				if err != nil {
					if !os.IsNotExist(err) {
						t.Fatalf("Unpexpected error: %+v", err)
					}
					continue
				}
				f, err := os.Open(fmt.Sprintf("%s%s", match, extension))
				require.NoError(t, err)
				defer func() { _ = f.Close() }()
				exp, err := ioutil.ReadAll(f)
				require.NoError(t, err)
				expected = append(expected, string(exp))
				langs = append(langs, lang)
			}
			rendered, _, err := renderer.RenderPrograms(body, langs)
			require.NoError(t, err)
			for i, lang := range langs {
				assert.Equal(t, rendered[lang], expected[i], match)
			}
		})

	}
}
