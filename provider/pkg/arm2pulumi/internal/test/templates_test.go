package test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/arm2pulumi"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testdataPath = filepath.Join("../testdata", "*/", "*.json")

func TestTemplateCoverage(t *testing.T) {
	matches, err := filepath.Glob(testdataPath)
	require.NoError(t, err)
	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			body, _, err := arm2pulumi.RenderFileIR(match)
			require.NoError(t, err)
			rendered, _, err := arm2pulumi.RenderPrograms(body, []string{"nodejs"})
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
