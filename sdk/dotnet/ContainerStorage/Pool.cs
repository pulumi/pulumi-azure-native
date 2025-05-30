// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerStorage
{
    /// <summary>
    /// Pool resource
    /// 
    /// Uses Azure REST API version 2023-07-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-07-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:containerstorage:Pool")]
    public partial class Pool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
        /// </summary>
        [Output("assignments")]
        public Output<ImmutableArray<Outputs.AssignmentResponse>> Assignments { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

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
        /// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
        /// </summary>
        [Output("poolType")]
        public Output<Outputs.PoolTypeResponse> PoolType { get; private set; } = null!;

        /// <summary>
        /// The status of the last operation.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
        /// </summary>
        [Output("reclaimPolicy")]
        public Output<string?> ReclaimPolicy { get; private set; } = null!;

        /// <summary>
        /// Resources represent the resources the pool should have.
        /// </summary>
        [Output("resources")]
        public Output<Outputs.ResourcesResponse?> Resources { get; private set; } = null!;

        /// <summary>
        /// The operational status of the resource
        /// </summary>
        [Output("status")]
        public Output<Outputs.ResourceOperationalStatusResponse> Status { get; private set; } = null!;

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
        /// List of availability zones that resources can be created in.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Pool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pool(string name, PoolArgs args, CustomResourceOptions? options = null)
            : base("azure-native:containerstorage:Pool", name, args ?? new PoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:containerstorage:Pool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:containerstorage/v20230701preview:Pool" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Pool(name, id, options);
        }
    }

    public sealed class PoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("assignments")]
        private InputList<Inputs.AssignmentArgs>? _assignments;

        /// <summary>
        /// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
        /// </summary>
        public InputList<Inputs.AssignmentArgs> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<Inputs.AssignmentArgs>());
            set => _assignments = value;
        }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Pool Object
        /// </summary>
        [Input("poolName")]
        public Input<string>? PoolName { get; set; }

        /// <summary>
        /// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
        /// </summary>
        [Input("poolType", required: true)]
        public Input<Inputs.PoolTypeArgs> PoolType { get; set; } = null!;

        /// <summary>
        /// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
        /// </summary>
        [Input("reclaimPolicy")]
        public InputUnion<string, Pulumi.AzureNative.ContainerStorage.ReclaimPolicy>? ReclaimPolicy { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Resources represent the resources the pool should have.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.ResourcesArgs>? Resources { get; set; }

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

        [Input("zones")]
        private InputList<Union<string, Pulumi.AzureNative.ContainerStorage.Zone>>? _zones;

        /// <summary>
        /// List of availability zones that resources can be created in.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.ContainerStorage.Zone>> Zones
        {
            get => _zones ?? (_zones = new InputList<Union<string, Pulumi.AzureNative.ContainerStorage.Zone>>());
            set => _zones = value;
        }

        public PoolArgs()
        {
        }
        public static new PoolArgs Empty => new PoolArgs();
    }
}
