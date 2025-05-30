// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets the operation link for the tag.
 *
 * Uses Azure REST API version 2022-09-01-preview.
 *
 * Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getTagOperationLink(args: GetTagOperationLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetTagOperationLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:apimanagement:getTagOperationLink", {
        "operationLinkId": args.operationLinkId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagOperationLinkArgs {
    /**
     * Tag-operation link identifier. Must be unique in the current API Management service instance.
     */
    operationLinkId: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    serviceName: string;
    /**
     * Tag identifier. Must be unique in the current API Management service instance.
     */
    tagId: string;
}

/**
 * Tag-operation link details.
 */
export interface GetTagOperationLinkResult {
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
     * Full resource Id of an API operation.
     */
    readonly operationId: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets the operation link for the tag.
 *
 * Uses Azure REST API version 2022-09-01-preview.
 *
 * Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getTagOperationLinkOutput(args: GetTagOperationLinkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTagOperationLinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:apimanagement:getTagOperationLink", {
        "operationLinkId": args.operationLinkId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagOperationLinkOutputArgs {
    /**
     * Tag-operation link identifier. Must be unique in the current API Management service instance.
     */
    operationLinkId: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Tag identifier. Must be unique in the current API Management service instance.
     */
    tagId: pulumi.Input<string>;
}
