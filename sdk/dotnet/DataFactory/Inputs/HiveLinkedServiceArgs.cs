// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Hive Server linked service.
    /// </summary>
    public sealed class HiveLinkedServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to require a CA-issued SSL certificate name to match the host name of the server when connecting over SSL. The default value is false.
        /// </summary>
        [Input("allowHostNameCNMismatch")]
        public Input<object>? AllowHostNameCNMismatch { get; set; }

        /// <summary>
        /// Specifies whether to allow self-signed certificates from the server. The default value is false.
        /// </summary>
        [Input("allowSelfSignedServerCert")]
        public Input<object>? AllowSelfSignedServerCert { get; set; }

        [Input("annotations")]
        private InputList<object>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public InputList<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// The authentication method used to access the Hive server.
        /// </summary>
        [Input("authenticationType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.HiveAuthenticationType> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether the connections to the server will validate server certificate, the default value is True. Only used for Version 2.0
        /// </summary>
        [Input("enableServerCertificateValidation")]
        public Input<object>? EnableServerCertificateValidation { get; set; }

        /// <summary>
        /// Specifies whether the connections to the server are encrypted using SSL. The default value is false.
        /// </summary>
        [Input("enableSsl")]
        public Input<object>? EnableSsl { get; set; }

        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string.
        /// </summary>
        [Input("encryptedCredential")]
        public Input<string>? EncryptedCredential { get; set; }

        /// <summary>
        /// IP address or host name of the Hive server, separated by ';' for multiple hosts (only when serviceDiscoveryMode is enable).
        /// </summary>
        [Input("host", required: true)]
        public Input<object> Host { get; set; } = null!;

        /// <summary>
        /// The partial URL corresponding to the Hive server.
        /// </summary>
        [Input("httpPath")]
        public Input<object>? HttpPath { get; set; }

        [Input("parameters")]
        private InputMap<Inputs.ParameterSpecificationArgs>? _parameters;

        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public InputMap<Inputs.ParameterSpecificationArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterSpecificationArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The password corresponding to the user name that you provided in the Username field
        /// </summary>
        [Input("password")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? Password { get; set; }

        /// <summary>
        /// The TCP port that the Hive server uses to listen for client connections.
        /// </summary>
        [Input("port")]
        public Input<object>? Port { get; set; }

        /// <summary>
        /// The type of Hive server.
        /// </summary>
        [Input("serverType")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.HiveServerType>? ServerType { get; set; }

        /// <summary>
        /// true to indicate using the ZooKeeper service, false not.
        /// </summary>
        [Input("serviceDiscoveryMode")]
        public Input<object>? ServiceDiscoveryMode { get; set; }

        /// <summary>
        /// The transport protocol to use in the Thrift layer.
        /// </summary>
        [Input("thriftTransportProtocol")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.HiveThriftTransportProtocol>? ThriftTransportProtocol { get; set; }

        /// <summary>
        /// The full path of the .pem file containing trusted CA certificates for verifying the server when connecting over SSL. This property can only be set when using SSL on self-hosted IR. The default value is the cacerts.pem file installed with the IR.
        /// </summary>
        [Input("trustedCertPath")]
        public Input<object>? TrustedCertPath { get; set; }

        /// <summary>
        /// Type of linked service.
        /// Expected value is 'Hive'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies whether the driver uses native HiveQL queries,or converts them into an equivalent form in HiveQL.
        /// </summary>
        [Input("useNativeQuery")]
        public Input<object>? UseNativeQuery { get; set; }

        /// <summary>
        /// Specifies whether to use a CA certificate from the system trust store or from a specified PEM file. The default value is false.
        /// </summary>
        [Input("useSystemTrustStore")]
        public Input<object>? UseSystemTrustStore { get; set; }

        /// <summary>
        /// The user name that you use to access Hive Server.
        /// </summary>
        [Input("username")]
        public Input<object>? Username { get; set; }

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The namespace on ZooKeeper under which Hive Server 2 nodes are added.
        /// </summary>
        [Input("zooKeeperNameSpace")]
        public Input<object>? ZooKeeperNameSpace { get; set; }

        public HiveLinkedServiceArgs()
        {
        }
        public static new HiveLinkedServiceArgs Empty => new HiveLinkedServiceArgs();
    }
}
