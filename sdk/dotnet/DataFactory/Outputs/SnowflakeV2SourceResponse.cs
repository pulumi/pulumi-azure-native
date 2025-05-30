// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// A copy activity snowflake source.
    /// </summary>
    [OutputType]
    public sealed class SnowflakeV2SourceResponse
    {
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DisableMetricsCollection;
        /// <summary>
        /// Snowflake export settings.
        /// </summary>
        public readonly Outputs.SnowflakeExportCopyCommandResponse ExportSettings;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// Snowflake Sql query. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Query;
        /// <summary>
        /// Source retry count. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? SourceRetryCount;
        /// <summary>
        /// Source retry wait. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? SourceRetryWait;
        /// <summary>
        /// Copy source type.
        /// Expected value is 'SnowflakeV2Source'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SnowflakeV2SourceResponse(
            object? disableMetricsCollection,

            Outputs.SnowflakeExportCopyCommandResponse exportSettings,

            object? maxConcurrentConnections,

            object? query,

            object? sourceRetryCount,

            object? sourceRetryWait,

            string type)
        {
            DisableMetricsCollection = disableMetricsCollection;
            ExportSettings = exportSettings;
            MaxConcurrentConnections = maxConcurrentConnections;
            Query = query;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            Type = type;
        }
    }
}
