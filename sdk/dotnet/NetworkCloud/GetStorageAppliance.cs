// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud
{
    public static class GetStorageAppliance
    {
        /// <summary>
        /// Get properties of the provided storage appliance.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetStorageApplianceResult> InvokeAsync(GetStorageApplianceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageApplianceResult>("azure-native:networkcloud:getStorageAppliance", args ?? new GetStorageApplianceArgs(), options.WithDefaults());

        /// <summary>
        /// Get properties of the provided storage appliance.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetStorageApplianceResult> Invoke(GetStorageApplianceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageApplianceResult>("azure-native:networkcloud:getStorageAppliance", args ?? new GetStorageApplianceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get properties of the provided storage appliance.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetStorageApplianceResult> Invoke(GetStorageApplianceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageApplianceResult>("azure-native:networkcloud:getStorageAppliance", args ?? new GetStorageApplianceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageApplianceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the storage appliance.
        /// </summary>
        [Input("storageApplianceName", required: true)]
        public string StorageApplianceName { get; set; } = null!;

        public GetStorageApplianceArgs()
        {
        }
        public static new GetStorageApplianceArgs Empty => new GetStorageApplianceArgs();
    }

    public sealed class GetStorageApplianceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the storage appliance.
        /// </summary>
        [Input("storageApplianceName", required: true)]
        public Input<string> StorageApplianceName { get; set; } = null!;

        public GetStorageApplianceInvokeArgs()
        {
        }
        public static new GetStorageApplianceInvokeArgs Empty => new GetStorageApplianceInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageApplianceResult
    {
        /// <summary>
        /// The credentials of the administrative interface on this storage appliance.
        /// </summary>
        public readonly Outputs.AdministrativeCredentialsResponse AdministratorCredentials;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The total capacity of the storage appliance. Measured in GiB.
        /// </summary>
        public readonly double Capacity;
        /// <summary>
        /// The amount of storage consumed.
        /// </summary>
        public readonly double CapacityUsed;
        /// <summary>
        /// The resource ID of the cluster this storage appliance is associated with. Measured in GiB.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The detailed status of the storage appliance.
        /// </summary>
        public readonly string DetailedStatus;
        /// <summary>
        /// The descriptive message about the current detailed status.
        /// </summary>
        public readonly string DetailedStatusMessage;
        /// <summary>
        /// Resource ETag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The extended location of the cluster associated with the resource.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The endpoint for the management interface of the storage appliance.
        /// </summary>
        public readonly string ManagementIpv4Address;
        /// <summary>
        /// The manufacturer of the storage appliance.
        /// </summary>
        public readonly string Manufacturer;
        /// <summary>
        /// The model of the storage appliance.
        /// </summary>
        public readonly string Model;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the storage appliance.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource ID of the rack where this storage appliance resides.
        /// </summary>
        public readonly string RackId;
        /// <summary>
        /// The slot the storage appliance is in the rack based on the BOM configuration.
        /// </summary>
        public readonly double RackSlot;
        /// <summary>
        /// The indicator of whether the storage appliance supports remote vendor management.
        /// </summary>
        public readonly string RemoteVendorManagementFeature;
        /// <summary>
        /// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
        /// </summary>
        public readonly string RemoteVendorManagementStatus;
        /// <summary>
        /// The list of statuses that represent secret rotation activity.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecretRotationStatusResponse> SecretRotationStatus;
        /// <summary>
        /// The serial number for the storage appliance.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The SKU for the storage appliance.
        /// </summary>
        public readonly string StorageApplianceSkuId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The version of the storage appliance.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetStorageApplianceResult(
            Outputs.AdministrativeCredentialsResponse administratorCredentials,

            string azureApiVersion,

            double capacity,

            double capacityUsed,

            string clusterId,

            string detailedStatus,

            string detailedStatusMessage,

            string etag,

            Outputs.ExtendedLocationResponse extendedLocation,

            string id,

            string location,

            string managementIpv4Address,

            string manufacturer,

            string model,

            string name,

            string provisioningState,

            string rackId,

            double rackSlot,

            string remoteVendorManagementFeature,

            string remoteVendorManagementStatus,

            ImmutableArray<Outputs.SecretRotationStatusResponse> secretRotationStatus,

            string serialNumber,

            string storageApplianceSkuId,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string version)
        {
            AdministratorCredentials = administratorCredentials;
            AzureApiVersion = azureApiVersion;
            Capacity = capacity;
            CapacityUsed = capacityUsed;
            ClusterId = clusterId;
            DetailedStatus = detailedStatus;
            DetailedStatusMessage = detailedStatusMessage;
            Etag = etag;
            ExtendedLocation = extendedLocation;
            Id = id;
            Location = location;
            ManagementIpv4Address = managementIpv4Address;
            Manufacturer = manufacturer;
            Model = model;
            Name = name;
            ProvisioningState = provisioningState;
            RackId = rackId;
            RackSlot = rackSlot;
            RemoteVendorManagementFeature = remoteVendorManagementFeature;
            RemoteVendorManagementStatus = remoteVendorManagementStatus;
            SecretRotationStatus = secretRotationStatus;
            SerialNumber = serialNumber;
            StorageApplianceSkuId = storageApplianceSkuId;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
