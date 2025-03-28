// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the Defender for Storage settings for the specified storage account.
 *
 * Uses Azure REST API version 2022-12-01-preview.
 *
 * Other available API versions: 2024-10-01-preview.
 */
export function getDefenderForStorage(args: GetDefenderForStorageArgs, opts?: pulumi.InvokeOptions): Promise<GetDefenderForStorageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:security:getDefenderForStorage", {
        "resourceId": args.resourceId,
        "settingName": args.settingName,
    }, opts);
}

export interface GetDefenderForStorageArgs {
    /**
     * The identifier of the resource.
     */
    resourceId: string;
    /**
     * Defender for Storage setting name.
     */
    settingName: string;
}

/**
 * The Defender for Storage resource.
 */
export interface GetDefenderForStorageResult {
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Defender for Storage resource properties.
     */
    readonly properties: outputs.security.DefenderForStorageSettingPropertiesResponse;
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Gets the Defender for Storage settings for the specified storage account.
 *
 * Uses Azure REST API version 2022-12-01-preview.
 *
 * Other available API versions: 2024-10-01-preview.
 */
export function getDefenderForStorageOutput(args: GetDefenderForStorageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDefenderForStorageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:security:getDefenderForStorage", {
        "resourceId": args.resourceId,
        "settingName": args.settingName,
    }, opts);
}

export interface GetDefenderForStorageOutputArgs {
    /**
     * The identifier of the resource.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Defender for Storage setting name.
     */
    settingName: pulumi.Input<string>;
}
