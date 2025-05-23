// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a specific addon by name.
 *
 * Uses Azure REST API version 2023-07-01.
 */
export function getArcAddon(args: GetArcAddonArgs, opts?: pulumi.InvokeOptions): Promise<GetArcAddonResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:databoxedge:getArcAddon", {
        "addonName": args.addonName,
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
        "roleName": args.roleName,
    }, opts);
}

export interface GetArcAddonArgs {
    /**
     * The addon name.
     */
    addonName: string;
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The role name.
     */
    roleName: string;
}

/**
 * Arc Addon.
 */
export interface GetArcAddonResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Host OS supported by the Arc addon.
     */
    readonly hostPlatform: string;
    /**
     * Platform where the runtime is hosted.
     */
    readonly hostPlatformType: string;
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * Addon type.
     * Expected value is 'ArcForKubernetes'.
     */
    readonly kind: "ArcForKubernetes";
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * Addon Provisioning State
     */
    readonly provisioningState: string;
    /**
     * Arc resource group name
     */
    readonly resourceGroupName: string;
    /**
     * Arc resource location
     */
    readonly resourceLocation: string;
    /**
     * Arc resource Name
     */
    readonly resourceName: string;
    /**
     * Arc resource subscription Id
     */
    readonly subscriptionId: string;
    /**
     * Metadata pertaining to creation and last modification of Addon
     */
    readonly systemData: outputs.databoxedge.SystemDataResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
    /**
     * Arc resource version
     */
    readonly version: string;
}
/**
 * Gets a specific addon by name.
 *
 * Uses Azure REST API version 2023-07-01.
 */
export function getArcAddonOutput(args: GetArcAddonOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetArcAddonResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:databoxedge:getArcAddon", {
        "addonName": args.addonName,
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
        "roleName": args.roleName,
    }, opts);
}

export interface GetArcAddonOutputArgs {
    /**
     * The addon name.
     */
    addonName: pulumi.Input<string>;
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The role name.
     */
    roleName: pulumi.Input<string>;
}
