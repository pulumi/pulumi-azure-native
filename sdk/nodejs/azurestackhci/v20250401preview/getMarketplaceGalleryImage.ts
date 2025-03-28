// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a marketplace gallery image
 */
export function getMarketplaceGalleryImage(args: GetMarketplaceGalleryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetMarketplaceGalleryImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurestackhci/v20250401preview:getMarketplaceGalleryImage", {
        "marketplaceGalleryImageName": args.marketplaceGalleryImageName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMarketplaceGalleryImageArgs {
    /**
     * Name of the marketplace gallery image
     */
    marketplaceGalleryImageName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * The marketplace gallery image resource definition.
 */
export interface GetMarketplaceGalleryImageResult {
    /**
     * Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
     */
    readonly cloudInitDataSource?: string;
    /**
     * Storage ContainerID of the storage container to be used for marketplace gallery image
     */
    readonly containerId?: string;
    /**
     * The extendedLocation of the resource.
     */
    readonly extendedLocation?: outputs.azurestackhci.v20250401preview.ExtendedLocationResponse;
    /**
     * The hypervisor generation of the Virtual Machine [V1, V2]
     */
    readonly hyperVGeneration?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * This is the gallery image definition identifier.
     */
    readonly identifier?: outputs.azurestackhci.v20250401preview.GalleryImageIdentifierResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Operating system type that the gallery image uses [Windows, Linux]
     */
    readonly osType: string;
    /**
     * Provisioning state of the marketplace gallery image.
     */
    readonly provisioningState: string;
    /**
     * The observed state of marketplace gallery images
     */
    readonly status: outputs.azurestackhci.v20250401preview.MarketplaceGalleryImageStatusResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.azurestackhci.v20250401preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Specifies information about the gallery image version that you want to create or update.
     */
    readonly version?: outputs.azurestackhci.v20250401preview.GalleryImageVersionResponse;
}
/**
 * Gets a marketplace gallery image
 */
export function getMarketplaceGalleryImageOutput(args: GetMarketplaceGalleryImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMarketplaceGalleryImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurestackhci/v20250401preview:getMarketplaceGalleryImage", {
        "marketplaceGalleryImageName": args.marketplaceGalleryImageName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMarketplaceGalleryImageOutputArgs {
    /**
     * Name of the marketplace gallery image
     */
    marketplaceGalleryImageName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
