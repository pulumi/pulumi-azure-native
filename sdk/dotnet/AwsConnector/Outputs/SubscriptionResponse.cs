// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of Subscription
    /// </summary>
    [OutputType]
    public sealed class SubscriptionResponse
    {
        /// <summary>
        /// The endpoint that receives notifications from the SNS topic. The endpoint value depends on the protocol that you specify. For more information, see the ``Endpoint`` parameter of the ``Subscribe`` action in the *API Reference*.
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// The subscription's protocol. For more information, see the ``Protocol`` parameter of the ``Subscribe`` action in the *API Reference*.
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private SubscriptionResponse(
            string? endpoint,

            string? protocol)
        {
            Endpoint = endpoint;
            Protocol = protocol;
        }
    }
}
