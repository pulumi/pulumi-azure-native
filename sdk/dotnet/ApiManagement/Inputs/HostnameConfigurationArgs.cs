// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Inputs
{

    /// <summary>
    /// Custom hostname configuration.
    /// </summary>
    public sealed class HostnameConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate information.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertificateInformationArgs>? Certificate { get; set; }

        /// <summary>
        /// Certificate Password.
        /// </summary>
        [Input("certificatePassword")]
        public Input<string>? CertificatePassword { get; set; }

        /// <summary>
        /// Certificate Source.
        /// </summary>
        [Input("certificateSource")]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.CertificateSource>? CertificateSource { get; set; }

        /// <summary>
        /// Certificate Status.
        /// </summary>
        [Input("certificateStatus")]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.CertificateStatus>? CertificateStatus { get; set; }

        /// <summary>
        /// Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate. If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The setting only applied to gateway Hostname Type.
        /// </summary>
        [Input("defaultSslBinding")]
        public Input<bool>? DefaultSslBinding { get; set; }

        /// <summary>
        /// Base64 Encoded certificate.
        /// </summary>
        [Input("encodedCertificate")]
        public Input<string>? EncodedCertificate { get; set; }

        /// <summary>
        /// Hostname to configure on the Api Management service.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
        /// </summary>
        [Input("identityClientId")]
        public Input<string>? IdentityClientId { get; set; }

        /// <summary>
        /// Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided, auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi. The secret should be of type *application/x-pkcs12*
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// Specify true to always negotiate client certificate on the hostname. Default Value is false.
        /// </summary>
        [Input("negotiateClientCertificate")]
        public Input<bool>? NegotiateClientCertificate { get; set; }

        /// <summary>
        /// Hostname type.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.HostnameType> Type { get; set; } = null!;

        public HostnameConfigurationArgs()
        {
            DefaultSslBinding = false;
            NegotiateClientCertificate = false;
        }
        public static new HostnameConfigurationArgs Empty => new HostnameConfigurationArgs();
    }
}
