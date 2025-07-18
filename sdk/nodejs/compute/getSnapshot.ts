// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets information about a snapshot.
 *
 * Uses Azure REST API version 2024-03-02.
 *
 * Other available API versions: 2022-07-02, 2023-01-02, 2023-04-02, 2023-10-02, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:compute:getSnapshot", {
        "resourceGroupName": args.resourceGroupName,
        "snapshotName": args.snapshotName,
    }, opts);
}

export interface GetSnapshotArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
     */
    snapshotName: string;
}

/**
 * Snapshot resource.
 */
export interface GetSnapshotResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Percentage complete for the background copy when a resource is created via the CopyStart operation.
     */
    readonly completionPercent?: number;
    /**
     * Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
     */
    readonly copyCompletionError?: outputs.compute.CopyCompletionErrorResponse;
    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    readonly creationData: outputs.compute.CreationDataResponse;
    /**
     * Additional authentication requirements when exporting or uploading to a disk or snapshot.
     */
    readonly dataAccessAuthMode?: string;
    /**
     * ARM id of the DiskAccess resource for using private endpoints on disks.
     */
    readonly diskAccessId?: string;
    /**
     * The size of the disk in bytes. This field is read only.
     */
    readonly diskSizeBytes: number;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    readonly diskSizeGB?: number;
    /**
     * The state of the snapshot.
     */
    readonly diskState: string;
    /**
     * Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
     */
    readonly encryption?: outputs.compute.EncryptionResponse;
    /**
     * Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    readonly encryptionSettingsCollection?: outputs.compute.EncryptionSettingsCollectionResponse;
    /**
     * The extended location where the snapshot will be created. Extended location cannot be changed.
     */
    readonly extendedLocation?: outputs.compute.ExtendedLocationResponse;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    readonly hyperVGeneration?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
     */
    readonly incremental?: boolean;
    /**
     * Incremental snapshots for a disk share an incremental snapshot family id. The Get Page Range Diff API can only be called on incremental snapshots with the same family id.
     */
    readonly incrementalSnapshotFamilyId: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Unused. Always Null.
     */
    readonly managedBy: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Policy for accessing the disk via network.
     */
    readonly networkAccessPolicy?: string;
    /**
     * The Operating System type.
     */
    readonly osType?: string;
    /**
     * The disk provisioning state.
     */
    readonly provisioningState: string;
    /**
     * Policy for controlling export on the disk.
     */
    readonly publicNetworkAccess?: string;
    /**
     * Purchase plan information for the image from which the source disk for the snapshot was originally created.
     */
    readonly purchasePlan?: outputs.compute.DiskPurchasePlanResponse;
    /**
     * Contains the security related information for the resource.
     */
    readonly securityProfile?: outputs.compute.DiskSecurityProfileResponse;
    /**
     * The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
     */
    readonly sku?: outputs.compute.SnapshotSkuResponse;
    /**
     * List of supported capabilities for the image from which the source disk from the snapshot was originally created.
     */
    readonly supportedCapabilities?: outputs.compute.SupportedCapabilitiesResponse;
    /**
     * Indicates the OS on a snapshot supports hibernation.
     */
    readonly supportsHibernation?: boolean;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.compute.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The time when the snapshot was created.
     */
    readonly timeCreated: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Unique Guid identifying the resource.
     */
    readonly uniqueId: string;
}
/**
 * Gets information about a snapshot.
 *
 * Uses Azure REST API version 2024-03-02.
 *
 * Other available API versions: 2022-07-02, 2023-01-02, 2023-04-02, 2023-10-02, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSnapshotOutput(args: GetSnapshotOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSnapshotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:compute:getSnapshot", {
        "resourceGroupName": args.resourceGroupName,
        "snapshotName": args.snapshotName,
    }, opts);
}

export interface GetSnapshotOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
     */
    snapshotName: pulumi.Input<string>;
}
