// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm.V20240601
{
    public static class GetGuestAgent
    {
        /// <summary>
        /// Implements GuestAgent GET method.
        /// </summary>
        public static Task<GetGuestAgentResult> InvokeAsync(GetGuestAgentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGuestAgentResult>("azure-native:scvmm/v20240601:getGuestAgent", args ?? new GetGuestAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Implements GuestAgent GET method.
        /// </summary>
        public static Output<GetGuestAgentResult> Invoke(GetGuestAgentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestAgentResult>("azure-native:scvmm/v20240601:getGuestAgent", args ?? new GetGuestAgentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements GuestAgent GET method.
        /// </summary>
        public static Output<GetGuestAgentResult> Invoke(GetGuestAgentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestAgentResult>("azure-native:scvmm/v20240601:getGuestAgent", args ?? new GetGuestAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGuestAgentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetGuestAgentArgs()
        {
        }
        public static new GetGuestAgentArgs Empty => new GetGuestAgentArgs();
    }

    public sealed class GetGuestAgentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetGuestAgentInvokeArgs()
        {
        }
        public static new GetGuestAgentInvokeArgs Empty => new GetGuestAgentInvokeArgs();
    }


    [OutputType]
    public sealed class GetGuestAgentResult
    {
        /// <summary>
        /// Username / Password Credentials to provision guest agent.
        /// </summary>
        public readonly Outputs.GuestCredentialResponse? Credentials;
        /// <summary>
        /// Gets the name of the corresponding resource in Kubernetes.
        /// </summary>
        public readonly string CustomResourceName;
        /// <summary>
        /// HTTP Proxy configuration for the VM.
        /// </summary>
        public readonly Outputs.HttpProxyConfigurationResponse? HttpProxyConfig;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource id of the private link scope this machine is assigned to, if any.
        /// </summary>
        public readonly string? PrivateLinkScopeResourceId;
        /// <summary>
        /// Gets or sets the guest agent provisioning action.
        /// </summary>
        public readonly string? ProvisioningAction;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Gets the guest agent status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets a unique identifier for this resource.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetGuestAgentResult(
            Outputs.GuestCredentialResponse? credentials,

            string customResourceName,

            Outputs.HttpProxyConfigurationResponse? httpProxyConfig,

            string id,

            string name,

            string? privateLinkScopeResourceId,

            string? provisioningAction,

            string provisioningState,

            string status,

            Outputs.SystemDataResponse systemData,

            string type,

            string uuid)
        {
            Credentials = credentials;
            CustomResourceName = customResourceName;
            HttpProxyConfig = httpProxyConfig;
            Id = id;
            Name = name;
            PrivateLinkScopeResourceId = privateLinkScopeResourceId;
            ProvisioningAction = provisioningAction;
            ProvisioningState = provisioningState;
            Status = status;
            SystemData = systemData;
            Type = type;
            Uuid = uuid;
        }
    }
}
