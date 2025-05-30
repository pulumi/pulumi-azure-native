// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CertificateRegistration.Outputs
{

    /// <summary>
    /// SSL certificate details.
    /// </summary>
    [OutputType]
    public sealed class CertificateDetailsResponse
    {
        /// <summary>
        /// Certificate Issuer.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Date Certificate is valid to.
        /// </summary>
        public readonly string NotAfter;
        /// <summary>
        /// Date Certificate is valid from.
        /// </summary>
        public readonly string NotBefore;
        /// <summary>
        /// Raw certificate data.
        /// </summary>
        public readonly string RawData;
        /// <summary>
        /// Certificate Serial Number.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// Certificate Signature algorithm.
        /// </summary>
        public readonly string SignatureAlgorithm;
        /// <summary>
        /// Certificate Subject.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// Certificate Thumbprint.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// Certificate Version.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private CertificateDetailsResponse(
            string issuer,

            string notAfter,

            string notBefore,

            string rawData,

            string serialNumber,

            string signatureAlgorithm,

            string subject,

            string thumbprint,

            int version)
        {
            Issuer = issuer;
            NotAfter = notAfter;
            NotBefore = notBefore;
            RawData = rawData;
            SerialNumber = serialNumber;
            SignatureAlgorithm = signatureAlgorithm;
            Subject = subject;
            Thumbprint = thumbprint;
            Version = version;
        }
    }
}
