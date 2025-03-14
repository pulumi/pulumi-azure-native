// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20230401.Outputs
{

    /// <summary>
    /// Properties of VPN client root certificate of VpnServerConfiguration.
    /// </summary>
    [OutputType]
    public sealed class VpnServerConfigVpnClientRootCertificateResponse
    {
        /// <summary>
        /// The certificate name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The certificate public data.
        /// </summary>
        public readonly string? PublicCertData;

        [OutputConstructor]
        private VpnServerConfigVpnClientRootCertificateResponse(
            string? name,

            string? publicCertData)
        {
            Name = name;
            PublicCertData = publicCertData;
        }
    }
}
