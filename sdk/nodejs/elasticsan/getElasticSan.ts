// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a ElasticSan.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2021-11-20-preview, 2022-12-01-preview, 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getElasticSan(args: GetElasticSanArgs, opts?: pulumi.InvokeOptions): Promise<GetElasticSanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:elasticsan:getElasticSan", {
        "elasticSanName": args.elasticSanName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetElasticSanArgs {
    /**
     * The name of the ElasticSan.
     */
    elasticSanName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Response for ElasticSan request.
 */
export interface GetElasticSanResult {
    /**
     * Logical zone for Elastic San resource; example: ["1"].
     */
    readonly availabilityZones?: string[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Base size of the Elastic San appliance in TiB.
     */
    readonly baseSizeTiB: number;
    /**
     * Extended size of the Elastic San appliance in TiB.
     */
    readonly extendedCapacitySizeTiB: number;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The list of Private Endpoint Connections.
     */
    readonly privateEndpointConnections: outputs.elasticsan.PrivateEndpointConnectionResponse[];
    /**
     * State of the operation on the resource.
     */
    readonly provisioningState: string;
    /**
     * Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
     */
    readonly publicNetworkAccess?: string;
    /**
     * resource sku
     */
    readonly sku: outputs.elasticsan.SkuResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.elasticsan.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Total Provisioned IOPS of the Elastic San appliance.
     */
    readonly totalIops: number;
    /**
     * Total Provisioned MBps Elastic San appliance.
     */
    readonly totalMBps: number;
    /**
     * Total size of the Elastic San appliance in TB.
     */
    readonly totalSizeTiB: number;
    /**
     * Total size of the provisioned Volumes in GiB.
     */
    readonly totalVolumeSizeGiB: number;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Total number of volume groups in this Elastic San appliance.
     */
    readonly volumeGroupCount: number;
}
/**
 * Get a ElasticSan.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2021-11-20-preview, 2022-12-01-preview, 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getElasticSanOutput(args: GetElasticSanOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetElasticSanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:elasticsan:getElasticSan", {
        "elasticSanName": args.elasticSanName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetElasticSanOutputArgs {
    /**
     * The name of the ElasticSan.
     */
    elasticSanName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
