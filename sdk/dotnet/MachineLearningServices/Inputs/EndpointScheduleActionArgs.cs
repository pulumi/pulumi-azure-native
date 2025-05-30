// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class EndpointScheduleActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// 
        /// Expected value is 'InvokeBatchEndpoint'.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// [Required] Defines Schedule action definition details.
        /// &lt;see href="TBD" /&gt;
        /// </summary>
        [Input("endpointInvocationDefinition", required: true)]
        public Input<object> EndpointInvocationDefinition { get; set; } = null!;

        public EndpointScheduleActionArgs()
        {
        }
        public static new EndpointScheduleActionArgs Empty => new EndpointScheduleActionArgs();
    }
}
