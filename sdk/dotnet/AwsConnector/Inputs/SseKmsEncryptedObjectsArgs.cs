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
    /// Definition of SseKmsEncryptedObjects
    /// </summary>
    public sealed class SseKmsEncryptedObjectsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.SseKmsEncryptedObjectsStatus>? Status { get; set; }

        public SseKmsEncryptedObjectsArgs()
        {
        }
        public static new SseKmsEncryptedObjectsArgs Empty => new SseKmsEncryptedObjectsArgs();
    }
}
