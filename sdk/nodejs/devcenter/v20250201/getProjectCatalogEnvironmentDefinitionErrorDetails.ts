// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets Environment Definition error details
 */
export function getProjectCatalogEnvironmentDefinitionErrorDetails(args: GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectCatalogEnvironmentDefinitionErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter/v20250201:getProjectCatalogEnvironmentDefinitionErrorDetails", {
        "catalogName": args.catalogName,
        "environmentDefinitionName": args.environmentDefinitionName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: string;
    /**
     * The name of the Environment Definition.
     */
    environmentDefinitionName: string;
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
 * List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
 */
export interface GetProjectCatalogEnvironmentDefinitionErrorDetailsResult {
    /**
     * Errors associated with resources synchronized from the catalog.
     */
    readonly errors: outputs.devcenter.v20250201.CatalogErrorDetailsResponse[];
}
/**
 * Gets Environment Definition error details
 */
export function getProjectCatalogEnvironmentDefinitionErrorDetailsOutput(args: GetProjectCatalogEnvironmentDefinitionErrorDetailsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectCatalogEnvironmentDefinitionErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter/v20250201:getProjectCatalogEnvironmentDefinitionErrorDetails", {
        "catalogName": args.catalogName,
        "environmentDefinitionName": args.environmentDefinitionName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectCatalogEnvironmentDefinitionErrorDetailsOutputArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: pulumi.Input<string>;
    /**
     * The name of the Environment Definition.
     */
    environmentDefinitionName: pulumi.Input<string>;
    /**
     * The name of the project.
     */
    projectName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
