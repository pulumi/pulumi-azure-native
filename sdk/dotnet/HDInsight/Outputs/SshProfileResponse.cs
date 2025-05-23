// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// The list of SSH public keys.
    /// </summary>
    [OutputType]
    public sealed class SshProfileResponse
    {
        /// <summary>
        /// The list of SSH public keys.
        /// </summary>
        public readonly ImmutableArray<Outputs.SshPublicKeyResponse> PublicKeys;

        [OutputConstructor]
        private SshProfileResponse(ImmutableArray<Outputs.SshPublicKeyResponse> publicKeys)
        {
            PublicKeys = publicKeys;
        }
    }
}
