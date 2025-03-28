// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * List available models from all connections.
 *
 * Uses Azure REST API version 2024-04-01-preview.
 */
export function listWorkspaceConnectionModels(args: ListWorkspaceConnectionModelsArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkspaceConnectionModelsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:machinelearningservices:listWorkspaceConnectionModels", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListWorkspaceConnectionModelsArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: string;
}

export interface ListWorkspaceConnectionModelsResult {
    /**
     * The link to the next page constructed using the continuationToken.  If null, there are no additional pages.
     */
    readonly nextLink?: string;
    /**
     * List of models.
     */
    readonly value?: outputs.machinelearningservices.AccountModelResponse[];
}
/**
 * List available models from all connections.
 *
 * Uses Azure REST API version 2024-04-01-preview.
 */
export function listWorkspaceConnectionModelsOutput(args: ListWorkspaceConnectionModelsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWorkspaceConnectionModelsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:machinelearningservices:listWorkspaceConnectionModels", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListWorkspaceConnectionModelsOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: pulumi.Input<string>;
}
