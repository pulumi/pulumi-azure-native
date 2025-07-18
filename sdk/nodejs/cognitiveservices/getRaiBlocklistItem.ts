// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the specified custom blocklist Item associated with the custom blocklist.
 *
 * Uses Azure REST API version 2024-10-01.
 *
 * Other available API versions: 2023-10-01-preview, 2024-04-01-preview, 2024-06-01-preview, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cognitiveservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getRaiBlocklistItem(args: GetRaiBlocklistItemArgs, opts?: pulumi.InvokeOptions): Promise<GetRaiBlocklistItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cognitiveservices:getRaiBlocklistItem", {
        "accountName": args.accountName,
        "raiBlocklistItemName": args.raiBlocklistItemName,
        "raiBlocklistName": args.raiBlocklistName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRaiBlocklistItemArgs {
    /**
     * The name of Cognitive Services account.
     */
    accountName: string;
    /**
     * The name of the RaiBlocklist Item associated with the custom blocklist
     */
    raiBlocklistItemName: string;
    /**
     * The name of the RaiBlocklist associated with the Cognitive Services Account
     */
    raiBlocklistName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Cognitive Services RaiBlocklist Item.
 */
export interface GetRaiBlocklistItemResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Resource Etag.
     */
    readonly etag: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Properties of Cognitive Services RaiBlocklist Item.
     */
    readonly properties: outputs.cognitiveservices.RaiBlocklistItemPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.cognitiveservices.SystemDataResponse;
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
 * Gets the specified custom blocklist Item associated with the custom blocklist.
 *
 * Uses Azure REST API version 2024-10-01.
 *
 * Other available API versions: 2023-10-01-preview, 2024-04-01-preview, 2024-06-01-preview, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cognitiveservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getRaiBlocklistItemOutput(args: GetRaiBlocklistItemOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRaiBlocklistItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cognitiveservices:getRaiBlocklistItem", {
        "accountName": args.accountName,
        "raiBlocklistItemName": args.raiBlocklistItemName,
        "raiBlocklistName": args.raiBlocklistName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRaiBlocklistItemOutputArgs {
    /**
     * The name of Cognitive Services account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the RaiBlocklist Item associated with the custom blocklist
     */
    raiBlocklistItemName: pulumi.Input<string>;
    /**
     * The name of the RaiBlocklist associated with the Cognitive Services Account
     */
    raiBlocklistName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
