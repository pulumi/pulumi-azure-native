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
    /// Customer Certificate used for https
    /// </summary>
    [OutputType]
    public sealed class CustomerCertificateParametersResponse
    {
        /// <summary>
        /// Certificate issuing authority.
        /// </summary>
        public readonly string CertificateAuthority;
        /// <summary>
        /// Certificate expiration date.
        /// </summary>
        public readonly string ExpirationDate;
        /// <summary>
        /// Resource reference to the Azure Key Vault certificate. Expected to be in format of /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{certificateName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
        /// </summary>
        public readonly Outputs.ResourceReferenceResponse SecretSource;
        /// <summary>
        /// Version of the secret to be used
        /// </summary>
        public readonly string? SecretVersion;
        /// <summary>
        /// Subject name in the certificate.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The list of SANs.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAlternativeNames;
        /// <summary>
        /// Certificate thumbprint.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The type of the secret resource.
        /// Expected value is 'CustomerCertificate'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Whether to use the latest version for the certificate
        /// </summary>
        public readonly bool? UseLatestVersion;

        [OutputConstructor]
        private CustomerCertificateParametersResponse(
            string certificateAuthority,

            string expirationDate,

            Outputs.ResourceReferenceResponse secretSource,

            string? secretVersion,

            string subject,

            ImmutableArray<string> subjectAlternativeNames,

            string thumbprint,

            string type,

            bool? useLatestVersion)
        {
            CertificateAuthority = certificateAuthority;
            ExpirationDate = expirationDate;
            SecretSource = secretSource;
            SecretVersion = secretVersion;
            Subject = subject;
            SubjectAlternativeNames = subjectAlternativeNames;
            Thumbprint = thumbprint;
            Type = type;
            UseLatestVersion = useLatestVersion;
        }
    }
}
