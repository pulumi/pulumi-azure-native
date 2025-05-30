// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Outputs
{

    /// <summary>
    /// Url signing key parameters
    /// </summary>
    [OutputType]
    public sealed class UrlSigningKeyParametersResponse
    {
        /// <summary>
        /// Defines the customer defined key Id. This id will exist in the incoming request to indicate the key used to form the hash.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Resource reference to the Azure Key Vault secret. Expected to be in format of /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{secretName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
        /// </summary>
        public readonly Outputs.ResourceReferenceResponse SecretSource;
        /// <summary>
        /// Version of the secret to be used
        /// </summary>
        public readonly string SecretVersion;
        /// <summary>
        /// The type of the secret resource.
        /// Expected value is 'UrlSigningKey'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private UrlSigningKeyParametersResponse(
            string keyId,

            Outputs.ResourceReferenceResponse secretSource,

            string secretVersion,

            string type)
        {
            KeyId = keyId;
            SecretSource = secretSource;
            SecretVersion = secretVersion;
            Type = type;
        }
    }
}
