// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// Defines how the Video Analyzer account is (optionally) encrypted.
    /// </summary>
    [OutputType]
    public sealed class AccountEncryptionResponse
    {
        /// <summary>
        /// The Key Vault identity.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponse? Identity;
        /// <summary>
        /// The properties of the key used to encrypt the account.
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponse? KeyVaultProperties;
        /// <summary>
        /// The current status of the Key Vault mapping.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of key used to encrypt the Account Key.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccountEncryptionResponse(
            Outputs.ResourceIdentityResponse? identity,

            Outputs.KeyVaultPropertiesResponse? keyVaultProperties,

            string status,

            string type)
        {
            Identity = identity;
            KeyVaultProperties = keyVaultProperties;
            Status = status;
            Type = type;
        }
    }
}
