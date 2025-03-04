// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20240801Preview.Outputs
{

    /// <summary>
    /// The cloudwatch to kinesis connection configuration
    /// </summary>
    [OutputType]
    public sealed class DefenderForContainersAwsOfferingResponseCloudWatchToKinesis
    {
        /// <summary>
        /// The cloud role ARN in AWS used by CloudWatch to transfer data into Kinesis
        /// </summary>
        public readonly string? CloudRoleArn;

        [OutputConstructor]
        private DefenderForContainersAwsOfferingResponseCloudWatchToKinesis(string? cloudRoleArn)
        {
            CloudRoleArn = cloudRoleArn;
        }
    }
}
