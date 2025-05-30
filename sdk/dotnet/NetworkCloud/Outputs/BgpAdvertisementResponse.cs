// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud.Outputs
{

    [OutputType]
    public sealed class BgpAdvertisementResponse
    {
        /// <summary>
        /// The indicator of if this advertisement is also made to the network fabric associated with the Network Cloud Cluster. This field is ignored if fabricPeeringEnabled is set to False.
        /// </summary>
        public readonly string? AdvertiseToFabric;
        /// <summary>
        /// The names of the BGP communities to be associated with the announcement, utilizing a BGP community string in 1234:1234 format.
        /// </summary>
        public readonly ImmutableArray<string> Communities;
        /// <summary>
        /// The names of the IP address pools associated with this announcement.
        /// </summary>
        public readonly ImmutableArray<string> IpAddressPools;
        /// <summary>
        /// The names of the BGP peers to limit this advertisement to. If no values are specified, all BGP peers will receive this advertisement.
        /// </summary>
        public readonly ImmutableArray<string> Peers;

        [OutputConstructor]
        private BgpAdvertisementResponse(
            string? advertiseToFabric,

            ImmutableArray<string> communities,

            ImmutableArray<string> ipAddressPools,

            ImmutableArray<string> peers)
        {
            AdvertiseToFabric = advertiseToFabric;
            Communities = communities;
            IpAddressPools = ipAddressPools;
            Peers = peers;
        }
    }
}
