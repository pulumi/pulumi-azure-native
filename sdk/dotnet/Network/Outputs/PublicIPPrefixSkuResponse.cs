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
    /// SKU of a public IP prefix.
    /// </summary>
    [OutputType]
    public sealed class PublicIPPrefixSkuResponse
    {
        /// <summary>
        /// Name of a public IP prefix SKU.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tier of a public IP prefix SKU.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private PublicIPPrefixSkuResponse(
            string? name,

            string? tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
