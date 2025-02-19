// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetRoleManagementPolicyArgs, GetRoleManagementPolicyResult, GetRoleManagementPolicyOutputArgs } from "./getRoleManagementPolicy";
export const getRoleManagementPolicy: typeof import("./getRoleManagementPolicy").getRoleManagementPolicy = null as any;
export const getRoleManagementPolicyOutput: typeof import("./getRoleManagementPolicy").getRoleManagementPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getRoleManagementPolicy","getRoleManagementPolicyOutput"], () => require("./getRoleManagementPolicy"));

export { GetRoleManagementPolicyAssignmentArgs, GetRoleManagementPolicyAssignmentResult, GetRoleManagementPolicyAssignmentOutputArgs } from "./getRoleManagementPolicyAssignment";
export const getRoleManagementPolicyAssignment: typeof import("./getRoleManagementPolicyAssignment").getRoleManagementPolicyAssignment = null as any;
export const getRoleManagementPolicyAssignmentOutput: typeof import("./getRoleManagementPolicyAssignment").getRoleManagementPolicyAssignmentOutput = null as any;
utilities.lazyLoad(exports, ["getRoleManagementPolicyAssignment","getRoleManagementPolicyAssignmentOutput"], () => require("./getRoleManagementPolicyAssignment"));

export { RoleManagementPolicyArgs } from "./roleManagementPolicy";
export type RoleManagementPolicy = import("./roleManagementPolicy").RoleManagementPolicy;
export const RoleManagementPolicy: typeof import("./roleManagementPolicy").RoleManagementPolicy = null as any;
utilities.lazyLoad(exports, ["RoleManagementPolicy"], () => require("./roleManagementPolicy"));

export { RoleManagementPolicyAssignmentArgs } from "./roleManagementPolicyAssignment";
export type RoleManagementPolicyAssignment = import("./roleManagementPolicyAssignment").RoleManagementPolicyAssignment;
export const RoleManagementPolicyAssignment: typeof import("./roleManagementPolicyAssignment").RoleManagementPolicyAssignment = null as any;
utilities.lazyLoad(exports, ["RoleManagementPolicyAssignment"], () => require("./roleManagementPolicyAssignment"));


// Export enums:
export * from "../../types/enums/authorization/v20240901preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:authorization/v20240901preview:RoleManagementPolicy":
                return new RoleManagementPolicy(name, <any>undefined, { urn })
            case "azure-native:authorization/v20240901preview:RoleManagementPolicyAssignment":
                return new RoleManagementPolicyAssignment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "authorization/v20240901preview", _module)
