// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Inputs
{

    /// <summary>
    /// Kafka endpoint Authentication properties. NOTE - only authentication property is allowed per entry
    /// </summary>
    public sealed class DataflowEndpointKafkaAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode of Authentication.
        /// </summary>
        [Input("method", required: true)]
        public InputUnion<string, Pulumi.AzureNative.IoTOperations.KafkaAuthMethod> Method { get; set; } = null!;

        /// <summary>
        /// SASL authentication.
        /// </summary>
        [Input("saslSettings")]
        public Input<Inputs.DataflowEndpointAuthenticationSaslArgs>? SaslSettings { get; set; }

        /// <summary>
        /// System-assigned managed identity authentication.
        /// </summary>
        [Input("systemAssignedManagedIdentitySettings")]
        public Input<Inputs.DataflowEndpointAuthenticationSystemAssignedManagedIdentityArgs>? SystemAssignedManagedIdentitySettings { get; set; }

        /// <summary>
        /// User-assigned managed identity authentication.
        /// </summary>
        [Input("userAssignedManagedIdentitySettings")]
        public Input<Inputs.DataflowEndpointAuthenticationUserAssignedManagedIdentityArgs>? UserAssignedManagedIdentitySettings { get; set; }

        /// <summary>
        /// X.509 certificate authentication.
        /// </summary>
        [Input("x509CertificateSettings")]
        public Input<Inputs.DataflowEndpointAuthenticationX509Args>? X509CertificateSettings { get; set; }

        public DataflowEndpointKafkaAuthenticationArgs()
        {
        }
        public static new DataflowEndpointKafkaAuthenticationArgs Empty => new DataflowEndpointKafkaAuthenticationArgs();
    }
}
