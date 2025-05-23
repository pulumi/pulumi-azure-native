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
    /// The aggregation expression to be used in the report.
    /// </summary>
    [OutputType]
    public sealed class ReportAggregationResponse
    {
        /// <summary>
        /// The name of the aggregation function to use.
        /// </summary>
        public readonly string Function;
        /// <summary>
        /// The name of the column to aggregate.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ReportAggregationResponse(
            string function,

            string name)
        {
            Function = function;
            Name = name;
        }
    }
}
