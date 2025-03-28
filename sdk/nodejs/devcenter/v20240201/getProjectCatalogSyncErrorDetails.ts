// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets project catalog synchronization error details
 */
export function getProjectCatalogSyncErrorDetails(args: GetProjectCatalogSyncErrorDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectCatalogSyncErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter/v20240201:getProjectCatalogSyncErrorDetails", {
        "catalogName": args.catalogName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectCatalogSyncErrorDetailsArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: string;
    /**
     * The name of the project.
     */
    projectName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Synchronization error details.
 */
export interface GetProjectCatalogSyncErrorDetailsResult {
    /**
     * Catalog items that have conflicting names.
     */
    readonly conflicts: outputs.devcenter.v20240201.CatalogConflictErrorResponse[];
    /**
     * Errors that occured during synchronization.
     */
    readonly errors: outputs.devcenter.v20240201.CatalogSyncErrorResponse[];
    /**
     * Error information for the overall synchronization operation.
     */
    readonly operationError: outputs.devcenter.v20240201.CatalogErrorDetailsResponse;
}
/**
 * Gets project catalog synchronization error details
 */
export function getProjectCatalogSyncErrorDetailsOutput(args: GetProjectCatalogSyncErrorDetailsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectCatalogSyncErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter/v20240201:getProjectCatalogSyncErrorDetails", {
        "catalogName": args.catalogName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectCatalogSyncErrorDetailsOutputArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: pulumi.Input<string>;
    /**
     * The name of the project.
     */
    projectName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
