// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Gets or sets the PrometheusOS provider properties.
    /// </summary>
    [OutputType]
    public sealed class PrometheusOsProviderInstancePropertiesResponse
    {
        /// <summary>
        /// URL of the Node Exporter endpoint
        /// </summary>
        public readonly string? PrometheusUrl;
        /// <summary>
        /// The provider type. For example, the value can be SapHana.
        /// Expected value is 'PrometheusOS'.
        /// </summary>
        public readonly string ProviderType;
        /// <summary>
        /// Gets or sets the SAP System Identifier
        /// </summary>
        public readonly string? SapSid;
        /// <summary>
        /// Gets or sets the blob URI to SSL certificate for the prometheus node exporter.
        /// </summary>
        public readonly string? SslCertificateUri;
        /// <summary>
        /// Gets or sets certificate preference if secure communication is enabled.
        /// </summary>
        public readonly string? SslPreference;

        [OutputConstructor]
        private PrometheusOsProviderInstancePropertiesResponse(
            string? prometheusUrl,

            string providerType,

            string? sapSid,

            string? sslCertificateUri,

            string? sslPreference)
        {
            PrometheusUrl = prometheusUrl;
            ProviderType = providerType;
            SapSid = sapSid;
            SslCertificateUri = sslCertificateUri;
            SslPreference = sslPreference;
        }
    }
}
