// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AnalysisServices.Outputs
{

    /// <summary>
    /// An array of firewall rules.
    /// </summary>
    [OutputType]
    public sealed class IPv4FirewallSettingsResponse
    {
        /// <summary>
        /// The indicator of enabling PBI service.
        /// </summary>
        public readonly bool? EnablePowerBIService;
        /// <summary>
        /// An array of firewall rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPv4FirewallRuleResponse> FirewallRules;

        [OutputConstructor]
        private IPv4FirewallSettingsResponse(
            bool? enablePowerBIService,

            ImmutableArray<Outputs.IPv4FirewallRuleResponse> firewallRules)
        {
            EnablePowerBIService = enablePowerBIService;
            FirewallRules = firewallRules;
        }
    }
}
