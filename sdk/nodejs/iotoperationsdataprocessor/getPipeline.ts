// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a Pipeline
 *
 * Uses Azure REST API version 2023-10-04-preview.
 */
export function getPipeline(args: GetPipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:iotoperationsdataprocessor:getPipeline", {
        "instanceName": args.instanceName,
        "pipelineName": args.pipelineName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPipelineArgs {
    /**
     * Name of instance.
     */
    instanceName: string;
    /**
     * Name of pipeline
     */
    pipelineName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A Pipeline resource belonging to an Instance resource.
 */
export interface GetPipelineResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Detailed description of the Pipeline.
     */
    readonly description?: string;
    /**
     * Flag indicating whether the pipeline should be running or not.
     */
    readonly enabled: boolean;
    /**
     * Edge location of the resource.
     */
    readonly extendedLocation: outputs.iotoperationsdataprocessor.ExtendedLocationResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Information about where to pull input data from.
     */
    readonly input: outputs.iotoperationsdataprocessor.PipelineInputResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * Map of stage ids to stage configurations for all pipeline processing and output stages.
     */
    readonly stages: {[key: string]: outputs.iotoperationsdataprocessor.PipelineStageResponse};
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.iotoperationsdataprocessor.SystemDataResponse;
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
 * Get a Pipeline
 *
 * Uses Azure REST API version 2023-10-04-preview.
 */
export function getPipelineOutput(args: GetPipelineOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPipelineResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:iotoperationsdataprocessor:getPipeline", {
        "instanceName": args.instanceName,
        "pipelineName": args.pipelineName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPipelineOutputArgs {
    /**
     * Name of instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Name of pipeline
     */
    pipelineName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
