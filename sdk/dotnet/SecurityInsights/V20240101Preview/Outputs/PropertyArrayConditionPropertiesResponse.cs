// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20240101Preview.Outputs
{

    /// <summary>
    /// Describes an automation rule condition that evaluates an array property's value
    /// </summary>
    [OutputType]
    public sealed class PropertyArrayConditionPropertiesResponse
    {
        public readonly Outputs.AutomationRulePropertyArrayValuesConditionResponse? ConditionProperties;
        /// <summary>
        /// 
        /// Expected value is 'PropertyArray'.
        /// </summary>
        public readonly string ConditionType;

        [OutputConstructor]
        private PropertyArrayConditionPropertiesResponse(
            Outputs.AutomationRulePropertyArrayValuesConditionResponse? conditionProperties,

            string conditionType)
        {
            ConditionProperties = conditionProperties;
            ConditionType = conditionType;
        }
    }
}
