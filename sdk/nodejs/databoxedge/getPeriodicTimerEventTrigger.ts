// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a specific trigger by name.
 *
 * Uses Azure REST API version 2023-07-01.
 */
export function getPeriodicTimerEventTrigger(args: GetPeriodicTimerEventTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetPeriodicTimerEventTriggerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:databoxedge:getPeriodicTimerEventTrigger", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPeriodicTimerEventTriggerArgs {
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The trigger name.
     */
    name: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * Trigger details.
 */
export interface GetPeriodicTimerEventTriggerResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
     */
    readonly customContextTag?: string;
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * Trigger Kind.
     * Expected value is 'PeriodicTimerEvent'.
     */
    readonly kind: "PeriodicTimerEvent";
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * Role Sink information.
     */
    readonly sinkInfo: outputs.databoxedge.RoleSinkInfoResponse;
    /**
     * Periodic timer details.
     */
    readonly sourceInfo: outputs.databoxedge.PeriodicTimerSourceInfoResponse;
    /**
     * Metadata pertaining to creation and last modification of Trigger
     */
    readonly systemData: outputs.databoxedge.SystemDataResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}
/**
 * Get a specific trigger by name.
 *
 * Uses Azure REST API version 2023-07-01.
 */
export function getPeriodicTimerEventTriggerOutput(args: GetPeriodicTimerEventTriggerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPeriodicTimerEventTriggerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:databoxedge:getPeriodicTimerEventTrigger", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPeriodicTimerEventTriggerOutputArgs {
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * The trigger name.
     */
    name: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
}
