// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.Inputs
{

    /// <summary>
    /// Properties to configure retention settings for the  eventhub
    /// </summary>
    public sealed class RetentionDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enumerates the possible values for cleanup policy
        /// </summary>
        [Input("cleanupPolicy")]
        public InputUnion<string, Pulumi.AzureNative.EventHub.CleanupPolicyRetentionDescription>? CleanupPolicy { get; set; }

        /// <summary>
        /// Number of hours to retain the events for this Event Hub. This value is only used when cleanupPolicy is Delete. If cleanupPolicy is Compact the returned value of this property is Long.MaxValue 
        /// </summary>
        [Input("retentionTimeInHours")]
        public Input<double>? RetentionTimeInHours { get; set; }

        /// <summary>
        /// Number of hours to retain the tombstone markers of a compacted Event Hub. This value is only used when cleanupPolicy is Compact. Consumer must complete reading the tombstone marker within this specified amount of time if consumer begins from starting offset to ensure they get a valid snapshot for the specific key described by the tombstone marker within the compacted Event Hub
        /// </summary>
        [Input("tombstoneRetentionTimeInHours")]
        public Input<int>? TombstoneRetentionTimeInHours { get; set; }

        public RetentionDescriptionArgs()
        {
        }
        public static new RetentionDescriptionArgs Empty => new RetentionDescriptionArgs();
    }
}
