// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieve an Azure ML commitment plan by its subscription, resource group and name.
 *
 * Uses Azure REST API version 2016-05-01-preview.
 */
export function getCommitmentPlan(args: GetCommitmentPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetCommitmentPlanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:machinelearning:getCommitmentPlan", {
        "commitmentPlanName": args.commitmentPlanName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCommitmentPlanArgs {
    /**
     * The Azure ML commitment plan name.
     */
    commitmentPlanName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * An Azure ML commitment plan resource.
 */
export interface GetCommitmentPlanResult {
    /**
     * An entity tag used to enforce optimistic concurrency.
     */
    readonly etag?: string;
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The commitment plan properties.
     */
    readonly properties: outputs.machinelearning.CommitmentPlanPropertiesResponse;
    /**
     * The commitment plan SKU.
     */
    readonly sku?: outputs.machinelearning.ResourceSkuResponse;
    /**
     * User-defined tags for the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Retrieve an Azure ML commitment plan by its subscription, resource group and name.
 *
 * Uses Azure REST API version 2016-05-01-preview.
 */
export function getCommitmentPlanOutput(args: GetCommitmentPlanOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCommitmentPlanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:machinelearning:getCommitmentPlan", {
        "commitmentPlanName": args.commitmentPlanName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCommitmentPlanOutputArgs {
    /**
     * The Azure ML commitment plan name.
     */
    commitmentPlanName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
}
