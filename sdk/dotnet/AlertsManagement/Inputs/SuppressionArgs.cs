// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement.Inputs
{

    /// <summary>
    /// Action rule with suppression configuration
    /// </summary>
    public sealed class SuppressionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// conditions on which alerts will be filtered
        /// </summary>
        [Input("conditions")]
        public Input<Inputs.ConditionsArgs>? Conditions { get; set; }

        /// <summary>
        /// Description of action rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// scope on which action rule will apply
        /// </summary>
        [Input("scope")]
        public Input<Inputs.ScopeArgs>? Scope { get; set; }

        /// <summary>
        /// Indicates if the given action rule is enabled or disabled
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.AlertsManagement.ActionRuleStatus>? Status { get; set; }

        /// <summary>
        /// suppression configuration for the action rule
        /// </summary>
        [Input("suppressionConfig", required: true)]
        public Input<Inputs.SuppressionConfigArgs> SuppressionConfig { get; set; } = null!;

        /// <summary>
        /// Indicates type of action rule
        /// Expected value is 'Suppression'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SuppressionArgs()
        {
        }
        public static new SuppressionArgs Empty => new SuppressionArgs();
    }
}
