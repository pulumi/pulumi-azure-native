// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Outputs
{

    /// <summary>
    /// Ingress profile for the container service cluster.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterIngressProfileResponse
    {
        /// <summary>
        /// Web App Routing settings for the ingress profile.
        /// </summary>
        public readonly Outputs.ManagedClusterIngressProfileWebAppRoutingResponse? WebAppRouting;

        [OutputConstructor]
        private ManagedClusterIngressProfileResponse(Outputs.ManagedClusterIngressProfileWebAppRoutingResponse? webAppRouting)
        {
            WebAppRouting = webAppRouting;
        }
    }
}
