package test

import (
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/arm2pulumi"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var testdataPath = filepath.Join("../testdata", "*", "*.json")

func TestTemplateCoverage(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	t.Log(wd)
	matches, err := filepath.Glob(testdataPath)
	require.NoError(t, err)
	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			body, _, err := arm2pulumi.RenderFileIR(match)
			require.NoError(t, err)
			fmt.Println(body)

			rendered, _, err := arm2pulumi.RenderPrograms(body, []string{"nodejs"})
			require.NoError(t, err)
			t.Log(rendered["nodejs"])
			time.Sleep(5 * time.Second)
		})

	}
}
