// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Description for Get function secrets for a function in a web site, or a deployment slot.
 */
export function listWebAppFunctionSecretsSlot(args: ListWebAppFunctionSecretsSlotArgs, opts?: pulumi.InvokeOptions): Promise<ListWebAppFunctionSecretsSlotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:web/v20231201:listWebAppFunctionSecretsSlot", {
        "functionName": args.functionName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface ListWebAppFunctionSecretsSlotArgs {
    /**
     * Function name.
     */
    functionName: string;
    /**
     * Site name.
     */
    name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: string;
    /**
     * Name of the deployment slot.
     */
    slot: string;
}

/**
 * Function secrets.
 */
export interface ListWebAppFunctionSecretsSlotResult {
    /**
     * Secret key.
     */
    readonly key?: string;
    /**
     * Trigger URL.
     */
    readonly triggerUrl?: string;
}
/**
 * Description for Get function secrets for a function in a web site, or a deployment slot.
 */
export function listWebAppFunctionSecretsSlotOutput(args: ListWebAppFunctionSecretsSlotOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWebAppFunctionSecretsSlotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:web/v20231201:listWebAppFunctionSecretsSlot", {
        "functionName": args.functionName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface ListWebAppFunctionSecretsSlotOutputArgs {
    /**
     * Function name.
     */
    functionName: pulumi.Input<string>;
    /**
     * Site name.
     */
    name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the deployment slot.
     */
    slot: pulumi.Input<string>;
}
