// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Nat Gateway resource.
    /// </summary>
    [OutputType]
    public sealed class NatGatewayResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The idle timeout of the nat gateway.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the NAT gateway resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// An array of public ip addresses associated with the nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> PublicIpAddresses;
        /// <summary>
        /// An array of public ip prefixes associated with the nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> PublicIpPrefixes;
        /// <summary>
        /// The resource GUID property of the NAT gateway resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// The nat gateway SKU.
        /// </summary>
        public readonly Outputs.NatGatewaySkuResponse? Sku;
        /// <summary>
        /// An array of references to the subnets using this nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> Subnets;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private NatGatewayResponse(
            string etag,

            string? id,

            int? idleTimeoutInMinutes,

            string? location,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.SubResourceResponse> publicIpAddresses,

            ImmutableArray<Outputs.SubResourceResponse> publicIpPrefixes,

            string resourceGuid,

            Outputs.NatGatewaySkuResponse? sku,

            ImmutableArray<Outputs.SubResourceResponse> subnets,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Etag = etag;
            Id = id;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            PublicIpAddresses = publicIpAddresses;
            PublicIpPrefixes = publicIpPrefixes;
            ResourceGuid = resourceGuid;
            Sku = sku;
            Subnets = subnets;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}
