// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Billing
{
    public static class GetBillingRoleAssignmentByDepartment
    {
        /// <summary>
        /// Gets a role assignment for the caller on a department. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2019-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native billing [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetBillingRoleAssignmentByDepartmentResult> InvokeAsync(GetBillingRoleAssignmentByDepartmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBillingRoleAssignmentByDepartmentResult>("azure-native:billing:getBillingRoleAssignmentByDepartment", args ?? new GetBillingRoleAssignmentByDepartmentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a role assignment for the caller on a department. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2019-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native billing [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBillingRoleAssignmentByDepartmentResult> Invoke(GetBillingRoleAssignmentByDepartmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingRoleAssignmentByDepartmentResult>("azure-native:billing:getBillingRoleAssignmentByDepartment", args ?? new GetBillingRoleAssignmentByDepartmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a role assignment for the caller on a department. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2019-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native billing [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBillingRoleAssignmentByDepartmentResult> Invoke(GetBillingRoleAssignmentByDepartmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingRoleAssignmentByDepartmentResult>("azure-native:billing:getBillingRoleAssignmentByDepartment", args ?? new GetBillingRoleAssignmentByDepartmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingRoleAssignmentByDepartmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public string BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a role assignment.
        /// </summary>
        [Input("billingRoleAssignmentName", required: true)]
        public string BillingRoleAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the department.
        /// </summary>
        [Input("departmentName", required: true)]
        public string DepartmentName { get; set; } = null!;

        public GetBillingRoleAssignmentByDepartmentArgs()
        {
        }
        public static new GetBillingRoleAssignmentByDepartmentArgs Empty => new GetBillingRoleAssignmentByDepartmentArgs();
    }

    public sealed class GetBillingRoleAssignmentByDepartmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a role assignment.
        /// </summary>
        [Input("billingRoleAssignmentName", required: true)]
        public Input<string> BillingRoleAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the department.
        /// </summary>
        [Input("departmentName", required: true)]
        public Input<string> DepartmentName { get; set; } = null!;

        public GetBillingRoleAssignmentByDepartmentInvokeArgs()
        {
        }
        public static new GetBillingRoleAssignmentByDepartmentInvokeArgs Empty => new GetBillingRoleAssignmentByDepartmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingRoleAssignmentByDepartmentResult
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
        /// The properties of the billing role assignment.
        /// </summary>
        public readonly Outputs.BillingRoleAssignmentPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Dictionary of metadata associated with the resource. It may not be populated for all resource types. Maximum key/value length supported of 256 characters. Keys/value should not empty value nor null. Keys can not contain &lt; &gt; % &amp; \ ? /
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBillingRoleAssignmentByDepartmentResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.BillingRoleAssignmentPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
