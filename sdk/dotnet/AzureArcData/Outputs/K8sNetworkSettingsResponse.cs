// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// The kubernetes network settings information.
    /// </summary>
    [OutputType]
    public sealed class K8sNetworkSettingsResponse
    {
        /// <summary>
        /// If 1, then SQL Server forces all connections to be encrypted. By default, this option is 0
        /// </summary>
        public readonly int? Forceencryption;
        /// <summary>
        /// Specifies which ciphers are allowed by SQL Server for TLS
        /// </summary>
        public readonly string? Tlsciphers;
        /// <summary>
        /// A comma-separated list of which TLS protocols are allowed by SQL Server
        /// </summary>
        public readonly string? Tlsprotocols;

        [OutputConstructor]
        private K8sNetworkSettingsResponse(
            int? forceencryption,

            string? tlsciphers,

            string? tlsprotocols)
        {
            Forceencryption = forceencryption;
            Tlsciphers = tlsciphers;
            Tlsprotocols = tlsprotocols;
        }
    }
}
