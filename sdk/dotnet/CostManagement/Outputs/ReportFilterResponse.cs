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
    /// The filter expression to be used in the report.
    /// </summary>
    [OutputType]
    public sealed class ReportFilterResponse
    {
        /// <summary>
        /// The logical "AND" expression. Must have at least 2 items.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportFilterResponse> And;
        /// <summary>
        /// Has comparison expression for a dimension
        /// </summary>
        public readonly Outputs.ReportComparisonExpressionResponse? Dimension;
        /// <summary>
        /// The logical "NOT" expression.
        /// </summary>
        public readonly Outputs.ReportFilterResponse? Not;
        /// <summary>
        /// The logical "OR" expression. Must have at least 2 items.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportFilterResponse> Or;
        /// <summary>
        /// Has comparison expression for a tag
        /// </summary>
        public readonly Outputs.ReportComparisonExpressionResponse? Tag;

        [OutputConstructor]
        private ReportFilterResponse(
            ImmutableArray<Outputs.ReportFilterResponse> and,

            Outputs.ReportComparisonExpressionResponse? dimension,

            Outputs.ReportFilterResponse? not,

            ImmutableArray<Outputs.ReportFilterResponse> or,

            Outputs.ReportComparisonExpressionResponse? tag)
        {
            And = and;
            Dimension = dimension;
            Not = not;
            Or = or;
            Tag = tag;
        }
    }
}
