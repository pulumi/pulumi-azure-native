// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetMongoDBResourceMongoUserDefinition
    {
        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMongoDBResourceMongoUserDefinitionResult> InvokeAsync(GetMongoDBResourceMongoUserDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMongoDBResourceMongoUserDefinitionResult>("azure-native:cosmosdb:getMongoDBResourceMongoUserDefinition", args ?? new GetMongoDBResourceMongoUserDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMongoDBResourceMongoUserDefinitionResult> Invoke(GetMongoDBResourceMongoUserDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoDBResourceMongoUserDefinitionResult>("azure-native:cosmosdb:getMongoDBResourceMongoUserDefinition", args ?? new GetMongoDBResourceMongoUserDefinitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMongoDBResourceMongoUserDefinitionResult> Invoke(GetMongoDBResourceMongoUserDefinitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoDBResourceMongoUserDefinitionResult>("azure-native:cosmosdb:getMongoDBResourceMongoUserDefinition", args ?? new GetMongoDBResourceMongoUserDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMongoDBResourceMongoUserDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The ID for the User Definition {dbName.userName}.
        /// </summary>
        [Input("mongoUserDefinitionId", required: true)]
        public string MongoUserDefinitionId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMongoDBResourceMongoUserDefinitionArgs()
        {
        }
        public static new GetMongoDBResourceMongoUserDefinitionArgs Empty => new GetMongoDBResourceMongoUserDefinitionArgs();
    }

    public sealed class GetMongoDBResourceMongoUserDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The ID for the User Definition {dbName.userName}.
        /// </summary>
        [Input("mongoUserDefinitionId", required: true)]
        public Input<string> MongoUserDefinitionId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMongoDBResourceMongoUserDefinitionInvokeArgs()
        {
        }
        public static new GetMongoDBResourceMongoUserDefinitionInvokeArgs Empty => new GetMongoDBResourceMongoUserDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetMongoDBResourceMongoUserDefinitionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A custom definition for the USer Definition.
        /// </summary>
        public readonly string? CustomData;
        /// <summary>
        /// The database name for which access is being granted for this User Definition.
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// The unique resource identifier of the database account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
        /// </summary>
        public readonly string? Mechanisms;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The password for User Definition. Response does not contain user password.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The set of roles inherited by the User Definition.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoleResponse> Roles;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user name for User Definition.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private GetMongoDBResourceMongoUserDefinitionResult(
            string azureApiVersion,

            string? customData,

            string? databaseName,

            string id,

            string? mechanisms,

            string name,

            string? password,

            ImmutableArray<Outputs.RoleResponse> roles,

            string type,

            string? userName)
        {
            AzureApiVersion = azureApiVersion;
            CustomData = customData;
            DatabaseName = databaseName;
            Id = id;
            Mechanisms = mechanisms;
            Name = name;
            Password = password;
            Roles = roles;
            Type = type;
            UserName = userName;
        }
    }
}
