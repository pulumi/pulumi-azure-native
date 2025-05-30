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
    /// Defines a managed rule group to use for exclusion.
    /// </summary>
    [OutputType]
    public sealed class ExclusionManagedRuleGroupResponse
    {
        /// <summary>
        /// The managed rule group for exclusion.
        /// </summary>
        public readonly string RuleGroupName;
        /// <summary>
        /// List of rules that will be excluded. If none specified, all rules in the group will be excluded.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExclusionManagedRuleResponse> Rules;

        [OutputConstructor]
        private ExclusionManagedRuleGroupResponse(
            string ruleGroupName,

            ImmutableArray<Outputs.ExclusionManagedRuleResponse> rules)
        {
            RuleGroupName = ruleGroupName;
            Rules = rules;
        }
    }
}
