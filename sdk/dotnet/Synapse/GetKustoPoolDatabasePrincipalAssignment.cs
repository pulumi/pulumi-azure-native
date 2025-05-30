// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    public static class GetKustoPoolDatabasePrincipalAssignment
    {
        /// <summary>
        /// Gets a Kusto pool database principalAssignment.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Task<GetKustoPoolDatabasePrincipalAssignmentResult> InvokeAsync(GetKustoPoolDatabasePrincipalAssignmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKustoPoolDatabasePrincipalAssignmentResult>("azure-native:synapse:getKustoPoolDatabasePrincipalAssignment", args ?? new GetKustoPoolDatabasePrincipalAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Kusto pool database principalAssignment.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<GetKustoPoolDatabasePrincipalAssignmentResult> Invoke(GetKustoPoolDatabasePrincipalAssignmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKustoPoolDatabasePrincipalAssignmentResult>("azure-native:synapse:getKustoPoolDatabasePrincipalAssignment", args ?? new GetKustoPoolDatabasePrincipalAssignmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Kusto pool database principalAssignment.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<GetKustoPoolDatabasePrincipalAssignmentResult> Invoke(GetKustoPoolDatabasePrincipalAssignmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetKustoPoolDatabasePrincipalAssignmentResult>("azure-native:synapse:getKustoPoolDatabasePrincipalAssignment", args ?? new GetKustoPoolDatabasePrincipalAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKustoPoolDatabasePrincipalAssignmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database in the Kusto pool.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public string KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto principalAssignment.
        /// </summary>
        [Input("principalAssignmentName", required: true)]
        public string PrincipalAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetKustoPoolDatabasePrincipalAssignmentArgs()
        {
        }
        public static new GetKustoPoolDatabasePrincipalAssignmentArgs Empty => new GetKustoPoolDatabasePrincipalAssignmentArgs();
    }

    public sealed class GetKustoPoolDatabasePrincipalAssignmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database in the Kusto pool.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public Input<string> KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto principalAssignment.
        /// </summary>
        [Input("principalAssignmentName", required: true)]
        public Input<string> PrincipalAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetKustoPoolDatabasePrincipalAssignmentInvokeArgs()
        {
        }
        public static new GetKustoPoolDatabasePrincipalAssignmentInvokeArgs Empty => new GetKustoPoolDatabasePrincipalAssignmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetKustoPoolDatabasePrincipalAssignmentResult
    {
        /// <summary>
        /// The service principal object id in AAD (Azure active directory)
        /// </summary>
        public readonly string AadObjectId;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The principal name
        /// </summary>
        public readonly string PrincipalName;
        /// <summary>
        /// Principal type.
        /// </summary>
        public readonly string PrincipalType;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Database principal role.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The tenant id of the principal
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The tenant name of the principal
        /// </summary>
        public readonly string TenantName;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetKustoPoolDatabasePrincipalAssignmentResult(
            string aadObjectId,

            string azureApiVersion,

            string id,

            string name,

            string principalId,

            string principalName,

            string principalType,

            string provisioningState,

            string role,

            Outputs.SystemDataResponse systemData,

            string? tenantId,

            string tenantName,

            string type)
        {
            AadObjectId = aadObjectId;
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            PrincipalId = principalId;
            PrincipalName = principalName;
            PrincipalType = principalType;
            ProvisioningState = provisioningState;
            Role = role;
            SystemData = systemData;
            TenantId = tenantId;
            TenantName = tenantName;
            Type = type;
        }
    }
}
