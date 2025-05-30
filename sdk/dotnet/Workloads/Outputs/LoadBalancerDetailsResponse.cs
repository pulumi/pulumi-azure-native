// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// The Load Balancer details such as Load Balancer ID.
    /// </summary>
    [OutputType]
    public sealed class LoadBalancerDetailsResponse
    {
        /// <summary>
        /// Fully qualified resource ID for the load balancer.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LoadBalancerDetailsResponse(string id)
        {
            Id = id;
        }
    }
}
