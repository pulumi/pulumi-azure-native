// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets details about the specified dedicated capacity.
 *
 * Uses Azure REST API version 2021-01-01.
 */
export function getCapacityDetails(args: GetCapacityDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetCapacityDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:powerbidedicated:getCapacityDetails", {
        "dedicatedCapacityName": args.dedicatedCapacityName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCapacityDetailsArgs {
    /**
     * The name of the dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
     */
    dedicatedCapacityName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Represents an instance of a Dedicated Capacity resource.
 */
export interface GetCapacityDetailsResult {
    /**
     * A collection of Dedicated capacity administrators
     */
    readonly administration?: outputs.powerbidedicated.DedicatedCapacityAdministratorsResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Capacity name
     */
    readonly friendlyName: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
     */
    readonly mode?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
     */
    readonly provisioningState: string;
    /**
     * The SKU of the PowerBI Dedicated capacity resource.
     */
    readonly sku: outputs.powerbidedicated.CapacitySkuResponse;
    /**
     * The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
     */
    readonly state: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.powerbidedicated.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Tenant ID for the capacity. Used for creating Pro Plus capacity.
     */
    readonly tenantId: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets details about the specified dedicated capacity.
 *
 * Uses Azure REST API version 2021-01-01.
 */
export function getCapacityDetailsOutput(args: GetCapacityDetailsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCapacityDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:powerbidedicated:getCapacityDetails", {
        "dedicatedCapacityName": args.dedicatedCapacityName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCapacityDetailsOutputArgs {
    /**
     * The name of the dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
     */
    dedicatedCapacityName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
