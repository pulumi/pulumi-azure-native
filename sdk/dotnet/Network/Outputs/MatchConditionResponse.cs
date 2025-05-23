// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Define match conditions.
    /// </summary>
    [OutputType]
    public sealed class MatchConditionResponse
    {
        /// <summary>
        /// Match value.
        /// </summary>
        public readonly ImmutableArray<string> MatchValues;
        /// <summary>
        /// List of match variables.
        /// </summary>
        public readonly ImmutableArray<Outputs.MatchVariableResponse> MatchVariables;
        /// <summary>
        /// Whether this is negate condition or not.
        /// </summary>
        public readonly bool? NegationConditon;
        /// <summary>
        /// The operator to be matched.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// List of transforms.
        /// </summary>
        public readonly ImmutableArray<string> Transforms;

        [OutputConstructor]
        private MatchConditionResponse(
            ImmutableArray<string> matchValues,

            ImmutableArray<Outputs.MatchVariableResponse> matchVariables,

            bool? negationConditon,

            string @operator,

            ImmutableArray<string> transforms)
        {
            MatchValues = matchValues;
            MatchVariables = matchVariables;
            NegationConditon = negationConditon;
            Operator = @operator;
            Transforms = transforms;
        }
    }
}
