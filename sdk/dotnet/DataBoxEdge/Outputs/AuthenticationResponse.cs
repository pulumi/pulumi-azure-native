// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Outputs
{

    /// <summary>
    /// Authentication mechanism for IoT devices.
    /// </summary>
    [OutputType]
    public sealed class AuthenticationResponse
    {
        /// <summary>
        /// Symmetric key for authentication.
        /// </summary>
        public readonly Outputs.SymmetricKeyResponse? SymmetricKey;

        [OutputConstructor]
        private AuthenticationResponse(Outputs.SymmetricKeyResponse? symmetricKey)
        {
            SymmetricKey = symmetricKey;
        }
    }
}
