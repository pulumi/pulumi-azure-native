// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Inputs
{

    /// <summary>
    /// MqttBridge RemoteBrokerConnection Authentication methods. NOTE - Enum only one is allowed to be passed.
    /// </summary>
    public sealed class MqttBridgeRemoteBrokerAuthenticationMethodsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed identity remote broker authentication method.
        /// </summary>
        [Input("systemAssignedManagedIdentity")]
        public Input<Inputs.ManagedIdentityAuthenticationArgs>? SystemAssignedManagedIdentity { get; set; }

        /// <summary>
        /// X509 remote broker authentication method.
        /// </summary>
        [Input("x509")]
        public Input<Inputs.MqttBridgeRemoteBrokerX509AuthenticationArgs>? X509 { get; set; }

        public MqttBridgeRemoteBrokerAuthenticationMethodsArgs()
        {
        }
        public static new MqttBridgeRemoteBrokerAuthenticationMethodsArgs Empty => new MqttBridgeRemoteBrokerAuthenticationMethodsArgs();
    }
}
