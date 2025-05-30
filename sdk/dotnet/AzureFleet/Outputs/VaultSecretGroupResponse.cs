// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Describes a set of certificates which are all in the same Key Vault.
    /// </summary>
    [OutputType]
    public sealed class VaultSecretGroupResponse
    {
        /// <summary>
        /// The relative URL of the Key Vault containing all of the certificates in
        /// VaultCertificates.
        /// </summary>
        public readonly Outputs.SubResourceResponse? SourceVault;
        /// <summary>
        /// The list of key vault references in SourceVault which contain certificates.
        /// </summary>
        public readonly ImmutableArray<Outputs.VaultCertificateResponse> VaultCertificates;

        [OutputConstructor]
        private VaultSecretGroupResponse(
            Outputs.SubResourceResponse? sourceVault,

            ImmutableArray<Outputs.VaultCertificateResponse> vaultCertificates)
        {
            SourceVault = sourceVault;
            VaultCertificates = vaultCertificates;
        }
    }
}
