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
    /// A copy activity source for OData source.
    /// </summary>
    [OutputType]
    public sealed class ODataSourceResponse
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
        /// The timeout (TimeSpan) to get an HTTP response. It is the timeout to get a response, not the timeout to read response data. Default value: 00:05:00. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? HttpRequestTimeout;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// OData query. For example, "$top=1". Type: string (or Expression with resultType string).
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
        /// Expected value is 'ODataSource'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ODataSourceResponse(
            object? additionalColumns,

            object? disableMetricsCollection,

            object? httpRequestTimeout,

            object? maxConcurrentConnections,

            object? query,

            object? sourceRetryCount,

            object? sourceRetryWait,

            string type)
        {
            AdditionalColumns = additionalColumns;
            DisableMetricsCollection = disableMetricsCollection;
            HttpRequestTimeout = httpRequestTimeout;
            MaxConcurrentConnections = maxConcurrentConnections;
            Query = query;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            Type = type;
        }
    }
}
