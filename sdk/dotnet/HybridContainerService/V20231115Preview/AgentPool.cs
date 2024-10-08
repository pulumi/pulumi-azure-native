// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview
{
    /// <summary>
    /// The agentPool resource definition
    /// </summary>
    [AzureNativeResourceType("azure-native:hybridcontainerservice/v20231115preview:AgentPool")]
    public partial class AgentPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        /// </summary>
        [Output("count")]
        public Output<int?> Count { get; private set; } = null!;

        /// <summary>
        /// Extended Location definition
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The version of node image
        /// </summary>
        [Output("nodeImageVersion")]
        public Output<string?> NodeImageVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when OSType is Windows.
        /// </summary>
        [Output("osSKU")]
        public Output<string?> OsSKU { get; private set; } = null!;

        /// <summary>
        /// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Defines the observed state of the agent pool
        /// </summary>
        [Output("status")]
        public Output<Outputs.AgentPoolProvisioningStatusResponseStatus?> Status { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VmSize - The size of the agent pool VMs.
        /// </summary>
        [Output("vmSize")]
        public Output<string?> VmSize { get; private set; } = null!;


        /// <summary>
        /// Create a AgentPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentPool(string name, AgentPoolArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hybridcontainerservice/v20231115preview:AgentPool", name, args ?? new AgentPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hybridcontainerservice/v20231115preview:AgentPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20231115preview:agentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20240101:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20240101:agentPool" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AgentPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentPool(name, id, options);
        }
    }

    public sealed class AgentPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter for the name of the agent pool in the provisioned cluster
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the connected cluster resource.
        /// </summary>
        [Input("connectedClusterResourceUri", required: true)]
        public Input<string> ConnectedClusterResourceUri { get; set; } = null!;

        /// <summary>
        /// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// Extended Location definition
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The version of node image
        /// </summary>
        [Input("nodeImageVersion")]
        public Input<string>? NodeImageVersion { get; set; }

        /// <summary>
        /// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when OSType is Windows.
        /// </summary>
        [Input("osSKU")]
        public InputUnion<string, Pulumi.AzureNative.HybridContainerService.V20231115Preview.OSSKU>? OsSKU { get; set; }

        /// <summary>
        /// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        /// </summary>
        [Input("osType")]
        public InputUnion<string, Pulumi.AzureNative.HybridContainerService.V20231115Preview.OsType>? OsType { get; set; }

        /// <summary>
        /// Defines the observed state of the agent pool
        /// </summary>
        [Input("status")]
        public Input<Inputs.AgentPoolProvisioningStatusStatusArgs>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// VmSize - The size of the agent pool VMs.
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        public AgentPoolArgs()
        {
            Count = 1;
        }
        public static new AgentPoolArgs Empty => new AgentPoolArgs();
    }
}
