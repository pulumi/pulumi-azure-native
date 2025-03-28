// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB
{
    /// <summary>
    /// A managed Cassandra data center.
    /// 
    /// Uses Azure REST API version 2023-04-15. In version 1.x of the Azure Native provider, it used API version 2021-03-01-preview.
    /// 
    /// Other available API versions: 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-11-15, 2024-12-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:documentdb:CassandraDataCenter")]
    public partial class CassandraDataCenter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the database account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of a managed Cassandra data center.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DataCenterResourceResponseProperties> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CassandraDataCenter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CassandraDataCenter(string name, CassandraDataCenterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:documentdb:CassandraDataCenter", name, args ?? new CassandraDataCenterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CassandraDataCenter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:documentdb:CassandraDataCenter", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20210301preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20210401preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20210701preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20211015:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20211015preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20211115preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20220215preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20220515:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20220515preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20220815:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20220815preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20221115:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20221115preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230301preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230315:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230315preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230415:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240215preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240815:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240901preview:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241115:CassandraDataCenter" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241201preview:CassandraDataCenter" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CassandraDataCenter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CassandraDataCenter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CassandraDataCenter(name, id, options);
        }
    }

    public sealed class CassandraDataCenterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed Cassandra cluster name.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Data center name in a managed Cassandra cluster.
        /// </summary>
        [Input("dataCenterName")]
        public Input<string>? DataCenterName { get; set; }

        /// <summary>
        /// Properties of a managed Cassandra data center.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.DataCenterResourcePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public CassandraDataCenterArgs()
        {
        }
        public static new CassandraDataCenterArgs Empty => new CassandraDataCenterArgs();
    }
}
