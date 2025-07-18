// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves information about a Shared Image Gallery.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGallery(args: GetGalleryArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:compute:getGallery", {
        "expand": args.expand,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
        "select": args.select,
    }, opts);
}

export interface GetGalleryArgs {
    /**
     * The expand query option to apply on the operation.
     */
    expand?: string;
    /**
     * The name of the Shared Image Gallery.
     */
    galleryName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The select expression to apply on the operation.
     */
    select?: string;
}

/**
 * Specifies information about the Shared Image Gallery that you want to create or update.
 */
export interface GetGalleryResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The description of this Shared Image Gallery resource. This property is updatable.
     */
    readonly description?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Describes the gallery unique name.
     */
    readonly identifier?: outputs.compute.GalleryIdentifierResponse;
    /**
     * The identity of the gallery, if configured.
     */
    readonly identity?: outputs.compute.GalleryIdentityResponse;
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
     * Profile for gallery sharing to subscription or tenant
     */
    readonly sharingProfile?: outputs.compute.SharingProfileResponse;
    /**
     * Sharing status of current gallery.
     */
    readonly sharingStatus: outputs.compute.SharingStatusResponse;
    /**
     * Contains information about the soft deletion policy of the gallery.
     */
    readonly softDeletePolicy?: outputs.compute.SoftDeletePolicyResponse;
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
 * Retrieves information about a Shared Image Gallery.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGalleryOutput(args: GetGalleryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGalleryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:compute:getGallery", {
        "expand": args.expand,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
        "select": args.select,
    }, opts);
}

export interface GetGalleryOutputArgs {
    /**
     * The expand query option to apply on the operation.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery.
     */
    galleryName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The select expression to apply on the operation.
     */
    select?: pulumi.Input<string>;
}
