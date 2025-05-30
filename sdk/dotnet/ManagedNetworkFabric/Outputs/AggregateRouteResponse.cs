// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Outputs
{

    /// <summary>
    /// aggregateIpv4Route model.
    /// </summary>
    [OutputType]
    public sealed class AggregateRouteResponse
    {
        /// <summary>
        /// IPv4 Prefix of the aggregate Ipv4Route.
        /// </summary>
        public readonly string Prefix;

        [OutputConstructor]
        private AggregateRouteResponse(string prefix)
        {
            Prefix = prefix;
        }
    }
}
