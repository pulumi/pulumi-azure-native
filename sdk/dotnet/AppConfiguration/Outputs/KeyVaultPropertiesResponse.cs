// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppConfiguration.Outputs
{

    /// <summary>
    /// Settings concerning key vault encryption for a configuration store.
    /// </summary>
    [OutputType]
    public sealed class KeyVaultPropertiesResponse
    {
        /// <summary>
        /// The client id of the identity which will be used to access key vault.
        /// </summary>
        public readonly string? IdentityClientId;
        /// <summary>
        /// The URI of the key vault key used to encrypt data.
        /// </summary>
        public readonly string? KeyIdentifier;

        [OutputConstructor]
        private KeyVaultPropertiesResponse(
            string? identityClientId,

            string? keyIdentifier)
        {
            IdentityClientId = identityClientId;
            KeyIdentifier = keyIdentifier;
        }
    }
}
