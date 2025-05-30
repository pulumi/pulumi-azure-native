// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kusto
{
    public static class GetClusterPrincipalAssignment
    {
        /// <summary>
        /// Gets a Kusto cluster principalAssignment.
        /// 
        /// Uses Azure REST API version 2024-04-13.
        /// 
        /// Other available API versions: 2019-11-09, 2020-02-15, 2020-06-14, 2020-09-18, 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetClusterPrincipalAssignmentResult> InvokeAsync(GetClusterPrincipalAssignmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterPrincipalAssignmentResult>("azure-native:kusto:getClusterPrincipalAssignment", args ?? new GetClusterPrincipalAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Kusto cluster principalAssignment.
        /// 
        /// Uses Azure REST API version 2024-04-13.
        /// 
        /// Other available API versions: 2019-11-09, 2020-02-15, 2020-06-14, 2020-09-18, 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterPrincipalAssignmentResult> Invoke(GetClusterPrincipalAssignmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterPrincipalAssignmentResult>("azure-native:kusto:getClusterPrincipalAssignment", args ?? new GetClusterPrincipalAssignmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Kusto cluster principalAssignment.
        /// 
        /// Uses Azure REST API version 2024-04-13.
        /// 
        /// Other available API versions: 2019-11-09, 2020-02-15, 2020-06-14, 2020-09-18, 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetClusterPrincipalAssignmentResult> Invoke(GetClusterPrincipalAssignmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterPrincipalAssignmentResult>("azure-native:kusto:getClusterPrincipalAssignment", args ?? new GetClusterPrincipalAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterPrincipalAssignmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

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

        public GetClusterPrincipalAssignmentArgs()
        {
        }
        public static new GetClusterPrincipalAssignmentArgs Empty => new GetClusterPrincipalAssignmentArgs();
    }

    public sealed class GetClusterPrincipalAssignmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

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

        public GetClusterPrincipalAssignmentInvokeArgs()
        {
        }
        public static new GetClusterPrincipalAssignmentInvokeArgs Empty => new GetClusterPrincipalAssignmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterPrincipalAssignmentResult
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
        /// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
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
        /// Cluster principal role.
        /// </summary>
        public readonly string Role;
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
        private GetClusterPrincipalAssignmentResult(
            string aadObjectId,

            string azureApiVersion,

            string id,

            string name,

            string principalId,

            string principalName,

            string principalType,

            string provisioningState,

            string role,

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
            TenantId = tenantId;
            TenantName = tenantName;
            Type = type;
        }
    }
}
