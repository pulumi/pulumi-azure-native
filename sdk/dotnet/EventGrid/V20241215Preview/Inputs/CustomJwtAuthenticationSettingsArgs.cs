// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20241215Preview.Inputs
{

    /// <summary>
    /// Custom JWT authentication settings for namespace resource.
    /// </summary>
    public sealed class CustomJwtAuthenticationSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("encodedIssuerCertificates")]
        private InputList<Inputs.EncodedIssuerCertificateInfoArgs>? _encodedIssuerCertificates;

        /// <summary>
        /// Information about the encoded public certificates that are used for custom authentication.
        /// </summary>
        public InputList<Inputs.EncodedIssuerCertificateInfoArgs> EncodedIssuerCertificates
        {
            get => _encodedIssuerCertificates ?? (_encodedIssuerCertificates = new InputList<Inputs.EncodedIssuerCertificateInfoArgs>());
            set => _encodedIssuerCertificates = value;
        }

        [Input("issuerCertificates")]
        private InputList<Inputs.IssuerCertificateInfoArgs>? _issuerCertificates;

        /// <summary>
        /// Information about the certificates that are used for token validation. We currently support maximum 2 certificates.
        /// </summary>
        public InputList<Inputs.IssuerCertificateInfoArgs> IssuerCertificates
        {
            get => _issuerCertificates ?? (_issuerCertificates = new InputList<Inputs.IssuerCertificateInfoArgs>());
            set => _issuerCertificates = value;
        }

        /// <summary>
        /// Expected JWT token issuer.
        /// </summary>
        [Input("tokenIssuer")]
        public Input<string>? TokenIssuer { get; set; }

        public CustomJwtAuthenticationSettingsArgs()
        {
        }
        public static new CustomJwtAuthenticationSettingsArgs Empty => new CustomJwtAuthenticationSettingsArgs();
    }
}
