import { v4 as uuidv4 } from 'uuid';

import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as authorization from "@pulumi/azure-native/authorization";

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("resourceGroup", {
    location: "westeurope",
});

// "Reports Reader" role, a harmless one with little security risk.
const roleId = "028f4ed7-e2a9-465e-a8f4-9c0ffdfdc027"; //"4a5d8f65-41da-4de4-8968-e035b65339cf";
const roleAssignment = new authorization.RoleAssignment("roleAssignment", {
    roleDefinitionId: pulumi.interpolate `${resourceGroup.id}/providers/Microsoft.Authorization/roleDefinitions/${roleId}`,
    principalId: "e61877e9-d26b-4890-9a3c-2287e77ca427", // user "testuser"
    scope: resourceGroup.id,
});

const res = new authorization.PimRoleEligibilitySchedule("res2", {
    scope: resourceGroup.id,
    roleDefinitionId: roleAssignment.roleDefinitionId,
    principalId: roleAssignment.principalId,
    justification: "why not",
    scheduleInfo: {
        startDateTime: new Date().toISOString(),
        expiration: {
            duration: "P1H",
            type: 'AfterDuration'
        },
    },
    // ticketInfo: {
    //     ticketNumber: "1234567890",
    //     ticketSystem: "Pulumi",
    // },
}, {
    dependsOn: [roleAssignment],
    ignoreChanges: ["scheduleInfo.startDateTime"]
});

export const roleAssignmentId = roleAssignment.id;
export const roleAssignmentPrincipalId = roleAssignment.principalId;
export const roleAssignmentScope = roleAssignment.scope;
export const roleAssignmentRoleDefinitionId = roleAssignment.roleDefinitionId;

export const resId = res.id;
export const resStatus = res.status