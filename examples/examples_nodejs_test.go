// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package examples

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestAccAppServiceTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "appservice"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccCosmosDBTs(t *testing.T) {
	t.Skip("Skipping due to CosmosDB failing with ServiceUnavailable due to high demand")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cosmosdb"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSimpleTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestDestroyTs(t *testing.T) {
	// Tests eventually-consistent deletion handling.
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "static-website"),
		})

	integration.ProgramTest(t, &test)
}

func TestImportTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "import"),
		})

	integration.ProgramTest(t, &test)
}

func TestPostgresTs(t *testing.T) {
	t.Skip("disabled due to a failure, see https://github.com/pulumi/pulumi-azure-native/issues/898")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "postgres"),
		})

	integration.ProgramTest(t, &test)
}

func TestMySqlTs(t *testing.T) {
	t.Skip("looks unreliable, getting errors with Code=ResourceNotFound Code=InternalServerError")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "mysql"),
		})

	integration.ProgramTest(t, &test)
}

func TestMessagingTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}

func TestSecretsTs(t *testing.T) {
	secretMessage := "secret message for testing"

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "secrets"),
			Config: map[string]string{
				"message": secretMessage,
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				assert.NotNil(t, stackInfo.Deployment)
				state, err := json.Marshal(stackInfo.Deployment)
				assert.NoError(t, err)

				assert.NotContains(t, string(state), secretMessage)
			},
		})

	integration.ProgramTest(t, &test)
}

func TestTimeSeriesTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "timeseries"),
		})

	integration.ProgramTest(t, &test)
}

func TestPublicIpUpdateTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "public-ip-update"),
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("public-ip-update", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestVnetSubnetsResolution(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "vnet-subnets-resolution"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("vnet-subnets-resolution", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/azure-native",
		},
	})

	return baseJS
}
