// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs.Inputs
{

    /// <summary>
    /// Description of a NotificationHub MpnsCredential.
    /// </summary>
    public sealed class MpnsCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the certificate key for this credential.
        /// </summary>
        [Input("certificateKey", required: true)]
        public Input<string> CertificateKey { get; set; } = null!;

        /// <summary>
        /// Gets or sets the MPNS certificate.
        /// </summary>
        [Input("mpnsCertificate", required: true)]
        public Input<string> MpnsCertificate { get; set; } = null!;

        /// <summary>
        /// Gets or sets the MPNS certificate Thumbprint
        /// </summary>
        [Input("thumbprint", required: true)]
        public Input<string> Thumbprint { get; set; } = null!;

        public MpnsCredentialArgs()
        {
        }
        public static new MpnsCredentialArgs Empty => new MpnsCredentialArgs();
    }
}
