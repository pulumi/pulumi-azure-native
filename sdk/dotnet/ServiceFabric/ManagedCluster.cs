// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric
{
    /// <summary>
    /// The managed cluster resource
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.
    /// 
    /// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:servicefabric:ManagedCluster")]
    public partial class ManagedCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of add-on features to enable on the cluster.
        /// </summary>
        [Output("addonFeatures")]
        public Output<ImmutableArray<string>> AddonFeatures { get; private set; } = null!;

        /// <summary>
        /// VM admin user password.
        /// </summary>
        [Output("adminPassword")]
        public Output<string?> AdminPassword { get; private set; } = null!;

        /// <summary>
        /// VM admin user name.
        /// </summary>
        [Output("adminUserName")]
        public Output<string> AdminUserName { get; private set; } = null!;

        /// <summary>
        /// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
        /// </summary>
        [Output("allowRdpAccess")]
        public Output<bool?> AllowRdpAccess { get; private set; } = null!;

        /// <summary>
        /// The policy used to clean up unused versions.
        /// </summary>
        [Output("applicationTypeVersionsCleanupPolicy")]
        public Output<Outputs.ApplicationTypeVersionsCleanupPolicyResponse?> ApplicationTypeVersionsCleanupPolicy { get; private set; } = null!;

        /// <summary>
        /// Auxiliary subnets for the cluster.
        /// </summary>
        [Output("auxiliarySubnets")]
        public Output<ImmutableArray<Outputs.SubnetResponse>> AuxiliarySubnets { get; private set; } = null!;

        /// <summary>
        /// The AAD authentication settings of the cluster.
        /// </summary>
        [Output("azureActiveDirectory")]
        public Output<Outputs.AzureActiveDirectoryResponse?> AzureActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The port used for client connections to the cluster.
        /// </summary>
        [Output("clientConnectionPort")]
        public Output<int?> ClientConnectionPort { get; private set; } = null!;

        /// <summary>
        /// Client certificates that are allowed to manage the cluster.
        /// </summary>
        [Output("clients")]
        public Output<ImmutableArray<Outputs.ClientCertificateResponse>> Clients { get; private set; } = null!;

        /// <summary>
        /// List of thumbprints of the cluster certificates.
        /// </summary>
        [Output("clusterCertificateThumbprints")]
        public Output<ImmutableArray<string>> ClusterCertificateThumbprints { get; private set; } = null!;

        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        [Output("clusterCodeVersion")]
        public Output<string?> ClusterCodeVersion { get; private set; } = null!;

        /// <summary>
        /// A service generated unique identifier for the cluster resource.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The current state of the cluster.
        /// </summary>
        [Output("clusterState")]
        public Output<string> ClusterState { get; private set; } = null!;

        /// <summary>
        /// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
        /// </summary>
        [Output("clusterUpgradeCadence")]
        public Output<string?> ClusterUpgradeCadence { get; private set; } = null!;

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// </summary>
        [Output("clusterUpgradeMode")]
        public Output<string?> ClusterUpgradeMode { get; private set; } = null!;

        /// <summary>
        /// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
        /// </summary>
        [Output("ddosProtectionPlanId")]
        public Output<string?> DdosProtectionPlanId { get; private set; } = null!;

        /// <summary>
        /// The cluster dns name.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
        /// </summary>
        [Output("enableAutoOSUpgrade")]
        public Output<bool?> EnableAutoOSUpgrade { get; private set; } = null!;

        /// <summary>
        /// If true, token-based authentication is not allowed on the HttpGatewayEndpoint. This is required to support TLS versions 1.3 and above. If token-based authentication is used, HttpGatewayTokenAuthConnectionPort must be defined.
        /// </summary>
        [Output("enableHttpGatewayExclusiveAuthMode")]
        public Output<bool?> EnableHttpGatewayExclusiveAuthMode { get; private set; } = null!;

        /// <summary>
        /// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
        /// </summary>
        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        /// <summary>
        /// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
        /// </summary>
        [Output("enableServicePublicIP")]
        public Output<bool?> EnableServicePublicIP { get; private set; } = null!;

        /// <summary>
        /// Azure resource etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        [Output("fabricSettings")]
        public Output<ImmutableArray<Outputs.SettingsSectionDescriptionResponse>> FabricSettings { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name associated with the public load balancer of the cluster.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The port used for HTTP connections to the cluster.
        /// </summary>
        [Output("httpGatewayConnectionPort")]
        public Output<int?> HttpGatewayConnectionPort { get; private set; } = null!;

        /// <summary>
        /// The port used for token-auth based HTTPS connections to the cluster. Cannot be set to the same port as HttpGatewayEndpoint.
        /// </summary>
        [Output("httpGatewayTokenAuthConnectionPort")]
        public Output<int?> HttpGatewayTokenAuthConnectionPort { get; private set; } = null!;

        /// <summary>
        /// The list of IP tags associated with the default public IP address of the cluster.
        /// </summary>
        [Output("ipTags")]
        public Output<ImmutableArray<Outputs.IpTagResponse>> IpTags { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address associated with the public load balancer of the cluster.
        /// </summary>
        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        /// <summary>
        /// IPv6 address for the cluster if IPv6 is enabled.
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// Load balancing rules that are applied to the public load balancer of the cluster.
        /// </summary>
        [Output("loadBalancingRules")]
        public Output<ImmutableArray<Outputs.LoadBalancingRuleResponse>> LoadBalancingRules { get; private set; } = null!;

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
        /// </summary>
        [Output("networkSecurityRules")]
        public Output<ImmutableArray<Outputs.NetworkSecurityRuleResponse>> NetworkSecurityRules { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the managed cluster resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Specify the resource id of a public IPv4 prefix that the load balancer will allocate a public IPv4 address from. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Output("publicIPPrefixId")]
        public Output<string?> PublicIPPrefixId { get; private set; } = null!;

        /// <summary>
        /// Specify the resource id of a public IPv6 prefix that the load balancer will allocate a public IPv6 address from. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Output("publicIPv6PrefixId")]
        public Output<string?> PublicIPv6PrefixId { get; private set; } = null!;

        /// <summary>
        /// Service endpoints for subnets in the cluster.
        /// </summary>
        [Output("serviceEndpoints")]
        public Output<ImmutableArray<Outputs.ServiceEndpointResponse>> ServiceEndpoints { get; private set; } = null!;

        /// <summary>
        /// The sku of the managed cluster
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        [Output("upgradeDescription")]
        public Output<Outputs.ClusterUpgradePolicyResponse?> UpgradeDescription { get; private set; } = null!;

        /// <summary>
        /// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
        /// </summary>
        [Output("useCustomVnet")]
        public Output<bool?> UseCustomVnet { get; private set; } = null!;

        /// <summary>
        /// Indicates if the cluster has zone resiliency.
        /// </summary>
        [Output("zonalResiliency")]
        public Output<bool?> ZonalResiliency { get; private set; } = null!;

        /// <summary>
        /// Indicates the update mode for Cross Az clusters.
        /// </summary>
        [Output("zonalUpdateMode")]
        public Output<string?> ZonalUpdateMode { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedCluster(string name, ManagedClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:ManagedCluster", name, args ?? new ManagedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:ManagedCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20200101preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210101preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210501:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210701preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210901privatepreview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20211101preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220101:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220201preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220601preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220801preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20221001preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230201preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230301preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230701preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230901preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231101preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231201preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240201preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240401:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240601preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240901preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20241101preview:ManagedCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20250301preview:ManagedCluster" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedCluster(name, id, options);
        }
    }

    public sealed class ManagedClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("addonFeatures")]
        private InputList<Union<string, Pulumi.AzureNative.ServiceFabric.ManagedClusterAddOnFeature>>? _addonFeatures;

        /// <summary>
        /// List of add-on features to enable on the cluster.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.ServiceFabric.ManagedClusterAddOnFeature>> AddonFeatures
        {
            get => _addonFeatures ?? (_addonFeatures = new InputList<Union<string, Pulumi.AzureNative.ServiceFabric.ManagedClusterAddOnFeature>>());
            set => _addonFeatures = value;
        }

        /// <summary>
        /// VM admin user password.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        /// <summary>
        /// VM admin user name.
        /// </summary>
        [Input("adminUserName", required: true)]
        public Input<string> AdminUserName { get; set; } = null!;

        /// <summary>
        /// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
        /// </summary>
        [Input("allowRdpAccess")]
        public Input<bool>? AllowRdpAccess { get; set; }

        /// <summary>
        /// The policy used to clean up unused versions.
        /// </summary>
        [Input("applicationTypeVersionsCleanupPolicy")]
        public Input<Inputs.ApplicationTypeVersionsCleanupPolicyArgs>? ApplicationTypeVersionsCleanupPolicy { get; set; }

        [Input("auxiliarySubnets")]
        private InputList<Inputs.SubnetArgs>? _auxiliarySubnets;

        /// <summary>
        /// Auxiliary subnets for the cluster.
        /// </summary>
        public InputList<Inputs.SubnetArgs> AuxiliarySubnets
        {
            get => _auxiliarySubnets ?? (_auxiliarySubnets = new InputList<Inputs.SubnetArgs>());
            set => _auxiliarySubnets = value;
        }

        /// <summary>
        /// The AAD authentication settings of the cluster.
        /// </summary>
        [Input("azureActiveDirectory")]
        public Input<Inputs.AzureActiveDirectoryArgs>? AzureActiveDirectory { get; set; }

        /// <summary>
        /// The port used for client connections to the cluster.
        /// </summary>
        [Input("clientConnectionPort")]
        public Input<int>? ClientConnectionPort { get; set; }

        [Input("clients")]
        private InputList<Inputs.ClientCertificateArgs>? _clients;

        /// <summary>
        /// Client certificates that are allowed to manage the cluster.
        /// </summary>
        public InputList<Inputs.ClientCertificateArgs> Clients
        {
            get => _clients ?? (_clients = new InputList<Inputs.ClientCertificateArgs>());
            set => _clients = value;
        }

        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        [Input("clusterCodeVersion")]
        public Input<string>? ClusterCodeVersion { get; set; }

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
        /// </summary>
        [Input("clusterUpgradeCadence")]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.ClusterUpgradeCadence>? ClusterUpgradeCadence { get; set; }

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// </summary>
        [Input("clusterUpgradeMode")]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.ClusterUpgradeMode>? ClusterUpgradeMode { get; set; }

        /// <summary>
        /// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
        /// </summary>
        [Input("ddosProtectionPlanId")]
        public Input<string>? DdosProtectionPlanId { get; set; }

        /// <summary>
        /// The cluster dns name.
        /// </summary>
        [Input("dnsName", required: true)]
        public Input<string> DnsName { get; set; } = null!;

        /// <summary>
        /// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
        /// </summary>
        [Input("enableAutoOSUpgrade")]
        public Input<bool>? EnableAutoOSUpgrade { get; set; }

        /// <summary>
        /// If true, token-based authentication is not allowed on the HttpGatewayEndpoint. This is required to support TLS versions 1.3 and above. If token-based authentication is used, HttpGatewayTokenAuthConnectionPort must be defined.
        /// </summary>
        [Input("enableHttpGatewayExclusiveAuthMode")]
        public Input<bool>? EnableHttpGatewayExclusiveAuthMode { get; set; }

        /// <summary>
        /// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
        /// </summary>
        [Input("enableServicePublicIP")]
        public Input<bool>? EnableServicePublicIP { get; set; }

        [Input("fabricSettings")]
        private InputList<Inputs.SettingsSectionDescriptionArgs>? _fabricSettings;

        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        public InputList<Inputs.SettingsSectionDescriptionArgs> FabricSettings
        {
            get => _fabricSettings ?? (_fabricSettings = new InputList<Inputs.SettingsSectionDescriptionArgs>());
            set => _fabricSettings = value;
        }

        /// <summary>
        /// The port used for HTTP connections to the cluster.
        /// </summary>
        [Input("httpGatewayConnectionPort")]
        public Input<int>? HttpGatewayConnectionPort { get; set; }

        /// <summary>
        /// The port used for token-auth based HTTPS connections to the cluster. Cannot be set to the same port as HttpGatewayEndpoint.
        /// </summary>
        [Input("httpGatewayTokenAuthConnectionPort")]
        public Input<int>? HttpGatewayTokenAuthConnectionPort { get; set; }

        [Input("ipTags")]
        private InputList<Inputs.IpTagArgs>? _ipTags;

        /// <summary>
        /// The list of IP tags associated with the default public IP address of the cluster.
        /// </summary>
        public InputList<Inputs.IpTagArgs> IpTags
        {
            get => _ipTags ?? (_ipTags = new InputList<Inputs.IpTagArgs>());
            set => _ipTags = value;
        }

        [Input("loadBalancingRules")]
        private InputList<Inputs.LoadBalancingRuleArgs>? _loadBalancingRules;

        /// <summary>
        /// Load balancing rules that are applied to the public load balancer of the cluster.
        /// </summary>
        public InputList<Inputs.LoadBalancingRuleArgs> LoadBalancingRules
        {
            get => _loadBalancingRules ?? (_loadBalancingRules = new InputList<Inputs.LoadBalancingRuleArgs>());
            set => _loadBalancingRules = value;
        }

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("networkSecurityRules")]
        private InputList<Inputs.NetworkSecurityRuleArgs>? _networkSecurityRules;

        /// <summary>
        /// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
        /// </summary>
        public InputList<Inputs.NetworkSecurityRuleArgs> NetworkSecurityRules
        {
            get => _networkSecurityRules ?? (_networkSecurityRules = new InputList<Inputs.NetworkSecurityRuleArgs>());
            set => _networkSecurityRules = value;
        }

        /// <summary>
        /// Specify the resource id of a public IPv4 prefix that the load balancer will allocate a public IPv4 address from. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Input("publicIPPrefixId")]
        public Input<string>? PublicIPPrefixId { get; set; }

        /// <summary>
        /// Specify the resource id of a public IPv6 prefix that the load balancer will allocate a public IPv6 address from. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Input("publicIPv6PrefixId")]
        public Input<string>? PublicIPv6PrefixId { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("serviceEndpoints")]
        private InputList<Inputs.ServiceEndpointArgs>? _serviceEndpoints;

        /// <summary>
        /// Service endpoints for subnets in the cluster.
        /// </summary>
        public InputList<Inputs.ServiceEndpointArgs> ServiceEndpoints
        {
            get => _serviceEndpoints ?? (_serviceEndpoints = new InputList<Inputs.ServiceEndpointArgs>());
            set => _serviceEndpoints = value;
        }

        /// <summary>
        /// The sku of the managed cluster
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        /// <summary>
        /// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        [Input("upgradeDescription")]
        public Input<Inputs.ClusterUpgradePolicyArgs>? UpgradeDescription { get; set; }

        /// <summary>
        /// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
        /// </summary>
        [Input("useCustomVnet")]
        public Input<bool>? UseCustomVnet { get; set; }

        /// <summary>
        /// Indicates if the cluster has zone resiliency.
        /// </summary>
        [Input("zonalResiliency")]
        public Input<bool>? ZonalResiliency { get; set; }

        /// <summary>
        /// Indicates the update mode for Cross Az clusters.
        /// </summary>
        [Input("zonalUpdateMode")]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.ZonalUpdateMode>? ZonalUpdateMode { get; set; }

        public ManagedClusterArgs()
        {
            ClientConnectionPort = 19000;
            HttpGatewayConnectionPort = 19080;
            ZonalResiliency = false;
        }
        public static new ManagedClusterArgs Empty => new ManagedClusterArgs();
    }
}
