// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud
{
    public static class GetClusterManager
    {
        /// <summary>
        /// Get the properties of the provided cluster manager.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetClusterManagerResult> InvokeAsync(GetClusterManagerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterManagerResult>("azure-native:networkcloud:getClusterManager", args ?? new GetClusterManagerArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of the provided cluster manager.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterManagerResult> Invoke(GetClusterManagerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterManagerResult>("azure-native:networkcloud:getClusterManager", args ?? new GetClusterManagerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of the provided cluster manager.
        /// 
        /// Uses Azure REST API version 2025-02-01.
        /// 
        /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterManagerResult> Invoke(GetClusterManagerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterManagerResult>("azure-native:networkcloud:getClusterManager", args ?? new GetClusterManagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterManagerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster manager.
        /// </summary>
        [Input("clusterManagerName", required: true)]
        public string ClusterManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterManagerArgs()
        {
        }
        public static new GetClusterManagerArgs Empty => new GetClusterManagerArgs();
    }

    public sealed class GetClusterManagerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster manager.
        /// </summary>
        [Input("clusterManagerName", required: true)]
        public Input<string> ClusterManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetClusterManagerInvokeArgs()
        {
        }
        public static new GetClusterManagerInvokeArgs Empty => new GetClusterManagerInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterManagerResult
    {
        /// <summary>
        /// The resource ID of the Log Analytics workspace that is used for the logs collection.
        /// </summary>
        public readonly string? AnalyticsWorkspaceId;
        /// <summary>
        /// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterAvailableVersionResponse> ClusterVersions;
        /// <summary>
        /// The detailed status that provides additional information about the cluster manager.
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
        /// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
        /// </summary>
        public readonly string FabricControllerId;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the cluster manager.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The configuration of the managed resource group associated with the resource.
        /// </summary>
        public readonly Outputs.ManagedResourceGroupConfigurationResponse? ManagedResourceGroupConfiguration;
        /// <summary>
        /// The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ManagerExtendedLocation;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the cluster manager.
        /// </summary>
        public readonly string ProvisioningState;
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
        /// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
        /// </summary>
        public readonly string? VmSize;

        [OutputConstructor]
        private GetClusterManagerResult(
            string? analyticsWorkspaceId,

            ImmutableArray<string> availabilityZones,

            string azureApiVersion,

            ImmutableArray<Outputs.ClusterAvailableVersionResponse> clusterVersions,

            string detailedStatus,

            string detailedStatusMessage,

            string etag,

            string fabricControllerId,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            Outputs.ManagedResourceGroupConfigurationResponse? managedResourceGroupConfiguration,

            Outputs.ExtendedLocationResponse managerExtendedLocation,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? vmSize)
        {
            AnalyticsWorkspaceId = analyticsWorkspaceId;
            AvailabilityZones = availabilityZones;
            AzureApiVersion = azureApiVersion;
            ClusterVersions = clusterVersions;
            DetailedStatus = detailedStatus;
            DetailedStatusMessage = detailedStatusMessage;
            Etag = etag;
            FabricControllerId = fabricControllerId;
            Id = id;
            Identity = identity;
            Location = location;
            ManagedResourceGroupConfiguration = managedResourceGroupConfiguration;
            ManagerExtendedLocation = managerExtendedLocation;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VmSize = vmSize;
        }
    }
}
