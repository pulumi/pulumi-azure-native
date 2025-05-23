// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// An Activity Log Alert rule condition that is met when all its member conditions are met.
    /// </summary>
    public sealed class AlertRuleAllOfConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allOf", required: true)]
        private InputList<Inputs.AlertRuleAnyOfOrLeafConditionArgs>? _allOf;

        /// <summary>
        /// The list of Activity Log Alert rule conditions.
        /// </summary>
        public InputList<Inputs.AlertRuleAnyOfOrLeafConditionArgs> AllOf
        {
            get => _allOf ?? (_allOf = new InputList<Inputs.AlertRuleAnyOfOrLeafConditionArgs>());
            set => _allOf = value;
        }

        public AlertRuleAllOfConditionArgs()
        {
        }
        public static new AlertRuleAllOfConditionArgs Empty => new AlertRuleAllOfConditionArgs();
    }
}
