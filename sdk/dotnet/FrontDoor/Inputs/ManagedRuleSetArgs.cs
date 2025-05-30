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
    /// Defines a managed rule set.
    /// </summary>
    public sealed class ManagedRuleSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("exclusions")]
        private InputList<Inputs.ManagedRuleExclusionArgs>? _exclusions;

        /// <summary>
        /// Describes the exclusions that are applied to all rules in the set.
        /// </summary>
        public InputList<Inputs.ManagedRuleExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.ManagedRuleExclusionArgs>());
            set => _exclusions = value;
        }

        [Input("ruleGroupOverrides")]
        private InputList<Inputs.ManagedRuleGroupOverrideArgs>? _ruleGroupOverrides;

        /// <summary>
        /// Defines the rule group overrides to apply to the rule set.
        /// </summary>
        public InputList<Inputs.ManagedRuleGroupOverrideArgs> RuleGroupOverrides
        {
            get => _ruleGroupOverrides ?? (_ruleGroupOverrides = new InputList<Inputs.ManagedRuleGroupOverrideArgs>());
            set => _ruleGroupOverrides = value;
        }

        /// <summary>
        /// Defines the rule set action.
        /// </summary>
        [Input("ruleSetAction")]
        public InputUnion<string, Pulumi.AzureNative.FrontDoor.ManagedRuleSetActionType>? RuleSetAction { get; set; }

        /// <summary>
        /// Defines the rule set type to use.
        /// </summary>
        [Input("ruleSetType", required: true)]
        public Input<string> RuleSetType { get; set; } = null!;

        /// <summary>
        /// Defines the version of the rule set to use.
        /// </summary>
        [Input("ruleSetVersion", required: true)]
        public Input<string> RuleSetVersion { get; set; } = null!;

        public ManagedRuleSetArgs()
        {
        }
        public static new ManagedRuleSetArgs Empty => new ManagedRuleSetArgs();
    }
}
