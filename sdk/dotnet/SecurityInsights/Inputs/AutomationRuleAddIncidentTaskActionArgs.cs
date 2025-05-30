// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Describes an automation rule action to add a task to an incident
    /// </summary>
    public sealed class AutomationRuleAddIncidentTaskActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes an automation rule action to add a task to an incident.
        /// </summary>
        [Input("actionConfiguration")]
        public Input<Inputs.AddIncidentTaskActionPropertiesArgs>? ActionConfiguration { get; set; }

        /// <summary>
        /// The type of the automation rule action.
        /// Expected value is 'AddIncidentTask'.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        public AutomationRuleAddIncidentTaskActionArgs()
        {
        }
        public static new AutomationRuleAddIncidentTaskActionArgs Empty => new AutomationRuleAddIncidentTaskActionArgs();
    }
}
