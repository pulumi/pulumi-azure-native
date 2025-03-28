// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20230602Preview
{
    /// <summary>
    /// A managed cluster snapshot resource.
    /// </summary>
    [AzureNativeResourceType("azure-native:containerservice/v20230602preview:ManagedClusterSnapshot")]
    public partial class ManagedClusterSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CreationData to be used to specify the source resource ID to create this snapshot.
        /// </summary>
        [Output("creationData")]
        public Output<Outputs.CreationDataResponse?> CreationData { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
        /// </summary>
        [Output("managedClusterPropertiesReadOnly")]
        public Output<Outputs.ManagedClusterPropertiesForSnapshotResponse> ManagedClusterPropertiesReadOnly { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of a snapshot. The default is NodePool.
        /// </summary>
        [Output("snapshotType")]
        public Output<string?> SnapshotType { get; private set; } = null!;

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
        /// Create a ManagedClusterSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedClusterSnapshot(string name, ManagedClusterSnapshotArgs args, CustomResourceOptions? options = null)
            : base("azure-native:containerservice/v20230602preview:ManagedClusterSnapshot", name, args ?? new ManagedClusterSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedClusterSnapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:containerservice/v20230602preview:ManagedClusterSnapshot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220202preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220302preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220402preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220502preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220602preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220702preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220802preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220803preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20220902preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20221002preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20221102preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230102preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230202preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230302preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230402preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230502preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230702preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230802preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20230902preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20231002preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20231102preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240102preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240202preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240302preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240402preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240502preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240602preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240702preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20240902preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice/v20241002preview:ManagedClusterSnapshot" },
                    new global::Pulumi.Alias { Type = "azure-native:containerservice:ManagedClusterSnapshot" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedClusterSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedClusterSnapshot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedClusterSnapshot(name, id, options);
        }
    }

    public sealed class ManagedClusterSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CreationData to be used to specify the source resource ID to create this snapshot.
        /// </summary>
        [Input("creationData")]
        public Input<Inputs.CreationDataArgs>? CreationData { get; set; }

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

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// The type of a snapshot. The default is NodePool.
        /// </summary>
        [Input("snapshotType")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20230602Preview.SnapshotType>? SnapshotType { get; set; }

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

        public ManagedClusterSnapshotArgs()
        {
        }
        public static new ManagedClusterSnapshotArgs Empty => new ManagedClusterSnapshotArgs();
    }
}
