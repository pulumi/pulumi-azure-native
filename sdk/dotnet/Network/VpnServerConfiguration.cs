// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// VpnServerConfiguration Resource.
    /// 
    /// Uses Azure REST API version 2023-02-01. In version 1.x of the Azure Native provider, it used API version 2020-11-01.
    /// 
    /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:VpnServerConfiguration")]
    public partial class VpnServerConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The set of aad vpn authentication parameters.
        /// </summary>
        [Output("aadAuthenticationParameters")]
        public Output<Outputs.AadAuthenticationParametersResponse?> AadAuthenticationParameters { get; private set; } = null!;

        /// <summary>
        /// List of all VpnServerConfigurationPolicyGroups.
        /// </summary>
        [Output("configurationPolicyGroups")]
        public Output<ImmutableArray<Outputs.VpnServerConfigurationPolicyGroupResponse>> ConfigurationPolicyGroups { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of references to P2SVpnGateways.
        /// </summary>
        [Output("p2SVpnGateways")]
        public Output<ImmutableArray<Outputs.P2SVpnGatewayResponse>> P2SVpnGateways { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Radius client root certificate of VpnServerConfiguration.
        /// </summary>
        [Output("radiusClientRootCertificates")]
        public Output<ImmutableArray<Outputs.VpnServerConfigRadiusClientRootCertificateResponse>> RadiusClientRootCertificates { get; private set; } = null!;

        /// <summary>
        /// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Output("radiusServerAddress")]
        public Output<string?> RadiusServerAddress { get; private set; } = null!;

        /// <summary>
        /// Radius Server root certificate of VpnServerConfiguration.
        /// </summary>
        [Output("radiusServerRootCertificates")]
        public Output<ImmutableArray<Outputs.VpnServerConfigRadiusServerRootCertificateResponse>> RadiusServerRootCertificates { get; private set; } = null!;

        /// <summary>
        /// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Output("radiusServerSecret")]
        public Output<string?> RadiusServerSecret { get; private set; } = null!;

        /// <summary>
        /// Multiple Radius Server configuration for VpnServerConfiguration.
        /// </summary>
        [Output("radiusServers")]
        public Output<ImmutableArray<Outputs.RadiusServerResponse>> RadiusServers { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VPN authentication types for the VpnServerConfiguration.
        /// </summary>
        [Output("vpnAuthenticationTypes")]
        public Output<ImmutableArray<string>> VpnAuthenticationTypes { get; private set; } = null!;

        /// <summary>
        /// VpnClientIpsecPolicies for VpnServerConfiguration.
        /// </summary>
        [Output("vpnClientIpsecPolicies")]
        public Output<ImmutableArray<Outputs.IpsecPolicyResponse>> VpnClientIpsecPolicies { get; private set; } = null!;

        /// <summary>
        /// VPN client revoked certificate of VpnServerConfiguration.
        /// </summary>
        [Output("vpnClientRevokedCertificates")]
        public Output<ImmutableArray<Outputs.VpnServerConfigVpnClientRevokedCertificateResponse>> VpnClientRevokedCertificates { get; private set; } = null!;

        /// <summary>
        /// VPN client root certificate of VpnServerConfiguration.
        /// </summary>
        [Output("vpnClientRootCertificates")]
        public Output<ImmutableArray<Outputs.VpnServerConfigVpnClientRootCertificateResponse>> VpnClientRootCertificates { get; private set; } = null!;

        /// <summary>
        /// VPN protocols for the VpnServerConfiguration.
        /// </summary>
        [Output("vpnProtocols")]
        public Output<ImmutableArray<string>> VpnProtocols { get; private set; } = null!;


        /// <summary>
        /// Create a VpnServerConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnServerConfiguration(string name, VpnServerConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:VpnServerConfiguration", name, args ?? new VpnServerConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnServerConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:VpnServerConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:VpnServerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:VpnServerConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnServerConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnServerConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnServerConfiguration(name, id, options);
        }
    }

    public sealed class VpnServerConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The set of aad vpn authentication parameters.
        /// </summary>
        [Input("aadAuthenticationParameters")]
        public Input<Inputs.AadAuthenticationParametersArgs>? AadAuthenticationParameters { get; set; }

        [Input("configurationPolicyGroups")]
        private InputList<Inputs.VpnServerConfigurationPolicyGroupArgs>? _configurationPolicyGroups;

        /// <summary>
        /// List of all VpnServerConfigurationPolicyGroups.
        /// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
        /// </summary>
        public InputList<Inputs.VpnServerConfigurationPolicyGroupArgs> ConfigurationPolicyGroups
        {
            get => _configurationPolicyGroups ?? (_configurationPolicyGroups = new InputList<Inputs.VpnServerConfigurationPolicyGroupArgs>());
            set => _configurationPolicyGroups = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the VpnServerConfiguration that is unique within a resource group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("radiusClientRootCertificates")]
        private InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs>? _radiusClientRootCertificates;

        /// <summary>
        /// Radius client root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs> RadiusClientRootCertificates
        {
            get => _radiusClientRootCertificates ?? (_radiusClientRootCertificates = new InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs>());
            set => _radiusClientRootCertificates = value;
        }

        /// <summary>
        /// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerAddress")]
        public Input<string>? RadiusServerAddress { get; set; }

        [Input("radiusServerRootCertificates")]
        private InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs>? _radiusServerRootCertificates;

        /// <summary>
        /// Radius Server root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs> RadiusServerRootCertificates
        {
            get => _radiusServerRootCertificates ?? (_radiusServerRootCertificates = new InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs>());
            set => _radiusServerRootCertificates = value;
        }

        /// <summary>
        /// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerSecret")]
        public Input<string>? RadiusServerSecret { get; set; }

        [Input("radiusServers")]
        private InputList<Inputs.RadiusServerArgs>? _radiusServers;

        /// <summary>
        /// Multiple Radius Server configuration for VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.RadiusServerArgs> RadiusServers
        {
            get => _radiusServers ?? (_radiusServers = new InputList<Inputs.RadiusServerArgs>());
            set => _radiusServers = value;
        }

        /// <summary>
        /// The resource group name of the VpnServerConfiguration.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        [Input("vpnAuthenticationTypes")]
        private InputList<Union<string, Pulumi.AzureNative.Network.VpnAuthenticationType>>? _vpnAuthenticationTypes;

        /// <summary>
        /// VPN authentication types for the VpnServerConfiguration.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Network.VpnAuthenticationType>> VpnAuthenticationTypes
        {
            get => _vpnAuthenticationTypes ?? (_vpnAuthenticationTypes = new InputList<Union<string, Pulumi.AzureNative.Network.VpnAuthenticationType>>());
            set => _vpnAuthenticationTypes = value;
        }

        [Input("vpnClientIpsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _vpnClientIpsecPolicies;

        /// <summary>
        /// VpnClientIpsecPolicies for VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> VpnClientIpsecPolicies
        {
            get => _vpnClientIpsecPolicies ?? (_vpnClientIpsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _vpnClientIpsecPolicies = value;
        }

        [Input("vpnClientRevokedCertificates")]
        private InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs>? _vpnClientRevokedCertificates;

        /// <summary>
        /// VPN client revoked certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs> VpnClientRevokedCertificates
        {
            get => _vpnClientRevokedCertificates ?? (_vpnClientRevokedCertificates = new InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs>());
            set => _vpnClientRevokedCertificates = value;
        }

        [Input("vpnClientRootCertificates")]
        private InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs>? _vpnClientRootCertificates;

        /// <summary>
        /// VPN client root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs> VpnClientRootCertificates
        {
            get => _vpnClientRootCertificates ?? (_vpnClientRootCertificates = new InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs>());
            set => _vpnClientRootCertificates = value;
        }

        [Input("vpnProtocols")]
        private InputList<Union<string, Pulumi.AzureNative.Network.VpnGatewayTunnelingProtocol>>? _vpnProtocols;

        /// <summary>
        /// VPN protocols for the VpnServerConfiguration.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Network.VpnGatewayTunnelingProtocol>> VpnProtocols
        {
            get => _vpnProtocols ?? (_vpnProtocols = new InputList<Union<string, Pulumi.AzureNative.Network.VpnGatewayTunnelingProtocol>>());
            set => _vpnProtocols = value;
        }

        /// <summary>
        /// The name of the VpnServerConfiguration being created or updated.
        /// </summary>
        [Input("vpnServerConfigurationName")]
        public Input<string>? VpnServerConfigurationName { get; set; }

        public VpnServerConfigurationArgs()
        {
        }
        public static new VpnServerConfigurationArgs Empty => new VpnServerConfigurationArgs();
    }
}
