// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Outputs
{

    /// <summary>
    /// Defines the common attributes for a custom rule that can be included in a waf policy
    /// </summary>
    [OutputType]
    public sealed class CustomRuleResponse
    {
        /// <summary>
        /// Describes what action to be applied when rule matches
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        /// </summary>
        public readonly string? EnabledState;
        /// <summary>
        /// List of match conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.MatchConditionResponse> MatchConditions;
        /// <summary>
        /// Defines the name of the custom rule
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines in what order this rule be evaluated in the overall list of custom rules
        /// </summary>
        public readonly int Priority;

        [OutputConstructor]
        private CustomRuleResponse(
            string action,

            string? enabledState,

            ImmutableArray<Outputs.MatchConditionResponse> matchConditions,

            string name,

            int priority)
        {
            Action = action;
            EnabledState = enabledState;
            MatchConditions = matchConditions;
            Name = name;
            Priority = priority;
        }
    }
}
