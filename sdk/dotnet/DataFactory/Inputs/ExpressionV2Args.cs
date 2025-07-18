// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Nested representation of a complex expression.
    /// </summary>
    public sealed class ExpressionV2Args : global::Pulumi.ResourceArgs
    {
        [Input("operands")]
        private InputList<Inputs.ExpressionV2Args>? _operands;

        /// <summary>
        /// List of nested expressions.
        /// </summary>
        public InputList<Inputs.ExpressionV2Args> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.ExpressionV2Args>());
            set => _operands = value;
        }

        [Input("operators")]
        private InputList<string>? _operators;

        /// <summary>
        /// Expression operator value Type: list of strings.
        /// </summary>
        public InputList<string> Operators
        {
            get => _operators ?? (_operators = new InputList<string>());
            set => _operators = value;
        }

        /// <summary>
        /// Type of expressions supported by the system. Type: string.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.ExpressionV2Type>? Type { get; set; }

        /// <summary>
        /// Value for Constant/Field Type: object.
        /// </summary>
        [Input("value")]
        public Input<object>? Value { get; set; }

        public ExpressionV2Args()
        {
        }
        public static new ExpressionV2Args Empty => new ExpressionV2Args();
    }
}
