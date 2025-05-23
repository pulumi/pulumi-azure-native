// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetDatabaseAccountMongoDBCollection
    {
        /// <summary>
        /// Gets the MongoDB collection under an existing Azure Cosmos DB database account.
        /// 
        /// Uses Azure REST API version 2016-03-31.
        /// 
        /// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDatabaseAccountMongoDBCollectionResult> InvokeAsync(GetDatabaseAccountMongoDBCollectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAccountMongoDBCollectionResult>("azure-native:cosmosdb:getDatabaseAccountMongoDBCollection", args ?? new GetDatabaseAccountMongoDBCollectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MongoDB collection under an existing Azure Cosmos DB database account.
        /// 
        /// Uses Azure REST API version 2016-03-31.
        /// 
        /// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseAccountMongoDBCollectionResult> Invoke(GetDatabaseAccountMongoDBCollectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseAccountMongoDBCollectionResult>("azure-native:cosmosdb:getDatabaseAccountMongoDBCollection", args ?? new GetDatabaseAccountMongoDBCollectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MongoDB collection under an existing Azure Cosmos DB database account.
        /// 
        /// Uses Azure REST API version 2016-03-31.
        /// 
        /// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseAccountMongoDBCollectionResult> Invoke(GetDatabaseAccountMongoDBCollectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseAccountMongoDBCollectionResult>("azure-native:cosmosdb:getDatabaseAccountMongoDBCollection", args ?? new GetDatabaseAccountMongoDBCollectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseAccountMongoDBCollectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB collection name.
        /// </summary>
        [Input("collectionName", required: true)]
        public string CollectionName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountMongoDBCollectionArgs()
        {
        }
        public static new GetDatabaseAccountMongoDBCollectionArgs Empty => new GetDatabaseAccountMongoDBCollectionArgs();
    }

    public sealed class GetDatabaseAccountMongoDBCollectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB collection name.
        /// </summary>
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountMongoDBCollectionInvokeArgs()
        {
        }
        public static new GetDatabaseAccountMongoDBCollectionInvokeArgs Empty => new GetDatabaseAccountMongoDBCollectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseAccountMongoDBCollectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The unique resource identifier of the database account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of index keys
        /// </summary>
        public readonly ImmutableArray<Outputs.MongoIndexResponse> Indexes;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A key-value pair of shard keys to be applied for the request.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ShardKey;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDatabaseAccountMongoDBCollectionResult(
            string azureApiVersion,

            string id,

            ImmutableArray<Outputs.MongoIndexResponse> indexes,

            string? location,

            string name,

            ImmutableDictionary<string, string>? shardKey,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Indexes = indexes;
            Location = location;
            Name = name;
            ShardKey = shardKey;
            Tags = tags;
            Type = type;
        }
    }
}
