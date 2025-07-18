// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab
{
    /// <summary>
    /// A virtual network.
    /// 
    /// Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.
    /// </summary>
    [AzureNativeResourceType("azure-native:devtestlab:VirtualNetwork")]
    public partial class VirtualNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allowed subnets of the virtual network.
        /// </summary>
        [Output("allowedSubnets")]
        public Output<ImmutableArray<Outputs.SubnetResponse>> AllowedSubnets { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The creation date of the virtual network.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The description of the virtual network.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Microsoft.Network resource identifier of the virtual network.
        /// </summary>
        [Output("externalProviderResourceId")]
        public Output<string?> ExternalProviderResourceId { get; private set; } = null!;

        /// <summary>
        /// The external subnet properties.
        /// </summary>
        [Output("externalSubnets")]
        public Output<ImmutableArray<Outputs.ExternalSubnetResponse>> ExternalSubnets { get; private set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The subnet overrides of the virtual network.
        /// </summary>
        [Output("subnetOverrides")]
        public Output<ImmutableArray<Outputs.SubnetOverrideResponse>> SubnetOverrides { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        [Output("uniqueIdentifier")]
        public Output<string> UniqueIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azure-native:devtestlab:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:devtestlab:VirtualNetwork", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:devtestlab/v20150521preview:VirtualNetwork" },
                    new global::Pulumi.Alias { Type = "azure-native:devtestlab/v20160515:VirtualNetwork" },
                    new global::Pulumi.Alias { Type = "azure-native:devtestlab/v20180915:VirtualNetwork" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetwork Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetwork(name, id, options);
        }
    }

    public sealed class VirtualNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedSubnets")]
        private InputList<Inputs.SubnetArgs>? _allowedSubnets;

        /// <summary>
        /// The allowed subnets of the virtual network.
        /// </summary>
        public InputList<Inputs.SubnetArgs> AllowedSubnets
        {
            get => _allowedSubnets ?? (_allowedSubnets = new InputList<Inputs.SubnetArgs>());
            set => _allowedSubnets = value;
        }

        /// <summary>
        /// The description of the virtual network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Microsoft.Network resource identifier of the virtual network.
        /// </summary>
        [Input("externalProviderResourceId")]
        public Input<string>? ExternalProviderResourceId { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the VirtualNetwork
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("subnetOverrides")]
        private InputList<Inputs.SubnetOverrideArgs>? _subnetOverrides;

        /// <summary>
        /// The subnet overrides of the virtual network.
        /// </summary>
        public InputList<Inputs.SubnetOverrideArgs> SubnetOverrides
        {
            get => _subnetOverrides ?? (_subnetOverrides = new InputList<Inputs.SubnetOverrideArgs>());
            set => _subnetOverrides = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VirtualNetworkArgs()
        {
        }
        public static new VirtualNetworkArgs Empty => new VirtualNetworkArgs();
    }
}
