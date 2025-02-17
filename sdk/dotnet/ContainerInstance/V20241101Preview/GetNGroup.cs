// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.V20241101Preview
{
    public static class GetNGroup
    {
        /// <summary>
        /// Get the properties of the specified NGroups resource.
        /// </summary>
        public static Task<GetNGroupResult> InvokeAsync(GetNGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNGroupResult>("azure-native:containerinstance/v20241101preview:getNGroup", args ?? new GetNGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of the specified NGroups resource.
        /// </summary>
        public static Output<GetNGroupResult> Invoke(GetNGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNGroupResult>("azure-native:containerinstance/v20241101preview:getNGroup", args ?? new GetNGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of the specified NGroups resource.
        /// </summary>
        public static Output<GetNGroupResult> Invoke(GetNGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNGroupResult>("azure-native:containerinstance/v20241101preview:getNGroup", args ?? new GetNGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The NGroups name.
        /// </summary>
        [Input("ngroupsName", required: true)]
        public string NgroupsName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNGroupArgs()
        {
        }
        public static new GetNGroupArgs Empty => new GetNGroupArgs();
    }

    public sealed class GetNGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The NGroups name.
        /// </summary>
        [Input("ngroupsName", required: true)]
        public Input<string> NgroupsName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNGroupInvokeArgs()
        {
        }
        public static new GetNGroupInvokeArgs Empty => new GetNGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetNGroupResult
    {
        /// <summary>
        /// The Container Group Profiles that could be used in the NGroups resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupProfileStubResponse> ContainerGroupProfiles;
        /// <summary>
        /// The elastic profile.
        /// </summary>
        public readonly Outputs.ElasticProfileResponse? ElasticProfile;
        /// <summary>
        /// The resource id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the NGroup, if configured.
        /// </summary>
        public readonly Outputs.NGroupIdentityResponse? Identity;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provides options w.r.t allocation and management w.r.t certain placement policies. These utilize capabilities provided by the underlying Azure infrastructure. They are typically used for high availability scenarios. E.g., distributing CGs across fault domains.
        /// </summary>
        public readonly Outputs.PlacementProfileResponse? PlacementProfile;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Used by the customer to specify the way to update the Container Groups in NGroup.
        /// </summary>
        public readonly Outputs.UpdateProfileResponse? UpdateProfile;
        /// <summary>
        /// The zones for the container group.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetNGroupResult(
            ImmutableArray<Outputs.ContainerGroupProfileStubResponse> containerGroupProfiles,

            Outputs.ElasticProfileResponse? elasticProfile,

            string id,

            Outputs.NGroupIdentityResponse? identity,

            string? location,

            string name,

            Outputs.PlacementProfileResponse? placementProfile,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.UpdateProfileResponse? updateProfile,

            ImmutableArray<string> zones)
        {
            ContainerGroupProfiles = containerGroupProfiles;
            ElasticProfile = elasticProfile;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            PlacementProfile = placementProfile;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            UpdateProfile = updateProfile;
            Zones = zones;
        }
    }
}
