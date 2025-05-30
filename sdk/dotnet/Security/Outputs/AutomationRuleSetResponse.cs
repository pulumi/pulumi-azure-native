// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// A rule set which evaluates all its rules upon an event interception. Only when all the included rules in the rule set will be evaluated as 'true', will the event trigger the defined actions. 
    /// </summary>
    [OutputType]
    public sealed class AutomationRuleSetResponse
    {
        public readonly ImmutableArray<Outputs.AutomationTriggeringRuleResponse> Rules;

        [OutputConstructor]
        private AutomationRuleSetResponse(ImmutableArray<Outputs.AutomationTriggeringRuleResponse> rules)
        {
            Rules = rules;
        }
    }
}
