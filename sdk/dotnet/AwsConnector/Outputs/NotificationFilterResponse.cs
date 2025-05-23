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
    /// Definition of NotificationFilter
    /// </summary>
    [OutputType]
    public sealed class NotificationFilterResponse
    {
        /// <summary>
        /// A container for object key name prefix and suffix filtering rules. A container for object key name prefix and suffix filtering rules. For more information about object key name filtering, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html) in the *Amazon S3 User Guide*.  The same type of filter rule cannot be used more than once. For example, you cannot specify two prefix rules.
        /// </summary>
        public readonly Outputs.S3KeyFilterResponse? S3Key;

        [OutputConstructor]
        private NotificationFilterResponse(Outputs.S3KeyFilterResponse? s3Key)
        {
            S3Key = s3Key;
        }
    }
}
