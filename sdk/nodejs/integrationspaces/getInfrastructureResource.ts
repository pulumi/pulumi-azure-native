// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a InfrastructureResource
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getInfrastructureResource(args: GetInfrastructureResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetInfrastructureResourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:integrationspaces:getInfrastructureResource", {
        "infrastructureResourceName": args.infrastructureResourceName,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetInfrastructureResourceArgs {
    /**
     * The name of the infrastructure resource in the space.
     */
    infrastructureResourceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the space
     */
    spaceName: string;
}

/**
 * An infrastructure resource under Space.
 */
export interface GetInfrastructureResourceResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * The id of the infrastructure resource.
     */
    readonly resourceId: string;
    /**
     * The type of the infrastructure resource.
     */
    readonly resourceType: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.integrationspaces.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a InfrastructureResource
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getInfrastructureResourceOutput(args: GetInfrastructureResourceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInfrastructureResourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:integrationspaces:getInfrastructureResource", {
        "infrastructureResourceName": args.infrastructureResourceName,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetInfrastructureResourceOutputArgs {
    /**
     * The name of the infrastructure resource in the space.
     */
    infrastructureResourceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the space
     */
    spaceName: pulumi.Input<string>;
}
