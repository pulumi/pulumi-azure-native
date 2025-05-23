// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Outputs
{

    [OutputType]
    public sealed class AccessControlResponse
    {
        /// <summary>
        /// The behavior for IP access control in Key Delivery.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// The IP allow list for access control in Key Delivery. If the default action is set to 'Allow', the IP allow list must be empty.
        /// </summary>
        public readonly ImmutableArray<string> IpAllowList;

        [OutputConstructor]
        private AccessControlResponse(
            string? defaultAction,

            ImmutableArray<string> ipAllowList)
        {
            DefaultAction = defaultAction;
            IpAllowList = ipAllowList;
        }
    }
}
