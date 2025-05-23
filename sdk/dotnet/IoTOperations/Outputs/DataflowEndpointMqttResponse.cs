// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// Broker endpoint properties
    /// </summary>
    [OutputType]
    public sealed class DataflowEndpointMqttResponse
    {
        /// <summary>
        /// authentication properties. DEFAULT: kubernetes.audience=aio-internal. NOTE - Enum field only property is allowed
        /// </summary>
        public readonly Outputs.DataflowEndpointMqttAuthenticationResponse Authentication;
        /// <summary>
        /// Client ID prefix. Client ID generated by the dataflow is &lt;prefix&gt;-TBD. Optional; no prefix if omitted.
        /// </summary>
        public readonly string? ClientIdPrefix;
        /// <summary>
        /// Cloud event mapping config.
        /// </summary>
        public readonly string? CloudEventAttributes;
        /// <summary>
        /// Host of the Broker in the form of &lt;hostname&gt;:&lt;port&gt;. Optional; connects to Broker if omitted.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Broker KeepAlive for connection in seconds.
        /// </summary>
        public readonly int? KeepAliveSeconds;
        /// <summary>
        /// The max number of messages to keep in flight. For subscribe, this is the receive maximum. For publish, this is the maximum number of messages to send before waiting for an ack.
        /// </summary>
        public readonly int? MaxInflightMessages;
        /// <summary>
        /// Enable or disable websockets.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Qos for Broker connection.
        /// </summary>
        public readonly int? Qos;
        /// <summary>
        /// Whether or not to keep the retain setting.
        /// </summary>
        public readonly string? Retain;
        /// <summary>
        /// Session expiry in seconds.
        /// </summary>
        public readonly int? SessionExpirySeconds;
        /// <summary>
        /// TLS configuration.
        /// </summary>
        public readonly Outputs.TlsPropertiesResponse? Tls;

        [OutputConstructor]
        private DataflowEndpointMqttResponse(
            Outputs.DataflowEndpointMqttAuthenticationResponse authentication,

            string? clientIdPrefix,

            string? cloudEventAttributes,

            string? host,

            int? keepAliveSeconds,

            int? maxInflightMessages,

            string? protocol,

            int? qos,

            string? retain,

            int? sessionExpirySeconds,

            Outputs.TlsPropertiesResponse? tls)
        {
            Authentication = authentication;
            ClientIdPrefix = clientIdPrefix;
            CloudEventAttributes = cloudEventAttributes;
            Host = host;
            KeepAliveSeconds = keepAliveSeconds;
            MaxInflightMessages = maxInflightMessages;
            Protocol = protocol;
            Qos = qos;
            Retain = retain;
            SessionExpirySeconds = sessionExpirySeconds;
            Tls = tls;
        }
    }
}
