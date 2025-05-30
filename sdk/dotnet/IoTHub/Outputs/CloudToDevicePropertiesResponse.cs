// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTHub.Outputs
{

    /// <summary>
    /// The IoT hub cloud-to-device messaging properties.
    /// </summary>
    [OutputType]
    public sealed class CloudToDevicePropertiesResponse
    {
        /// <summary>
        /// The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        public readonly string? DefaultTtlAsIso8601;
        /// <summary>
        /// The properties of the feedback queue for cloud-to-device messages.
        /// </summary>
        public readonly Outputs.FeedbackPropertiesResponse? Feedback;
        /// <summary>
        /// The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        public readonly int? MaxDeliveryCount;

        [OutputConstructor]
        private CloudToDevicePropertiesResponse(
            string? defaultTtlAsIso8601,

            Outputs.FeedbackPropertiesResponse? feedback,

            int? maxDeliveryCount)
        {
            DefaultTtlAsIso8601 = defaultTtlAsIso8601;
            Feedback = feedback;
            MaxDeliveryCount = maxDeliveryCount;
        }
    }
}
