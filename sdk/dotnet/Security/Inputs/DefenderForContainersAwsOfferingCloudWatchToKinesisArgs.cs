// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The cloudwatch to kinesis connection configuration
    /// </summary>
    public sealed class DefenderForContainersAwsOfferingCloudWatchToKinesisArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cloud role ARN in AWS used by CloudWatch to transfer data into Kinesis
        /// </summary>
        [Input("cloudRoleArn")]
        public Input<string>? CloudRoleArn { get; set; }

        public DefenderForContainersAwsOfferingCloudWatchToKinesisArgs()
        {
        }
        public static new DefenderForContainersAwsOfferingCloudWatchToKinesisArgs Empty => new DefenderForContainersAwsOfferingCloudWatchToKinesisArgs();
    }
}
