// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// The configuration settings of the Azure Active Directory app registration.
    /// </summary>
    public sealed class AzureActiveDirectoryRegistrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of this relying party application, known as the client_id.
        /// This setting is required for enabling OpenID Connection authentication with Azure Active Directory or 
        /// other 3rd party OpenID Connect providers.
        /// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// An alternative to the client secret thumbprint, that is the issuer of a certificate used for signing purposes. This property acts as
        /// a replacement for the Client Secret Certificate Thumbprint. It is also optional.
        /// </summary>
        [Input("clientSecretCertificateIssuer")]
        public Input<string>? ClientSecretCertificateIssuer { get; set; }

        /// <summary>
        /// An alternative to the client secret thumbprint, that is the subject alternative name of a certificate used for signing purposes. This property acts as
        /// a replacement for the Client Secret Certificate Thumbprint. It is also optional.
        /// </summary>
        [Input("clientSecretCertificateSubjectAlternativeName")]
        public Input<string>? ClientSecretCertificateSubjectAlternativeName { get; set; }

        /// <summary>
        /// An alternative to the client secret, that is the thumbprint of a certificate used for signing purposes. This property acts as
        /// a replacement for the Client Secret. It is also optional.
        /// </summary>
        [Input("clientSecretCertificateThumbprint")]
        public Input<string>? ClientSecretCertificateThumbprint { get; set; }

        /// <summary>
        /// The app setting name that contains the client secret of the relying party application.
        /// </summary>
        [Input("clientSecretSettingName")]
        public Input<string>? ClientSecretSettingName { get; set; }

        /// <summary>
        /// The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application.
        /// When using Azure Active Directory, this value is the URI of the directory tenant, e.g. `https://login.microsoftonline.com/v2.0/{tenant-guid}/`.
        /// This URI is a case-sensitive identifier for the token issuer.
        /// More information on OpenID Connect Discovery: http://openid.net/specs/openid-connect-discovery-1_0.html
        /// </summary>
        [Input("openIdIssuer")]
        public Input<string>? OpenIdIssuer { get; set; }

        public AzureActiveDirectoryRegistrationArgs()
        {
        }
        public static new AzureActiveDirectoryRegistrationArgs Empty => new AzureActiveDirectoryRegistrationArgs();
    }
}
