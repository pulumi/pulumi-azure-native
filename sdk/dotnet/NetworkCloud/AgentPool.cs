// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud
{
    /// <summary>
    /// Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.
    /// 
    /// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:networkcloud:AgentPool")]
    public partial class AgentPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrator credentials to be used for the nodes in this agent pool.
        /// </summary>
        [Output("administratorConfiguration")]
        public Output<Outputs.AdministratorConfigurationResponse?> AdministratorConfiguration { get; private set; } = null!;

        /// <summary>
        /// The configurations that will be applied to each agent in this agent pool.
        /// </summary>
        [Output("agentOptions")]
        public Output<Outputs.AgentOptionsResponse?> AgentOptions { get; private set; } = null!;

        /// <summary>
        /// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
        /// </summary>
        [Output("attachedNetworkConfiguration")]
        public Output<Outputs.AttachedNetworkConfigurationResponse?> AttachedNetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The number of virtual machines that use this configuration.
        /// </summary>
        [Output("count")]
        public Output<double> Count { get; private set; } = null!;

        /// <summary>
        /// The current status of the agent pool.
        /// </summary>
        [Output("detailedStatus")]
        public Output<string> DetailedStatus { get; private set; } = null!;

        /// <summary>
        /// The descriptive message about the current detailed status.
        /// </summary>
        [Output("detailedStatusMessage")]
        public Output<string> DetailedStatusMessage { get; private set; } = null!;

        /// <summary>
        /// Resource ETag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The extended location of the cluster associated with the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes version running in this agent pool.
        /// </summary>
        [Output("kubernetesVersion")]
        public Output<string> KubernetesVersion { get; private set; } = null!;

        /// <summary>
        /// The labels applied to the nodes in this agent pool.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.KubernetesLabelResponse>> Labels { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the agent pool.
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
        /// The taints applied to the nodes in this agent pool.
        /// </summary>
        [Output("taints")]
        public Output<ImmutableArray<Outputs.KubernetesLabelResponse>> Taints { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The configuration of the agent pool.
        /// </summary>
        [Output("upgradeSettings")]
        public Output<Outputs.AgentPoolUpgradeSettingsResponse?> UpgradeSettings { get; private set; } = null!;

        /// <summary>
        /// The name of the VM SKU that determines the size of resources allocated for node VMs.
        /// </summary>
        [Output("vmSkuName")]
        public Output<string> VmSkuName { get; private set; } = null!;


        /// <summary>
        /// Create a AgentPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentPool(string name, AgentPoolArgs args, CustomResourceOptions? options = null)
            : base("azure-native:networkcloud:AgentPool", name, args ?? new AgentPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:networkcloud:AgentPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20230701:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20231001preview:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20240601preview:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20240701:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20241001preview:AgentPool" },
                    new global::Pulumi.Alias { Type = "azure-native:networkcloud/v20250201:AgentPool" },
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
        /// The administrator credentials to be used for the nodes in this agent pool.
        /// </summary>
        [Input("administratorConfiguration")]
        public Input<Inputs.AdministratorConfigurationArgs>? AdministratorConfiguration { get; set; }

        /// <summary>
        /// The configurations that will be applied to each agent in this agent pool.
        /// </summary>
        [Input("agentOptions")]
        public Input<Inputs.AgentOptionsArgs>? AgentOptions { get; set; }

        /// <summary>
        /// The name of the Kubernetes cluster agent pool.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
        /// </summary>
        [Input("attachedNetworkConfiguration")]
        public Input<Inputs.AttachedNetworkConfigurationArgs>? AttachedNetworkConfiguration { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The number of virtual machines that use this configuration.
        /// </summary>
        [Input("count", required: true)]
        public Input<double> Count { get; set; } = null!;

        /// <summary>
        /// The extended location of the cluster associated with the resource.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The name of the Kubernetes cluster.
        /// </summary>
        [Input("kubernetesClusterName", required: true)]
        public Input<string> KubernetesClusterName { get; set; } = null!;

        [Input("labels")]
        private InputList<Inputs.KubernetesLabelArgs>? _labels;

        /// <summary>
        /// The labels applied to the nodes in this agent pool.
        /// </summary>
        public InputList<Inputs.KubernetesLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.KubernetesLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
        /// </summary>
        [Input("mode", required: true)]
        public InputUnion<string, Pulumi.AzureNative.NetworkCloud.AgentPoolMode> Mode { get; set; } = null!;

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

        [Input("taints")]
        private InputList<Inputs.KubernetesLabelArgs>? _taints;

        /// <summary>
        /// The taints applied to the nodes in this agent pool.
        /// </summary>
        public InputList<Inputs.KubernetesLabelArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.KubernetesLabelArgs>());
            set => _taints = value;
        }

        /// <summary>
        /// The configuration of the agent pool.
        /// </summary>
        [Input("upgradeSettings")]
        public Input<Inputs.AgentPoolUpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The name of the VM SKU that determines the size of resources allocated for node VMs.
        /// </summary>
        [Input("vmSkuName", required: true)]
        public Input<string> VmSkuName { get; set; } = null!;

        public AgentPoolArgs()
        {
        }
        public static new AgentPoolArgs Empty => new AgentPoolArgs();
    }
}
