// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20200101Preview.Inputs
{

    /// <summary>
    /// Client Certificate definition.
    /// </summary>
    public sealed class ClientCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate Common name.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// Whether the certificate is admin or not.
        /// </summary>
        [Input("isAdmin", required: true)]
        public Input<bool> IsAdmin { get; set; } = null!;

        /// <summary>
        /// Issuer thumbprint for the certificate. Only used together with CommonName.
        /// </summary>
        [Input("issuerThumbprint")]
        public Input<string>? IssuerThumbprint { get; set; }

        /// <summary>
        /// Certificate Thumbprint.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public ClientCertificateArgs()
        {
        }
        public static new ClientCertificateArgs Empty => new ClientCertificateArgs();
    }
}
