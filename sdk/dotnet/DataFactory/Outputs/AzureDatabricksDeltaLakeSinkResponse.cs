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
    /// A copy activity Azure Databricks Delta Lake sink.
    /// </summary>
    [OutputType]
    public sealed class AzureDatabricksDeltaLakeSinkResponse
    {
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DisableMetricsCollection;
        /// <summary>
        /// Azure Databricks Delta Lake import settings.
        /// </summary>
        public readonly Outputs.AzureDatabricksDeltaLakeImportCommandResponse? ImportSettings;
        /// <summary>
        /// The maximum concurrent connection count for the sink data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// SQL pre-copy script. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PreCopyScript;
        /// <summary>
        /// Sink retry count. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? SinkRetryCount;
        /// <summary>
        /// Sink retry wait. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? SinkRetryWait;
        /// <summary>
        /// Copy sink type.
        /// Expected value is 'AzureDatabricksDeltaLakeSink'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Write batch size. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public readonly object? WriteBatchSize;
        /// <summary>
        /// Write batch timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? WriteBatchTimeout;

        [OutputConstructor]
        private AzureDatabricksDeltaLakeSinkResponse(
            object? disableMetricsCollection,

            Outputs.AzureDatabricksDeltaLakeImportCommandResponse? importSettings,

            object? maxConcurrentConnections,

            object? preCopyScript,

            object? sinkRetryCount,

            object? sinkRetryWait,

            string type,

            object? writeBatchSize,

            object? writeBatchTimeout)
        {
            DisableMetricsCollection = disableMetricsCollection;
            ImportSettings = importSettings;
            MaxConcurrentConnections = maxConcurrentConnections;
            PreCopyScript = preCopyScript;
            SinkRetryCount = sinkRetryCount;
            SinkRetryWait = sinkRetryWait;
            Type = type;
            WriteBatchSize = writeBatchSize;
            WriteBatchTimeout = writeBatchTimeout;
        }
    }
}
