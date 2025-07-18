// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This api gets ispCacheNode resource tls certificate histrory information
 *
 * Uses Azure REST API version 2024-11-30-preview.
 */
export function getEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistory(args: GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:connectedcache:getEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistory", {
        "cacheNodeResourceName": args.cacheNodeResourceName,
        "customerResourceName": args.customerResourceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryArgs {
    /**
     * Name of the ConnectedCache resource
     */
    cacheNodeResourceName: string;
    /**
     * Name of the Customer resource
     */
    customerResourceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Mcc cache node resource Tls certificate history details.
 */
export interface GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryResult {
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
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
     * Mcc cache node resource Tls certificate details.
     */
    readonly properties: outputs.connectedcache.MccCacheNodeTlsCertificatePropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.connectedcache.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * This api gets ispCacheNode resource tls certificate histrory information
 *
 * Uses Azure REST API version 2024-11-30-preview.
 */
export function getEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryOutput(args: GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:connectedcache:getEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistory", {
        "cacheNodeResourceName": args.cacheNodeResourceName,
        "customerResourceName": args.customerResourceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnterpriseMccCacheNodesOperationCacheNodeTlsCertificateHistoryOutputArgs {
    /**
     * Name of the ConnectedCache resource
     */
    cacheNodeResourceName: pulumi.Input<string>;
    /**
     * Name of the Customer resource
     */
    customerResourceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
