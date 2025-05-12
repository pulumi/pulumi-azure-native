package customresources

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const cdnCustomDomainHttpsPath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/customHttps"

// Define certificate source type constants
const (
	CertificateSourceCdn         = "Cdn"
	CertificateSourceKeyVault    = "AzureKeyVault"
	ProtocolTypeServerNameIndication = "ServerNameIndication"
	ProtocolTypeIPBased = "IPBased"
)

// Certificate source parameters types
var cdnCertificateSourceParametersProperties = map[string]schema.PropertySpec{
	"certificateType": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Certificate type for CDN managed certificates. One of: 'Shared', 'Dedicated'.",
	},
}

var keyVaultCertificateSourceParametersProperties = map[string]schema.PropertySpec{
	"subscriptionId": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Azure subscription ID containing the Key Vault.",
	},
	"resourceGroupName": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Resource group containing the Key Vault.",
	},
	"vaultName": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the Key Vault.",
	},
	"secretName": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the secret in Key Vault.",
	},
	"secretVersion": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Version of the secret in Key Vault.",
	},
}

// Define type specifications
var keyVaultCertificateSourceParametersTypeSpec = schema.ComplexTypeSpec{
	ObjectTypeSpec: schema.ObjectTypeSpec{
		Type:        "object",
		Description: "Key Vault certificate parameters for custom domain HTTPS configuration.",
		Properties:  keyVaultCertificateSourceParametersProperties,
		Required:    []string{"subscriptionId", "resourceGroupName", "vaultName", "secretName"},
	},
}

var cdnCertificateSourceParametersTypeSpec = schema.ComplexTypeSpec{
	ObjectTypeSpec: schema.ObjectTypeSpec{
		Type:        "object",
		Description: "CDN certificate parameters for custom domain HTTPS configuration.",
		Properties:  cdnCertificateSourceParametersProperties,
	},
}

// HTTPS configuration properties
var httpsConfigurationProperties = map[string]schema.PropertySpec{
	"certificateSource": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Certificate source for the custom domain. One of: 'Cdn', 'AzureKeyVault'.",
	},
	"protocolType": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "TLS extension protocol. One of: 'ServerNameIndication', 'IPBased'.",
	},
	"minimumTlsVersion": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "TLS protocol version that will be used for HTTPS. One of: 'None', 'TLS10', 'TLS12'.",
	},
	"cdnCertificateSourceParameters": {
		TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:cdn:CdnCertificateSourceParameters"},
		Description: "Parameters required for using CDN managed certificates.",
	},
	"keyVaultCertificateSourceParameters": {
		TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:cdn:KeyVaultCertificateSourceParameters"},
		Description: "Parameters required for using certificates stored in Azure KeyVault.",
	},
}

// Common properties required for the resource
var cdnCustomDomainHttpsProperties = map[string]schema.PropertySpec{
	"customDomainId": {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "The resource ID of the CDN Custom Domain.",
	},
	"httpsEnabled": {
		TypeSpec:    schema.TypeSpec{Type: "boolean"},
		Description: "Whether HTTPS is enabled on the custom domain.",
	},
	"httpsConfiguration": {
		TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:cdn:HttpsConfiguration"},
		Description: "HTTPS configuration when HTTPS is enabled.",
	},
}

func cdnCustomDomainHttps(azureClient azure.AzureClient) *CustomResource {
	apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("Cdn", "CustomDomain")
	if !ok {
		// default as of 2024-05-01
		apiVersion = "2024-05-01"
		logging.V(3).Infof("Warning: could not find default API version for Cdn:CustomDomain. Using %s", apiVersion)
	}

	httpsConfigType := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Type:        "object",
			Description: "HTTPS configuration for CDN custom domain.",
			Properties:  httpsConfigurationProperties,
			Required:    []string{"certificateSource", "protocolType"},
		},
	}

	return &CustomResource{
		tok:  "azure-native:cdn:CustomDomainHttps",
		path: cdnCustomDomainHttpsPath,
		Types: map[string]schema.ComplexTypeSpec{
			"azure-native:cdn:CdnCertificateSourceParameters":      cdnCertificateSourceParametersTypeSpec,
			"azure-native:cdn:KeyVaultCertificateSourceParameters": keyVaultCertificateSourceParametersTypeSpec,
			"azure-native:cdn:HttpsConfiguration":                  httpsConfigType,
		},
		LegacySchema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Enables or disables HTTPS on a CDN custom domain.",
				Type:        "object",
				Properties:  cdnCustomDomainHttpsProperties,
			},
			InputProperties: cdnCustomDomainHttpsProperties,
			RequiredInputs: []string{"customDomainId", "httpsEnabled"},
		},
		Meta: &AzureAPIResource{
			Path: cdnCustomDomainHttpsPath,
			PutParameters: []AzureAPIParameter{
				{Name: subscriptionId, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: "profileName", Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: "endpointName", Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: "customDomainName", Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{
					Name:     "properties",
					Location: "body",
					Body: &AzureAPIType{
						Properties: map[string]AzureAPIProperty{
							"httpsEnabled": {
								Type: "boolean",
							},
							"certificateSource": {
								Type: "string",
							},
							"protocolType": {
								Type: "string",
							},
							"minimumTlsVersion": {
								Type: "string",
							},
						},
						RequiredProperties: []string{"httpsEnabled"},
					},
				},
			},
		},
		Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, error) {
			// For create, ensure we have either a valid customDomainId in properties or a valid id path
			props := properties.Mappable()
			if _, hasId := props["customDomainId"].(string); !hasId && !strings.Contains(id, "/customDomains/") {
				return nil, errors.New("missing required property 'customDomainId' or invalid resource ID")
			}
			return updateCdnCustomDomainHttps(ctx, id, properties, azureClient, apiVersion)
		},
		Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, bool, error) {
			// Extract the custom domain ID from the CustomDomainHttps ID
			cdnCustomDomainId := strings.TrimSuffix(id, "/customHttps")

			// Read the custom domain to get its HTTPS settings
			customDomain, err := azureClient.Get(ctx, cdnCustomDomainId, apiVersion, nil)
			if err != nil {
				if azure.IsNotFound(err) {
					return nil, false, nil
				}
				return nil, false, err
			}

			// Extract HTTPS settings
			result := map[string]any{
				"customDomainId": cdnCustomDomainId,
				"httpsEnabled":   false,
			}

			props, ok := customDomain["properties"].(map[string]any)
			if !ok {
				return result, true, nil
			}

			provisioningState, ok := props["customHttpsProvisioningState"].(string)
			if !ok {
				return result, true, nil
			}

			// If HTTPS is enabled
			if provisioningState == "Enabled" {
				result["httpsEnabled"] = true

				// Extract HTTPS parameters if available
				httpsConfig := map[string]any{}

				if httpsParams, ok := props["customHttpsParameters"].(map[string]any); ok {
					httpsConfig["certificateSource"] = httpsParams["certificateSource"]
					httpsConfig["protocolType"] = httpsParams["protocolType"]

					if minTls, ok := httpsParams["minimumTlsVersion"]; ok {
						httpsConfig["minimumTlsVersion"] = minTls
					}

					if certParams, ok := httpsParams["certificateSourceParameters"].(map[string]any); ok {
						if typeName, ok := certParams["typeName"].(string); ok {
							if typeName == "CdnCertificateSourceParameters" {
								cdnParams := map[string]any{}
								if certType, ok := certParams["certificateType"].(string); ok {
									cdnParams["certificateType"] = certType
								}
								httpsConfig["cdnCertificateSourceParameters"] = cdnParams
							} else if typeName == "KeyVaultCertificateSourceParameters" {
								kvParams := map[string]any{}
								if val, ok := certParams["subscriptionId"].(string); ok {
									kvParams["subscriptionId"] = val
								}
								if val, ok := certParams["resourceGroupName"].(string); ok {
									kvParams["resourceGroupName"] = val
								}
								if val, ok := certParams["vaultName"].(string); ok {
									kvParams["vaultName"] = val
								}
								if val, ok := certParams["secretName"].(string); ok {
									kvParams["secretName"] = val
								}
								if val, ok := certParams["secretVersion"].(string); ok {
									kvParams["secretVersion"] = val
								}
								httpsConfig["keyVaultCertificateSourceParameters"] = kvParams
							}
						}
					}
				}

				result["httpsConfiguration"] = httpsConfig
			}

			return result, true, nil
		},
		Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
			return updateCdnCustomDomainHttps(ctx, id, news, azureClient, apiVersion)
		},
		Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
			// Extract the custom domain ID
			customDomainId := strings.TrimSuffix(id, "/customHttps")

			// For delete, we'll disable HTTPS
			path := customDomainId + "/disableCustomHttps"

			queryParams := map[string]any{
				"api-version": apiVersion,
			}

			_, err := azureClient.Post(ctx, path, map[string]any{}, queryParams)
			return err
		},
	}
}

func updateCdnCustomDomainHttps(ctx context.Context, id string, properties resource.PropertyMap, azureClient azure.AzureClient, apiVersion string) (map[string]any, error) {
	props := properties.Mappable()

	// Extract or get the custom domain ID
	customDomainId, ok := props["customDomainId"].(string)
	if !ok {
		// If not provided directly, extract from the resource ID
		customDomainId = strings.TrimSuffix(id, "/customHttps")
	}

	httpsEnabled, ok := props["httpsEnabled"].(bool)
	if !ok {
		return nil, errors.New("missing required property 'httpsEnabled'")
	}

	queryParams := map[string]any{
		"api-version": apiVersion,
	}

	// If HTTPS is being disabled
	if !httpsEnabled {
		path := customDomainId + "/disableCustomHttps"
		_, err := azureClient.Post(ctx, path, map[string]any{}, queryParams)
		if err != nil {
			return nil, err
		}
		return props, nil
	}

	// HTTPS is being enabled
	path := customDomainId + "/enableCustomHttps"

	// Get the HTTPS configuration
	httpsConfig, ok := props["httpsConfiguration"].(map[string]any)
	if !ok {
		return nil, errors.New("missing required property 'httpsConfiguration' when httpsEnabled is true")
	}

	// Certificate source is required
	certificateSource, ok := httpsConfig["certificateSource"].(string)
	if !ok {
		return nil, errors.New("missing required property 'certificateSource' in httpsConfiguration")
	}

	// Protocol type is required
	protocolType, ok := httpsConfig["protocolType"].(string)
	if !ok {
		return nil, errors.New("missing required property 'protocolType' in httpsConfiguration")
	}

	// Prepare the request body for enabling HTTPS
	body := map[string]any{
		"certificateSource": certificateSource,
		"protocolType":      protocolType,
	}

	// Optional TLS version
	if minimumTlsVersion, ok := httpsConfig["minimumTlsVersion"].(string); ok {
		body["minimumTlsVersion"] = minimumTlsVersion
	}

	// Configure certificate source parameters based on certificate source
	certificateSourceParams := map[string]any{}

	if certificateSource == CertificateSourceCdn {
		cdnParams, hasCdnParams := httpsConfig["cdnCertificateSourceParameters"].(map[string]any)
		if hasCdnParams {
			if certificateType, ok := cdnParams["certificateType"].(string); ok {
				certificateSourceParams["certificateType"] = certificateType
			} else {
				// Default to Shared if not specified
				certificateSourceParams["certificateType"] = "Shared"
			}
		} else {
			// Default to Shared if not specified
			certificateSourceParams["certificateType"] = "Shared"
		}
		certificateSourceParams["typeName"] = "CdnCertificateSourceParameters"
	} else if certificateSource == CertificateSourceKeyVault {
		// Get KeyVault parameters
		kvParams, hasKvParams := httpsConfig["keyVaultCertificateSourceParameters"].(map[string]any)
		if !hasKvParams {
			return nil, errors.New("missing required property 'keyVaultCertificateSourceParameters' for AzureKeyVault certificate source")
		}

		// Validate required parameters for Key Vault
		requiredKvParams := []string{
			"subscriptionId",
			"resourceGroupName",
			"vaultName",
			"secretName",
		}

		for _, param := range requiredKvParams {
			if _, ok := kvParams[param].(string); !ok {
				return nil, errors.New("missing required property '" + param + "' in keyVaultCertificateSourceParameters")
			}
		}

		// Copy all parameters to the request
		for k, v := range kvParams {
			certificateSourceParams[k] = v
		}

		// Default update and delete rules if not specified
		if _, hasUpdateRule := certificateSourceParams["updateRule"]; !hasUpdateRule {
			certificateSourceParams["updateRule"] = "NoAction"
		}

		if _, hasDeleteRule := certificateSourceParams["deleteRule"]; !hasDeleteRule {
			certificateSourceParams["deleteRule"] = "NoAction"
		}

		certificateSourceParams["typeName"] = "KeyVaultCertificateSourceParameters"
	}

	body["certificateSourceParameters"] = certificateSourceParams

	_, err := azureClient.Post(ctx, path, body, queryParams)
	if err != nil {
		return nil, err
	}

	// Return the updated properties
	return map[string]any{
		"customDomainId":      customDomainId,
		"httpsEnabled":        httpsEnabled,
		"httpsConfiguration":  httpsConfig,
	}, nil
}