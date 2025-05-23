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
    /// A copy activity source for an Office 365 service.
    /// </summary>
    [OutputType]
    public sealed class Office365SourceResponse
    {
        /// <summary>
        /// The groups containing all the users. Type: array of strings (or Expression with resultType array of strings).
        /// </summary>
        public readonly object? AllowedGroups;
        /// <summary>
        /// The Column to apply the &lt;paramref name="StartTime"/&gt; and &lt;paramref name="EndTime"/&gt;. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? DateFilterColumn;
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DisableMetricsCollection;
        /// <summary>
        /// End time of the requested range for this dataset. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? EndTime;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// The columns to be read out from the Office 365 table. Type: array of objects (or Expression with resultType array of objects). itemType: OutputColumn. Example: [ { "name": "Id" }, { "name": "CreatedDateTime" } ]
        /// </summary>
        public readonly object? OutputColumns;
        /// <summary>
        /// Source retry count. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? SourceRetryCount;
        /// <summary>
        /// Source retry wait. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? SourceRetryWait;
        /// <summary>
        /// Start time of the requested range for this dataset. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? StartTime;
        /// <summary>
        /// Copy source type.
        /// Expected value is 'Office365Source'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user scope uri. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? UserScopeFilterUri;

        [OutputConstructor]
        private Office365SourceResponse(
            object? allowedGroups,

            object? dateFilterColumn,

            object? disableMetricsCollection,

            object? endTime,

            object? maxConcurrentConnections,

            object? outputColumns,

            object? sourceRetryCount,

            object? sourceRetryWait,

            object? startTime,

            string type,

            object? userScopeFilterUri)
        {
            AllowedGroups = allowedGroups;
            DateFilterColumn = dateFilterColumn;
            DisableMetricsCollection = disableMetricsCollection;
            EndTime = endTime;
            MaxConcurrentConnections = maxConcurrentConnections;
            OutputColumns = outputColumns;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            StartTime = startTime;
            Type = type;
            UserScopeFilterUri = userScopeFilterUri;
        }
    }
}
