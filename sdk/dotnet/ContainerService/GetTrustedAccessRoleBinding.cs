// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService
{
    public static class GetTrustedAccessRoleBinding
    {
        /// <summary>
        /// Defines binding between a resource and role
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-02-preview, 2025-01-01, 2025-01-02-preview, 2025-02-01, 2025-02-02-preview, 2025-03-01, 2025-03-02-preview, 2025-04-01, 2025-04-02-preview, 2025-05-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTrustedAccessRoleBindingResult> InvokeAsync(GetTrustedAccessRoleBindingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrustedAccessRoleBindingResult>("azure-native:containerservice:getTrustedAccessRoleBinding", args ?? new GetTrustedAccessRoleBindingArgs(), options.WithDefaults());

        /// <summary>
        /// Defines binding between a resource and role
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-02-preview, 2025-01-01, 2025-01-02-preview, 2025-02-01, 2025-02-02-preview, 2025-03-01, 2025-03-02-preview, 2025-04-01, 2025-04-02-preview, 2025-05-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTrustedAccessRoleBindingResult> Invoke(GetTrustedAccessRoleBindingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrustedAccessRoleBindingResult>("azure-native:containerservice:getTrustedAccessRoleBinding", args ?? new GetTrustedAccessRoleBindingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Defines binding between a resource and role
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-02-preview, 2025-01-01, 2025-01-02-preview, 2025-02-01, 2025-02-02-preview, 2025-03-01, 2025-03-02-preview, 2025-04-01, 2025-04-02-preview, 2025-05-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTrustedAccessRoleBindingResult> Invoke(GetTrustedAccessRoleBindingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrustedAccessRoleBindingResult>("azure-native:containerservice:getTrustedAccessRoleBinding", args ?? new GetTrustedAccessRoleBindingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrustedAccessRoleBindingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        /// <summary>
        /// The name of trusted access role binding.
        /// </summary>
        [Input("trustedAccessRoleBindingName", required: true)]
        public string TrustedAccessRoleBindingName { get; set; } = null!;

        public GetTrustedAccessRoleBindingArgs()
        {
        }
        public static new GetTrustedAccessRoleBindingArgs Empty => new GetTrustedAccessRoleBindingArgs();
    }

    public sealed class GetTrustedAccessRoleBindingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// The name of trusted access role binding.
        /// </summary>
        [Input("trustedAccessRoleBindingName", required: true)]
        public Input<string> TrustedAccessRoleBindingName { get; set; } = null!;

        public GetTrustedAccessRoleBindingInvokeArgs()
        {
        }
        public static new GetTrustedAccessRoleBindingInvokeArgs Empty => new GetTrustedAccessRoleBindingInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrustedAccessRoleBindingResult
    {
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
        /// The current provisioning state of trusted access role binding.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// The ARM resource ID of source resource that trusted access is configured for.
        /// </summary>
        public readonly string SourceResourceId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTrustedAccessRoleBindingResult(
            string azureApiVersion,

            string id,

            string name,

            string provisioningState,

            ImmutableArray<string> roles,

            string sourceResourceId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            Roles = roles;
            SourceResourceId = sourceResourceId;
            SystemData = systemData;
            Type = type;
        }
    }
}
