// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceNetworking.V20250101
{
    public static class GetSecurityPoliciesInterface
    {
        /// <summary>
        /// Get a SecurityPolicy
        /// </summary>
        public static Task<GetSecurityPoliciesInterfaceResult> InvokeAsync(GetSecurityPoliciesInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityPoliciesInterfaceResult>("azure-native:servicenetworking/v20250101:getSecurityPoliciesInterface", args ?? new GetSecurityPoliciesInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SecurityPolicy
        /// </summary>
        public static Output<GetSecurityPoliciesInterfaceResult> Invoke(GetSecurityPoliciesInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPoliciesInterfaceResult>("azure-native:servicenetworking/v20250101:getSecurityPoliciesInterface", args ?? new GetSecurityPoliciesInterfaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a SecurityPolicy
        /// </summary>
        public static Output<GetSecurityPoliciesInterfaceResult> Invoke(GetSecurityPoliciesInterfaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPoliciesInterfaceResult>("azure-native:servicenetworking/v20250101:getSecurityPoliciesInterface", args ?? new GetSecurityPoliciesInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityPoliciesInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SecurityPolicy
        /// </summary>
        [Input("securityPolicyName", required: true)]
        public string SecurityPolicyName { get; set; } = null!;

        /// <summary>
        /// traffic controller name for path
        /// </summary>
        [Input("trafficControllerName", required: true)]
        public string TrafficControllerName { get; set; } = null!;

        public GetSecurityPoliciesInterfaceArgs()
        {
        }
        public static new GetSecurityPoliciesInterfaceArgs Empty => new GetSecurityPoliciesInterfaceArgs();
    }

    public sealed class GetSecurityPoliciesInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SecurityPolicy
        /// </summary>
        [Input("securityPolicyName", required: true)]
        public Input<string> SecurityPolicyName { get; set; } = null!;

        /// <summary>
        /// traffic controller name for path
        /// </summary>
        [Input("trafficControllerName", required: true)]
        public Input<string> TrafficControllerName { get; set; } = null!;

        public GetSecurityPoliciesInterfaceInvokeArgs()
        {
        }
        public static new GetSecurityPoliciesInterfaceInvokeArgs Empty => new GetSecurityPoliciesInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityPoliciesInterfaceResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the Traffic Controller Security Policy
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// Provisioning State of Traffic Controller SecurityPolicy Resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Web Application Firewall Policy of the Traffic Controller Security Policy. Single Security Policy can have only one policy type set.
        /// </summary>
        public readonly Outputs.WafPolicyResponse? WafPolicy;

        [OutputConstructor]
        private GetSecurityPoliciesInterfaceResult(
            string id,

            string location,

            string name,

            string policyType,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.WafPolicyResponse? wafPolicy)
        {
            Id = id;
            Location = location;
            Name = name;
            PolicyType = policyType;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            WafPolicy = wafPolicy;
        }
    }
}
