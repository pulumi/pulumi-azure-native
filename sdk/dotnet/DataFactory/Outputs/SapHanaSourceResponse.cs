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
    /// A copy activity source for SAP HANA source.
    /// </summary>
    [OutputType]
    public sealed class SapHanaSourceResponse
    {
        /// <summary>
        /// Specifies the additional columns to be added to source data. Type: array of objects(AdditionalColumns) (or Expression with resultType array of objects).
        /// </summary>
        public readonly object? AdditionalColumns;
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DisableMetricsCollection;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// The packet size of data read from SAP HANA. Type: integer(or Expression with resultType integer).
        /// </summary>
        public readonly object? PacketSize;
        /// <summary>
        /// The partition mechanism that will be used for SAP HANA read in parallel. Possible values include: "None", "PhysicalPartitionsOfTable", "SapHanaDynamicRange". 
        /// </summary>
        public readonly object? PartitionOption;
        /// <summary>
        /// The settings that will be leveraged for SAP HANA source partitioning.
        /// </summary>
        public readonly Outputs.SapHanaPartitionSettingsResponse? PartitionSettings;
        /// <summary>
        /// SAP HANA Sql query. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Query;
        /// <summary>
        /// Query timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? QueryTimeout;
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
        /// Expected value is 'SapHanaSource'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SapHanaSourceResponse(
            object? additionalColumns,

            object? disableMetricsCollection,

            object? maxConcurrentConnections,

            object? packetSize,

            object? partitionOption,

            Outputs.SapHanaPartitionSettingsResponse? partitionSettings,

            object? query,

            object? queryTimeout,

            object? sourceRetryCount,

            object? sourceRetryWait,

            string type)
        {
            AdditionalColumns = additionalColumns;
            DisableMetricsCollection = disableMetricsCollection;
            MaxConcurrentConnections = maxConcurrentConnections;
            PacketSize = packetSize;
            PartitionOption = partitionOption;
            PartitionSettings = partitionSettings;
            Query = query;
            QueryTimeout = queryTimeout;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            Type = type;
        }
    }
}
