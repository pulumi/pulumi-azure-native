// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.FrontDoor.Inputs
{

    /// <summary>
    /// Define a match condition
    /// </summary>
    public sealed class RulesEngineMatchConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes if this is negate condition or not
        /// </summary>
        [Input("negateCondition")]
        public Input<bool>? NegateCondition { get; set; }

        [Input("rulesEngineMatchValue", required: true)]
        private InputList<string>? _rulesEngineMatchValue;

        /// <summary>
        /// Match values to match against. The operator will apply to each value in here with OR semantics. If any of them match the variable with the given operator this match condition is considered a match.
        /// </summary>
        public InputList<string> RulesEngineMatchValue
        {
            get => _rulesEngineMatchValue ?? (_rulesEngineMatchValue = new InputList<string>());
            set => _rulesEngineMatchValue = value;
        }

        /// <summary>
        /// Match Variable
        /// </summary>
        [Input("rulesEngineMatchVariable", required: true)]
        public InputUnion<string, Pulumi.AzureNative.FrontDoor.RulesEngineMatchVariable> RulesEngineMatchVariable { get; set; } = null!;

        /// <summary>
        /// Describes operator to apply to the match condition.
        /// </summary>
        [Input("rulesEngineOperator", required: true)]
        public InputUnion<string, Pulumi.AzureNative.FrontDoor.RulesEngineOperator> RulesEngineOperator { get; set; } = null!;

        /// <summary>
        /// Name of selector in RequestHeader or RequestBody to be matched
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        [Input("transforms")]
        private InputList<Union<string, Pulumi.AzureNative.FrontDoor.Transform>>? _transforms;

        /// <summary>
        /// List of transforms
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.FrontDoor.Transform>> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Union<string, Pulumi.AzureNative.FrontDoor.Transform>>());
            set => _transforms = value;
        }

        public RulesEngineMatchConditionArgs()
        {
        }
        public static new RulesEngineMatchConditionArgs Empty => new RulesEngineMatchConditionArgs();
    }
}
