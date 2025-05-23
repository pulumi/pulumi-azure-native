// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Outputs
{

    /// <summary>
    /// AddressSpace contains an array of IP address ranges that can be used by subnets of the virtual network.
    /// </summary>
    [OutputType]
    public sealed class AddressSpaceResponse
    {
        /// <summary>
        /// A list of address blocks reserved for this virtual network in CIDR notation.
        /// </summary>
        public readonly ImmutableArray<string> AddressPrefixes;

        [OutputConstructor]
        private AddressSpaceResponse(ImmutableArray<string> addressPrefixes)
        {
            AddressPrefixes = addressPrefixes;
        }
    }
}
