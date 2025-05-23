// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AgFoodPlatform.Outputs
{

    /// <summary>
    /// Properties of the key vault.
    /// </summary>
    [OutputType]
    public sealed class KeyVaultPropertiesResponse
    {
        /// <summary>
        /// Name of Key Vault key.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// Uri of the key vault.
        /// </summary>
        public readonly string KeyVaultUri;
        /// <summary>
        /// Version of Key Vault key.
        /// </summary>
        public readonly string KeyVersion;

        [OutputConstructor]
        private KeyVaultPropertiesResponse(
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
