// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The soft delete policy for a container registry
    /// </summary>
    [OutputType]
    public sealed class SoftDeletePolicyResponse
    {
        /// <summary>
        /// The timestamp when the policy was last updated.
        /// </summary>
        public readonly string LastUpdatedTime;
        /// <summary>
        /// The number of days after which a soft-deleted item is permanently deleted.
        /// </summary>
        public readonly int? RetentionDays;
        /// <summary>
        /// The value that indicates whether the policy is enabled or not.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SoftDeletePolicyResponse(
            string lastUpdatedTime,

            int? retentionDays,

            string? status)
        {
            LastUpdatedTime = lastUpdatedTime;
            RetentionDays = retentionDays;
            Status = status;
        }
    }
}
