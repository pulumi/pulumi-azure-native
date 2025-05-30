// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scheduler.Inputs
{

    public sealed class ServiceBusBrokeredMessagePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the content type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Gets or sets the correlation ID.
        /// </summary>
        [Input("correlationId")]
        public Input<string>? CorrelationId { get; set; }

        /// <summary>
        /// Gets or sets the force persistence.
        /// </summary>
        [Input("forcePersistence")]
        public Input<bool>? ForcePersistence { get; set; }

        /// <summary>
        /// Gets or sets the label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Gets or sets the message ID.
        /// </summary>
        [Input("messageId")]
        public Input<string>? MessageId { get; set; }

        /// <summary>
        /// Gets or sets the partition key.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        /// <summary>
        /// Gets or sets the reply to.
        /// </summary>
        [Input("replyTo")]
        public Input<string>? ReplyTo { get; set; }

        /// <summary>
        /// Gets or sets the reply to session ID.
        /// </summary>
        [Input("replyToSessionId")]
        public Input<string>? ReplyToSessionId { get; set; }

        /// <summary>
        /// Gets or sets the scheduled enqueue time UTC.
        /// </summary>
        [Input("scheduledEnqueueTimeUtc")]
        public Input<string>? ScheduledEnqueueTimeUtc { get; set; }

        /// <summary>
        /// Gets or sets the session ID.
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Gets or sets the time to live.
        /// </summary>
        [Input("timeToLive")]
        public Input<string>? TimeToLive { get; set; }

        /// <summary>
        /// Gets or sets the to.
        /// </summary>
        [Input("to")]
        public Input<string>? To { get; set; }

        /// <summary>
        /// Gets or sets the via partition key.
        /// </summary>
        [Input("viaPartitionKey")]
        public Input<string>? ViaPartitionKey { get; set; }

        public ServiceBusBrokeredMessagePropertiesArgs()
        {
        }
        public static new ServiceBusBrokeredMessagePropertiesArgs Empty => new ServiceBusBrokeredMessagePropertiesArgs();
    }
}
