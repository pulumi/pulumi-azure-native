// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    /// <summary>
    /// The Network Fabric Controller resource definition.
    /// 
    /// Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.
    /// 
    /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:managednetworkfabric:NetworkFabricController")]
    public partial class NetworkFabricController : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        [Output("annotation")]
        public Output<string?> Annotation { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// As part of an update, the Infrastructure ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Infrastructure services. (This is a Mandatory attribute)
        /// </summary>
        [Output("infrastructureExpressRouteConnections")]
        public Output<ImmutableArray<Outputs.ExpressRouteConnectionInformationResponse>> InfrastructureExpressRouteConnections { get; private set; } = null!;

        /// <summary>
        /// InfrastructureServices IP ranges.
        /// </summary>
        [Output("infrastructureServices")]
        public Output<Outputs.ControllerServicesResponse> InfrastructureServices { get; private set; } = null!;

        /// <summary>
        /// IPv4 Network Fabric Controller Address Space.
        /// </summary>
        [Output("ipv4AddressSpace")]
        public Output<string?> Ipv4AddressSpace { get; private set; } = null!;

        /// <summary>
        /// IPv6 Network Fabric Controller Address Space.
        /// </summary>
        [Output("ipv6AddressSpace")]
        public Output<string?> Ipv6AddressSpace { get; private set; } = null!;

        /// <summary>
        /// A workload management network is required for all the tenant (workload) traffic. This traffic is only dedicated for Tenant workloads which are required to access internet or any other MSFT/Public endpoints.
        /// </summary>
        [Output("isWorkloadManagementNetworkEnabled")]
        public Output<string?> IsWorkloadManagementNetworkEnabled { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Managed Resource Group configuration properties.
        /// </summary>
        [Output("managedResourceGroupConfiguration")]
        public Output<Outputs.ManagedResourceGroupConfigurationResponse?> ManagedResourceGroupConfiguration { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The NF-ID will be an input parameter used by the NF to link and get associated with the parent NFC Service.
        /// </summary>
        [Output("networkFabricIds")]
        public Output<ImmutableArray<string>> NetworkFabricIds { get; private set; } = null!;

        /// <summary>
        /// Network Fabric Controller SKU.
        /// </summary>
        [Output("nfcSku")]
        public Output<string?> NfcSku { get; private set; } = null!;

        /// <summary>
        /// Provides you the latest status of the NFC service, whether it is Accepted, updating, Succeeded or Failed. During this process, the states keep changing based on the status of NFC provisioning.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// List of tenant InternetGateway resource IDs
        /// </summary>
        [Output("tenantInternetGatewayIds")]
        public Output<ImmutableArray<string>> TenantInternetGatewayIds { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// As part of an update, the workload ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Workload services. (This is a Mandatory attribute).
        /// </summary>
        [Output("workloadExpressRouteConnections")]
        public Output<ImmutableArray<Outputs.ExpressRouteConnectionInformationResponse>> WorkloadExpressRouteConnections { get; private set; } = null!;

        /// <summary>
        /// A workload management network is required for all the tenant (workload) traffic. This traffic is only dedicated for Tenant workloads which are required to access internet or any other MSFT/Public endpoints. This is used for the backward compatibility.
        /// </summary>
        [Output("workloadManagementNetwork")]
        public Output<bool> WorkloadManagementNetwork { get; private set; } = null!;

        /// <summary>
        /// WorkloadServices IP ranges.
        /// </summary>
        [Output("workloadServices")]
        public Output<Outputs.ControllerServicesResponse> WorkloadServices { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkFabricController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkFabricController(string name, NetworkFabricControllerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:NetworkFabricController", name, args ?? new NetworkFabricControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkFabricController(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:NetworkFabricController", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230201preview:NetworkFabricController" },
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230615:NetworkFabricController" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkFabricController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkFabricController Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkFabricController(name, id, options);
        }
    }

    public sealed class NetworkFabricControllerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        [Input("annotation")]
        public Input<string>? Annotation { get; set; }

        [Input("infrastructureExpressRouteConnections")]
        private InputList<Inputs.ExpressRouteConnectionInformationArgs>? _infrastructureExpressRouteConnections;

        /// <summary>
        /// As part of an update, the Infrastructure ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Infrastructure services. (This is a Mandatory attribute)
        /// </summary>
        public InputList<Inputs.ExpressRouteConnectionInformationArgs> InfrastructureExpressRouteConnections
        {
            get => _infrastructureExpressRouteConnections ?? (_infrastructureExpressRouteConnections = new InputList<Inputs.ExpressRouteConnectionInformationArgs>());
            set => _infrastructureExpressRouteConnections = value;
        }

        /// <summary>
        /// IPv4 Network Fabric Controller Address Space.
        /// </summary>
        [Input("ipv4AddressSpace")]
        public Input<string>? Ipv4AddressSpace { get; set; }

        /// <summary>
        /// IPv6 Network Fabric Controller Address Space.
        /// </summary>
        [Input("ipv6AddressSpace")]
        public Input<string>? Ipv6AddressSpace { get; set; }

        /// <summary>
        /// A workload management network is required for all the tenant (workload) traffic. This traffic is only dedicated for Tenant workloads which are required to access internet or any other MSFT/Public endpoints.
        /// </summary>
        [Input("isWorkloadManagementNetworkEnabled")]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.IsWorkloadManagementNetworkEnabled>? IsWorkloadManagementNetworkEnabled { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Managed Resource Group configuration properties.
        /// </summary>
        [Input("managedResourceGroupConfiguration")]
        public Input<Inputs.ManagedResourceGroupConfigurationArgs>? ManagedResourceGroupConfiguration { get; set; }

        /// <summary>
        /// Name of the Network Fabric Controller.
        /// </summary>
        [Input("networkFabricControllerName")]
        public Input<string>? NetworkFabricControllerName { get; set; }

        /// <summary>
        /// Network Fabric Controller SKU.
        /// </summary>
        [Input("nfcSku")]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.NfcSku>? NfcSku { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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

        [Input("workloadExpressRouteConnections")]
        private InputList<Inputs.ExpressRouteConnectionInformationArgs>? _workloadExpressRouteConnections;

        /// <summary>
        /// As part of an update, the workload ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Workload services. (This is a Mandatory attribute).
        /// </summary>
        public InputList<Inputs.ExpressRouteConnectionInformationArgs> WorkloadExpressRouteConnections
        {
            get => _workloadExpressRouteConnections ?? (_workloadExpressRouteConnections = new InputList<Inputs.ExpressRouteConnectionInformationArgs>());
            set => _workloadExpressRouteConnections = value;
        }

        public NetworkFabricControllerArgs()
        {
            Ipv4AddressSpace = "10.0.0.0/19";
            Ipv6AddressSpace = "FC00::/59";
            IsWorkloadManagementNetworkEnabled = "True";
            NfcSku = "Standard";
        }
        public static new NetworkFabricControllerArgs Empty => new NetworkFabricControllerArgs();
    }
}
