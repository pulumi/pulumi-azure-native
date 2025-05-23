// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerPlatform.Outputs
{

    /// <summary>
    /// Settings concerning network injection.
    /// </summary>
    [OutputType]
    public sealed class PropertiesResponseNetworkInjection
    {
        /// <summary>
        /// Network injection configuration
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualNetworkPropertiesResponse> VirtualNetworks;

        [OutputConstructor]
        private PropertiesResponseNetworkInjection(ImmutableArray<Outputs.VirtualNetworkPropertiesResponse> virtualNetworks)
        {
            VirtualNetworks = virtualNetworks;
        }
    }
}
