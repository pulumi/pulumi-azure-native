// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the details of the protected item.
 */
export function getProtectedItem(args: GetProtectedItemArgs, opts?: pulumi.InvokeOptions): Promise<GetProtectedItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:datareplication/v20210216preview:getProtectedItem", {
        "protectedItemName": args.protectedItemName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectedItemArgs {
    /**
     * The protected item name.
     */
    protectedItemName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The vault name.
     */
    vaultName: string;
}

/**
 * Protected item model.
 */
export interface GetProtectedItemResult {
    /**
     * Gets or sets the Id of the resource.
     */
    readonly id: string;
    /**
     * Gets or sets the name of the resource.
     */
    readonly name: string;
    /**
     * Protected item model properties.
     */
    readonly properties: outputs.datareplication.v20210216preview.ProtectedItemModelPropertiesResponse;
    readonly systemData: outputs.datareplication.v20210216preview.ProtectedItemModelResponseSystemData;
    /**
     * Gets or sets the type of the resource.
     */
    readonly type: string;
}
/**
 * Gets the details of the protected item.
 */
export function getProtectedItemOutput(args: GetProtectedItemOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProtectedItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:datareplication/v20210216preview:getProtectedItem", {
        "protectedItemName": args.protectedItemName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectedItemOutputArgs {
    /**
     * The protected item name.
     */
    protectedItemName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The vault name.
     */
    vaultName: pulumi.Input<string>;
}
