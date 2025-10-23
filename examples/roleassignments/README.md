# Azure Native RoleAssignment Example

This example demonstrates how to create Azure role assignments using Pulumi, including both standard and cross-tenant scenarios.

## Overview

This example demonstrates two scenarios:
1. **Scenario 1**: A standard role assignment to a user (testuser) with the Reports Reader role
2. **Scenario 2** (optional): A cross-tenant role assignment to a managed identity with the Reports Reader role

## Prerequisites

- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
- [Node.js](https://nodejs.org/)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)
- Azure subscription(s) with appropriate permissions

## Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Configure your Azure credentials:
   ```bash
   az login
   ```

3. Initialize a new Pulumi stack:
   ```bash
   pulumi stack init dev
   ```

## Running the Example

### Scenario 1: Standard Role Assignment Only

To run with just the standard testuser assignment:

```bash
pulumi up
```

This will create:
- A resource group in West Europe
- A role assignment granting the testuser the Reports Reader role on the resource group

### Scenario 2: Cross-Tenant Role Assignment (Optional)

To include the cross-tenant scenario, you need to create a managed identity in a different Azure subscription/tenant first.

1. Create a managed identity in the delegated subscription:
   ```bash
   az identity create \
     --name cross-tenant-mi \
     --resource-group <delegated-resource-group> \
     --subscription <delegated-subscription-id>
   ```

2. Get the resource ID and principal ID of the managed identity:
   ```bash
   az identity show \
     --name cross-tenant-mi \
     --resource-group <delegated-resource-group> \
     --subscription <delegated-subscription-id> \
     --query '{resourceId:id, principalId:principalId}' \
     --output json
   ```

3. Configure the Pulumi stack with the managed identity details:
   ```bash
   pulumi config set delegatedManagedIdentityResourceId "<resource-id-from-step-2>"
   pulumi config set delegatedManagedIdentityPrincipalId "<principal-id-from-step-2>"
   ```

4. Run the Pulumi program:
   ```bash
   pulumi up
   ```

## Understanding Cross-Tenant Role Assignments

Cross-tenant role assignments allow you to grant permissions to identities from a different Azure AD tenant. This is particularly useful in:

- Multi-tenant SaaS applications
- Partner/vendor access scenarios
- Cross-subscription delegated access

The `delegatedManagedIdentityResourceId` property tells Azure that the principal belongs to a different tenant, and Azure will make the appropriate cross-tenant API calls to validate and create the assignment.

## Outputs

The example exports the following outputs:

- `resourceGroupId`: The ID of the created resource group
- `resourceGroupName`: The name of the created resource group
- `testUserAssignmentId`: The ID of the testuser role assignment
- `crossTenantAssignmentId`: The ID of the cross-tenant role assignment (if configured)

## Clean Up

To destroy the resources:

```bash
pulumi destroy
```

## Notes

- The Reports Reader role (`028f4ed7-e2a9-465e-a8f4-9c0ffdfdc027`) is used as it has minimal security impact
- The testuser principal ID (`e61877e9-d26b-4890-9a3c-2287e77ca427`) is a placeholder and should be replaced with an actual user in your tenant
- Cross-tenant assignments require that the delegated tenant has authorized your tenant for cross-tenant access
