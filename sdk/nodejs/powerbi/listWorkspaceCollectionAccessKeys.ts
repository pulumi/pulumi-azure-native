// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
 *
 * Uses Azure REST API version 2016-01-29.
 */
export function listWorkspaceCollectionAccessKeys(args: ListWorkspaceCollectionAccessKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkspaceCollectionAccessKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:powerbi:listWorkspaceCollectionAccessKeys", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceCollectionName": args.workspaceCollectionName,
    }, opts);
}

export interface ListWorkspaceCollectionAccessKeysArgs {
    /**
     * Azure resource group
     */
    resourceGroupName: string;
    /**
     * Power BI Embedded Workspace Collection name
     */
    workspaceCollectionName: string;
}

export interface ListWorkspaceCollectionAccessKeysResult {
    /**
     * Access key 1
     */
    readonly key1?: string;
    /**
     * Access key 2
     */
    readonly key2?: string;
}
/**
 * Retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
 *
 * Uses Azure REST API version 2016-01-29.
 */
export function listWorkspaceCollectionAccessKeysOutput(args: ListWorkspaceCollectionAccessKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWorkspaceCollectionAccessKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:powerbi:listWorkspaceCollectionAccessKeys", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceCollectionName": args.workspaceCollectionName,
    }, opts);
}

export interface ListWorkspaceCollectionAccessKeysOutputArgs {
    /**
     * Azure resource group
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Power BI Embedded Workspace Collection name
     */
    workspaceCollectionName: pulumi.Input<string>;
}
