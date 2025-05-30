// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Inputs
{

    /// <summary>
    /// Defines the common attributes for a custom rule that can be included in a waf policy
    /// </summary>
    public sealed class CustomRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes what action to be applied when rule matches
        /// </summary>
        [Input("action", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cdn.ActionType> Action { get; set; } = null!;

        /// <summary>
        /// Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        /// </summary>
        [Input("enabledState")]
        public InputUnion<string, Pulumi.AzureNative.Cdn.CustomRuleEnabledState>? EnabledState { get; set; }

        [Input("matchConditions", required: true)]
        private InputList<Inputs.MatchConditionArgs>? _matchConditions;

        /// <summary>
        /// List of match conditions.
        /// </summary>
        public InputList<Inputs.MatchConditionArgs> MatchConditions
        {
            get => _matchConditions ?? (_matchConditions = new InputList<Inputs.MatchConditionArgs>());
            set => _matchConditions = value;
        }

        /// <summary>
        /// Defines the name of the custom rule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Defines in what order this rule be evaluated in the overall list of custom rules
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public CustomRuleArgs()
        {
        }
        public static new CustomRuleArgs Empty => new CustomRuleArgs();
    }
}
