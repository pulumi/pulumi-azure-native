// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight
{
    /// <summary>
    /// The cluster.
    /// 
    /// Uses Azure REST API version 2024-05-01-preview.
    /// 
    /// Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:hdinsight:ClusterPoolCluster")]
    public partial class ClusterPoolCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Cluster profile.
        /// </summary>
        [Output("clusterProfile")]
        public Output<Outputs.ClusterProfileResponse> ClusterProfile { get; private set; } = null!;

        /// <summary>
        /// The type of cluster.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// The compute profile.
        /// </summary>
        [Output("computeProfile")]
        public Output<Outputs.ClusterPoolComputeProfileResponse> ComputeProfile { get; private set; } = null!;

        /// <summary>
        /// A unique id generated by the RP to identify the resource.
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

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
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Business status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

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
        /// Create a ClusterPoolCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterPoolCluster(string name, ClusterPoolClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hdinsight:ClusterPoolCluster", name, args ?? new ClusterPoolClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterPoolCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hdinsight:ClusterPoolCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20230601preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20230601preview:ClusterPoolCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20231101preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20231101preview:ClusterPoolCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20240501preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:hdinsight/v20240501preview:ClusterPoolCluster" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterPoolCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterPoolCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterPoolCluster(name, id, options);
        }
    }

    public sealed class ClusterPoolClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the HDInsight cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The name of the cluster pool.
        /// </summary>
        [Input("clusterPoolName", required: true)]
        public Input<string> ClusterPoolName { get; set; } = null!;

        /// <summary>
        /// Cluster profile.
        /// </summary>
        [Input("clusterProfile", required: true)]
        public Input<Inputs.ClusterProfileArgs> ClusterProfile { get; set; } = null!;

        /// <summary>
        /// The type of cluster.
        /// </summary>
        [Input("clusterType", required: true)]
        public Input<string> ClusterType { get; set; } = null!;

        /// <summary>
        /// The compute profile.
        /// </summary>
        [Input("computeProfile", required: true)]
        public Input<Inputs.ClusterPoolComputeProfileArgs> ComputeProfile { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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

        public ClusterPoolClusterArgs()
        {
        }
        public static new ClusterPoolClusterArgs Empty => new ClusterPoolClusterArgs();
    }
}
