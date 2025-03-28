// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * This method gets the unencrypted secrets related to the job.
 */
export function listJobCredentials(args: ListJobCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<ListJobCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:databox/v20250201:listJobCredentials", {
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListJobCredentialsArgs {
    /**
     * The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
     */
    jobName: string;
    /**
     * The Resource Group Name
     */
    resourceGroupName: string;
}

/**
 * List of unencrypted credentials for accessing device.
 */
export interface ListJobCredentialsResult {
    /**
     * Link for the next set of unencrypted credentials.
     */
    readonly nextLink?: string;
    /**
     * List of unencrypted credentials.
     */
    readonly value?: outputs.databox.v20250201.UnencryptedCredentialsResponse[];
}
/**
 * This method gets the unencrypted secrets related to the job.
 */
export function listJobCredentialsOutput(args: ListJobCredentialsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListJobCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:databox/v20250201:listJobCredentials", {
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListJobCredentialsOutputArgs {
    /**
     * The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
     */
    jobName: pulumi.Input<string>;
    /**
     * The Resource Group Name
     */
    resourceGroupName: pulumi.Input<string>;
}
