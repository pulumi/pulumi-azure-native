// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Inputs
{

    /// <summary>
    /// Defines how the Video Analyzer account is (optionally) encrypted.
    /// </summary>
    public sealed class AccountEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Key Vault identity.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The properties of the key used to encrypt the account.
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        /// <summary>
        /// The type of key used to encrypt the Account Key.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.VideoAnalyzer.AccountEncryptionKeyType> Type { get; set; } = null!;

        public AccountEncryptionArgs()
        {
        }
        public static new AccountEncryptionArgs Empty => new AccountEncryptionArgs();
    }
}
