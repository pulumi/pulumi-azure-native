// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The logic app action that should be triggered. To learn more about Microsoft Defender for Cloud's Workflow Automation capabilities, visit https://aka.ms/ASCWorkflowAutomationLearnMore
    /// </summary>
    public sealed class AutomationActionLogicAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the action that will be triggered by the Automation
        /// Expected value is 'LogicApp'.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// The triggered Logic App Azure Resource ID. This can also reside on other subscriptions, given that you have permissions to trigger the Logic App
        /// </summary>
        [Input("logicAppResourceId")]
        public Input<string>? LogicAppResourceId { get; set; }

        /// <summary>
        /// The Logic App trigger URI endpoint (it will not be included in any response).
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public AutomationActionLogicAppArgs()
        {
        }
        public static new AutomationActionLogicAppArgs Empty => new AutomationActionLogicAppArgs();
    }
}
