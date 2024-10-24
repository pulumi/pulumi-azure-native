// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.V20230701Preview.Outputs
{

    /// <summary>
    /// OutboundIP represents a desired outbound IP resource for the cluster load balancer.
    /// </summary>
    [OutputType]
    public sealed class OutboundIPResponse
    {
        /// <summary>
        /// The fully qualified Azure resource id of the IP address resource.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private OutboundIPResponse(string? id)
        {
            Id = id;
        }
    }
}
