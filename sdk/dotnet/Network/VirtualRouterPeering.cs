// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// Virtual Router Peering resource.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:VirtualRouterPeering")]
    public partial class VirtualRouterPeering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Name of the virtual router peering that is unique within a virtual router.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Peer ASN.
        /// </summary>
        [Output("peerAsn")]
        public Output<double?> PeerAsn { get; private set; } = null!;

        /// <summary>
        /// Peer IP.
        /// </summary>
        [Output("peerIp")]
        public Output<string?> PeerIp { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Peering type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualRouterPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualRouterPeering(string name, VirtualRouterPeeringArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:VirtualRouterPeering", name, args ?? new VirtualRouterPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualRouterPeering(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:VirtualRouterPeering", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190701:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:VirtualRouterPeering" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:VirtualRouterPeering" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualRouterPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualRouterPeering Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualRouterPeering(name, id, options);
        }
    }

    public sealed class VirtualRouterPeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the virtual router peering that is unique within a virtual router.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Peer ASN.
        /// </summary>
        [Input("peerAsn")]
        public Input<double>? PeerAsn { get; set; }

        /// <summary>
        /// Peer IP.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        /// <summary>
        /// The name of the Virtual Router Peering.
        /// </summary>
        [Input("peeringName")]
        public Input<string>? PeeringName { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Virtual Router.
        /// </summary>
        [Input("virtualRouterName", required: true)]
        public Input<string> VirtualRouterName { get; set; } = null!;

        public VirtualRouterPeeringArgs()
        {
        }
        public static new VirtualRouterPeeringArgs Empty => new VirtualRouterPeeringArgs();
    }
}
