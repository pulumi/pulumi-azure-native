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

const res = new authorization.RoleEligibilitySchedule("resr", {
    scope: resourceGroup.id,
    roleDefinitionId: roleAssignment.roleDefinitionId,
    principalId: roleAssignment.principalId,
    justification: "why not",
    scheduleInfo: {
        startDateTime: new Date().toISOString(),
        expiration: {
            duration: "P1D",
            type: 'AfterDuration'
        },
    },
    // ticketInfo: {
    //     ticketNumber: "1234567890",
    //     ticketSystem: "Pulumi",
    // },
}, { dependsOn: [roleAssignment] });

export const roleAssignmentId = roleAssignment.id;
export const roleAssignmentPrincipalId = roleAssignment.principalId;
export const roleAssignmentScope = roleAssignment.scope;
export const roleAssignmentRoleDefinitionId = roleAssignment.roleDefinitionId;

export const resId = res.id;
export const resStatus = res.status