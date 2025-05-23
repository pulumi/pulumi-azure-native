// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet
{
    public static class GetFleet
    {
        /// <summary>
        /// Get a Fleet
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurefleet [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetFleetResult> InvokeAsync(GetFleetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFleetResult>("azure-native:azurefleet:getFleet", args ?? new GetFleetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Fleet
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurefleet [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFleetResult> Invoke(GetFleetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFleetResult>("azure-native:azurefleet:getFleet", args ?? new GetFleetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Fleet
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurefleet [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFleetResult> Invoke(GetFleetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFleetResult>("azure-native:azurefleet:getFleet", args ?? new GetFleetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFleetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Compute Fleet
        /// </summary>
        [Input("fleetName", required: true)]
        public string FleetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFleetArgs()
        {
        }
        public static new GetFleetArgs Empty => new GetFleetArgs();
    }

    public sealed class GetFleetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Compute Fleet
        /// </summary>
        [Input("fleetName", required: true)]
        public Input<string> FleetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFleetInvokeArgs()
        {
        }
        public static new GetFleetInvokeArgs Empty => new GetFleetInvokeArgs();
    }


    [OutputType]
    public sealed class GetFleetResult
    {
        /// <summary>
        /// Represents the configuration for additional locations where Fleet resources may be deployed.
        /// </summary>
        public readonly Outputs.AdditionalLocationsProfileResponse? AdditionalLocationsProfile;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Compute Profile to use for running user's workloads.
        /// </summary>
        public readonly Outputs.ComputeProfileResponse ComputeProfile;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Details of the resource plan.
        /// </summary>
        public readonly Outputs.PlanResponse? Plan;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Configuration Options for Regular instances in Compute Fleet.
        /// </summary>
        public readonly Outputs.RegularPriorityProfileResponse? RegularPriorityProfile;
        /// <summary>
        /// Configuration Options for Spot instances in Compute Fleet.
        /// </summary>
        public readonly Outputs.SpotPriorityProfileResponse? SpotPriorityProfile;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies the time at which the Compute Fleet is created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies the ID which uniquely identifies a Compute Fleet.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// Attribute based Fleet.
        /// </summary>
        public readonly Outputs.VMAttributesResponse? VmAttributes;
        /// <summary>
        /// List of VM sizes supported for Compute Fleet
        /// </summary>
        public readonly ImmutableArray<Outputs.VmSizeProfileResponse> VmSizesProfile;
        /// <summary>
        /// Zones in which the Compute Fleet is available
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetFleetResult(
            Outputs.AdditionalLocationsProfileResponse? additionalLocationsProfile,

            string azureApiVersion,

            Outputs.ComputeProfileResponse computeProfile,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.PlanResponse? plan,

            string provisioningState,

            Outputs.RegularPriorityProfileResponse? regularPriorityProfile,

            Outputs.SpotPriorityProfileResponse? spotPriorityProfile,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string timeCreated,

            string type,

            string uniqueId,

            Outputs.VMAttributesResponse? vmAttributes,

            ImmutableArray<Outputs.VmSizeProfileResponse> vmSizesProfile,

            ImmutableArray<string> zones)
        {
            AdditionalLocationsProfile = additionalLocationsProfile;
            AzureApiVersion = azureApiVersion;
            ComputeProfile = computeProfile;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Plan = plan;
            ProvisioningState = provisioningState;
            RegularPriorityProfile = regularPriorityProfile;
            SpotPriorityProfile = spotPriorityProfile;
            SystemData = systemData;
            Tags = tags;
            TimeCreated = timeCreated;
            Type = type;
            UniqueId = uniqueId;
            VmAttributes = vmAttributes;
            VmSizesProfile = vmSizesProfile;
            Zones = zones;
        }
    }
}
