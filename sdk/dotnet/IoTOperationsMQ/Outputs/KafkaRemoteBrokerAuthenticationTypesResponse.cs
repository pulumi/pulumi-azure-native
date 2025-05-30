// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Outputs
{

    /// <summary>
    /// Kafka RemoteBrokerConnection Authentication types. NOTE - Enum only one method is allowed to be passed.
    /// </summary>
    [OutputType]
    public sealed class KafkaRemoteBrokerAuthenticationTypesResponse
    {
        /// <summary>
        /// Sasl remote broker authentication method.
        /// </summary>
        public readonly Outputs.SaslRemoteBrokerBasicAuthenticationResponse? Sasl;
        /// <summary>
        /// Managed identity remote broker authentication method.
        /// </summary>
        public readonly Outputs.ManagedIdentityAuthenticationResponse? SystemAssignedManagedIdentity;
        /// <summary>
        /// X509 remote broker authentication method.
        /// </summary>
        public readonly Outputs.KafkaX509AuthenticationResponse? X509;

        [OutputConstructor]
        private KafkaRemoteBrokerAuthenticationTypesResponse(
            Outputs.SaslRemoteBrokerBasicAuthenticationResponse? sasl,

            Outputs.ManagedIdentityAuthenticationResponse? systemAssignedManagedIdentity,

            Outputs.KafkaX509AuthenticationResponse? x509)
        {
            Sasl = sasl;
            SystemAssignedManagedIdentity = systemAssignedManagedIdentity;
            X509 = x509;
        }
    }
}
