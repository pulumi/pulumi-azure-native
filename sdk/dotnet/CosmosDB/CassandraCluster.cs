// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    /// <summary>
    /// Representation of a managed Cassandra cluster.
    /// 
    /// Uses Azure REST API version 2024-11-15.
    /// 
    /// Other available API versions: 2021-03-01-preview, 2021-04-01-preview, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cosmosdb:CassandraCluster")]
    public partial class CassandraCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Identity for the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedCassandraManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of a managed Cassandra cluster.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ClusterResourceResponseProperties> Properties { get; private set; } = null!;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CassandraCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CassandraCluster(string name, CassandraClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:CassandraCluster", name, args ?? new CassandraClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CassandraCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:CassandraCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210301preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210401preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210701preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211015:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211015preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211115preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220215preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220515:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220515preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220815:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220815preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20221115:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20221115preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230301preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230315:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230315preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230415:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230915:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230915preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20231115:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20231115preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240215preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240515:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240515preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240815:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240901preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20241115:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20241201preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20250415:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20250501preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20210701preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230415:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240215preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240815:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240901preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241115:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241201preview:CassandraCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb:CassandraCluster" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CassandraCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CassandraCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CassandraCluster(name, id, options);
        }
    }

    public sealed class CassandraClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed Cassandra cluster name.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Identity for the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedCassandraManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Properties of a managed Cassandra cluster.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ClusterResourcePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CassandraClusterArgs()
        {
        }
        public static new CassandraClusterArgs Empty => new CassandraClusterArgs();
    }
}
