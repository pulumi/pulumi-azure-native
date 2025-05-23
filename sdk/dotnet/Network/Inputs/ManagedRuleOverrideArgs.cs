// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Defines a managed rule group override setting.
    /// </summary>
    public sealed class ManagedRuleOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the override action to be applied when rule matches.
        /// </summary>
        [Input("action")]
        public InputUnion<string, Pulumi.AzureNative.Network.ActionType>? Action { get; set; }

        /// <summary>
        /// Identifier for the managed rule.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// Describes the override sensitivity to be applied when rule matches.
        /// </summary>
        [Input("sensitivity")]
        public InputUnion<string, Pulumi.AzureNative.Network.SensitivityType>? Sensitivity { get; set; }

        /// <summary>
        /// The state of the managed rule. Defaults to Disabled if not specified.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.Network.ManagedRuleEnabledState>? State { get; set; }

        public ManagedRuleOverrideArgs()
        {
        }
        public static new ManagedRuleOverrideArgs Empty => new ManagedRuleOverrideArgs();
    }
}
