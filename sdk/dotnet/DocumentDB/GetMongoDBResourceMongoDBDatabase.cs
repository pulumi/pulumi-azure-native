// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB
{
    public static class GetMongoDBResourceMongoDBDatabase
    {
        /// <summary>
        /// Gets the MongoDB databases under an existing Azure Cosmos DB database account with the provided name.
        /// 
        /// Uses Azure REST API version 2023-04-15.
        /// 
        /// Other available API versions: 2023-03-15-preview, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-11-15, 2024-12-01-preview.
        /// </summary>
        public static Task<GetMongoDBResourceMongoDBDatabaseResult> InvokeAsync(GetMongoDBResourceMongoDBDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMongoDBResourceMongoDBDatabaseResult>("azure-native:documentdb:getMongoDBResourceMongoDBDatabase", args ?? new GetMongoDBResourceMongoDBDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MongoDB databases under an existing Azure Cosmos DB database account with the provided name.
        /// 
        /// Uses Azure REST API version 2023-04-15.
        /// 
        /// Other available API versions: 2023-03-15-preview, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-11-15, 2024-12-01-preview.
        /// </summary>
        public static Output<GetMongoDBResourceMongoDBDatabaseResult> Invoke(GetMongoDBResourceMongoDBDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoDBResourceMongoDBDatabaseResult>("azure-native:documentdb:getMongoDBResourceMongoDBDatabase", args ?? new GetMongoDBResourceMongoDBDatabaseInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MongoDB databases under an existing Azure Cosmos DB database account with the provided name.
        /// 
        /// Uses Azure REST API version 2023-04-15.
        /// 
        /// Other available API versions: 2023-03-15-preview, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-11-15, 2024-12-01-preview.
        /// </summary>
        public static Output<GetMongoDBResourceMongoDBDatabaseResult> Invoke(GetMongoDBResourceMongoDBDatabaseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoDBResourceMongoDBDatabaseResult>("azure-native:documentdb:getMongoDBResourceMongoDBDatabase", args ?? new GetMongoDBResourceMongoDBDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMongoDBResourceMongoDBDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMongoDBResourceMongoDBDatabaseArgs()
        {
        }
        public static new GetMongoDBResourceMongoDBDatabaseArgs Empty => new GetMongoDBResourceMongoDBDatabaseArgs();
    }

    public sealed class GetMongoDBResourceMongoDBDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMongoDBResourceMongoDBDatabaseInvokeArgs()
        {
        }
        public static new GetMongoDBResourceMongoDBDatabaseInvokeArgs Empty => new GetMongoDBResourceMongoDBDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetMongoDBResourceMongoDBDatabaseResult
    {
        /// <summary>
        /// The unique resource identifier of the ARM resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.MongoDBDatabaseGetPropertiesResponseOptions? Options;
        public readonly Outputs.MongoDBDatabaseGetPropertiesResponseResource? Resource;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMongoDBResourceMongoDBDatabaseResult(
            string id,

            string? location,

            string name,

            Outputs.MongoDBDatabaseGetPropertiesResponseOptions? options,

            Outputs.MongoDBDatabaseGetPropertiesResponseResource? resource,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            Options = options;
            Resource = resource;
            Tags = tags;
            Type = type;
        }
    }
}
