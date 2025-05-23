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
    /// Configuration for static routes on this HubVnetConnectionConfiguration for static routes on this HubVnetConnection.
    /// </summary>
    [OutputType]
    public sealed class StaticRoutesConfigResponse
    {
        /// <summary>
        /// Boolean indicating whether static routes on this connection are automatically propagate to route tables which this connection propagates to.
        /// </summary>
        public readonly bool PropagateStaticRoutes;
        /// <summary>
        /// Parameter determining whether NVA in spoke vnet is bypassed for traffic with destination in spoke.
        /// </summary>
        public readonly string? VnetLocalRouteOverrideCriteria;

        [OutputConstructor]
        private StaticRoutesConfigResponse(
            bool propagateStaticRoutes,

            string? vnetLocalRouteOverrideCriteria)
        {
            PropagateStaticRoutes = propagateStaticRoutes;
            VnetLocalRouteOverrideCriteria = vnetLocalRouteOverrideCriteria;
        }
    }
}
