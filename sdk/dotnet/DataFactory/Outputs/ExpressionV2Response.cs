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
    /// Nested representation of a complex expression.
    /// </summary>
    [OutputType]
    public sealed class ExpressionV2Response
    {
        /// <summary>
        /// List of nested expressions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressionV2Response> Operands;
        /// <summary>
        /// Expression operator value Type: list of strings.
        /// </summary>
        public readonly ImmutableArray<string> Operators;
        /// <summary>
        /// Type of expressions supported by the system. Type: string.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Value for Constant/Field Type: object.
        /// </summary>
        public readonly object? Value;

        [OutputConstructor]
        private ExpressionV2Response(
            ImmutableArray<Outputs.ExpressionV2Response> operands,

            ImmutableArray<string> operators,

            string? type,

            object? value)
        {
            Operands = operands;
            Operators = operators;
            Type = type;
            Value = value;
        }
    }
}
