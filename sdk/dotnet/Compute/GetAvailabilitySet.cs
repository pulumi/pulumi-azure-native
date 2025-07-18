// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute
{
    public static class GetAvailabilitySet
    {
        /// <summary>
        /// Retrieves information about an availability set.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAvailabilitySetResult> InvokeAsync(GetAvailabilitySetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilitySetResult>("azure-native:compute:getAvailabilitySet", args ?? new GetAvailabilitySetArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about an availability set.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvailabilitySetResult> Invoke(GetAvailabilitySetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilitySetResult>("azure-native:compute:getAvailabilitySet", args ?? new GetAvailabilitySetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about an availability set.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAvailabilitySetResult> Invoke(GetAvailabilitySetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilitySetResult>("azure-native:compute:getAvailabilitySet", args ?? new GetAvailabilitySetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailabilitySetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the availability set.
        /// </summary>
        [Input("availabilitySetName", required: true)]
        public string AvailabilitySetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetArgs()
        {
        }
        public static new GetAvailabilitySetArgs Empty => new GetAvailabilitySetArgs();
    }

    public sealed class GetAvailabilitySetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the availability set.
        /// </summary>
        [Input("availabilitySetName", required: true)]
        public Input<string> AvailabilitySetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetInvokeArgs()
        {
        }
        public static new GetAvailabilitySetInvokeArgs Empty => new GetAvailabilitySetInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailabilitySetResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Fault Domain count.
        /// </summary>
        public readonly int? PlatformFaultDomainCount;
        /// <summary>
        /// Update Domain count.
        /// </summary>
        public readonly int? PlatformUpdateDomainCount;
        /// <summary>
        /// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? ProximityPlacementGroup;
        /// <summary>
        /// Specifies Redeploy, Reboot and ScheduledEventsAdditionalPublishingTargets Scheduled Event related configurations for the availability set.
        /// </summary>
        public readonly Outputs.ScheduledEventsPolicyResponse? ScheduledEventsPolicy;
        /// <summary>
        /// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponse> Statuses;
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
        /// Describes the migration properties on the Availability Set.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetMigrationInfoResponse VirtualMachineScaleSetMigrationInfo;
        /// <summary>
        /// A list of references to all virtual machines in the availability set.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> VirtualMachines;

        [OutputConstructor]
        private GetAvailabilitySetResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            int? platformFaultDomainCount,

            int? platformUpdateDomainCount,

            Outputs.SubResourceResponse? proximityPlacementGroup,

            Outputs.ScheduledEventsPolicyResponse? scheduledEventsPolicy,

            Outputs.SkuResponse? sku,

            ImmutableArray<Outputs.InstanceViewStatusResponse> statuses,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.VirtualMachineScaleSetMigrationInfoResponse virtualMachineScaleSetMigrationInfo,

            ImmutableArray<Outputs.SubResourceResponse> virtualMachines)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            PlatformFaultDomainCount = platformFaultDomainCount;
            PlatformUpdateDomainCount = platformUpdateDomainCount;
            ProximityPlacementGroup = proximityPlacementGroup;
            ScheduledEventsPolicy = scheduledEventsPolicy;
            Sku = sku;
            Statuses = statuses;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VirtualMachineScaleSetMigrationInfo = virtualMachineScaleSetMigrationInfo;
            VirtualMachines = virtualMachines;
        }
    }
}
