// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.V20240601Preview.Outputs
{

    /// <summary>
    /// Server-side certificate used for mTLS validation
    /// </summary>
    [OutputType]
    public sealed class AfdSecretMtlsCertificateChainResponse
    {
        /// <summary>
        /// Soonest expiration date among certificates in customer's certificate chain in ISO 8601 compliant format yyyy-MM-ddTHH:mm:ss.fffffffK in UTC
        /// </summary>
        public readonly string ExpirationDate;
        /// <summary>
        /// Resource reference to the Azure Key Vault secret. Expected to be in format of /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}
        /// </summary>
        public readonly Outputs.ResourceReferenceResponse SecretSource;
        /// <summary>
        /// Version of the secret to be used
        /// </summary>
        public readonly string SecretVersion;
        /// <summary>
        /// The type of the secret resource.
        /// Expected value is 'MtlsCertificateChain'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AfdSecretMtlsCertificateChainResponse(
            string expirationDate,

            Outputs.ResourceReferenceResponse secretSource,

            string secretVersion,

            string type)
        {
            ExpirationDate = expirationDate;
            SecretSource = secretSource;
            SecretVersion = secretVersion;
            Type = type;
        }
    }
}
