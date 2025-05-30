// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * List springbootservers resource.
 *
 * Uses Azure REST API version 2024-04-01-preview.
 *
 * Other available API versions: 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native offazurespringboot [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSpringbootserver(args: GetSpringbootserverArgs, opts?: pulumi.InvokeOptions): Promise<GetSpringbootserverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:offazurespringboot:getSpringbootserver", {
        "resourceGroupName": args.resourceGroupName,
        "siteName": args.siteName,
        "springbootserversName": args.springbootserversName,
    }, opts);
}

export interface GetSpringbootserverArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The springbootsites name.
     */
    siteName: string;
    /**
     * The springbootservers name.
     */
    springbootserversName: string;
}

/**
 * The springbootservers envelope resource definition.
 */
export interface GetSpringbootserverResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The springbootservers resource definition.
     */
    readonly properties: outputs.offazurespringboot.SpringbootserversPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.offazurespringboot.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * List springbootservers resource.
 *
 * Uses Azure REST API version 2024-04-01-preview.
 *
 * Other available API versions: 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native offazurespringboot [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSpringbootserverOutput(args: GetSpringbootserverOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSpringbootserverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:offazurespringboot:getSpringbootserver", {
        "resourceGroupName": args.resourceGroupName,
        "siteName": args.siteName,
        "springbootserversName": args.springbootserversName,
    }, opts);
}

export interface GetSpringbootserverOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The springbootsites name.
     */
    siteName: pulumi.Input<string>;
    /**
     * The springbootservers name.
     */
    springbootserversName: pulumi.Input<string>;
}
