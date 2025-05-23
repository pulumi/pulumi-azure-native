// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedVMwarevSphere
{
    public static class GetGuestAgent
    {
        /// <summary>
        /// Implements GuestAgent GET method.
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2022-07-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetGuestAgentResult> InvokeAsync(GetGuestAgentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGuestAgentResult>("azure-native:connectedvmwarevsphere:getGuestAgent", args ?? new GetGuestAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Implements GuestAgent GET method.
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2022-07-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetGuestAgentResult> Invoke(GetGuestAgentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestAgentResult>("azure-native:connectedvmwarevsphere:getGuestAgent", args ?? new GetGuestAgentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements GuestAgent GET method.
        /// 
        /// Uses Azure REST API version 2023-03-01-preview.
        /// 
        /// Other available API versions: 2022-07-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetGuestAgentResult> Invoke(GetGuestAgentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestAgentResult>("azure-native:connectedvmwarevsphere:getGuestAgent", args ?? new GetGuestAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGuestAgentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the GuestAgent.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the vm.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public string VirtualMachineName { get; set; } = null!;

        public GetGuestAgentArgs()
        {
        }
        public static new GetGuestAgentArgs Empty => new GetGuestAgentArgs();
    }

    public sealed class GetGuestAgentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the GuestAgent.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the vm.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public Input<string> VirtualMachineName { get; set; } = null!;

        public GetGuestAgentInvokeArgs()
        {
        }
        public static new GetGuestAgentInvokeArgs Empty => new GetGuestAgentInvokeArgs();
    }


    [OutputType]
    public sealed class GetGuestAgentResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
        /// Gets the provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Gets or sets the guest agent status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceStatusResponse> Statuses;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets or sets a unique identifier for this resource.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetGuestAgentResult(
            string azureApiVersion,

            Outputs.GuestCredentialResponse? credentials,

            string customResourceName,

            Outputs.HttpProxyConfigurationResponse? httpProxyConfig,

            string id,

            string name,

            string? privateLinkScopeResourceId,

            string? provisioningAction,

            string provisioningState,

            string status,

            ImmutableArray<Outputs.ResourceStatusResponse> statuses,

            Outputs.SystemDataResponse systemData,

            string type,

            string uuid)
        {
            AzureApiVersion = azureApiVersion;
            Credentials = credentials;
            CustomResourceName = customResourceName;
            HttpProxyConfig = httpProxyConfig;
            Id = id;
            Name = name;
            PrivateLinkScopeResourceId = privateLinkScopeResourceId;
            ProvisioningAction = provisioningAction;
            ProvisioningState = provisioningState;
            Status = status;
            Statuses = statuses;
            SystemData = systemData;
            Type = type;
            Uuid = uuid;
        }
    }
}
