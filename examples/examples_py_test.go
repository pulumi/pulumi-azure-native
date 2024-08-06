// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build python || all
// +build python all

package examples

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSimplePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccLoadBalancer(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-loadbalancer"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccLoadBalancerSubResource(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-loadbalancer-subresource"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccNSG(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-nsg"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccUserAssignedIdentityPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-user-assigned-identity"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccWebAppWithSiteConfigPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-webapp"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				require.NotNil(t, stackInfo.Deployment)
				require.NotNil(t, stackInfo.Deployment.Resources)

				found := map[string]struct{}{}
				for _, resource := range stackInfo.Deployment.Resources {
					typeStr := string(resource.Type)
					if typeStr != "azure-native:web:WebApp" && typeStr != "azure-native:web:WebAppSlot" {
						continue
					}
					webApp := &resource
					kind := strings.TrimPrefix(typeStr, "azure-native:web:")
					found[kind] = struct{}{}

					require.Contains(t, webApp.Outputs, "siteConfig", kind)
					siteConfig, ok := webApp.Outputs["siteConfig"].(map[string]any)
					require.True(t, ok, kind)

					require.Contains(t, siteConfig, "ipSecurityRestrictions", kind)
					restrictions, ok := siteConfig["ipSecurityRestrictions"].([]any)
					require.True(t, ok, kind)
					require.Len(t, restrictions, 2, kind) // one from our program, one default "deny all"
					restriction, ok := restrictions[0].(map[string]any)
					require.True(t, ok, kind)
					require.Contains(t, restriction, "ipAddress", kind)
					assert.Equal(t, "198.51.100.0/22", restriction["ipAddress"], kind)
					assert.Equal(t, "pulumi", restriction["name"], kind)

					require.Contains(t, siteConfig, "defaultDocuments", kind)
					defaultDocs, ok := siteConfig["defaultDocuments"].([]any)
					require.True(t, ok, kind)
					assert.Contains(t, defaultDocs, "pulumi.html", kind)
				}

				assert.Contains(t, found, "WebApp")
				assert.Contains(t, found, "WebAppSlot")
			},
		})

	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	sdkPath := filepath.Join("..", "sdk", "python", "bin")
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		t.Fatalf("python SDK not found at %s, run `make build_python` first", sdkPath)
	}
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{sdkPath},
	})

	return basePy
}
