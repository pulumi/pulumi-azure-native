// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Webhook details specific for Azure DevOps
    /// </summary>
    [OutputType]
    public sealed class AzureDevOpsWebhookResponse
    {
        /// <summary>
        /// Send callback on a specified notification event
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// Enum to determine the webhook callback service type.
        /// Expected value is 'AzureDevOps'.
        /// </summary>
        public readonly string WebhookType;

        [OutputConstructor]
        private AzureDevOpsWebhookResponse(
            string? eventType,

            string webhookType)
        {
            EventType = eventType;
            WebhookType = webhookType;
        }
    }
}
