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
    /// Mqtt to Kafka route properties
    /// </summary>
    [OutputType]
    public sealed class MqttToKafkaRoutesResponse
    {
        /// <summary>
        /// The kafka acks to use.
        /// </summary>
        public readonly string KafkaAcks;
        /// <summary>
        /// The kafka topic to publish to.
        /// </summary>
        public readonly string KafkaTopic;
        /// <summary>
        /// The mqtt topic to pull from.
        /// </summary>
        public readonly string MqttTopic;
        /// <summary>
        /// The name of the route.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The qos to use for mqtt.
        /// </summary>
        public readonly int? Qos;
        /// <summary>
        /// The properties for shared subscription.
        /// </summary>
        public readonly Outputs.KafkaSharedSubscriptionPropertiesResponse? SharedSubscription;

        [OutputConstructor]
        private MqttToKafkaRoutesResponse(
            string kafkaAcks,

            string kafkaTopic,

            string mqttTopic,

            string name,

            int? qos,

            Outputs.KafkaSharedSubscriptionPropertiesResponse? sharedSubscription)
        {
            KafkaAcks = kafkaAcks;
            KafkaTopic = kafkaTopic;
            MqttTopic = mqttTopic;
            Name = name;
            Qos = qos;
            SharedSubscription = sharedSubscription;
        }
    }
}
