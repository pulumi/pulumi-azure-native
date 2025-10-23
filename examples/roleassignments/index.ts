// Copyright 2025, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as authorization from "@pulumi/azure-native/authorization";

const config = new pulumi.Config();

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("resourceGroup", {
    location: "westeurope",
});

// "Reports Reader" role - a harmless role with little security risk
const reportsReaderRoleId = "028f4ed7-e2a9-465e-a8f4-9c0ffdfdc027";

// Scenario 1: Standard role assignment to testuser
const testUserAssignment = new authorization.RoleAssignment("testUserAssignment", {
    roleDefinitionId: pulumi.interpolate`${resourceGroup.id}/providers/Microsoft.Authorization/roleDefinitions/${reportsReaderRoleId}`,
    principalId: "e61877e9-d26b-4890-9a3c-2287e77ca427", // testuser
    principalType: authorization.PrincipalType.User,
    scope: resourceGroup.id,
});

// Scenario 2: Cross-tenant assignment to a managed identity (optional)
// The managed identity should be created separately using:
// az identity create --name cross-tenant-mi --resource-group <rg-name> --subscription <delegated-subscription-id>
const delegatedIdentityResourceId = config.get("delegatedManagedIdentityResourceId");
const delegatedIdentityPrincipalId = config.get("delegatedManagedIdentityPrincipalId");

const crossTenantAssignment = delegatedIdentityResourceId && delegatedIdentityPrincipalId
    ? new authorization.RoleAssignment("crossTenantAssignment", {
        roleDefinitionId: pulumi.interpolate`${resourceGroup.id}/providers/Microsoft.Authorization/roleDefinitions/${reportsReaderRoleId}`,
        principalId: delegatedIdentityPrincipalId, // Principal ID from the managed identity
        principalType: authorization.PrincipalType.ServicePrincipal,
        scope: resourceGroup.id,
        delegatedManagedIdentityResourceId: delegatedIdentityResourceId,
    })
    : undefined;

// Export outputs
export const resourceGroupId = resourceGroup.id;
export const resourceGroupName = resourceGroup.name;

export const testUserAssignmentId = testUserAssignment.id;
export const testUserAssignmentPrincipalId = testUserAssignment.principalId;
export const testUserAssignmentScope = testUserAssignment.scope;

export const crossTenantAssignmentId = crossTenantAssignment?.id;
export const crossTenantAssignmentPrincipalId = crossTenantAssignment?.principalId;
export const crossTenantAssignmentScope = crossTenantAssignment?.scope;
