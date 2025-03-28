// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork
{
    /// <summary>
    /// Network function resource response.
    /// 
    /// Uses Azure REST API version 2022-01-01-preview. In version 1.x of the Azure Native provider, it used API version 2020-01-01-preview.
    /// 
    /// Other available API versions: 2023-09-01, 2024-04-15.
    /// </summary>
    [AzureNativeResourceType("azure-native:hybridnetwork:NetworkFunction")]
    public partial class NetworkFunction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The reference to the device resource. Once set, it cannot be updated.
        /// </summary>
        [Output("device")]
        public Output<Outputs.SubResourceResponse?> Device { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource URI of the managed application.
        /// </summary>
        [Output("managedApplication")]
        public Output<Outputs.SubResourceResponse> ManagedApplication { get; private set; } = null!;

        /// <summary>
        /// The parameters for the managed application.
        /// </summary>
        [Output("managedApplicationParameters")]
        public Output<object?> ManagedApplicationParameters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network function container configurations from the user.
        /// </summary>
        [Output("networkFunctionContainerConfigurations")]
        public Output<object?> NetworkFunctionContainerConfigurations { get; private set; } = null!;

        /// <summary>
        /// The network function configurations from the user.
        /// </summary>
        [Output("networkFunctionUserConfigurations")]
        public Output<ImmutableArray<Outputs.NetworkFunctionUserConfigurationResponse>> NetworkFunctionUserConfigurations { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the network function resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The service key for the network function resource.
        /// </summary>
        [Output("serviceKey")]
        public Output<string> ServiceKey { get; private set; } = null!;

        /// <summary>
        /// The sku name for the network function. Once set, it cannot be updated.
        /// </summary>
        [Output("skuName")]
        public Output<string?> SkuName { get; private set; } = null!;

        /// <summary>
        /// The sku type for the network function.
        /// </summary>
        [Output("skuType")]
        public Output<string> SkuType { get; private set; } = null!;

        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The vendor name for the network function. Once set, it cannot be updated.
        /// </summary>
        [Output("vendorName")]
        public Output<string?> VendorName { get; private set; } = null!;

        /// <summary>
        /// The vendor provisioning state for the network function resource.
        /// </summary>
        [Output("vendorProvisioningState")]
        public Output<string> VendorProvisioningState { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkFunction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkFunction(string name, NetworkFunctionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hybridnetwork:NetworkFunction", name, args ?? new NetworkFunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkFunction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hybridnetwork:NetworkFunction", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hybridnetwork/v20200101preview:NetworkFunction" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridnetwork/v20210501:NetworkFunction" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridnetwork/v20220101preview:NetworkFunction" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridnetwork/v20230901:NetworkFunction" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridnetwork/v20240415:NetworkFunction" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkFunction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkFunction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkFunction(name, id, options);
        }
    }

    public sealed class NetworkFunctionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference to the device resource. Once set, it cannot be updated.
        /// </summary>
        [Input("device")]
        public Input<Inputs.SubResourceArgs>? Device { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The parameters for the managed application.
        /// </summary>
        [Input("managedApplicationParameters")]
        public Input<object>? ManagedApplicationParameters { get; set; }

        /// <summary>
        /// The network function container configurations from the user.
        /// </summary>
        [Input("networkFunctionContainerConfigurations")]
        public Input<object>? NetworkFunctionContainerConfigurations { get; set; }

        /// <summary>
        /// Resource name for the network function resource.
        /// </summary>
        [Input("networkFunctionName")]
        public Input<string>? NetworkFunctionName { get; set; }

        [Input("networkFunctionUserConfigurations")]
        private InputList<Inputs.NetworkFunctionUserConfigurationArgs>? _networkFunctionUserConfigurations;

        /// <summary>
        /// The network function configurations from the user.
        /// </summary>
        public InputList<Inputs.NetworkFunctionUserConfigurationArgs> NetworkFunctionUserConfigurations
        {
            get => _networkFunctionUserConfigurations ?? (_networkFunctionUserConfigurations = new InputList<Inputs.NetworkFunctionUserConfigurationArgs>());
            set => _networkFunctionUserConfigurations = value;
        }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name for the network function. Once set, it cannot be updated.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The vendor name for the network function. Once set, it cannot be updated.
        /// </summary>
        [Input("vendorName")]
        public Input<string>? VendorName { get; set; }

        public NetworkFunctionArgs()
        {
        }
        public static new NetworkFunctionArgs Empty => new NetworkFunctionArgs();
    }
}
