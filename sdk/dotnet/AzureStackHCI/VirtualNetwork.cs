// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI
{
    /// <summary>
    /// The virtual network resource definition.
    /// 
    /// Uses Azure REST API version 2023-07-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-12-15-preview.
    /// 
    /// Other available API versions: 2022-12-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:azurestackhci:VirtualNetwork")]
    public partial class VirtualNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for a subnet overrides VNET DHCP options.
        /// </summary>
        [Output("dhcpOptions")]
        public Output<Outputs.VirtualNetworkPropertiesResponseDhcpOptions?> DhcpOptions { get; private set; } = null!;

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of the network
        /// </summary>
        [Output("networkType")]
        public Output<string?> NetworkType { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the virtual network.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The observed state of virtual networks
        /// </summary>
        [Output("status")]
        public Output<Outputs.VirtualNetworkStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Subnet - list of subnets under the virtual network
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.VirtualNetworkPropertiesResponseSubnets>> Subnets { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
        /// name of the network switch to be used for VMs
        /// </summary>
        [Output("vmSwitchName")]
        public Output<string?> VmSwitchName { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci:VirtualNetwork", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210701preview:VirtualNetwork" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210901preview:VirtualNetwork" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210901preview:VirtualnetworkRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:VirtualNetwork" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230701preview:VirtualNetwork" },
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
        /// <summary>
        /// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for a subnet overrides VNET DHCP options.
        /// </summary>
        [Input("dhcpOptions")]
        public Input<Inputs.VirtualNetworkPropertiesDhcpOptionsArgs>? DhcpOptions { get; set; }

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Type of the network
        /// </summary>
        [Input("networkType")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.NetworkTypeEnum>? NetworkType { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("subnets")]
        private InputList<Inputs.VirtualNetworkPropertiesSubnetsArgs>? _subnets;

        /// <summary>
        /// Subnet - list of subnets under the virtual network
        /// </summary>
        public InputList<Inputs.VirtualNetworkPropertiesSubnetsArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.VirtualNetworkPropertiesSubnetsArgs>());
            set => _subnets = value;
        }

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
        /// Name of the virtual network
        /// </summary>
        [Input("virtualNetworkName")]
        public Input<string>? VirtualNetworkName { get; set; }

        /// <summary>
        /// name of the network switch to be used for VMs
        /// </summary>
        [Input("vmSwitchName")]
        public Input<string>? VmSwitchName { get; set; }

        public VirtualNetworkArgs()
        {
        }
        public static new VirtualNetworkArgs Empty => new VirtualNetworkArgs();
    }
}
