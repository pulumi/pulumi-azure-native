// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Inputs
{

    /// <summary>
    /// Properties for a certificate stored in a Key Vault.
    /// </summary>
    public sealed class CertificateKeyVaultPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of a managed identity to authenticate with Azure Key Vault, or System to use a system-assigned identity.
        /// </summary>
        [Input("identity")]
        public Input<string>? Identity { get; set; }

        /// <summary>
        /// URL pointing to the Azure Key Vault secret that holds the certificate.
        /// </summary>
        [Input("keyVaultUrl")]
        public Input<string>? KeyVaultUrl { get; set; }

        public CertificateKeyVaultPropertiesArgs()
        {
        }
        public static new CertificateKeyVaultPropertiesArgs Empty => new CertificateKeyVaultPropertiesArgs();
    }
}
