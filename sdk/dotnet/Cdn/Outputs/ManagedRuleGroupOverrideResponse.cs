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
    /// Defines a managed rule group override setting.
    /// </summary>
    [OutputType]
    public sealed class ManagedRuleGroupOverrideResponse
    {
        /// <summary>
        /// Describes the managed rule group within the rule set to override
        /// </summary>
        public readonly string RuleGroupName;
        /// <summary>
        /// List of rules that will be enabled. If none specified, all rules in the group will be disabled.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedRuleOverrideResponse> Rules;

        [OutputConstructor]
        private ManagedRuleGroupOverrideResponse(
            string ruleGroupName,

            ImmutableArray<Outputs.ManagedRuleOverrideResponse> rules)
        {
            RuleGroupName = ruleGroupName;
            Rules = rules;
        }
    }
}
