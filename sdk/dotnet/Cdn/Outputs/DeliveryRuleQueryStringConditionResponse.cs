// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Outputs
{

    /// <summary>
    /// Defines the QueryString condition for the delivery rule.
    /// </summary>
    [OutputType]
    public sealed class DeliveryRuleQueryStringConditionResponse
    {
        /// <summary>
        /// The name of the condition for the delivery rule.
        /// Expected value is 'QueryString'.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the parameters for the condition.
        /// </summary>
        public readonly Outputs.QueryStringMatchConditionParametersResponse Parameters;

        [OutputConstructor]
        private DeliveryRuleQueryStringConditionResponse(
            string name,

            Outputs.QueryStringMatchConditionParametersResponse parameters)
        {
            Name = name;
            Parameters = parameters;
        }
    }
}
