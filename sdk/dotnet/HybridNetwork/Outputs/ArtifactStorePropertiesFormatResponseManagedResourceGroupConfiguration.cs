// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Outputs
{

    [OutputType]
    public sealed class ArtifactStorePropertiesFormatResponseManagedResourceGroupConfiguration
    {
        /// <summary>
        /// The managed resource group location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The managed resource group name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ArtifactStorePropertiesFormatResponseManagedResourceGroupConfiguration(
            string? location,

            string? name)
        {
            Location = location;
            Name = name;
        }
    }
}
