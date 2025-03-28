// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    /// <summary>
    /// The NetworkToNetworkInterconnect resource definition.
    /// 
    /// Uses Azure REST API version 2023-02-01-preview. In version 1.x of the Azure Native provider, it used API version 2023-02-01-preview.
    /// 
    /// Other available API versions: 2023-06-15.
    /// </summary>
    [AzureNativeResourceType("azure-native:managednetworkfabric:NetworkToNetworkInterconnect")]
    public partial class NetworkToNetworkInterconnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the administrativeState of the resource. Example -Enabled/Disabled
        /// </summary>
        [Output("administrativeState")]
        public Output<string> AdministrativeState { get; private set; } = null!;

        /// <summary>
        /// Configuration to use NNI for Infrastructure Management. Example: True/False.
        /// </summary>
        [Output("isManagementType")]
        public Output<string> IsManagementType { get; private set; } = null!;

        /// <summary>
        /// Common properties for Layer2Configuration.
        /// </summary>
        [Output("layer2Configuration")]
        public Output<Outputs.Layer2ConfigurationResponse?> Layer2Configuration { get; private set; } = null!;

        /// <summary>
        /// Common properties for Layer3Configuration.
        /// </summary>
        [Output("layer3Configuration")]
        public Output<Outputs.Layer3ConfigurationResponse?> Layer3Configuration { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of NNI used. Example: CE | NPB
        /// </summary>
        [Output("nniType")]
        public Output<string?> NniType { get; private set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the resource.
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
        /// Based on this parameter the layer2/layer3 is made as mandatory. Example: True/False
        /// </summary>
        [Output("useOptionB")]
        public Output<string> UseOptionB { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkToNetworkInterconnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkToNetworkInterconnect(string name, NetworkToNetworkInterconnectArgs args, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:NetworkToNetworkInterconnect", name, args ?? new NetworkToNetworkInterconnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkToNetworkInterconnect(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:NetworkToNetworkInterconnect", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230201preview:NetworkToNetworkInterconnect" },
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230615:NetworkToNetworkInterconnect" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkToNetworkInterconnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkToNetworkInterconnect Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkToNetworkInterconnect(name, id, options);
        }
    }

    public sealed class NetworkToNetworkInterconnectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration to use NNI for Infrastructure Management. Example: True/False.
        /// </summary>
        [Input("isManagementType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.BooleanEnumProperty> IsManagementType { get; set; } = null!;

        /// <summary>
        /// Common properties for Layer2Configuration.
        /// </summary>
        [Input("layer2Configuration")]
        public Input<Inputs.Layer2ConfigurationArgs>? Layer2Configuration { get; set; }

        /// <summary>
        /// Common properties for Layer3Configuration.
        /// </summary>
        [Input("layer3Configuration")]
        public Input<Inputs.Layer3ConfigurationArgs>? Layer3Configuration { get; set; }

        /// <summary>
        /// Name of the NetworkFabric.
        /// </summary>
        [Input("networkFabricName", required: true)]
        public Input<string> NetworkFabricName { get; set; } = null!;

        /// <summary>
        /// Name of the NetworkToNetworkInterconnectName
        /// </summary>
        [Input("networkToNetworkInterconnectName")]
        public Input<string>? NetworkToNetworkInterconnectName { get; set; }

        /// <summary>
        /// Type of NNI used. Example: CE | NPB
        /// </summary>
        [Input("nniType")]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.NniType>? NniType { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Based on this parameter the layer2/layer3 is made as mandatory. Example: True/False
        /// </summary>
        [Input("useOptionB", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ManagedNetworkFabric.BooleanEnumProperty> UseOptionB { get; set; } = null!;

        public NetworkToNetworkInterconnectArgs()
        {
            NniType = "CE";
        }
        public static new NetworkToNetworkInterconnectArgs Empty => new NetworkToNetworkInterconnectArgs();
    }
}
