// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get custom image.
 *
 * Uses Azure REST API version 2018-09-15.
 */
export function getCustomImage(args: GetCustomImageArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devtestlab:getCustomImage", {
        "expand": args.expand,
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCustomImageArgs {
    /**
     * Specify the $expand query. Example: 'properties($select=vm)'
     */
    expand?: string;
    /**
     * The name of the lab.
     */
    labName: string;
    /**
     * The name of the CustomImage
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A custom image.
 */
export interface GetCustomImageResult {
    /**
     * The author of the custom image.
     */
    readonly author?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The creation date of the custom image.
     */
    readonly creationDate: string;
    /**
     * Storage information about the plan related to this custom image
     */
    readonly customImagePlan?: outputs.devtestlab.CustomImagePropertiesFromPlanResponse;
    /**
     * Storage information about the data disks present in the custom image
     */
    readonly dataDiskStorageInfo?: outputs.devtestlab.DataDiskStorageTypeInfoResponse[];
    /**
     * The description of the custom image.
     */
    readonly description?: string;
    /**
     * The identifier of the resource.
     */
    readonly id: string;
    /**
     * Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
     */
    readonly isPlanAuthorized?: boolean;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The Managed Image Id backing the custom image.
     */
    readonly managedImageId?: string;
    /**
     * The Managed Snapshot Id backing the custom image.
     */
    readonly managedSnapshotId?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier: string;
    /**
     * The VHD from which the image is to be created.
     */
    readonly vhd?: outputs.devtestlab.CustomImagePropertiesCustomResponse;
    /**
     * The virtual machine from which the image is to be created.
     */
    readonly vm?: outputs.devtestlab.CustomImagePropertiesFromVmResponse;
}
/**
 * Get custom image.
 *
 * Uses Azure REST API version 2018-09-15.
 */
export function getCustomImageOutput(args: GetCustomImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCustomImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devtestlab:getCustomImage", {
        "expand": args.expand,
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCustomImageOutputArgs {
    /**
     * Specify the $expand query. Example: 'properties($select=vm)'
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    labName: pulumi.Input<string>;
    /**
     * The name of the CustomImage
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
