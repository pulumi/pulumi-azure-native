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
    /// Exclude variables from managed rule evaluation.
    /// </summary>
    public sealed class ManagedRuleExclusionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The variable type to be excluded.
        /// </summary>
        [Input("matchVariable", required: true)]
        public InputUnion<string, Pulumi.AzureNative.FrontDoor.ManagedRuleExclusionMatchVariable> MatchVariable { get; set; } = null!;

        /// <summary>
        /// Selector value for which elements in the collection this exclusion applies to.
        /// </summary>
        [Input("selector", required: true)]
        public Input<string> Selector { get; set; } = null!;

        /// <summary>
        /// Comparison operator to apply to the selector when specifying which elements in the collection this exclusion applies to.
        /// </summary>
        [Input("selectorMatchOperator", required: true)]
        public InputUnion<string, Pulumi.AzureNative.FrontDoor.ManagedRuleExclusionSelectorMatchOperator> SelectorMatchOperator { get; set; } = null!;

        public ManagedRuleExclusionArgs()
        {
        }
        public static new ManagedRuleExclusionArgs Empty => new ManagedRuleExclusionArgs();
    }
}
