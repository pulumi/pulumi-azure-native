// Copyright 2021, Pulumi Corporation.  All rights reserved.

package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testdataPath = filepath.Join("../testdata", "templates", "*/", "*.json")

var renderOptionsOverride = map[string][]arm2pulumi.RenderOption{
	"../testdata/templates/armexport/rancher.json": {arm2pulumi.DisableResourceLinking()},
}

func loadMetadata(t *testing.T) *resources.AzureAPIMetadata {
	bytes, err := ioutil.ReadFile("../../../../cmd/pulumi-resource-azure-native/metadata.json")
	require.NoError(t, err)
	var metadata resources.AzureAPIMetadata
	require.NoError(t, json.Unmarshal(bytes, &metadata))
	return &metadata
}

func loadSchema(t *testing.T) *schema.PackageSpec {
	bytes, err := ioutil.ReadFile("../../../../cmd/pulumi-resource-azure-native/schema-full.json")
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
			body, diags, err := renderer.RenderFileIR(match, renderOptionsOverride[match]...)
			if err != nil {
				t.Logf("%+v", err)
			}
			require.NoError(t, err, "%+v", err)
			for k, v := range diags {
				t.Logf("Diagnostics for %s", k)
				for _, diag := range v {
					t.Logf("[%s] '%s' @%s - %s", diag.Severity, diag.SourceToken, diag.SourceElement, diag.Description)
				}
			}
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
				assert.Equal(t, expected[i], rendered[lang], match)
			}
		})
	}
}
