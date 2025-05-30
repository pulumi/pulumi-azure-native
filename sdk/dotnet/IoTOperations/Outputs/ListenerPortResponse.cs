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
    /// Defines a TCP port on which a `BrokerListener` listens.
    /// </summary>
    [OutputType]
    public sealed class ListenerPortResponse
    {
        /// <summary>
        /// Reference to client authentication settings. Omit to disable authentication.
        /// </summary>
        public readonly string? AuthenticationRef;
        /// <summary>
        /// Reference to client authorization settings. Omit to disable authorization.
        /// </summary>
        public readonly string? AuthorizationRef;
        /// <summary>
        /// Kubernetes node port. Only relevant when this port is associated with a `NodePort` listener.
        /// </summary>
        public readonly int? NodePort;
        /// <summary>
        /// TCP port for accepting client connections.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Protocol to use for client connections.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// TLS server certificate settings for this port. Omit to disable TLS.
        /// </summary>
        public readonly Outputs.TlsCertMethodResponse? Tls;

        [OutputConstructor]
        private ListenerPortResponse(
            string? authenticationRef,

            string? authorizationRef,

            int? nodePort,

            int port,

            string? protocol,

            Outputs.TlsCertMethodResponse? tls)
        {
            AuthenticationRef = authenticationRef;
            AuthorizationRef = authorizationRef;
            NodePort = nodePort;
            Port = port;
            Protocol = protocol;
            Tls = tls;
        }
    }
}
