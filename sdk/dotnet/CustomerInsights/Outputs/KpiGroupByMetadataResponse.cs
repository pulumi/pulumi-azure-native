// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Outputs
{

    /// <summary>
    /// The KPI GroupBy field metadata.
    /// </summary>
    [OutputType]
    public sealed class KpiGroupByMetadataResponse
    {
        /// <summary>
        /// The display name.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// The name of the field.
        /// </summary>
        public readonly string? FieldName;
        /// <summary>
        /// The type of the field.
        /// </summary>
        public readonly string? FieldType;

        [OutputConstructor]
        private KpiGroupByMetadataResponse(
            ImmutableDictionary<string, string>? displayName,

            string? fieldName,

            string? fieldType)
        {
            DisplayName = displayName;
            FieldName = fieldName;
            FieldType = fieldType;
        }
    }
}
