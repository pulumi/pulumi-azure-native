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
    /// Defines the RequestBody condition for the delivery rule.
    /// </summary>
    [OutputType]
    public sealed class DeliveryRuleRequestBodyConditionResponse
    {
        /// <summary>
        /// The name of the condition for the delivery rule.
        /// Expected value is 'RequestBody'.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the parameters for the condition.
        /// </summary>
        public readonly Outputs.RequestBodyMatchConditionParametersResponse Parameters;

        [OutputConstructor]
        private DeliveryRuleRequestBodyConditionResponse(
            string name,

            Outputs.RequestBodyMatchConditionParametersResponse parameters)
        {
            Name = name;
            Parameters = parameters;
        }
    }
}
