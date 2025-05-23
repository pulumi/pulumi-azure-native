// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Outputs
{

    /// <summary>
    /// Key Vault input properties for encryption.
    /// </summary>
    [OutputType]
    public sealed class ManagedDiskEncryptionResponseKeyVaultProperties
    {
        /// <summary>
        /// The name of KeyVault key.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The URI of KeyVault.
        /// </summary>
        public readonly string KeyVaultUri;
        /// <summary>
        /// The version of KeyVault key.
        /// </summary>
        public readonly string KeyVersion;

        [OutputConstructor]
        private ManagedDiskEncryptionResponseKeyVaultProperties(
            string keyName,

            string keyVaultUri,

            string keyVersion)
        {
            KeyName = keyName;
            KeyVaultUri = keyVaultUri;
            KeyVersion = keyVersion;
        }
    }
}
