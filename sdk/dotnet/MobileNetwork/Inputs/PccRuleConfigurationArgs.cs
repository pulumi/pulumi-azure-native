// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// Data flow policy rule configuration
    /// </summary>
    public sealed class PccRuleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the rule. This must be unique within the parent service. You must not use any of the following reserved strings - `default`, `requested` or `service`.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        /// <summary>
        /// A precedence value that is used to decide between data flow policy rules when identifying the QoS values to use for a particular SIM. A lower value means a higher priority. This value should be unique among all data flow policy rules configured in the mobile network.
        /// </summary>
        [Input("rulePrecedence", required: true)]
        public Input<int> RulePrecedence { get; set; } = null!;

        /// <summary>
        /// The QoS policy to use for packets matching this rule. If this field is null then the parent service will define the QoS settings.
        /// </summary>
        [Input("ruleQosPolicy")]
        public Input<Inputs.PccRuleQosPolicyArgs>? RuleQosPolicy { get; set; }

        [Input("serviceDataFlowTemplates", required: true)]
        private InputList<Inputs.ServiceDataFlowTemplateArgs>? _serviceDataFlowTemplates;

        /// <summary>
        /// The set of data flow templates to use for this data flow policy rule.
        /// </summary>
        public InputList<Inputs.ServiceDataFlowTemplateArgs> ServiceDataFlowTemplates
        {
            get => _serviceDataFlowTemplates ?? (_serviceDataFlowTemplates = new InputList<Inputs.ServiceDataFlowTemplateArgs>());
            set => _serviceDataFlowTemplates = value;
        }

        /// <summary>
        /// Determines whether flows that match this data flow policy rule are permitted.
        /// </summary>
        [Input("trafficControl")]
        public InputUnion<string, Pulumi.AzureNative.MobileNetwork.TrafficControlPermission>? TrafficControl { get; set; }

        public PccRuleConfigurationArgs()
        {
            TrafficControl = "Enabled";
        }
        public static new PccRuleConfigurationArgs Empty => new PccRuleConfigurationArgs();
    }
}
