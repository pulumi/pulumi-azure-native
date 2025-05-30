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
    /// Kafka RemoteBrokerConnection Authentication methods
    /// </summary>
    public sealed class KafkaRemoteBrokerAuthenticationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of authentication to use for Kafka remote broker.
        /// </summary>
        [Input("authType")]
        public Input<Inputs.KafkaRemoteBrokerAuthenticationTypesArgs>? AuthType { get; set; }

        /// <summary>
        /// If authentication is enabled for Kafka remote broker.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public KafkaRemoteBrokerAuthenticationPropertiesArgs()
        {
        }
        public static new KafkaRemoteBrokerAuthenticationPropertiesArgs Empty => new KafkaRemoteBrokerAuthenticationPropertiesArgs();
    }
}
