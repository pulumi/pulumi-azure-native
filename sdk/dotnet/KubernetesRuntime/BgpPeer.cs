// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesRuntime
{
    /// <summary>
    /// A BgpPeer resource for an Arc connected cluster (Microsoft.Kubernetes/connectedClusters)
    /// 
    /// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2024-03-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:kubernetesruntime:BgpPeer")]
    public partial class BgpPeer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// My ASN
        /// </summary>
        [Output("myAsn")]
        public Output<int> MyAsn { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Peer Address
        /// </summary>
        [Output("peerAddress")]
        public Output<string> PeerAddress { get; private set; } = null!;

        /// <summary>
        /// Peer ASN
        /// </summary>
        [Output("peerAsn")]
        public Output<int> PeerAsn { get; private set; } = null!;

        /// <summary>
        /// Resource provision state
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BgpPeer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpPeer(string name, BgpPeerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:kubernetesruntime:BgpPeer", name, args ?? new BgpPeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpPeer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:kubernetesruntime:BgpPeer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:kubernetesruntime/v20231001preview:BgpPeer" },
                    new global::Pulumi.Alias { Type = "azure-native:kubernetesruntime/v20240301:BgpPeer" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BgpPeer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpPeer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BgpPeer(name, id, options);
        }
    }

    public sealed class BgpPeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the BgpPeer
        /// </summary>
        [Input("bgpPeerName")]
        public Input<string>? BgpPeerName { get; set; }

        /// <summary>
        /// My ASN
        /// </summary>
        [Input("myAsn", required: true)]
        public Input<int> MyAsn { get; set; } = null!;

        /// <summary>
        /// Peer Address
        /// </summary>
        [Input("peerAddress", required: true)]
        public Input<string> PeerAddress { get; set; } = null!;

        /// <summary>
        /// Peer ASN
        /// </summary>
        [Input("peerAsn", required: true)]
        public Input<int> PeerAsn { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public BgpPeerArgs()
        {
        }
        public static new BgpPeerArgs Empty => new BgpPeerArgs();
    }
}
