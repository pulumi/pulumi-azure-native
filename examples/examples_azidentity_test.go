// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package examples

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzidentity(t *testing.T) {

	type validationOpts struct {
		ExtraValidation func(t *testing.T, clientConfig map[string]interface{}, clientToken jwt.MapClaims)
	}

	validate := func(opts validationOpts) func(*testing.T, integration.RuntimeValidationStackInfo) {
		return func(t *testing.T, up integration.RuntimeValidationStackInfo) {
			// validate clientConfig
			clientConfig, ok := up.Outputs["clientConfig"].(map[string]interface{})
			require.True(t, ok, "expected clientConfig output")
			clientConfigJSON, _ := json.Marshal(clientConfig)
			t.Logf("clientConfig: %s", clientConfigJSON)

			assert.Contains(t, clientConfig, "clientId")
			assert.Contains(t, clientConfig, "objectId")
			assert.Contains(t, clientConfig, "subscriptionId")
			assert.Contains(t, clientConfig, "tenantId")

			// validate clientToken
			clientToken, ok := up.Outputs["clientToken"].(map[string]interface{})
			require.True(t, ok, "expected clientToken output")
			claims, err := parseJwtUnverified(clientToken["token"].(string))
			require.NoError(t, err)
			claimsJSON, _ := json.Marshal(claims)
			t.Logf("clientToken: %s", claimsJSON)

			if opts.ExtraValidation != nil {
				opts.ExtraValidation(t, clientConfig, claims)
			}
		}
	}

	baseOptions := integration.ProgramTestOptions{
		NoParallel:             true,
		SkipPreview:            true,
		SkipRefresh:            true,
		SkipExportImport:       true,
		SkipEmptyPreviewUpdate: true,
	}

	t.Run("OIDC", func(t *testing.T) {
		oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
		if oidcClientId == "" {
			t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
		}

		test := getJSBaseOptions(t).
			With(baseOptions).
			With(integration.ProgramTestOptions{
				Dir: filepath.Join(getCwd(t), "azidentity"),
				Env: []string{
					"ARM_USE_OIDC=true",
					"ARM_CLIENT_ID=" + oidcClientId,
					// not strictly necessary but making sure we test the OIDC path
					"ARM_CLIENT_SECRET=",
				},
			})

		integration.ProgramTest(t, &test)
	})

	t.Run("SP_clientsecret", func(t *testing.T) {
		skipIfShort(t)

		clientSecret := os.Getenv("ARM_CLIENT_SECRET")
		if clientSecret == "" {
			t.Skip("Skipping SP test without ARM_CLIENT_SECRET")
		}

		test := getJSBaseOptions(t).
			With(baseOptions).
			With(integration.ProgramTestOptions{
				Dir: filepath.Join(getCwd(t), "azidentity"),
				Env: []string{
					"ARM_CLIENT_ID=" + os.Getenv("ARM_CLIENT_ID"),
					"ARM_CLIENT_SECRET=" + clientSecret,
					// Make sure we test the client secret path
					"ACTIONS_ID_TOKEN_REQUEST_TOKEN=",
					"ACTIONS_ID_TOKEN_REQUEST_URL=",
					"ARM_CLIENT_CERTIFICATE_PATH=",
					"ARM_CLIENT_CERTIFICATE_PASSWORD=",
				},
				ExtraRuntimeValidation: validate(validationOpts{
					ExtraValidation: func(t *testing.T, clientConfig map[string]interface{}, token jwt.MapClaims) {
						assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
						assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), token["appid"])
						assert.Equal(t, "app", token["idtyp"])
					},
				}),
			})

		integration.ProgramTest(t, &test)
	})

	t.Run("SP_clientcert", func(t *testing.T) {
		skipIfShort(t)

		certPath := os.Getenv("ARM_CLIENT_CERTIFICATE_PATH_FOR_TEST")
		if certPath == "" {
			t.Skip("Skipping SP test without ARM_CLIENT_CERTIFICATE_PATH_FOR_TEST")
		}

		test := getJSBaseOptions(t).
			With(baseOptions).
			With(integration.ProgramTestOptions{
				Dir: filepath.Join(getCwd(t), "azidentity"),
				Env: []string{
					"ARM_CLIENT_ID=" + os.Getenv("ARM_CLIENT_ID"),
					"ARM_CLIENT_CERTIFICATE_PATH=" + certPath,
					"ARM_CLIENT_CERTIFICATE_PASSWORD=" + os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD_FOR_TEST"),
					// Make sure we test the client cert path
					"ACTIONS_ID_TOKEN_REQUEST_TOKEN=",
					"ACTIONS_ID_TOKEN_REQUEST_URL=",
					"ARM_CLIENT_SECRET=",
				},
				ExtraRuntimeValidation: validate(validationOpts{
					ExtraValidation: func(t *testing.T, clientConfig map[string]interface{}, token jwt.MapClaims) {
						assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
						assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), token["appid"])
						assert.Equal(t, "app", token["idtyp"])
					},
				}),
			})

		integration.ProgramTest(t, &test)
	})

	t.Run("CLI", func(t *testing.T) {
		skipIfShort(t)

		// AZURE_CONFIG_DIR_FOR_TEST is set by the GH workflow build-test.yml
		// to provide an isolated configuration directory for the Azure CLI.
		configDir := os.Getenv("AZURE_CONFIG_DIR_FOR_TEST")
		if configDir == "" {
			t.Skip("Skipping CLI test without AZURE_CONFIG_DIR_FOR_TEST")
		}
		t.Setenv("AZURE_CONFIG_DIR", configDir)

		test := getJSBaseOptions(t).
			With(baseOptions).
			With(integration.ProgramTestOptions{
				Dir: filepath.Join(getCwd(t), "azidentity"),
				Env: []string{
					// Unset auth variables to make sure we're testing the CLI auth path
					"ARM_CLIENT_ID=",
					"ARM_CLIENT_SECRET=",
					"ARM_CLIENT_CERTIFICATE_PATH=",
					"ARM_USE_MSI=false",
					"ARM_USE_OIDC=false",
				},
				ExtraRuntimeValidation: validate(validationOpts{
					ExtraValidation: func(t *testing.T, clientConfig map[string]interface{}, token jwt.MapClaims) {
						assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", clientConfig["clientId"])
						assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", token["appid"])
						assert.Equal(t, "user", token["idtyp"])
					},
				}),
			})

		integration.ProgramTest(t, &test)
	})
}

func parseJwtUnverified(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}
