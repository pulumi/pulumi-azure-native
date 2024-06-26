// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20240401.Outputs
{

    /// <summary>
    /// Access profile for the Fleet hub API server.
    /// </summary>
    [OutputType]
    public sealed class APIServerAccessProfileResponse
    {
        /// <summary>
        /// Whether to create the Fleet hub as a private cluster or not.
        /// </summary>
        public readonly bool? EnablePrivateCluster;

        [OutputConstructor]
        private APIServerAccessProfileResponse(bool? enablePrivateCluster)
        {
            EnablePrivateCluster = enablePrivateCluster;
        }
    }
}
