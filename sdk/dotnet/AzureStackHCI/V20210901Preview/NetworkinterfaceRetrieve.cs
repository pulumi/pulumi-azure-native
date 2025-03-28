// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20210901Preview
{
    /// <summary>
    /// The network interface resource definition.
    /// </summary>
    [AzureNativeResourceType("azure-native:azurestackhci/v20210901preview:NetworkinterfaceRetrieve")]
    public partial class NetworkinterfaceRetrieve : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS Settings for the interface
        /// </summary>
        [Output("dnsSettings")]
        public Output<Outputs.InterfaceDNSSettingsResponse?> DnsSettings { get; private set; } = null!;

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// IPConfigurations - A list of IPConfigurations of the network interface.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.IpConfigurationResponse>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// MacAddress - The MAC address of the network interface.
        /// </summary>
        [Output("macAddress")]
        public Output<string?> MacAddress { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// name of the object to be used in moc
        /// </summary>
        [Output("resourceName")]
        public Output<string?> ResourceName { get; private set; } = null!;

        /// <summary>
        /// NetworkInterfaceStatus defines the observed state of network interfaces
        /// </summary>
        [Output("status")]
        public Output<Outputs.NetworkInterfaceStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
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
        /// Create a NetworkinterfaceRetrieve resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkinterfaceRetrieve(string name, NetworkinterfaceRetrieveArgs args, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci/v20210901preview:NetworkinterfaceRetrieve", name, args ?? new NetworkinterfaceRetrieveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkinterfaceRetrieve(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci/v20210901preview:NetworkinterfaceRetrieve", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210701preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210901preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230701preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230701preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230901preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230901preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240101:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240101:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240201preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240201preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240501preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240501preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240715preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240715preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240801preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240801preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20241001preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20241001preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250201preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250201preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250401preview:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250401preview:networkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci:NetworkinterfaceRetrieve" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci:networkinterfaceRetrieve" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkinterfaceRetrieve resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkinterfaceRetrieve Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkinterfaceRetrieve(name, id, options);
        }
    }

    public sealed class NetworkinterfaceRetrieveArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS Settings for the interface
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.InterfaceDNSSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.IpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// IPConfigurations - A list of IPConfigurations of the network interface.
        /// </summary>
        public InputList<Inputs.IpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.IpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// MacAddress - The MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("networkinterfacesName")]
        public Input<string>? NetworkinterfacesName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// name of the object to be used in moc
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

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

        public NetworkinterfaceRetrieveArgs()
        {
        }
        public static new NetworkinterfaceRetrieveArgs Empty => new NetworkinterfaceRetrieveArgs();
    }
}
