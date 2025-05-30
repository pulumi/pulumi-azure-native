// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedVMwarevSphere
{
    public static class GetCluster
    {
        /// <summary>
        /// Implements cluster GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure-native:connectedvmwarevsphere:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Implements cluster GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure-native:connectedvmwarevsphere:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements cluster GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure-native:connectedvmwarevsphere:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets the name of the corresponding resource in Kubernetes.
        /// </summary>
        public readonly string CustomResourceName;
        /// <summary>
        /// Gets the datastore ARM ids.
        /// </summary>
        public readonly ImmutableArray<string> DatastoreIds;
        /// <summary>
        /// Gets or sets the extended location.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Gets or sets the Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets the inventory Item ID for the cluster.
        /// </summary>
        public readonly string? InventoryItemId;
        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Gets or sets the location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Gets or sets the vCenter Managed Object name for the cluster.
        /// </summary>
        public readonly string MoName;
        /// <summary>
        /// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the cluster.
        /// </summary>
        public readonly string? MoRefId;
        /// <summary>
        /// Gets or sets the name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets the network ARM ids.
        /// </summary>
        public readonly ImmutableArray<string> NetworkIds;
        /// <summary>
        /// Gets the provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceStatusResponse> Statuses;
        /// <summary>
        /// The system data.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Gets or sets the Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets the max CPU usage across all cores on the cluster in MHz.
        /// </summary>
        public readonly double TotalCpuMHz;
        /// <summary>
        /// Gets the total amount of physical memory on the cluster in GB.
        /// </summary>
        public readonly double TotalMemoryGB;
        /// <summary>
        /// Gets or sets the type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets the used CPU usage across all cores on the cluster in MHz.
        /// </summary>
        public readonly double UsedCpuMHz;
        /// <summary>
        /// Gets the used physical memory on the cluster in GB.
        /// </summary>
        public readonly double UsedMemoryGB;
        /// <summary>
        /// Gets or sets a unique identifier for this resource.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// Gets or sets the ARM Id of the vCenter resource in which this cluster resides.
        /// </summary>
        public readonly string? VCenterId;

        [OutputConstructor]
        private GetClusterResult(
            string azureApiVersion,

            string customResourceName,

            ImmutableArray<string> datastoreIds,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string id,

            string? inventoryItemId,

            string? kind,

            string location,

            string moName,

            string? moRefId,

            string name,

            ImmutableArray<string> networkIds,

            string provisioningState,

            ImmutableArray<Outputs.ResourceStatusResponse> statuses,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            double totalCpuMHz,

            double totalMemoryGB,

            string type,

            double usedCpuMHz,

            double usedMemoryGB,

            string uuid,

            string? vCenterId)
        {
            AzureApiVersion = azureApiVersion;
            CustomResourceName = customResourceName;
            DatastoreIds = datastoreIds;
            ExtendedLocation = extendedLocation;
            Id = id;
            InventoryItemId = inventoryItemId;
            Kind = kind;
            Location = location;
            MoName = moName;
            MoRefId = moRefId;
            Name = name;
            NetworkIds = networkIds;
            ProvisioningState = provisioningState;
            Statuses = statuses;
            SystemData = systemData;
            Tags = tags;
            TotalCpuMHz = totalCpuMHz;
            TotalMemoryGB = totalMemoryGB;
            Type = type;
            UsedCpuMHz = usedCpuMHz;
            UsedMemoryGB = usedMemoryGB;
            Uuid = uuid;
            VCenterId = vCenterId;
        }
    }
}
