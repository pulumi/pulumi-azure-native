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
    /// Definition of S3WordsList
    /// </summary>
    [OutputType]
    public sealed class S3WordsListResponse
    {
        /// <summary>
        /// Property bucketName
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// Property objectKey
        /// </summary>
        public readonly string? ObjectKey;

        [OutputConstructor]
        private S3WordsListResponse(
            string? bucketName,

            string? objectKey)
        {
            BucketName = bucketName;
            ObjectKey = objectKey;
        }
    }
}
