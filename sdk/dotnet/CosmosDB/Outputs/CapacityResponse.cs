// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Outputs
{

    /// <summary>
    /// The object that represents all properties related to capacity enforcement on an account.
    /// </summary>
    [OutputType]
    public sealed class CapacityResponse
    {
        /// <summary>
        /// The total throughput limit imposed on the account. A totalThroughputLimit of 2000 imposes a strict limit of max throughput that can be provisioned on that account to be 2000. A totalThroughputLimit of -1 indicates no limits on provisioning of throughput.
        /// </summary>
        public readonly int? TotalThroughputLimit;

        [OutputConstructor]
        private CapacityResponse(int? totalThroughputLimit)
        {
            TotalThroughputLimit = totalThroughputLimit;
        }
    }
}
