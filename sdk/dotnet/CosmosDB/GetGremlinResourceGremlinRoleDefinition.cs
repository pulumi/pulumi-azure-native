// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetGremlinResourceGremlinRoleDefinition
    {
        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Gremlin Role Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2025-05-01-preview.
        /// </summary>
        public static Task<GetGremlinResourceGremlinRoleDefinitionResult> InvokeAsync(GetGremlinResourceGremlinRoleDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGremlinResourceGremlinRoleDefinitionResult>("azure-native:cosmosdb:getGremlinResourceGremlinRoleDefinition", args ?? new GetGremlinResourceGremlinRoleDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Gremlin Role Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2025-05-01-preview.
        /// </summary>
        public static Output<GetGremlinResourceGremlinRoleDefinitionResult> Invoke(GetGremlinResourceGremlinRoleDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGremlinResourceGremlinRoleDefinitionResult>("azure-native:cosmosdb:getGremlinResourceGremlinRoleDefinition", args ?? new GetGremlinResourceGremlinRoleDefinitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Gremlin Role Definition with the given Id.
        /// 
        /// Uses Azure REST API version 2025-05-01-preview.
        /// </summary>
        public static Output<GetGremlinResourceGremlinRoleDefinitionResult> Invoke(GetGremlinResourceGremlinRoleDefinitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGremlinResourceGremlinRoleDefinitionResult>("azure-native:cosmosdb:getGremlinResourceGremlinRoleDefinition", args ?? new GetGremlinResourceGremlinRoleDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGremlinResourceGremlinRoleDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The GUID for the Role Definition.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public string RoleDefinitionId { get; set; } = null!;

        public GetGremlinResourceGremlinRoleDefinitionArgs()
        {
        }
        public static new GetGremlinResourceGremlinRoleDefinitionArgs Empty => new GetGremlinResourceGremlinRoleDefinitionArgs();
    }

    public sealed class GetGremlinResourceGremlinRoleDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The GUID for the Role Definition.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public Input<string> RoleDefinitionId { get; set; } = null!;

        public GetGremlinResourceGremlinRoleDefinitionInvokeArgs()
        {
        }
        public static new GetGremlinResourceGremlinRoleDefinitionInvokeArgs Empty => new GetGremlinResourceGremlinRoleDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetGremlinResourceGremlinRoleDefinitionResult
    {
        /// <summary>
        /// A set of fully qualified Scopes at or below which Gremlin Role Assignments may be created using this Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Must have at least one element. Scopes higher than Database account are not enforceable as assignable Scopes. Note that resources referenced in assignable Scopes need not exist.
        /// </summary>
        public readonly ImmutableArray<string> AssignableScopes;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The set of operations allowed through this Role Definition.
        /// </summary>
        public readonly ImmutableArray<Outputs.PermissionResponse> Permissions;
        /// <summary>
        /// A user-friendly name for the Role Definition. Must be unique for the database account.
        /// </summary>
        public readonly string? RoleName;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGremlinResourceGremlinRoleDefinitionResult(
            ImmutableArray<string> assignableScopes,

            string azureApiVersion,

            string id,

            string name,

            ImmutableArray<Outputs.PermissionResponse> permissions,

            string? roleName,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AssignableScopes = assignableScopes;
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Permissions = permissions;
            RoleName = roleName;
            SystemData = systemData;
            Type = type;
        }
    }
}
