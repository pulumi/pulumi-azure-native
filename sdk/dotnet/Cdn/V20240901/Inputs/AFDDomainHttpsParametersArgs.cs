// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.V20240901.Inputs
{

    /// <summary>
    /// The JSON object that contains the properties to secure a domain.
    /// </summary>
    public sealed class AFDDomainHttpsParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the source of the SSL certificate.
        /// </summary>
        [Input("certificateType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cdn.V20240901.AfdCertificateType> CertificateType { get; set; } = null!;

        /// <summary>
        /// TLS protocol version that will be used for Https
        /// </summary>
        [Input("minimumTlsVersion")]
        public Input<Pulumi.AzureNative.Cdn.V20240901.AfdMinimumTlsVersion>? MinimumTlsVersion { get; set; }

        /// <summary>
        /// Resource reference to the secret. ie. subs/rg/profile/secret
        /// </summary>
        [Input("secret")]
        public Input<Inputs.ResourceReferenceArgs>? Secret { get; set; }

        public AFDDomainHttpsParametersArgs()
        {
        }
        public static new AFDDomainHttpsParametersArgs Empty => new AFDDomainHttpsParametersArgs();
    }
}
