// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * DC Access code in the case of Self Managed Shipping.
 */
export function listOrderDCAccessCode(args: ListOrderDCAccessCodeArgs, opts?: pulumi.InvokeOptions): Promise<ListOrderDCAccessCodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:databoxedge/v20220301:listOrderDCAccessCode", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListOrderDCAccessCodeArgs {
    /**
     * The device name
     */
    deviceName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * DC Access code in the case of Self Managed Shipping.
 */
export interface ListOrderDCAccessCodeResult {
    /**
     * DCAccess Code for the Self Managed shipment.
     */
    readonly authCode?: string;
}
/**
 * DC Access code in the case of Self Managed Shipping.
 */
export function listOrderDCAccessCodeOutput(args: ListOrderDCAccessCodeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListOrderDCAccessCodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:databoxedge/v20220301:listOrderDCAccessCode", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListOrderDCAccessCodeOutputArgs {
    /**
     * The device name
     */
    deviceName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
}
