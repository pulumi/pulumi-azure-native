// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230401Preview.Outputs
{

    [OutputType]
    public sealed class DatabaseSourceResponse
    {
        /// <summary>
        /// Workspace connection for data import source storage
        /// </summary>
        public readonly string? Connection;
        /// <summary>
        /// SQL Query statement for data import Database source
        /// </summary>
        public readonly string? Query;
        /// <summary>
        /// Enum to determine the type of data.
        /// Expected value is 'database'.
        /// </summary>
        public readonly string SourceType;
        /// <summary>
        /// SQL StoredProcedure on data import Database source
        /// </summary>
        public readonly string? StoredProcedure;
        /// <summary>
        /// SQL StoredProcedure parameters
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> StoredProcedureParams;
        /// <summary>
        /// Name of the table on data import Database source
        /// </summary>
        public readonly string? TableName;

        [OutputConstructor]
        private DatabaseSourceResponse(
            string? connection,

            string? query,

            string sourceType,

            string? storedProcedure,

            ImmutableArray<ImmutableDictionary<string, string>> storedProcedureParams,

            string? tableName)
        {
            Connection = connection;
            Query = query;
            SourceType = sourceType;
            StoredProcedure = storedProcedure;
            StoredProcedureParams = storedProcedureParams;
            TableName = tableName;
        }
    }
}
