// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover.Inputs
{

    /// <summary>
    /// The Azure Key Vault secret URIs which store the credentials.
    /// </summary>
    public sealed class AzureKeyVaultSmbCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Key Vault secret URI which stores the password. Use empty string to clean-up existing value.
        /// </summary>
        [Input("passwordUri")]
        public Input<string>? PasswordUri { get; set; }

        /// <summary>
        /// The Credentials type.
        /// Expected value is 'AzureKeyVaultSmb'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The Azure Key Vault secret URI which stores the username. Use empty string to clean-up existing value.
        /// </summary>
        [Input("usernameUri")]
        public Input<string>? UsernameUri { get; set; }

        public AzureKeyVaultSmbCredentialsArgs()
        {
        }
        public static new AzureKeyVaultSmbCredentialsArgs Empty => new AzureKeyVaultSmbCredentialsArgs();
    }
}
