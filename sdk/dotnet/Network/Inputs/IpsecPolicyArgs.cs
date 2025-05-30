// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// An IPSec Policy configuration for a virtual network gateway connection.
    /// </summary>
    public sealed class IpsecPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DH Group used in IKE Phase 1 for initial SA.
        /// </summary>
        [Input("dhGroup", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.DhGroup> DhGroup { get; set; } = null!;

        /// <summary>
        /// The IKE encryption algorithm (IKE phase 2).
        /// </summary>
        [Input("ikeEncryption", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.IkeEncryption> IkeEncryption { get; set; } = null!;

        /// <summary>
        /// The IKE integrity algorithm (IKE phase 2).
        /// </summary>
        [Input("ikeIntegrity", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.IkeIntegrity> IkeIntegrity { get; set; } = null!;

        /// <summary>
        /// The IPSec encryption algorithm (IKE phase 1).
        /// </summary>
        [Input("ipsecEncryption", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.IpsecEncryption> IpsecEncryption { get; set; } = null!;

        /// <summary>
        /// The IPSec integrity algorithm (IKE phase 1).
        /// </summary>
        [Input("ipsecIntegrity", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.IpsecIntegrity> IpsecIntegrity { get; set; } = null!;

        /// <summary>
        /// The Pfs Group used in IKE Phase 2 for new child SA.
        /// </summary>
        [Input("pfsGroup", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.PfsGroup> PfsGroup { get; set; } = null!;

        /// <summary>
        /// The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
        /// </summary>
        [Input("saDataSizeKilobytes", required: true)]
        public Input<int> SaDataSizeKilobytes { get; set; } = null!;

        /// <summary>
        /// The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.
        /// </summary>
        [Input("saLifeTimeSeconds", required: true)]
        public Input<int> SaLifeTimeSeconds { get; set; } = null!;

        public IpsecPolicyArgs()
        {
        }
        public static new IpsecPolicyArgs Empty => new IpsecPolicyArgs();
    }
}
