// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Webhook details specific for Azure DevOps
    /// </summary>
    public sealed class AzureDevOpsWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Send callback on a specified notification event
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Enum to determine the webhook callback service type.
        /// Expected value is 'AzureDevOps'.
        /// </summary>
        [Input("webhookType", required: true)]
        public Input<string> WebhookType { get; set; } = null!;

        public AzureDevOpsWebhookArgs()
        {
        }
        public static new AzureDevOpsWebhookArgs Empty => new AzureDevOpsWebhookArgs();
    }
}
