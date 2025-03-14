// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230901Preview.Inputs
{

    /// <summary>
    /// Describes an automation rule condition that evaluates an array property's value
    /// </summary>
    public sealed class PropertyArrayConditionPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditionProperties")]
        public Input<Inputs.AutomationRulePropertyArrayValuesConditionArgs>? ConditionProperties { get; set; }

        /// <summary>
        /// 
        /// Expected value is 'PropertyArray'.
        /// </summary>
        [Input("conditionType", required: true)]
        public Input<string> ConditionType { get; set; } = null!;

        public PropertyArrayConditionPropertiesArgs()
        {
        }
        public static new PropertyArrayConditionPropertiesArgs Empty => new PropertyArrayConditionPropertiesArgs();
    }
}
