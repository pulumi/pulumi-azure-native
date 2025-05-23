// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Outputs
{

    /// <summary>
    /// List of network fabric controller ids.
    /// </summary>
    [OutputType]
    public sealed class ArtifactStoreNetworkFabricControllerEndPointsResponse
    {
        /// <summary>
        /// list of network fabric controllers.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReferencedResourceResponse> NetworkFabricControllerIds;

        [OutputConstructor]
        private ArtifactStoreNetworkFabricControllerEndPointsResponse(ImmutableArray<Outputs.ReferencedResourceResponse> networkFabricControllerIds)
        {
            NetworkFabricControllerIds = networkFabricControllerIds;
        }
    }
}
