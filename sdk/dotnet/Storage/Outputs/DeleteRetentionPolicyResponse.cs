// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// The service properties for soft delete.
    /// </summary>
    [OutputType]
    public sealed class DeleteRetentionPolicyResponse
    {
        /// <summary>
        /// This property when set to true allows deletion of the soft deleted blob versions and snapshots. This property cannot be used blob restore policy. This property only applies to blob service and does not apply to containers or file share.
        /// </summary>
        public readonly bool? AllowPermanentDelete;
        /// <summary>
        /// Indicates the number of days that the deleted item should be retained. The minimum specified value can be 1 and the maximum value can be 365.
        /// </summary>
        public readonly int? Days;
        /// <summary>
        /// Indicates whether DeleteRetentionPolicy is enabled.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private DeleteRetentionPolicyResponse(
            bool? allowPermanentDelete,

            int? days,

            bool? enabled)
        {
            AllowPermanentDelete = allowPermanentDelete;
            Days = days;
            Enabled = enabled;
        }
    }
}
