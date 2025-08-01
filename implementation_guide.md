# Implementation Guide: Enabling SSL on CDN Custom Domains

## Problem

Currently, the Pulumi Azure Native provider lacks the ability to enable HTTPS (SSL) on CDN custom domains. The Azure REST API supports this feature through the following endpoint:

```
POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/enableCustomHttps
```

However, this functionality is not exposed in the Pulumi Azure Native SDK because it's a POST action rather than a standard resource property that can be set via PUT operations.

## Solution Overview

We need to implement a custom resource in the provider that will handle enabling and disabling HTTPS on CDN custom domains. This solution will be similar to other custom resources in the provider, such as `BlobContainerLegalHold`.

The solution will:

1. Create a new `CdnCustomDomainHttps` custom resource
2. Implement POST operations to enable and disable HTTPS
3. Add appropriate schema definitions to expose the resource in the Pulumi SDK
4. Implement necessary CRUD operations
5. Register the custom resource with the provider

## Implementation Steps

### 1. Create Custom Resource Definition

Create a new file at `provider/pkg/resources/customresources/custom_cdn_domain_https.go` that will define the custom resource:

```go
package customresources

import (
    "context"
    "strings"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
    . "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
    "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
    "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
    "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
    "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
    "github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const cdnCustomDomainHttpsPath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/customHttps"

// Common properties required for the resource
var cdnCustomDomainHttpsProperties = map[string]schema.PropertySpec{
    resourceGroupName: {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the resource group that contains the CDN profile.",
    },
    "profileName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the CDN profile.",
    },
    "endpointName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the endpoint under the profile.",
    },
    "customDomainName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the custom domain.",
    },
    "httpsEnabled": {
        TypeSpec:    schema.TypeSpec{Type: "boolean"},
        Description: "Whether HTTPS is enabled on the custom domain.",
    },
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
    "certificateType": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Certificate type (applicable for CDN-managed certificates). One of: 'Shared', 'Dedicated'.",
    },
    // KeyVault certificate parameters (if using AzureKeyVault as certificate source)
    "keyVaultSubscriptionId": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Azure subscription ID containing the Key Vault.",
    },
    "keyVaultResourceGroupName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Resource group containing the Key Vault.",
    },
    "keyVaultVaultName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the Key Vault.",
    },
    "keyVaultSecretName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the secret in Key Vault.",
    },
    "keyVaultSecretVersion": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Version of the secret in Key Vault.",
    },
}

func cdnCustomDomainHttps(azureClient azure.AzureClient) *CustomResource {
    apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("Cdn", "CustomDomain")
    if !ok {
        // default as of 2024-05-01
        apiVersion = "2024-05-01"
        logging.V(3).Infof("Warning: could not find default API version for Cdn:CustomDomain. Using %s", apiVersion)
    }

    return &CustomResource{
        tok:  "azure-native:cdn:CustomDomainHttps",
        path: cdnCustomDomainHttpsPath,
        LegacySchema: &schema.ResourceSpec{
            ObjectTypeSpec: schema.ObjectTypeSpec{
                Description: "Enables or disables HTTPS on a CDN custom domain.",
                Type:        "object",
                Properties:  cdnCustomDomainHttpsProperties,
            },
            InputProperties: cdnCustomDomainHttpsProperties,
            RequiredInputs: []string{
                resourceGroupName, "profileName", "endpointName", "customDomainName", 
                "httpsEnabled", "certificateSource", "protocolType",
            },
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
                            "certificateType": {
                                Type: "string",
                            },
                            "keyVaultSubscriptionId": {
                                Type: "string",
                            },
                            "keyVaultResourceGroupName": {
                                Type: "string",
                            },
                            "keyVaultVaultName": {
                                Type: "string",
                            },
                            "keyVaultSecretName": {
                                Type: "string",
                            },
                            "keyVaultSecretVersion": {
                                Type: "string",
                            },
                        },
                        RequiredProperties: []string{"httpsEnabled"},
                    },
                },
            },
        },
        Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, error) {
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
                "httpsEnabled": false,
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
                if httpsParams, ok := props["customHttpsParameters"].(map[string]any); ok {
                    result["certificateSource"] = httpsParams["certificateSource"]
                    result["protocolType"] = httpsParams["protocolType"]
                    
                    if minTls, ok := httpsParams["minimumTlsVersion"]; ok {
                        result["minimumTlsVersion"] = minTls
                    }
                    
                    if certParams, ok := httpsParams["certificateSourceParameters"].(map[string]any); ok {
                        if typeName, ok := certParams["typeName"].(string); ok {
                            if typeName == "CdnCertificateSourceParameters" {
                                if certType, ok := certParams["certificateType"].(string); ok {
                                    result["certificateType"] = certType
                                }
                            } else if typeName == "KeyVaultCertificateSourceParameters" {
                                if kv, ok := certParams["subscriptionId"].(string); ok {
                                    result["keyVaultSubscriptionId"] = kv
                                }
                                if kv, ok := certParams["resourceGroupName"].(string); ok {
                                    result["keyVaultResourceGroupName"] = kv
                                }
                                if kv, ok := certParams["vaultName"].(string); ok {
                                    result["keyVaultVaultName"] = kv
                                }
                                if kv, ok := certParams["secretName"].(string); ok {
                                    result["keyVaultSecretName"] = kv
                                }
                                if kv, ok := certParams["secretVersion"].(string); ok {
                                    result["keyVaultSecretVersion"] = kv
                                }
                            }
                        }
                    }
                }
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
    
    // Extract the custom domain ID
    customDomainId := strings.TrimSuffix(id, "/customHttps")
    
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
    
    certificateSource, ok := props["certificateSource"].(string)
    if !ok {
        return nil, errors.New("missing required property 'certificateSource'")
    }
    
    protocolType, ok := props["protocolType"].(string)
    if !ok {
        return nil, errors.New("missing required property 'protocolType'")
    }
    
    // Prepare the request body for enabling HTTPS
    body := map[string]any{
        "certificateSource": certificateSource,
        "protocolType": protocolType,
    }
    
    if minimumTlsVersion, ok := props["minimumTlsVersion"].(string); ok {
        body["minimumTlsVersion"] = minimumTlsVersion
    }
    
    // Configure certificate source parameters based on certificate source
    certificateSourceParams := map[string]any{}
    
    if certificateSource == "Cdn" {
        if certificateType, ok := props["certificateType"].(string); ok {
            certificateSourceParams["certificateType"] = certificateType
        } else {
            // Default to Shared if not specified
            certificateSourceParams["certificateType"] = "Shared"
        }
    } else if certificateSource == "AzureKeyVault" {
        // Validate required parameters for Key Vault
        requiredParams := []string{
            "keyVaultSubscriptionId", 
            "keyVaultResourceGroupName", 
            "keyVaultVaultName", 
            "keyVaultSecretName",
        }
        
        for _, param := range requiredParams {
            if _, ok := props[param].(string); !ok {
                return nil, errors.New("missing required property '" + param + "' for AzureKeyVault certificate source")
            }
        }
        
        certificateSourceParams["subscriptionId"] = props["keyVaultSubscriptionId"]
        certificateSourceParams["resourceGroupName"] = props["keyVaultResourceGroupName"]
        certificateSourceParams["vaultName"] = props["keyVaultVaultName"]
        certificateSourceParams["secretName"] = props["keyVaultSecretName"]
        
        if secretVersion, ok := props["keyVaultSecretVersion"].(string); ok {
            certificateSourceParams["secretVersion"] = secretVersion
        }
        
        // Default update and delete rules
        certificateSourceParams["updateRule"] = "NoAction"
        certificateSourceParams["deleteRule"] = "NoAction"
    }
    
    body["certificateSourceParameters"] = certificateSourceParams
    
    _, err := azureClient.Post(ctx, path, body, queryParams)
    if err != nil {
        return nil, err
    }
    
    return props, nil
}
```

### 2. Register the Custom Resource

Modify `provider/pkg/resources/customresources/customresources.go` to add the new custom resource to the `BuildCustomResources` function:

```go
func BuildCustomResources(env *azureEnv.Environment,
    azureClient azure.AzureClient,
    lookupResource ResourceLookupFunc,
    crudClientFactory crud.ResourceCrudClientFactory,
    subscriptionID string,
    bearerAuth autorest.Authorizer,
    tokenAuth autorest.Authorizer,
    kvBearerAuth autorest.Authorizer,
    userAgent string,
    tokenCred azcore.TokenCredential) (map[string]*CustomResource, error) {
    
    // ... existing code ...
    
    resources := []*CustomResource{
        keyVaultAccessPolicy(armKVClient),

        // Customization of regular resources
        blobContainerLegalHold(azureClient),
        portalDashboard(),
        customWebApp,
        customWebAppSlot,
        postgresConf,
        protectedItem,
        pimRoleManagementPolicy,
        pimRoleEligibilitySchedule,
        // Add the new custom resource
        cdnCustomDomainHttps(azureClient),
    }
    
    // ... rest of the function ...
}
```

### 3. Add Example Code

Create an example in the `examples` directory to demonstrate usage of the new resource:

```typescript
// examples/cdn-custom-domain-https/index.ts
import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as cdn from "@pulumi/azure-native/cdn";

// Create a resource group
const resourceGroup = new resources.ResourceGroup("resourceGroup");

// Create a CDN profile
const cdnProfile = new cdn.Profile("cdnProfile", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    sku: {
        name: "Standard_Microsoft",
    },
});

// Create a CDN endpoint
const cdnEndpoint = new cdn.Endpoint("cdnEndpoint", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    location: resourceGroup.location,
    originHostHeader: "www.contoso.com",
    origins: [{
        name: "contoso",
        hostName: "www.contoso.com",
    }],
});

// Create a custom domain
const customDomain = new cdn.CustomDomain("customDomain", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    endpointName: cdnEndpoint.name,
    hostName: "cdn.contoso.com",
});

// Enable HTTPS on the custom domain using CDN-managed certificate
const customDomainHttps = new cdn.CustomDomainHttps("customDomainHttps", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    endpointName: cdnEndpoint.name,
    customDomainName: customDomain.name,
    httpsEnabled: true,
    certificateSource: "Cdn",
    protocolType: "ServerNameIndication",
    certificateType: "Shared",
    minimumTlsVersion: "TLS12",
});

// Export the custom domain name
export const customDomainName = customDomain.name;
export const httpsEnabled = customDomainHttps.httpsEnabled;
```

## Testing Plan

1. Unit Tests:
   - Create tests for the CustomDomainHttps resource in the provider
   - Test both CDN-managed and Key Vault-based certificate scenarios
   - Test enabling and disabling HTTPS

2. Integration Tests:
   - Test the full workflow in a real Azure environment
   - Verify correct HTTPS provisioning state

## Notes

- The implementation follows the pattern used by other custom resources in the provider, particularly the BlobContainerLegalHold resource
- The resource is designed to be idempotent, handling both the creation and updating of HTTPS settings
- Error handling ensures proper validation of required parameters based on certificate source
- Since the actual HTTPS provisioning may take time, the resource doesn't wait for completion - users should check the CustomDomain's `customHttpsProvisioningState` property to monitor the process

## Future Enhancements

1. Add a way to track the provisioning status of the HTTPS enablement process
2. Consider supporting additional certificate sources if Azure adds them
3. Improve error handling for certificate validation failures

---

Once this implementation is complete, users will be able to enable and disable HTTPS on CDN custom domains directly through Pulumi, eliminating the need for external scripts or manual Azure portal steps.