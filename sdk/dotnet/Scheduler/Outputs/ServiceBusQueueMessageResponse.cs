// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scheduler.Outputs
{

    [OutputType]
    public sealed class ServiceBusQueueMessageResponse
    {
        /// <summary>
        /// Gets or sets the Service Bus authentication.
        /// </summary>
        public readonly Outputs.ServiceBusAuthenticationResponse? Authentication;
        /// <summary>
        /// Gets or sets the brokered message properties.
        /// </summary>
        public readonly Outputs.ServiceBusBrokeredMessagePropertiesResponse? BrokeredMessageProperties;
        /// <summary>
        /// Gets or sets the custom message properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomMessageProperties;
        /// <summary>
        /// Gets or sets the message.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Gets or sets the namespace.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// Gets or sets the queue name.
        /// </summary>
        public readonly string? QueueName;
        /// <summary>
        /// Gets or sets the transport type.
        /// </summary>
        public readonly string? TransportType;

        [OutputConstructor]
        private ServiceBusQueueMessageResponse(
            Outputs.ServiceBusAuthenticationResponse? authentication,

            Outputs.ServiceBusBrokeredMessagePropertiesResponse? brokeredMessageProperties,

            ImmutableDictionary<string, string>? customMessageProperties,

            string? message,

            string? @namespace,

            string? queueName,

            string? transportType)
        {
            Authentication = authentication;
            BrokeredMessageProperties = brokeredMessageProperties;
            CustomMessageProperties = customMessageProperties;
            Message = message;
            Namespace = @namespace;
            QueueName = queueName;
            TransportType = transportType;
        }
    }
}
