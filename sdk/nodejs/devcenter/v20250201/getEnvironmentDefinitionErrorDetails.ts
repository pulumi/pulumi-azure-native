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
export function getEnvironmentDefinitionErrorDetails(args: GetEnvironmentDefinitionErrorDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentDefinitionErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter/v20250201:getEnvironmentDefinitionErrorDetails", {
        "catalogName": args.catalogName,
        "devCenterName": args.devCenterName,
        "environmentDefinitionName": args.environmentDefinitionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnvironmentDefinitionErrorDetailsArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: string;
    /**
     * The name of the devcenter.
     */
    devCenterName: string;
    /**
     * The name of the Environment Definition.
     */
    environmentDefinitionName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
 */
export interface GetEnvironmentDefinitionErrorDetailsResult {
    /**
     * Errors associated with resources synchronized from the catalog.
     */
    readonly errors: outputs.devcenter.v20250201.CatalogErrorDetailsResponse[];
}
/**
 * Gets Environment Definition error details
 */
export function getEnvironmentDefinitionErrorDetailsOutput(args: GetEnvironmentDefinitionErrorDetailsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEnvironmentDefinitionErrorDetailsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter/v20250201:getEnvironmentDefinitionErrorDetails", {
        "catalogName": args.catalogName,
        "devCenterName": args.devCenterName,
        "environmentDefinitionName": args.environmentDefinitionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnvironmentDefinitionErrorDetailsOutputArgs {
    /**
     * The name of the Catalog.
     */
    catalogName: pulumi.Input<string>;
    /**
     * The name of the devcenter.
     */
    devCenterName: pulumi.Input<string>;
    /**
     * The name of the Environment Definition.
     */
    environmentDefinitionName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
