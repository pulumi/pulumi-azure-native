// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// Custom hostname configuration.
    /// </summary>
    [OutputType]
    public sealed class HostnameConfigurationResponse
    {
        /// <summary>
        /// Certificate information.
        /// </summary>
        public readonly Outputs.CertificateInformationResponse? Certificate;
        /// <summary>
        /// Certificate Password.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// Certificate Source.
        /// </summary>
        public readonly string? CertificateSource;
        /// <summary>
        /// Certificate Status.
        /// </summary>
        public readonly string? CertificateStatus;
        /// <summary>
        /// Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate. If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The setting only applied to gateway Hostname Type.
        /// </summary>
        public readonly bool? DefaultSslBinding;
        /// <summary>
        /// Base64 Encoded certificate.
        /// </summary>
        public readonly string? EncodedCertificate;
        /// <summary>
        /// Hostname to configure on the Api Management service.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
        /// </summary>
        public readonly string? IdentityClientId;
        /// <summary>
        /// Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided, auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi. The secret should be of type *application/x-pkcs12*
        /// </summary>
        public readonly string? KeyVaultId;
        /// <summary>
        /// Specify true to always negotiate client certificate on the hostname. Default Value is false.
        /// </summary>
        public readonly bool? NegotiateClientCertificate;
        /// <summary>
        /// Hostname type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private HostnameConfigurationResponse(
            Outputs.CertificateInformationResponse? certificate,

            string? certificatePassword,

            string? certificateSource,

            string? certificateStatus,

            bool? defaultSslBinding,

            string? encodedCertificate,

            string hostName,

            string? identityClientId,

            string? keyVaultId,

            bool? negotiateClientCertificate,

            string type)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            CertificateSource = certificateSource;
            CertificateStatus = certificateStatus;
            DefaultSslBinding = defaultSslBinding;
            EncodedCertificate = encodedCertificate;
            HostName = hostName;
            IdentityClientId = identityClientId;
            KeyVaultId = keyVaultId;
            NegotiateClientCertificate = negotiateClientCertificate;
            Type = type;
        }
    }
}
