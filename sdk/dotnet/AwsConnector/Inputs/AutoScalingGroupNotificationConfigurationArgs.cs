// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of NotificationConfiguration
    /// </summary>
    public sealed class AutoScalingGroupNotificationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("notificationTypes")]
        private InputList<string>? _notificationTypes;

        /// <summary>
        /// A list of event types that send a notification. Event types can include any of the following types.  *Allowed values*:  +   ``autoscaling:EC2_INSTANCE_LAUNCH``   +   ``autoscaling:EC2_INSTANCE_LAUNCH_ERROR``   +   ``autoscaling:EC2_INSTANCE_TERMINATE``   +   ``autoscaling:EC2_INSTANCE_TERMINATE_ERROR``   +   ``autoscaling:TEST_NOTIFICATION``
        /// </summary>
        public InputList<string> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<string>());
            set => _notificationTypes = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon SNS topic.
        /// </summary>
        [Input("topicARN")]
        public Input<string>? TopicARN { get; set; }

        public AutoScalingGroupNotificationConfigurationArgs()
        {
        }
        public static new AutoScalingGroupNotificationConfigurationArgs Empty => new AutoScalingGroupNotificationConfigurationArgs();
    }
}
