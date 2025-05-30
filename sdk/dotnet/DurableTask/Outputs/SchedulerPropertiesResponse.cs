// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DurableTask.Outputs
{

    /// <summary>
    /// Details of the Scheduler
    /// </summary>
    [OutputType]
    public sealed class SchedulerPropertiesResponse
    {
        /// <summary>
        /// URL of the durable task scheduler
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// IP allow list for durable task scheduler. Values can be IPv4, IPv6 or CIDR
        /// </summary>
        public readonly ImmutableArray<string> IpAllowlist;
        /// <summary>
        /// The status of the last operation
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// SKU of the durable task scheduler
        /// </summary>
        public readonly Outputs.SchedulerSkuResponse Sku;

        [OutputConstructor]
        private SchedulerPropertiesResponse(
            string endpoint,

            ImmutableArray<string> ipAllowlist,

            string provisioningState,

            Outputs.SchedulerSkuResponse sku)
        {
            Endpoint = endpoint;
            IpAllowlist = ipAllowlist;
            ProvisioningState = provisioningState;
            Sku = sku;
        }
    }
}
