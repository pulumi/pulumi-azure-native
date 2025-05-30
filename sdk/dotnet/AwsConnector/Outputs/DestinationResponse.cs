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
    /// Definition of Destination
    /// </summary>
    [OutputType]
    public sealed class DestinationResponse
    {
        /// <summary>
        /// The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not validated before exporting data.   Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
        /// </summary>
        public readonly string? BucketAccountId;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the bucket to which data is exported.
        /// </summary>
        public readonly string? BucketArn;
        /// <summary>
        /// Specifies the file format used when exporting data to Amazon S3.  *Allowed values*: ``CSV`` | ``ORC`` | ``Parquet``
        /// </summary>
        public readonly string? Format;
        /// <summary>
        /// The prefix to use when exporting data. The prefix is prepended to all results.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private DestinationResponse(
            string? bucketAccountId,

            string? bucketArn,

            string? format,

            string? prefix)
        {
            BucketAccountId = bucketAccountId;
            BucketArn = bucketArn;
            Format = format;
            Prefix = prefix;
        }
    }
}
