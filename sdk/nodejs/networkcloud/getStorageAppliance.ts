// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get properties of the provided storage appliance.
 *
 * Uses Azure REST API version 2025-02-01.
 *
 * Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getStorageAppliance(args: GetStorageApplianceArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageApplianceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:networkcloud:getStorageAppliance", {
        "resourceGroupName": args.resourceGroupName,
        "storageApplianceName": args.storageApplianceName,
    }, opts);
}

export interface GetStorageApplianceArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the storage appliance.
     */
    storageApplianceName: string;
}

export interface GetStorageApplianceResult {
    /**
     * The credentials of the administrative interface on this storage appliance.
     */
    readonly administratorCredentials: outputs.networkcloud.AdministrativeCredentialsResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The total capacity of the storage appliance. Measured in GiB.
     */
    readonly capacity: number;
    /**
     * The amount of storage consumed.
     */
    readonly capacityUsed: number;
    /**
     * The resource ID of the cluster this storage appliance is associated with. Measured in GiB.
     */
    readonly clusterId: string;
    /**
     * The detailed status of the storage appliance.
     */
    readonly detailedStatus: string;
    /**
     * The descriptive message about the current detailed status.
     */
    readonly detailedStatusMessage: string;
    /**
     * Resource ETag.
     */
    readonly etag: string;
    /**
     * The extended location of the cluster associated with the resource.
     */
    readonly extendedLocation: outputs.networkcloud.ExtendedLocationResponse;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The endpoint for the management interface of the storage appliance.
     */
    readonly managementIpv4Address: string;
    /**
     * The manufacturer of the storage appliance.
     */
    readonly manufacturer: string;
    /**
     * The model of the storage appliance.
     */
    readonly model: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state of the storage appliance.
     */
    readonly provisioningState: string;
    /**
     * The resource ID of the rack where this storage appliance resides.
     */
    readonly rackId: string;
    /**
     * The slot the storage appliance is in the rack based on the BOM configuration.
     */
    readonly rackSlot: number;
    /**
     * The indicator of whether the storage appliance supports remote vendor management.
     */
    readonly remoteVendorManagementFeature: string;
    /**
     * The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
     */
    readonly remoteVendorManagementStatus: string;
    /**
     * The list of statuses that represent secret rotation activity.
     */
    readonly secretRotationStatus: outputs.networkcloud.SecretRotationStatusResponse[];
    /**
     * The serial number for the storage appliance.
     */
    readonly serialNumber: string;
    /**
     * The SKU for the storage appliance.
     */
    readonly storageApplianceSkuId: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.networkcloud.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The version of the storage appliance.
     */
    readonly version: string;
}
/**
 * Get properties of the provided storage appliance.
 *
 * Uses Azure REST API version 2025-02-01.
 *
 * Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getStorageApplianceOutput(args: GetStorageApplianceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStorageApplianceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:networkcloud:getStorageAppliance", {
        "resourceGroupName": args.resourceGroupName,
        "storageApplianceName": args.storageApplianceName,
    }, opts);
}

export interface GetStorageApplianceOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the storage appliance.
     */
    storageApplianceName: pulumi.Input<string>;
}
