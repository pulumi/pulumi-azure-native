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
    /// Definition of LoggingConfig
    /// </summary>
    public sealed class LoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set this property to filter the application logs for your function that Lambda sends to CloudWatch. Lambda only sends application logs at the selected level of detail and lower, where ``TRACE`` is the highest level and ``FATAL`` is the lowest.
        /// </summary>
        [Input("applicationLogLevel")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.LoggingConfigApplicationLogLevel>? ApplicationLogLevel { get; set; }

        /// <summary>
        /// Property failureFeedbackRoleArn
        /// </summary>
        [Input("failureFeedbackRoleArn")]
        public Input<string>? FailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The format in which Lambda sends your function's application and system logs to CloudWatch. Select between plain text and structured JSON.
        /// </summary>
        [Input("logFormat")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.LoggingConfigLogFormat>? LogFormat { get; set; }

        /// <summary>
        /// The name of the Amazon CloudWatch log group the function sends logs to. By default, Lambda functions send logs to a default log group named ``/aws/lambda/&lt;function name&gt;``. To use a different log group, enter an existing log group or enter a new log group name.
        /// </summary>
        [Input("logGroup")]
        public Input<string>? LogGroup { get; set; }

        /// <summary>
        /// Property protocol
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.LoggingConfigProtocol>? Protocol { get; set; }

        /// <summary>
        /// Property successFeedbackRoleArn
        /// </summary>
        [Input("successFeedbackRoleArn")]
        public Input<string>? SuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Property successFeedbackSampleRate
        /// </summary>
        [Input("successFeedbackSampleRate")]
        public Input<string>? SuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
        /// </summary>
        [Input("systemLogLevel")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.LoggingConfigSystemLogLevel>? SystemLogLevel { get; set; }

        public LoggingConfigArgs()
        {
        }
        public static new LoggingConfigArgs Empty => new LoggingConfigArgs();
    }
}
