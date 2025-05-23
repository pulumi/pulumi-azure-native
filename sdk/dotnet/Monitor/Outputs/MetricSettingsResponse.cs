// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
    /// </summary>
    [OutputType]
    public sealed class MetricSettingsResponse
    {
        /// <summary>
        /// Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// a value indicating whether this category is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// the retention policy for this category.
        /// </summary>
        public readonly Outputs.RetentionPolicyResponse? RetentionPolicy;
        /// <summary>
        /// the timegrain of the metric in ISO8601 format.
        /// </summary>
        public readonly string? TimeGrain;

        [OutputConstructor]
        private MetricSettingsResponse(
            string? category,

            bool enabled,

            Outputs.RetentionPolicyResponse? retentionPolicy,

            string? timeGrain)
        {
            Category = category;
            Enabled = enabled;
            RetentionPolicy = retentionPolicy;
            TimeGrain = timeGrain;
        }
    }
}
