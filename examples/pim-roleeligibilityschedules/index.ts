import { v4 as uuidv4 } from 'uuid';

import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as authorization from "@pulumi/azure-native/authorization";

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("resourceGroup");

const roleAssignment = new authorization.RoleAssignment("roleAssignment", {
    // Azure RBAC Contributor role.
    roleDefinitionId: pulumi.interpolate `${resourceGroup.id}/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c`,
    principalId: "db8fa9c3-36b4-4474-91b0-aa96bcc706bc", // Thomas K.
    scope: resourceGroup.id,
});

const resr = new authorization.RoleEligibilityScheduleRequest("resr", {
    requestType: authorization.RequestType.AdminAssign,
    roleEligibilityScheduleRequestName: uuidv4(),
    // https://github.com/Azure/azure-rest-api-specs/blob/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/RoleEligibilityScheduleRequest.json:
    // "The scope of the role eligibility schedule request to create. The scope can be any REST resource instance. For
    // example, use '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/' for a subscription,
    // '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a
    // resource group, and '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
    // for a resource."
    scope: resourceGroup.id,
    roleDefinitionId: roleAssignment.roleDefinitionId,
    principalId: roleAssignment.principalId,
    scheduleInfo: {
        startDateTime: new Date().toISOString(),
        expiration: {
            duration: "P1D",
        },
    },
    ticketInfo: {
        ticketNumber: "1234567890",
        ticketSystem: "Pulumi",
    },
}, { dependsOn: [roleAssignment] });

export const roleAssignmentId = roleAssignment.id;
export const roleAssignmentPrincipalId = roleAssignment.principalId;
export const roleAssignmentScope = roleAssignment.scope;
export const roleAssignmentRoleDefinitionId = roleAssignment.roleDefinitionId;