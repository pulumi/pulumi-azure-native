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
    /// Definition of ReplicaRegion
    /// </summary>
    [OutputType]
    public sealed class ReplicaRegionResponse
    {
        /// <summary>
        /// The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// A string that represents a ``Region``, for example 'us-east-1'.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private ReplicaRegionResponse(
            string? kmsKeyId,

            string? region)
        {
            KmsKeyId = kmsKeyId;
            Region = region;
        }
    }
}
