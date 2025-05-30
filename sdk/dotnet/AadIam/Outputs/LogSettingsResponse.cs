// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AadIam.Outputs
{

    /// <summary>
    /// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
    /// </summary>
    [OutputType]
    public sealed class LogSettingsResponse
    {
        /// <summary>
        /// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// A value indicating whether this log is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The retention policy for this log.
        /// </summary>
        public readonly Outputs.RetentionPolicyResponse? RetentionPolicy;

        [OutputConstructor]
        private LogSettingsResponse(
            string? category,

            bool enabled,

            Outputs.RetentionPolicyResponse? retentionPolicy)
        {
            Category = category;
            Enabled = enabled;
            RetentionPolicy = retentionPolicy;
        }
    }
}
