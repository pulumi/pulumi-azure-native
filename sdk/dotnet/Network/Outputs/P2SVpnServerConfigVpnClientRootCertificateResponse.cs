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
    /// VPN client root certificate of P2SVpnServerConfiguration.
    /// </summary>
    [OutputType]
    public sealed class P2SVpnServerConfigVpnClientRootCertificateResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the VPN client root certificate resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The certificate public data.
        /// </summary>
        public readonly string PublicCertData;

        [OutputConstructor]
        private P2SVpnServerConfigVpnClientRootCertificateResponse(
            string? etag,

            string? id,

            string? name,

            string provisioningState,

            string publicCertData)
        {
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            PublicCertData = publicCertData;
        }
    }
}
