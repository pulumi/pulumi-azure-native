// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    /// <summary>
    /// The definition of data present in the report.
    /// </summary>
    [OutputType]
    public sealed class ReportConfigDatasetResponse
    {
        /// <summary>
        /// Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ReportConfigAggregationResponse>? Aggregation;
        /// <summary>
        /// Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
        /// </summary>
        public readonly Outputs.ReportConfigDatasetConfigurationResponse? Configuration;
        /// <summary>
        /// Has filter expression to use in the report.
        /// </summary>
        public readonly Outputs.ReportConfigFilterResponse? Filter;
        /// <summary>
        /// The granularity of rows in the report.
        /// </summary>
        public readonly string? Granularity;
        /// <summary>
        /// Array of group by expression to use in the report. Report can have up to 2 group by clauses.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportConfigGroupingResponse> Grouping;
        /// <summary>
        /// Array of order by expression to use in the report.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportConfigSortingResponse> Sorting;

        [OutputConstructor]
        private ReportConfigDatasetResponse(
            ImmutableDictionary<string, Outputs.ReportConfigAggregationResponse>? aggregation,

            Outputs.ReportConfigDatasetConfigurationResponse? configuration,

            Outputs.ReportConfigFilterResponse? filter,

            string? granularity,

            ImmutableArray<Outputs.ReportConfigGroupingResponse> grouping,

            ImmutableArray<Outputs.ReportConfigSortingResponse> sorting)
        {
            Aggregation = aggregation;
            Configuration = configuration;
            Filter = filter;
            Granularity = granularity;
            Grouping = grouping;
            Sorting = sorting;
        }
    }
}
