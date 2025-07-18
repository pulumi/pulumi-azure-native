// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves information about a gallery Application Version.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGalleryApplicationVersion(args: GetGalleryApplicationVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryApplicationVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:compute:getGalleryApplicationVersion", {
        "expand": args.expand,
        "galleryApplicationName": args.galleryApplicationName,
        "galleryApplicationVersionName": args.galleryApplicationVersionName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryApplicationVersionArgs {
    /**
     * The expand expression to apply on the operation.
     */
    expand?: string;
    /**
     * The name of the gallery Application Definition to be retrieved.
     */
    galleryApplicationName: string;
    /**
     * The name of the gallery Application Version to be retrieved.
     */
    galleryApplicationVersionName: string;
    /**
     * The name of the Shared Image Gallery.
     */
    galleryName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Specifies information about the gallery Application Version that you want to create or update.
 */
export interface GetGalleryApplicationVersionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * The publishing profile of a gallery image version.
     */
    readonly publishingProfile: outputs.compute.GalleryApplicationVersionPublishingProfileResponse;
    /**
     * This is the replication status of the gallery image version.
     */
    readonly replicationStatus: outputs.compute.ReplicationStatusResponse;
    /**
     * The safety profile of the Gallery Application Version.
     */
    readonly safetyProfile?: outputs.compute.GalleryApplicationVersionSafetyProfileResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.compute.SystemDataResponse;
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
 * Retrieves information about a gallery Application Version.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGalleryApplicationVersionOutput(args: GetGalleryApplicationVersionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGalleryApplicationVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:compute:getGalleryApplicationVersion", {
        "expand": args.expand,
        "galleryApplicationName": args.galleryApplicationName,
        "galleryApplicationVersionName": args.galleryApplicationVersionName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryApplicationVersionOutputArgs {
    /**
     * The expand expression to apply on the operation.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the gallery Application Definition to be retrieved.
     */
    galleryApplicationName: pulumi.Input<string>;
    /**
     * The name of the gallery Application Version to be retrieved.
     */
    galleryApplicationVersionName: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery.
     */
    galleryName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
