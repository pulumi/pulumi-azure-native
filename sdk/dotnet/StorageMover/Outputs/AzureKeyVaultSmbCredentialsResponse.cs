// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover.Outputs
{

    /// <summary>
    /// The Azure Key Vault secret URIs which store the credentials.
    /// </summary>
    [OutputType]
    public sealed class AzureKeyVaultSmbCredentialsResponse
    {
        /// <summary>
        /// The Azure Key Vault secret URI which stores the password. Use empty string to clean-up existing value.
        /// </summary>
        public readonly string? PasswordUri;
        /// <summary>
        /// The Credentials type.
        /// Expected value is 'AzureKeyVaultSmb'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The Azure Key Vault secret URI which stores the username. Use empty string to clean-up existing value.
        /// </summary>
        public readonly string? UsernameUri;

        [OutputConstructor]
        private AzureKeyVaultSmbCredentialsResponse(
            string? passwordUri,

            string type,

            string? usernameUri)
        {
            PasswordUri = passwordUri;
            Type = type;
            UsernameUri = usernameUri;
        }
    }
}
