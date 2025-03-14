// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20180701.Outputs
{

    /// <summary>
    /// Policies for vpn gateway.
    /// </summary>
    [OutputType]
    public sealed class PoliciesResponse
    {
        /// <summary>
        /// True if branch to branch traffic is allowed.
        /// </summary>
        public readonly bool? AllowBranchToBranchTraffic;
        /// <summary>
        /// True if Vnet to Vnet traffic is allowed.
        /// </summary>
        public readonly bool? AllowVnetToVnetTraffic;

        [OutputConstructor]
        private PoliciesResponse(
            bool? allowBranchToBranchTraffic,

            bool? allowVnetToVnetTraffic)
        {
            AllowBranchToBranchTraffic = allowBranchToBranchTraffic;
            AllowVnetToVnetTraffic = allowVnetToVnetTraffic;
        }
    }
}
