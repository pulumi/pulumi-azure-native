// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedVMwarevSphere
{
    public static class GetHost
    {
        /// <summary>
        /// Implements host GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetHostResult> InvokeAsync(GetHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("azure-native:connectedvmwarevsphere:getHost", args ?? new GetHostArgs(), options.WithDefaults());

        /// <summary>
        /// Implements host GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetHostResult> Invoke(GetHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostResult>("azure-native:connectedvmwarevsphere:getHost", args ?? new GetHostInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements host GET method.
        /// 
        /// Uses Azure REST API version 2023-12-01.
        /// 
        /// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetHostResult> Invoke(GetHostInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostResult>("azure-native:connectedvmwarevsphere:getHost", args ?? new GetHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the host.
        /// </summary>
        [Input("hostName", required: true)]
        public string HostName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHostArgs()
        {
        }
        public static new GetHostArgs Empty => new GetHostArgs();
    }

    public sealed class GetHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the host.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetHostInvokeArgs()
        {
        }
        public static new GetHostInvokeArgs Empty => new GetHostInvokeArgs();
    }


    [OutputType]
    public sealed class GetHostResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets the max CPU usage across all cores in MHz.
        /// </summary>
        public readonly double CpuMhz;
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
        /// Gets or sets the inventory Item ID for the host.
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
        /// Gets the total amount of physical memory on the host in GB.
        /// </summary>
        public readonly double MemorySizeGB;
        /// <summary>
        /// Gets or sets the vCenter Managed Object name for the host.
        /// </summary>
        public readonly string MoName;
        /// <summary>
        /// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
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
        /// Gets the used CPU usage across all cores in MHz.
        /// </summary>
        public readonly double OverallCpuUsageMHz;
        /// <summary>
        /// Gets the used physical memory on the host in GB.
        /// </summary>
        public readonly double OverallMemoryUsageGB;
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
        /// Gets or sets the type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets or sets a unique identifier for this resource.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// Gets or sets the ARM Id of the vCenter resource in which this host resides.
        /// </summary>
        public readonly string? VCenterId;

        [OutputConstructor]
        private GetHostResult(
            string azureApiVersion,

            double cpuMhz,

            string customResourceName,

            ImmutableArray<string> datastoreIds,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string id,

            string? inventoryItemId,

            string? kind,

            string location,

            double memorySizeGB,

            string moName,

            string? moRefId,

            string name,

            ImmutableArray<string> networkIds,

            double overallCpuUsageMHz,

            double overallMemoryUsageGB,

            string provisioningState,

            ImmutableArray<Outputs.ResourceStatusResponse> statuses,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string uuid,

            string? vCenterId)
        {
            AzureApiVersion = azureApiVersion;
            CpuMhz = cpuMhz;
            CustomResourceName = customResourceName;
            DatastoreIds = datastoreIds;
            ExtendedLocation = extendedLocation;
            Id = id;
            InventoryItemId = inventoryItemId;
            Kind = kind;
            Location = location;
            MemorySizeGB = memorySizeGB;
            MoName = moName;
            MoRefId = moRefId;
            Name = name;
            NetworkIds = networkIds;
            OverallCpuUsageMHz = overallCpuUsageMHz;
            OverallMemoryUsageGB = overallMemoryUsageGB;
            ProvisioningState = provisioningState;
            Statuses = statuses;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            Uuid = uuid;
            VCenterId = vCenterId;
        }
    }
}
