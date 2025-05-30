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
    /// Defines contents of custom rules
    /// </summary>
    [OutputType]
    public sealed class CustomRuleListResponse
    {
        /// <summary>
        /// List of rules
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomRuleResponse> Rules;

        [OutputConstructor]
        private CustomRuleListResponse(ImmutableArray<Outputs.CustomRuleResponse> rules)
        {
            Rules = rules;
        }
    }
}
