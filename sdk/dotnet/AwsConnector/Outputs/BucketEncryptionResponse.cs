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
    /// Definition of BucketEncryption
    /// </summary>
    [OutputType]
    public sealed class BucketEncryptionResponse
    {
        /// <summary>
        /// Specifies the default server-side-encryption configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerSideEncryptionRuleResponse> ServerSideEncryptionConfiguration;

        [OutputConstructor]
        private BucketEncryptionResponse(ImmutableArray<Outputs.ServerSideEncryptionRuleResponse> serverSideEncryptionConfiguration)
        {
            ServerSideEncryptionConfiguration = serverSideEncryptionConfiguration;
        }
    }
}
