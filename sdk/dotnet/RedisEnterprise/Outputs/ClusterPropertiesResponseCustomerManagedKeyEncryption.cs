// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedisEnterprise.Outputs
{

    /// <summary>
    /// All Customer-managed key encryption properties for the resource. Set this to an empty object to use Microsoft-managed key encryption.
    /// </summary>
    [OutputType]
    public sealed class ClusterPropertiesResponseCustomerManagedKeyEncryption
    {
        /// <summary>
        /// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
        /// </summary>
        public readonly Outputs.ClusterPropertiesResponseKeyEncryptionKeyIdentity? KeyEncryptionKeyIdentity;
        /// <summary>
        /// Key encryption key Url, versioned only. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
        /// </summary>
        public readonly string? KeyEncryptionKeyUrl;

        [OutputConstructor]
        private ClusterPropertiesResponseCustomerManagedKeyEncryption(
            Outputs.ClusterPropertiesResponseKeyEncryptionKeyIdentity? keyEncryptionKeyIdentity,

            string? keyEncryptionKeyUrl)
        {
            KeyEncryptionKeyIdentity = keyEncryptionKeyIdentity;
            KeyEncryptionKeyUrl = keyEncryptionKeyUrl;
        }
    }
}
