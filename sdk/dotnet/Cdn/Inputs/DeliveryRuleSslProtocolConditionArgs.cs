// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Inputs
{

    /// <summary>
    /// Defines the SslProtocol condition for the delivery rule.
    /// </summary>
    public sealed class DeliveryRuleSslProtocolConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the condition for the delivery rule.
        /// Expected value is 'SslProtocol'.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Defines the parameters for the condition.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<Inputs.SslProtocolMatchConditionParametersArgs> Parameters { get; set; } = null!;

        public DeliveryRuleSslProtocolConditionArgs()
        {
        }
        public static new DeliveryRuleSslProtocolConditionArgs Empty => new DeliveryRuleSslProtocolConditionArgs();
    }
}
