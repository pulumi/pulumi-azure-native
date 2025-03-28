// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.V20250201
{
    public static class GetProjectPolicy
    {
        /// <summary>
        /// Gets a specific project policy.
        /// </summary>
        public static Task<GetProjectPolicyResult> InvokeAsync(GetProjectPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectPolicyResult>("azure-native:devcenter/v20250201:getProjectPolicy", args ?? new GetProjectPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a specific project policy.
        /// </summary>
        public static Output<GetProjectPolicyResult> Invoke(GetProjectPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectPolicyResult>("azure-native:devcenter/v20250201:getProjectPolicy", args ?? new GetProjectPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a specific project policy.
        /// </summary>
        public static Output<GetProjectPolicyResult> Invoke(GetProjectPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectPolicyResult>("azure-native:devcenter/v20250201:getProjectPolicy", args ?? new GetProjectPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public string DevCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the project policy.
        /// </summary>
        [Input("projectPolicyName", required: true)]
        public string ProjectPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetProjectPolicyArgs()
        {
        }
        public static new GetProjectPolicyArgs Empty => new GetProjectPolicyArgs();
    }

    public sealed class GetProjectPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public Input<string> DevCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the project policy.
        /// </summary>
        [Input("projectPolicyName", required: true)]
        public Input<string> ProjectPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetProjectPolicyInvokeArgs()
        {
        }
        public static new GetProjectPolicyInvokeArgs Empty => new GetProjectPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectPolicyResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource policies that are a part of this project policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourcePolicyResponse> ResourcePolicies;
        /// <summary>
        /// Resources that have access to the shared resources that are a part of this project policy.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProjectPolicyResult(
            string id,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.ResourcePolicyResponse> resourcePolicies,

            ImmutableArray<string> scopes,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ResourcePolicies = resourcePolicies;
            Scopes = scopes;
            SystemData = systemData;
            Type = type;
        }
    }
}
