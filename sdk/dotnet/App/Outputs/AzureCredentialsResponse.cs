// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Container App credentials.
    /// </summary>
    [OutputType]
    public sealed class AzureCredentialsResponse
    {
        /// <summary>
        /// Subscription Id.
        /// </summary>
        public readonly string? SubscriptionId;

        [OutputConstructor]
        private AzureCredentialsResponse(string? subscriptionId)
        {
            SubscriptionId = subscriptionId;
        }
    }
}
