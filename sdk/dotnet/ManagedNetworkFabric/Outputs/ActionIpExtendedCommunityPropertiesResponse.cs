// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Outputs
{

    /// <summary>
    /// IP Extended Community Properties.
    /// </summary>
    [OutputType]
    public sealed class ActionIpExtendedCommunityPropertiesResponse
    {
        /// <summary>
        /// List of IP Extended Community IDs.
        /// </summary>
        public readonly Outputs.IpExtendedCommunityIdListResponse? Add;
        /// <summary>
        /// List of IP Extended Community IDs.
        /// </summary>
        public readonly Outputs.IpExtendedCommunityIdListResponse? Delete;
        /// <summary>
        /// List of IP Extended Community IDs.
        /// </summary>
        public readonly Outputs.IpExtendedCommunityIdListResponse? Set;

        [OutputConstructor]
        private ActionIpExtendedCommunityPropertiesResponse(
            Outputs.IpExtendedCommunityIdListResponse? add,

            Outputs.IpExtendedCommunityIdListResponse? delete,

            Outputs.IpExtendedCommunityIdListResponse? set)
        {
            Add = add;
            Delete = delete;
            Set = set;
        }
    }
}
