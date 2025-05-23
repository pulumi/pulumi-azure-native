// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.Outputs
{

    /// <summary>
    /// SKU details
    /// </summary>
    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// SKU family name
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// SKU name to specify whether the key vault is a standard vault or a premium vault.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SkuResponse(
            string family,

            string name)
        {
            Family = family;
            Name = name;
        }
    }
}
