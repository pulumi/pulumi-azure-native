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
    /// Symmetric key for authentication.
    /// </summary>
    [OutputType]
    public sealed class SymmetricKeyResponse
    {
        /// <summary>
        /// Connection string based on the symmetric key.
        /// </summary>
        public readonly Outputs.AsymmetricEncryptedSecretResponse? ConnectionString;

        [OutputConstructor]
        private SymmetricKeyResponse(Outputs.AsymmetricEncryptedSecretResponse? connectionString)
        {
            ConnectionString = connectionString;
        }
    }
}
