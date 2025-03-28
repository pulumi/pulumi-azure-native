// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Defines binding between a resource and role
 */
export function getTrustedAccessRoleBinding(args: GetTrustedAccessRoleBindingArgs, opts?: pulumi.InvokeOptions): Promise<GetTrustedAccessRoleBindingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:containerservice/v20230902preview:getTrustedAccessRoleBinding", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
        "trustedAccessRoleBindingName": args.trustedAccessRoleBindingName,
    }, opts);
}

export interface GetTrustedAccessRoleBindingArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the managed cluster resource.
     */
    resourceName: string;
    /**
     * The name of trusted access role binding.
     */
    trustedAccessRoleBindingName: string;
}

/**
 * Defines binding between a resource and role
 */
export interface GetTrustedAccessRoleBindingResult {
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The current provisioning state of trusted access role binding.
     */
    readonly provisioningState: string;
    /**
     * A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
     */
    readonly roles: string[];
    /**
     * The ARM resource ID of source resource that trusted access is configured for.
     */
    readonly sourceResourceId: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.containerservice.v20230902preview.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Defines binding between a resource and role
 */
export function getTrustedAccessRoleBindingOutput(args: GetTrustedAccessRoleBindingOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTrustedAccessRoleBindingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:containerservice/v20230902preview:getTrustedAccessRoleBinding", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
        "trustedAccessRoleBindingName": args.trustedAccessRoleBindingName,
    }, opts);
}

export interface GetTrustedAccessRoleBindingOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the managed cluster resource.
     */
    resourceName: pulumi.Input<string>;
    /**
     * The name of trusted access role binding.
     */
    trustedAccessRoleBindingName: pulumi.Input<string>;
}
