// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves the properties of an existing Azure Cosmos DB FleetAnalytics under a fleet
 *
 * Uses Azure REST API version 2025-05-01-preview.
 */
export function getFleetAnalytic(args: GetFleetAnalyticArgs, opts?: pulumi.InvokeOptions): Promise<GetFleetAnalyticResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cosmosdb:getFleetAnalytic", {
        "fleetAnalyticsName": args.fleetAnalyticsName,
        "fleetName": args.fleetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFleetAnalyticArgs {
    /**
     * Cosmos DB fleetAnalytics name.
     */
    fleetAnalyticsName: string;
    /**
     * Cosmos DB fleet name. Needs to be unique under a subscription.
     */
    fleetName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * An Azure Cosmos DB FleetAnalytics.
 */
export interface GetFleetAnalyticResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * A provisioning state of the FleetAnalytics.
     */
    readonly provisioningState: string;
    /**
     * The type of the fleet analytics resource.
     */
    readonly storageLocationType?: string;
    /**
     * The unique identifier of the fleet analytics resource.
     */
    readonly storageLocationUri?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.cosmosdb.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Retrieves the properties of an existing Azure Cosmos DB FleetAnalytics under a fleet
 *
 * Uses Azure REST API version 2025-05-01-preview.
 */
export function getFleetAnalyticOutput(args: GetFleetAnalyticOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFleetAnalyticResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cosmosdb:getFleetAnalytic", {
        "fleetAnalyticsName": args.fleetAnalyticsName,
        "fleetName": args.fleetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFleetAnalyticOutputArgs {
    /**
     * Cosmos DB fleetAnalytics name.
     */
    fleetAnalyticsName: pulumi.Input<string>;
    /**
     * Cosmos DB fleet name. Needs to be unique under a subscription.
     */
    fleetName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
