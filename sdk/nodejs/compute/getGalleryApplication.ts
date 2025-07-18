// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves information about a gallery Application Definition.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGalleryApplication(args: GetGalleryApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:compute:getGalleryApplication", {
        "galleryApplicationName": args.galleryApplicationName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryApplicationArgs {
    /**
     * The name of the gallery Application Definition to be retrieved.
     */
    galleryApplicationName: string;
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
 * Specifies information about the gallery Application Definition that you want to create or update.
 */
export interface GetGalleryApplicationResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * A list of custom actions that can be performed with all of the Gallery Application Versions within this Gallery Application.
     */
    readonly customActions?: outputs.compute.GalleryApplicationCustomActionResponse[];
    /**
     * The description of this gallery Application Definition resource. This property is updatable.
     */
    readonly description?: string;
    /**
     * The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    readonly endOfLifeDate?: string;
    /**
     * The Eula agreement for the gallery Application Definition.
     */
    readonly eula?: string;
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
     * The privacy statement uri.
     */
    readonly privacyStatementUri?: string;
    /**
     * The release note uri.
     */
    readonly releaseNoteUri?: string;
    /**
     * This property allows you to specify the supported type of the OS that application is built for. Possible values are: **Windows,** **Linux.**
     */
    readonly supportedOSType: string;
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
 * Retrieves information about a gallery Application Definition.
 *
 * Uses Azure REST API version 2024-03-03.
 *
 * Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGalleryApplicationOutput(args: GetGalleryApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGalleryApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:compute:getGalleryApplication", {
        "galleryApplicationName": args.galleryApplicationName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryApplicationOutputArgs {
    /**
     * The name of the gallery Application Definition to be retrieved.
     */
    galleryApplicationName: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery.
     */
    galleryName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
