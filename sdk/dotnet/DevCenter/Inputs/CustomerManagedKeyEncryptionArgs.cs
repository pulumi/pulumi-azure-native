// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.Inputs
{

    /// <summary>
    /// All Customer-managed key encryption properties for the resource.
    /// </summary>
    public sealed class CustomerManagedKeyEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
        /// </summary>
        [Input("keyEncryptionKeyIdentity")]
        public Input<Inputs.CustomerManagedKeyEncryptionKeyEncryptionKeyIdentityArgs>? KeyEncryptionKeyIdentity { get; set; }

        /// <summary>
        /// key encryption key Url, versioned or non-versioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78 or https://contosovault.vault.azure.net/keys/contosokek.
        /// </summary>
        [Input("keyEncryptionKeyUrl")]
        public Input<string>? KeyEncryptionKeyUrl { get; set; }

        public CustomerManagedKeyEncryptionArgs()
        {
        }
        public static new CustomerManagedKeyEncryptionArgs Empty => new CustomerManagedKeyEncryptionArgs();
    }
}
