// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * List primary and secondary keys for a specific key name
 *
 * Uses Azure REST API version 2023-03-01-preview.
 *
 * Other available API versions: 2017-08-21-preview, 2017-11-15, 2018-01-22, 2020-01-01, 2020-03-01, 2020-09-01-preview, 2021-10-15, 2022-02-05, 2022-12-12, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceprovisioningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listIotDpsResourceKeysForKeyName(args: ListIotDpsResourceKeysForKeyNameArgs, opts?: pulumi.InvokeOptions): Promise<ListIotDpsResourceKeysForKeyNameResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:deviceprovisioningservices:listIotDpsResourceKeysForKeyName", {
        "keyName": args.keyName,
        "provisioningServiceName": args.provisioningServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIotDpsResourceKeysForKeyNameArgs {
    /**
     * Logical key name to get key-values for.
     */
    keyName: string;
    /**
     * Name of the provisioning service.
     */
    provisioningServiceName: string;
    /**
     * The name of the resource group that contains the provisioning service.
     */
    resourceGroupName: string;
}

/**
 * Description of the shared access key.
 */
export interface ListIotDpsResourceKeysForKeyNameResult {
    /**
     * Name of the key.
     */
    readonly keyName: string;
    /**
     * Primary SAS key value.
     */
    readonly primaryKey?: string;
    /**
     * Rights that this key has.
     */
    readonly rights: string;
    /**
     * Secondary SAS key value.
     */
    readonly secondaryKey?: string;
}
/**
 * List primary and secondary keys for a specific key name
 *
 * Uses Azure REST API version 2023-03-01-preview.
 *
 * Other available API versions: 2017-08-21-preview, 2017-11-15, 2018-01-22, 2020-01-01, 2020-03-01, 2020-09-01-preview, 2021-10-15, 2022-02-05, 2022-12-12, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceprovisioningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listIotDpsResourceKeysForKeyNameOutput(args: ListIotDpsResourceKeysForKeyNameOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListIotDpsResourceKeysForKeyNameResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:deviceprovisioningservices:listIotDpsResourceKeysForKeyName", {
        "keyName": args.keyName,
        "provisioningServiceName": args.provisioningServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIotDpsResourceKeysForKeyNameOutputArgs {
    /**
     * Logical key name to get key-values for.
     */
    keyName: pulumi.Input<string>;
    /**
     * Name of the provisioning service.
     */
    provisioningServiceName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the provisioning service.
     */
    resourceGroupName: pulumi.Input<string>;
}
