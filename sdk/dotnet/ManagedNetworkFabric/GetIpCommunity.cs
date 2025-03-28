// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    public static class GetIpCommunity
    {
        /// <summary>
        /// Implements an IP Community GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Task<GetIpCommunityResult> InvokeAsync(GetIpCommunityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpCommunityResult>("azure-native:managednetworkfabric:getIpCommunity", args ?? new GetIpCommunityArgs(), options.WithDefaults());

        /// <summary>
        /// Implements an IP Community GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Output<GetIpCommunityResult> Invoke(GetIpCommunityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpCommunityResult>("azure-native:managednetworkfabric:getIpCommunity", args ?? new GetIpCommunityInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements an IP Community GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Output<GetIpCommunityResult> Invoke(GetIpCommunityInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpCommunityResult>("azure-native:managednetworkfabric:getIpCommunity", args ?? new GetIpCommunityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpCommunityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the IP Community.
        /// </summary>
        [Input("ipCommunityName", required: true)]
        public string IpCommunityName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIpCommunityArgs()
        {
        }
        public static new GetIpCommunityArgs Empty => new GetIpCommunityArgs();
    }

    public sealed class GetIpCommunityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the IP Community.
        /// </summary>
        [Input("ipCommunityName", required: true)]
        public Input<string> IpCommunityName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetIpCommunityInvokeArgs()
        {
        }
        public static new GetIpCommunityInvokeArgs Empty => new GetIpCommunityInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpCommunityResult
    {
        /// <summary>
        /// Action to be taken on the configuration. Example: Permit | Deny.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        public readonly string? Annotation;
        /// <summary>
        /// List the communityMembers of IP Community .
        /// </summary>
        public readonly ImmutableArray<string> CommunityMembers;
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
        /// Gets the provisioning state of the resource.
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
        /// Supported well known Community List.
        /// </summary>
        public readonly ImmutableArray<string> WellKnownCommunities;

        [OutputConstructor]
        private GetIpCommunityResult(
            string action,

            string? annotation,

            ImmutableArray<string> communityMembers,

            string id,

            string location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> wellKnownCommunities)
        {
            Action = action;
            Annotation = annotation;
            CommunityMembers = communityMembers;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            WellKnownCommunities = wellKnownCommunities;
        }
    }
}
