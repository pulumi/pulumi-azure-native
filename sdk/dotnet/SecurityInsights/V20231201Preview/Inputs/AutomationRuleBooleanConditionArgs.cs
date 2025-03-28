// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20231201Preview.Inputs
{

    public sealed class AutomationRuleBooleanConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("innerConditions")]
        private InputList<object>? _innerConditions;
        public InputList<object> InnerConditions
        {
            get => _innerConditions ?? (_innerConditions = new InputList<object>());
            set => _innerConditions = value;
        }

        [Input("operator")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.V20231201Preview.AutomationRuleBooleanConditionSupportedOperator>? Operator { get; set; }

        public AutomationRuleBooleanConditionArgs()
        {
        }
        public static new AutomationRuleBooleanConditionArgs Empty => new AutomationRuleBooleanConditionArgs();
    }
}
