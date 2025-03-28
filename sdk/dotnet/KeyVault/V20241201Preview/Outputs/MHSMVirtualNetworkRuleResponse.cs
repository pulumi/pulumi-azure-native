// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.V20241201Preview.Outputs
{

    /// <summary>
    /// A rule governing the accessibility of a managed hsm pool from a specific virtual network.
    /// </summary>
    [OutputType]
    public sealed class MHSMVirtualNetworkRuleResponse
    {
        /// <summary>
        /// Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private MHSMVirtualNetworkRuleResponse(string id)
        {
            Id = id;
        }
    }
}
