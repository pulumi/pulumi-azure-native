# Implementation Guide: Semi-Automated Integration with Azure API Specs

## Overview

This guide describes how to integrate a new or updated Azure API into the Pulumi Azure Native provider in a semi-automated way by leveraging the existing schema generation tools. This approach is ideal for adding support for new Azure services or updating existing ones to support new functionality.

## Prerequisites

1. **Environment Setup**:
   - Go 1.21+
   - NodeJS 18.X.X+
   - Python 3.10+ (for testing)
   - Azure REST API specs submodule properly initialized

2. **Understanding the Azure Specs**:
   - Familiarity with the [Azure REST API specifications](https://github.com/Azure/azure-rest-api-specs)
   - Understanding of the API version format (e.g., `2020-01-01`, `2020-01-01-preview`)

## High-Level Process

The process of integrating a new Azure API or updating an existing one consists of these steps:

1. **Update the API specs submodule** (if needed)
2. **Configure version selection**
3. **Generate the schema**
4. **Handle custom behaviors** (if needed)
5. **Test and validate**

This guide will walk through each step in detail.

## Detailed Steps

### 1. Update API Specs Submodule

If you're working with a new API or a more recent version than what's currently included:

```bash
cd azure-rest-api-specs
git checkout main
git pull origin main
cd ..
```

After updating, run:

```bash
make update_submodules
```

This command updates the API specs and refreshes the Azure provider list.

### 2. Configure Version Selection

The Pulumi Azure Native provider is selective about which API versions to include. This selection is managed in the `versions/` directory.

#### 2.1 Identify New API Versions

To check if new versions are available for a service:

```bash
bin/pulumi-gen-azure-native schema
```

This will generate reports in the `reports/` directory. Check:
- `allResourceVersionsByResource.json` - All available API versions by resource
- `allResourcesByVersion.json` - All resources available in each API version

#### 2.2 Controlling Which Specs Are Generated

There are several ways to control which API specs are processed:

##### Using Environment Variables for Targeted Schema Generation

For development and testing, you can use environment variables to limit which API specs are processed:

```bash
# Filter by a specific namespace (e.g., Compute, Storage, Network)
DEBUG_CODEGEN_NAMESPACES=Compute make schema

# Filter by multiple namespaces (comma-separated)
DEBUG_CODEGEN_NAMESPACES=Compute,Storage make schema

# Filter by API version pattern
DEBUG_CODEGEN_APIVERSIONS="2023*" make schema

# Combine namespace and version filters
DEBUG_CODEGEN_NAMESPACES=Compute DEBUG_CODEGEN_APIVERSIONS="2023*" make schema
```

These environment variables are processed in the `pulumi-gen-azure-native` tool and help you focus on specific parts of the API when developing.

##### Version Configuration Files

For permanent configuration, use the following files in the `versions/` directory:

1. **Base Version Configuration**: `versions/v{MAJOR_VERSION}.yaml`
   - Contains baseline configuration for each major version
   - Generated automatically by the tooling
   - Don't edit this file directly

2. **Default Version Configuration**: `versions/v{MAJOR_VERSION}-spec.yaml`
   - Defines which API version to use for each resource
   - Controls which API version is the default when users don't specify a version

3. **Custom Configuration**: `versions/v{MAJOR_VERSION}-config.yaml`
   - Where you make your customizations to include additional versions
   - Add new services or resources here

4. **Exclusion Configuration**: `versions/v{MAJOR_VERSION}-removed.json`
   - Defines resources to exclude from generation
   - Used for problematic or deprecated resources

###### Structure of Configuration Files

The configuration files use a hierarchical structure:

```yaml
# Add specific versions for a namespace
additionalVersions:
  Storage:                 # Namespace
    "2023-01-01":          # API Version
      resources:           # Resources to include
        - BlobContainer
        - StorageAccount
        
  Compute:
    "2023-03-01":
      resources:
        - VirtualMachine
      # You can also exclude specific resources within a version
      exclude:
        - DiskEncryptionSet

# Change default versions for resources
defaultVersionUpdates:
  Storage:
    BlobContainer: "2023-01-01"
    StorageAccount: "2023-01-01"

# Exclude resources entirely
remove:
  Storage:
    - Queue   # Exclude Queue resource from all versions
```

##### Resource Filtering Options

You can also control which resources are included or excluded from specific API versions:

```yaml
additionalVersions:
  Network:
    "2022-11-01":
      # Include only these resources
      resources:
        - VirtualNetwork
        - NetworkInterface
        - PublicIPAddress
        
      # Exclude these specific resources
      exclude:
        - RouteTable
        - RouteFilter
        
  # Include all resources from a specific version
  Security:
    "2023-01-01":
      includeAllResources: true  # Include all resources from this version
```

When working with additional versions, keep these points in mind:
- Use `includeAllResources: true` with caution as it may include many resources
- Always test thoroughly after adding new versions
- Be mindful of breaking changes between API versions

#### 2.3 Update Version Configuration

Based on your needs, edit the appropriate version configuration file:

- For the current major version, edit `versions/v{MAJOR_VERSION}-config.yaml`
- Follow the existing format to add your service/resource

Example configuration entry:

```yaml
# Add support for new storage API
additionalVersions:
  Storage:
    "2023-01-01":
      resources:
        - BlobContainer
        - StorageAccount
```

To make a version the default for new resources:

```yaml
defaultVersionUpdates:
  Storage:
    BlobContainer: "2023-01-01"
    StorageAccount: "2023-01-01"
```

### 3. Generate the Schema

After configuration changes, generate the schema:

```bash
make generate_schema generate_docs
```

This will:
1. Process the Azure API specs
2. Apply your version configurations
3. Generate the Pulumi schema and documentation
4. Create SDK resources for all selected API versions

#### 3.1 Schema Generation Process

The schema generation process follows these steps:

1. The `pulumi-gen-azure-native schema` command:
   - Reads the Azure REST API specifications
   - Applies version configuration from the `versions/` directory
   - Filters resources based on your configuration
   - Generates a unified Pulumi schema
   - Creates metadata for the provider

2. It outputs several files:
   - `bin/schema-full.json` - Complete schema with all included versions
   - `bin/schema-default-versions.json` - Schema with only default versions
   - `bin/metadata-compact.json` - Metadata for the provider

3. These files are then used to generate language-specific SDKs and documentation.

#### 3.2 Regenerating After Changes

When you make configuration changes, use these commands to regenerate:

```bash
# Regenerate all schemas
make generate_schema

# Regenerate schemas and documentation
make generate_schema generate_docs

# Regenerate and build all SDKs
make
```

If you're only interested in specific SDKs, you can use:

```bash
# Generate schemas and TypeScript SDK
make generate_schema generate_nodejs

# Generate schemas and Python SDK 
make generate_schema generate_python
```

### 4. Handle Custom Behaviors (If Needed)

Some Azure APIs require special handling due to limitations in the OpenAPI specifications or to support specific Pulumi patterns.

#### 4.1 Recognize When Custom Code Is Needed

Custom code may be needed when:
- The API uses POST operations to perform actions (like enabling/disabling features)
- Resources have special parent/child relationships
- Special data transformations are required
- Error handling needs customization

#### If the functionality can be represented as a standard resource, prefer to implement a custom resource that provides a familiar CRUD interface to Pulumi users, even if the underlying API uses different operations.

## Action-Based Custom Resources

Many Azure APIs expose functionality through POST operations rather than standard PUT/GET/DELETE operations. These are often referred to as "actions" and require special handling in the Pulumi provider.

### Recognizing Action-Based APIs

Action-based APIs have these characteristics:
- Functionality exposed via POST endpoints (often ending with verbs like `/enable`, `/disable`, `/start`, etc.)
- Operations that don't map cleanly to CRUD lifecycle
- Often operate on existing resources rather than creating new ones
- Configuration is provided in the request body rather than as path parameters

Common examples include:
- Enabling features on a resource (like HTTPS on CDN domains)
- Regenerating access keys or secrets
- Starting/stopping/restarting services
- Triggering operations on resources (like backups, syncs, validations)

### Implementing Action-Based Custom Resources

Action-based APIs are implemented as custom resources that use the underlying POST operations but expose a standard resource interface to Pulumi users.

#### Step 1: Research the API

Start by understanding the API from the Azure documentation and OpenAPI specs:

1. Identify the base resource path and action endpoints
2. Understand the request/response structure
3. Determine how to map user-provided properties to API parameters
4. Check for long-running operations or async patterns

For example, with CDN Custom Domain HTTPS:
- Base resource path: `/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}`
- Action endpoints: `/enableCustomHttps` and `/disableCustomHttps`
- Property mapping: `certificateSource`, `protocolType`, etc. in the request body

#### Step 2: Design the Custom Resource Interface

Design a Pulumi resource that presents a clean, declarative interface:

1. Define the resource token and path
2. Create input/output properties that make sense to Pulumi users
3. Choose appropriate property types and validation
4. Decide how to represent the desired state

For example, `CdnCustomDomainHttps` uses boolean `httpsEnabled` property to control enabling/disabling.

#### Step 3: Implement the Custom Resource Code

Create a new file in `provider/pkg/resources/customresources/`, following this pattern:

```go
package customresources

import (
    "context"
    "strings"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
    . "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
    "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
    "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
    "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Define the base path and action paths
const myCustomResourceBasePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Service/resources/{resourceName}"
const myCustomEnableActionPath = myCustomResourceBasePath + "/enableFeature"
const myCustomDisableActionPath = myCustomResourceBasePath + "/disableFeature"

// Define input/output properties
var myCustomResourceProperties = map[string]schema.PropertySpec{
    resourceGroupName: {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the resource group.",
    },
    "resourceName": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the resource.",
    },
    "featureEnabled": {
        TypeSpec:    schema.TypeSpec{Type: "boolean"},
        Description: "Whether the feature is enabled.",
    },
    // Additional properties specific to your resource
}

// Implement the custom resource
func myCustomResource(azureClient azure.AzureClient) *CustomResource {
    apiVersion := "2023-01-01"
    
    return &CustomResource{
        tok:  "azure-native:myservice:MyCustomFeature",
        path: myCustomResourceBasePath + "/customFeature", // Logical path for Pulumi resource ID
        LegacySchema: &schema.ResourceSpec{
            ObjectTypeSpec: schema.ObjectTypeSpec{
                Description: "Controls the custom feature settings for a resource.",
                Type:        "object",
                Properties:  myCustomResourceProperties,
            },
            InputProperties: myCustomResourceProperties,
            RequiredInputs: []string{
                resourceGroupName, "resourceName", "featureEnabled",
            },
        },
        Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, error) {
            return updateMyCustomResource(ctx, id, properties, azureClient, apiVersion)
        },
        Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, bool, error) {
            // Extract the base resource ID from the custom feature ID
            baseResourceId := strings.TrimSuffix(id, "/customFeature")
            
            // Read the base resource to determine the current state
            baseResource, err := azureClient.Get(ctx, baseResourceId, apiVersion, nil)
            if err != nil {
                if azure.IsNotFound(err) {
                    return nil, false, nil
                }
                return nil, false, err
            }
            
            // Extract feature state from the resource properties
            result := map[string]any{
                "featureEnabled": false, // Default state
            }
            
            props, ok := baseResource["properties"].(map[string]any)
            if !ok {
                return result, true, nil
            }
            
            // Set actual state based on resource properties
            if featureState, ok := props["featureState"].(string); ok && featureState == "Enabled" {
                result["featureEnabled"] = true
            }
            
            return result, true, nil
        },
        Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
            return updateMyCustomResource(ctx, id, news, azureClient, apiVersion)
        },
        Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
            // For most action-based resources, delete means reverting to a default state
            // Extract the base resource ID
            baseResourceId := strings.TrimSuffix(id, "/customFeature")
            
            // Disable the feature
            path := strings.TrimSuffix(id, "/customFeature") + "/disableFeature"
            
            queryParams := map[string]any{
                "api-version": apiVersion,
            }
            
            _, err := azureClient.Post(ctx, path, map[string]any{}, queryParams)
            return err
        },
    }
}

// Helper function for Create and Update operations
func updateMyCustomResource(ctx context.Context, id string, properties resource.PropertyMap, azureClient azure.AzureClient, apiVersion string) (map[string]any, error) {
    props := properties.Mappable()
    
    // Extract the base resource ID
    baseResourceId := strings.TrimSuffix(id, "/customFeature")
    
    featureEnabled, ok := props["featureEnabled"].(bool)
    if !ok {
        return nil, errors.New("missing required property 'featureEnabled'")
    }
    
    queryParams := map[string]any{
        "api-version": apiVersion,
    }
    
    // Determine whether to enable or disable based on featureEnabled flag
    var path string
    var body map[string]any
    
    if featureEnabled {
        path = baseResourceId + "/enableFeature"
        
        // Build the request body for enabling
        body = map[string]any{
            // Include required properties from the user input
            "property1": props["property1"],
            "property2": props["property2"],
        }
    } else {
        path = baseResourceId + "/disableFeature"
        // Usually disable operations take no parameters or minimal ones
        body = map[string]any{}
    }
    
    // Make the POST request to the appropriate action endpoint
    _, err := azureClient.Post(ctx, path, body, queryParams)
    if err != nil {
        return nil, err
    }
    
    return props, nil
}
```

#### Step 4: Register the Custom Resource

Add your resource to the `BuildCustomResources` function in `provider/pkg/resources/customresources/customresources.go`:

```go
func BuildCustomResources(env *azureEnv.Environment,
    azureClient azure.AzureClient,
    // ... other parameters ...
) (map[string]*CustomResource, error) {
    
    // ... existing code ...
    
    resources := []*CustomResource{
        keyVaultAccessPolicy(armKVClient),
        // ... other resources ...
        
        // Add your new custom resource
        myCustomResource(azureClient),
    }
    
    // ... rest of the function ...
}
```

### Real-World Example: CDN Custom Domain HTTPS

Let's examine how the Pulumi Azure Native provider implements the custom resource for enabling HTTPS on CDN custom domains.

#### Understanding the API

The Azure API exposes two POST endpoints for managing HTTPS on CDN custom domains:
- `/enableCustomHttps` - Enables HTTPS with certificate settings
- `/disableCustomHttps` - Disables HTTPS

These don't map cleanly to standard CRUD operations, but users expect a declarative resource that handles the underlying complexity.

#### Resource Implementation

The implementation consists of:

1. **Resource Definition**:
```go
const cdnCustomDomainHttpsPath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/customHttps"

var cdnCustomDomainHttpsProperties = map[string]schema.PropertySpec{
    resourceGroupName: {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the resource group that contains the CDN profile.",
    },
    // ... other properties ...
    "httpsEnabled": {
        TypeSpec:    schema.TypeSpec{Type: "boolean"},
        Description: "Whether HTTPS is enabled on the custom domain.",
    },
    "certificateSource": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Certificate source for the custom domain. One of: 'Cdn', 'AzureKeyVault'.",
    },
    // ... additional certificate properties ...
}
```

2. **CRUD Operations**:
   - **Create/Update**: Calls `enableCustomHttps` or `disableCustomHttps` POST endpoints based on the `httpsEnabled` property
   - **Read**: Reads the parent `CustomDomain` resource and checks its HTTPS provisioning state
   - **Delete**: Calls the `disableCustomHttps` endpoint to revert to the default state

3. **ID Composition**:
   - The resource ID is the base custom domain path with "/customHttps" appended
   - Operations extract the base resource ID when needed to call the appropriate endpoints

4. **Implementation Details**:
```go
func updateCdnCustomDomainHttps(ctx context.Context, id string, properties resource.PropertyMap, azureClient azure.AzureClient, apiVersion string) (map[string]any, error) {
    props := properties.Mappable()
    
    // Extract the custom domain ID
    customDomainId := strings.TrimSuffix(id, "/customHttps")
    
    httpsEnabled, ok := props["httpsEnabled"].(bool)
    if !ok {
        return nil, errors.New("missing required property 'httpsEnabled'")
    }
    
    // Choose the appropriate endpoint based on httpsEnabled
    if !httpsEnabled {
        path := customDomainId + "/disableCustomHttps"
        _, err := azureClient.Post(ctx, path, map[string]any{}, queryParams)
        // ...
    } else {
        path := customDomainId + "/enableCustomHttps"
        // Build the request body with certificate information
        body := map[string]any{
            "certificateSource": certificateSource,
            "protocolType": protocolType,
            // ...
        }
        _, err := azureClient.Post(ctx, path, body, queryParams)
        // ...
    }
    
    return props, nil
}
```

### Handling Long-Running Operations

Many Azure action APIs are asynchronous and return a 202 Accepted response with an operation location header. For these operations:

1. Use the `azure.LongRunningOperation` interface to track operation status
2. In the Read operation, check for transitional states
3. Use the parent resource's status properties to determine the current state

Example for handling async operations:

```go
// In Create/Update function
operation, err := azureClient.PostLongRunning(ctx, path, body, queryParams)
if err != nil {
    return nil, err
}

// Wait for the operation to complete (optional, based on your needs)
if err = operation.WaitForCompletion(ctx); err != nil {
    return nil, err
}

// In Read function
// Check the provisioning state of the parent resource
if provisioningState, ok := props["provisioningState"].(string); ok {
    // Handle transitional states
    if provisioningState == "Updating" || provisioningState == "Provisioning" {
        // Resource is still being updated
        // You may want to return the last known state
    }
}
```

### Best Practices for Action-Based Resources

1. **Use Boolean Toggles**: When possible, use boolean properties like `enabled` to control action-based features
2. **Map to Familiar Patterns**: Make the resource feel like a standard resource to users
3. **Handle Transitional States**: Check for and handle in-progress operations
4. **Detailed Error Messages**: Provide clear error messages for API-specific failures
5. **Document Thoroughly**: Explain the behavior and expected states in the resource documentation

### Example Usage in Pulumi

The implementation allows users to use a declarative pattern:

```typescript
// Example usage of the CDN Custom Domain HTTPS resource
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
```

### Summary

Action-based custom resources allow you to present a standard resource interface for Azure APIs that use POST operations. By hiding the complexity of these operations, you provide a more consistent and intuitive experience for Pulumi users.

## Creating Standard Custom Resources

If your API conforms more closely to standard Azure resource patterns with PUT/GET/DELETE operations, you can implement a custom resource with a more straightforward mapping.

#### 4.2 Create a Standard Custom Resource

For standard resource patterns:

1. Create a new file in `provider/pkg/resources/customresources/`
2. Implement the custom resource following this pattern:

```go
// Define token and path
const myCustomResourcePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/..."

// Define properties
var myCustomResourceProperties = map[string]schema.PropertySpec{
    // Standard properties
    resourceGroupName: {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Name of the resource group.",
    },
    // Custom properties for your resource
    "myProperty": {
        TypeSpec:    schema.TypeSpec{Type: "string"},
        Description: "Description of the property.",
    },
}

// Create the custom resource
func myCustomResource(azureClient azure.AzureClient) *CustomResource {
    apiVersion := "2023-01-01" // Use appropriate version
    
    return &CustomResource{
        tok:  "azure-native:myservice:MyCustomResource", // Use appropriate module and name
        path: myCustomResourcePath,
        LegacySchema: &schema.ResourceSpec{
            ObjectTypeSpec: schema.ObjectTypeSpec{
                Description: "Manages a custom Azure resource.",
                Type:        "object",
                Properties:  myCustomResourceProperties,
            },
            InputProperties: myCustomResourceProperties,
            RequiredInputs: []string{
                resourceGroupName, "requiredProperty1", "requiredProperty2",
            },
        },
        // Implement CRUD operations
        Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, error) {
            // Implementation details
        },
        Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, bool, error) {
            // Implementation details
        },
        Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
            // Implementation details
        },
        Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
            // Implementation details
        },
    }
}
```

3. Register the custom resource in `provider/pkg/resources/customresources/customresources.go`

### 5. Test and Validate

#### 5.1 Build the Provider and SDKs

```bash
make
```

#### 5.2 Create Integration Tests

Create an example that uses your new resource:

1. Add a new example directory in `examples/`
2. Create the basic resource configuration
3. Test the example using `make test`

For a specific test:

```bash
make test TEST_NAME=TestMyExample
```

#### 5.3 Test With Local Provider

To test with the local provider:

```bash
make install_provider
cd examples/my-example
pulumi up
```

## Advanced Topics

### Adding Custom Documentation

For important resources, add documentation in `docs/resources/`:

1. Create a file following the pattern `[module]-[Resource].md`
2. Include proper documentation with examples
3. Regenerate docs with `make generate generate_docs`

### API Version Lifecycle Management

When managing multiple API versions:

1. Keep the previous version as a fallback
2. Test backward compatibility
3. Consider using `versions/v{MAJOR_VERSION}-removed.json` if resources need to be excluded

### Handling API Changes

When Azure makes breaking changes:

1. Record the changes in `versions/v{MAJOR_VERSION}-removed.json`
2. Create custom adapters if needed for compatibility
3. Document breaking changes for users

## Troubleshooting

### Common Issues

1. **Schema Generation Failures**:
   - Ensure the specs submodule is properly initialized
   - Check for syntax errors in version configuration files
   - Look at error messages for specific resources or API versions

2. **Resource Not Found**:
   - Confirm the resource exists in the selected API version
   - Check resource naming in the version configuration
   - Verify the correct module name is being used

3. **Custom Resource Issues**:
   - Debug paths and parameter mappings
   - Check REST API documentation for any undocumented requirements
   - Test direct HTTP requests using the Azure CLI or Postman

4. **Action-Based Resource Issues**:
   - Ensure you're using the correct action endpoints
   - Check for required properties in the request body
   - Verify the parent resource exists before attempting actions
   - Handle asynchronous operations correctly

## References

- Azure REST API Specifications: https://github.com/Azure/azure-rest-api-specs
- Azure SDK Generation Process: https://github.com/Azure/azure-sdk-for-go
- Pulumi Schema Reference: https://www.pulumi.com/docs/reference/pkg/