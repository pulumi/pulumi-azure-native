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
    /// A copy activity Salesforce V2 source.
    /// </summary>
    [OutputType]
    public sealed class SalesforceV2SourceResponse
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
        /// This property control whether query result contains Deleted objects. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? IncludeDeletedObjects;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// Page size for each http request, too large pageSize will caused timeout, default 300,000. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? PageSize;
        /// <summary>
        /// You can only use Salesforce Object Query Language (SOQL) query with limitations. For SOQL limitations, see this article: https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/queries.htm#SOQL%20Considerations. If query is not specified, all the data of the Salesforce object specified in ObjectApiName/reportId in dataset will be retrieved. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Query;
        /// <summary>
        /// Query timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? QueryTimeout;
        /// <summary>
        /// Deprecating, please use 'query' property instead. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? SOQLQuery;
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
        /// Expected value is 'SalesforceV2Source'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SalesforceV2SourceResponse(
            object? additionalColumns,

            object? disableMetricsCollection,

            object? includeDeletedObjects,

            object? maxConcurrentConnections,

            object? pageSize,

            object? query,

            object? queryTimeout,

            object? sOQLQuery,

            object? sourceRetryCount,

            object? sourceRetryWait,

            string type)
        {
            AdditionalColumns = additionalColumns;
            DisableMetricsCollection = disableMetricsCollection;
            IncludeDeletedObjects = includeDeletedObjects;
            MaxConcurrentConnections = maxConcurrentConnections;
            PageSize = pageSize;
            Query = query;
            QueryTimeout = queryTimeout;
            SOQLQuery = sOQLQuery;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            Type = type;
        }
    }
}
