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
    /// Definition of AbortIncompleteMultipartUpload
    /// </summary>
    [OutputType]
    public sealed class AbortIncompleteMultipartUploadResponse
    {
        /// <summary>
        /// Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.
        /// </summary>
        public readonly int? DaysAfterInitiation;

        [OutputConstructor]
        private AbortIncompleteMultipartUploadResponse(int? daysAfterInitiation)
        {
            DaysAfterInitiation = daysAfterInitiation;
        }
    }
}
