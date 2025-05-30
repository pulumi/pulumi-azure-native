// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// The comparison rule.
    /// </summary>
    [OutputType]
    public sealed class ComparisonRuleResponse
    {
        /// <summary>
        /// The comparison operator.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Threshold setting.
        /// </summary>
        public readonly double Threshold;

        [OutputConstructor]
        private ComparisonRuleResponse(
            string @operator,

            double threshold)
        {
            Operator = @operator;
            Threshold = threshold;
        }
    }
}
