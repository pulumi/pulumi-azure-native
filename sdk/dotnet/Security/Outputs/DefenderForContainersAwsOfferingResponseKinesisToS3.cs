// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// The kinesis to s3 connection configuration
    /// </summary>
    [OutputType]
    public sealed class DefenderForContainersAwsOfferingResponseKinesisToS3
    {
        /// <summary>
        /// The cloud role ARN in AWS used by Kinesis to transfer data into S3
        /// </summary>
        public readonly string? CloudRoleArn;

        [OutputConstructor]
        private DefenderForContainersAwsOfferingResponseKinesisToS3(string? cloudRoleArn)
        {
            CloudRoleArn = cloudRoleArn;
        }
    }
}
