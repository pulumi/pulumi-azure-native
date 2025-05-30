// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataLakeStore.Inputs
{

    /// <summary>
    /// The encryption configuration for the account.
    /// </summary>
    public sealed class EncryptionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Key Vault information for connecting to user managed encryption keys.
        /// </summary>
        [Input("keyVaultMetaInfo")]
        public Input<Inputs.KeyVaultMetaInfoArgs>? KeyVaultMetaInfo { get; set; }

        /// <summary>
        /// The type of encryption configuration being used. Currently the only supported types are 'UserManaged' and 'ServiceManaged'.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AzureNative.DataLakeStore.EncryptionConfigType> Type { get; set; } = null!;

        public EncryptionConfigArgs()
        {
        }
        public static new EncryptionConfigArgs Empty => new EncryptionConfigArgs();
    }
}
