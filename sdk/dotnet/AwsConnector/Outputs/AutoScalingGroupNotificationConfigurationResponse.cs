// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of NotificationConfiguration
    /// </summary>
    [OutputType]
    public sealed class AutoScalingGroupNotificationConfigurationResponse
    {
        /// <summary>
        /// A list of event types that send a notification. Event types can include any of the following types.  *Allowed values*:  +   ``autoscaling:EC2_INSTANCE_LAUNCH``   +   ``autoscaling:EC2_INSTANCE_LAUNCH_ERROR``   +   ``autoscaling:EC2_INSTANCE_TERMINATE``   +   ``autoscaling:EC2_INSTANCE_TERMINATE_ERROR``   +   ``autoscaling:TEST_NOTIFICATION``
        /// </summary>
        public readonly ImmutableArray<string> NotificationTypes;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon SNS topic.
        /// </summary>
        public readonly string? TopicARN;

        [OutputConstructor]
        private AutoScalingGroupNotificationConfigurationResponse(
            ImmutableArray<string> notificationTypes,

            string? topicARN)
        {
            NotificationTypes = notificationTypes;
            TopicARN = topicARN;
        }
    }
}
