// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Properties of the Radius client root certificate of VpnServerConfiguration.
    /// </summary>
    [OutputType]
    public sealed class VpnServerConfigRadiusClientRootCertificateResponse
    {
        /// <summary>
        /// The certificate name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The Radius client root certificate thumbprint.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private VpnServerConfigRadiusClientRootCertificateResponse(
            string? name,

            string? thumbprint)
        {
            Name = name;
            Thumbprint = thumbprint;
        }
    }
}
