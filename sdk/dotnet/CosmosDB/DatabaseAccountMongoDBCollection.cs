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
    /// An Azure Cosmos DB MongoDB collection.
    /// 
    /// Uses Azure REST API version 2016-03-31.
    /// 
    /// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cosmosdb:DatabaseAccountMongoDBCollection")]
    public partial class DatabaseAccountMongoDBCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// List of index keys
        /// </summary>
        [Output("indexes")]
        public Output<ImmutableArray<Outputs.MongoIndexResponse>> Indexes { get; private set; } = null!;

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the database account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A key-value pair of shard keys to be applied for the request.
        /// </summary>
        [Output("shardKey")]
        public Output<ImmutableDictionary<string, string>?> ShardKey { get; private set; } = null!;

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
        /// Create a DatabaseAccountMongoDBCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseAccountMongoDBCollection(string name, DatabaseAccountMongoDBCollectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:DatabaseAccountMongoDBCollection", name, args ?? new DatabaseAccountMongoDBCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseAccountMongoDBCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:DatabaseAccountMongoDBCollection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20150401:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20150408:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20151106:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20160319:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20160331:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20190801:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20191212:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20200301:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20200401:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20200601preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20200901:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210115:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210301preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210315:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210401preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210415:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210515:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210615:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20210701preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211015:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211015preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20211115preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220215preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220515:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220515preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220815:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20220815preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20221115:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20221115preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230301preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230315:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230315preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230415:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230915:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20230915preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20231115:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20231115preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240215preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240515:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240515preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240815:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20240901preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20241115:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20241201preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20250415:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20250501preview:DatabaseAccountMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230315preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230415:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20230915preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20231115preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240215preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240515preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240815:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20240901preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241115:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb/v20241201preview:MongoDBResourceMongoDBCollection" },
                    new global::Pulumi.Alias { Type = "azure-native:documentdb:MongoDBResourceMongoDBCollection" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseAccountMongoDBCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseAccountMongoDBCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabaseAccountMongoDBCollection(name, id, options);
        }
    }

    public sealed class DatabaseAccountMongoDBCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB collection name.
        /// </summary>
        [Input("collectionName")]
        public Input<string>? CollectionName { get; set; }

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("options", required: true)]
        private InputMap<string>? _options;

        /// <summary>
        /// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// The standard JSON format of a MongoDB collection
        /// </summary>
        [Input("resource", required: true)]
        public Input<Inputs.MongoDBCollectionResourceArgs> Resource { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DatabaseAccountMongoDBCollectionArgs()
        {
        }
        public static new DatabaseAccountMongoDBCollectionArgs Empty => new DatabaseAccountMongoDBCollectionArgs();
    }
}
