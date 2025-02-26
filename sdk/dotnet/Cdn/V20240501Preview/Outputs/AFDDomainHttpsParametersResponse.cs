// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.V20240501Preview.Outputs
{

    /// <summary>
    /// The JSON object that contains the properties to secure a domain.
    /// </summary>
    [OutputType]
    public sealed class AFDDomainHttpsParametersResponse
    {
        /// <summary>
        /// Defines the source of the SSL certificate.
        /// </summary>
        public readonly string CertificateType;
        /// <summary>
        /// cipher suite set type that will be used for Https
        /// </summary>
        public readonly string? CipherSuiteSetType;
        /// <summary>
        /// Customized cipher suites object that will be used for Https when cipherSuiteSetType is Customized.
        /// </summary>
        public readonly Outputs.AFDDomainHttpsCustomizedCipherSuiteSetResponse? CustomizedCipherSuiteSet;
        /// <summary>
        /// TLS protocol version that will be used for Https when cipherSuiteSetType is Customized.
        /// </summary>
        public readonly string? MinimumTlsVersion;
        /// <summary>
        /// Resource reference to the secret. ie. subs/rg/profile/secret
        /// </summary>
        public readonly Outputs.ResourceReferenceResponse? Secret;

        [OutputConstructor]
        private AFDDomainHttpsParametersResponse(
            string certificateType,

            string? cipherSuiteSetType,

            Outputs.AFDDomainHttpsCustomizedCipherSuiteSetResponse? customizedCipherSuiteSet,

            string? minimumTlsVersion,

            Outputs.ResourceReferenceResponse? secret)
        {
            CertificateType = certificateType;
            CipherSuiteSetType = cipherSuiteSetType;
            CustomizedCipherSuiteSet = customizedCipherSuiteSet;
            MinimumTlsVersion = minimumTlsVersion;
            Secret = secret;
        }
    }
}
