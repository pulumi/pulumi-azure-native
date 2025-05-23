// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets the specified Data Lake Analytics compute policy.
 *
 * Uses Azure REST API version 2019-11-01-preview.
 */
export function getComputePolicy(args: GetComputePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetComputePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:datalakeanalytics:getComputePolicy", {
        "accountName": args.accountName,
        "computePolicyName": args.computePolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetComputePolicyArgs {
    /**
     * The name of the Data Lake Analytics account.
     */
    accountName: string;
    /**
     * The name of the compute policy to retrieve.
     */
    computePolicyName: string;
    /**
     * The name of the Azure resource group.
     */
    resourceGroupName: string;
}

/**
 * Data Lake Analytics compute policy information.
 */
export interface GetComputePolicyResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The resource identifier.
     */
    readonly id: string;
    /**
     * The maximum degree of parallelism per job this user can use to submit jobs.
     */
    readonly maxDegreeOfParallelismPerJob: number;
    /**
     * The minimum priority per job this user can use to submit jobs.
     */
    readonly minPriorityPerJob: number;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The AAD object identifier for the entity to create a policy for.
     */
    readonly objectId: string;
    /**
     * The type of AAD object the object identifier refers to.
     */
    readonly objectType: string;
    /**
     * The resource type.
     */
    readonly type: string;
}
/**
 * Gets the specified Data Lake Analytics compute policy.
 *
 * Uses Azure REST API version 2019-11-01-preview.
 */
export function getComputePolicyOutput(args: GetComputePolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetComputePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:datalakeanalytics:getComputePolicy", {
        "accountName": args.accountName,
        "computePolicyName": args.computePolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetComputePolicyOutputArgs {
    /**
     * The name of the Data Lake Analytics account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the compute policy to retrieve.
     */
    computePolicyName: pulumi.Input<string>;
    /**
     * The name of the Azure resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
