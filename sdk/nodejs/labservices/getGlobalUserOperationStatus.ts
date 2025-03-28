// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets the status of long running operation
 *
 * Uses Azure REST API version 2018-10-15.
 */
export function getGlobalUserOperationStatus(args: GetGlobalUserOperationStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalUserOperationStatusResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:labservices:getGlobalUserOperationStatus", {
        "operationUrl": args.operationUrl,
        "userName": args.userName,
    }, opts);
}

export interface GetGlobalUserOperationStatusArgs {
    /**
     * The operation url of long running operation
     */
    operationUrl: string;
    /**
     * The name of the user.
     */
    userName: string;
}

/**
 * Status Details of the long running operation for an environment
 */
export interface GetGlobalUserOperationStatusResult {
    /**
     * status of the long running operation for an environment
     */
    readonly status: string;
}
/**
 * Gets the status of long running operation
 *
 * Uses Azure REST API version 2018-10-15.
 */
export function getGlobalUserOperationStatusOutput(args: GetGlobalUserOperationStatusOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGlobalUserOperationStatusResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:labservices:getGlobalUserOperationStatus", {
        "operationUrl": args.operationUrl,
        "userName": args.userName,
    }, opts);
}

export interface GetGlobalUserOperationStatusOutputArgs {
    /**
     * The operation url of long running operation
     */
    operationUrl: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    userName: pulumi.Input<string>;
}
