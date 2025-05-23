// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SqlVirtualMachine.Outputs
{

    /// <summary>
    /// SQL VM Troubleshoot UnhealthyReplica scenario information.
    /// </summary>
    [OutputType]
    public sealed class UnhealthyReplicaInfoResponse
    {
        /// <summary>
        /// The name of the availability group
        /// </summary>
        public readonly string? AvailabilityGroupName;

        [OutputConstructor]
        private UnhealthyReplicaInfoResponse(string? availabilityGroupName)
        {
            AvailabilityGroupName = availabilityGroupName;
        }
    }
}
