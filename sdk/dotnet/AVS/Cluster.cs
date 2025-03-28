// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS
{
    /// <summary>
    /// A cluster resource
    /// 
    /// Uses Azure REST API version 2022-05-01. In version 1.x of the Azure Native provider, it used API version 2020-03-20.
    /// 
    /// Other available API versions: 2020-03-20, 2021-06-01, 2023-03-01, 2023-09-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The identity
        /// </summary>
        [Output("clusterId")]
        public Output<int> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The cluster size
        /// </summary>
        [Output("clusterSize")]
        public Output<int?> ClusterSize { get; private set; } = null!;

        /// <summary>
        /// The hosts
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<string>> Hosts { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The state of the cluster provisioning
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The cluster SKU
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20200320:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20200717preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210101preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:Cluster" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cluster in the private cloud
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The cluster size
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// The hosts
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// The name of the private cloud.
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The cluster SKU
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }
}
