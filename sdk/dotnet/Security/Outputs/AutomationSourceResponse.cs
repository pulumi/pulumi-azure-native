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
    /// The source event types which evaluate the security automation set of rules. For example - security alerts and security assessments. To learn more about the supported security events data models schemas - please visit https://aka.ms/ASCAutomationSchemas.
    /// </summary>
    [OutputType]
    public sealed class AutomationSourceResponse
    {
        /// <summary>
        /// A valid event source type.
        /// </summary>
        public readonly string? EventSource;
        /// <summary>
        /// A set of rules which evaluate upon event interception. A logical disjunction is applied between defined rule sets (logical 'or').
        /// </summary>
        public readonly ImmutableArray<Outputs.AutomationRuleSetResponse> RuleSets;

        [OutputConstructor]
        private AutomationSourceResponse(
            string? eventSource,

            ImmutableArray<Outputs.AutomationRuleSetResponse> ruleSets)
        {
            EventSource = eventSource;
            RuleSets = ruleSets;
        }
    }
}
