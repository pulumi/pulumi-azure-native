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
    /// Application gateway web application firewall configuration.
    /// </summary>
    [OutputType]
    public sealed class ApplicationGatewayWebApplicationFirewallConfigurationResponse
    {
        /// <summary>
        /// The disabled rule groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFirewallDisabledRuleGroupResponse> DisabledRuleGroups;
        /// <summary>
        /// Whether the web application firewall is enabled or not.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The exclusion list.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFirewallExclusionResponse> Exclusions;
        /// <summary>
        /// Maximum file upload size in Mb for WAF.
        /// </summary>
        public readonly int? FileUploadLimitInMb;
        /// <summary>
        /// Web application firewall mode.
        /// </summary>
        public readonly string FirewallMode;
        /// <summary>
        /// Maximum request body size for WAF.
        /// </summary>
        public readonly int? MaxRequestBodySize;
        /// <summary>
        /// Maximum request body size in Kb for WAF.
        /// </summary>
        public readonly int? MaxRequestBodySizeInKb;
        /// <summary>
        /// Whether allow WAF to check request Body.
        /// </summary>
        public readonly bool? RequestBodyCheck;
        /// <summary>
        /// The type of the web application firewall rule set. Possible values are: 'OWASP'.
        /// </summary>
        public readonly string RuleSetType;
        /// <summary>
        /// The version of the rule set type.
        /// </summary>
        public readonly string RuleSetVersion;

        [OutputConstructor]
        private ApplicationGatewayWebApplicationFirewallConfigurationResponse(
            ImmutableArray<Outputs.ApplicationGatewayFirewallDisabledRuleGroupResponse> disabledRuleGroups,

            bool enabled,

            ImmutableArray<Outputs.ApplicationGatewayFirewallExclusionResponse> exclusions,

            int? fileUploadLimitInMb,

            string firewallMode,

            int? maxRequestBodySize,

            int? maxRequestBodySizeInKb,

            bool? requestBodyCheck,

            string ruleSetType,

            string ruleSetVersion)
        {
            DisabledRuleGroups = disabledRuleGroups;
            Enabled = enabled;
            Exclusions = exclusions;
            FileUploadLimitInMb = fileUploadLimitInMb;
            FirewallMode = firewallMode;
            MaxRequestBodySize = maxRequestBodySize;
            MaxRequestBodySizeInKb = maxRequestBodySizeInKb;
            RequestBodyCheck = requestBodyCheck;
            RuleSetType = ruleSetType;
            RuleSetVersion = ruleSetVersion;
        }
    }
}
