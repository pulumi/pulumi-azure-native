// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService
{
    public static class GetLoadBalancer
    {
        /// <summary>
        /// The configurations regarding multiple standard load balancers. If not supplied, single load balancer mode will be used. Multiple standard load balancers mode will be used if at lease one configuration is supplied. There has to be a configuration named `kubernetes`.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("azure-native:containerservice:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithDefaults());

        /// <summary>
        /// The configurations regarding multiple standard load balancers. If not supplied, single load balancer mode will be used. Multiple standard load balancers mode will be used if at lease one configuration is supplied. There has to be a configuration named `kubernetes`.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancerResult>("azure-native:containerservice:getLoadBalancer", args ?? new GetLoadBalancerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The configurations regarding multiple standard load balancers. If not supplied, single load balancer mode will be used. Multiple standard load balancers mode will be used if at lease one configuration is supplied. There has to be a configuration named `kubernetes`.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancerResult>("azure-native:containerservice:getLoadBalancer", args ?? new GetLoadBalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public string LoadBalancerName { get; set; } = null!;

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

        public GetLoadBalancerArgs()
        {
        }
        public static new GetLoadBalancerArgs Empty => new GetLoadBalancerArgs();
    }

    public sealed class GetLoadBalancerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

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

        public GetLoadBalancerInvokeArgs()
        {
        }
        public static new GetLoadBalancerInvokeArgs Empty => new GetLoadBalancerInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        /// <summary>
        /// Whether to automatically place services on the load balancer. If not supplied, the default value is true. If set to false manually, both of the external and the internal load balancer will not be selected for services unless they explicitly target it.
        /// </summary>
        public readonly bool? AllowServicePlacement;
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
        /// Nodes that match this selector will be possible members of this load balancer.
        /// </summary>
        public readonly Outputs.LabelSelectorResponse? NodeSelector;
        /// <summary>
        /// Required field. A string value that must specify the ID of an existing agent pool. All nodes in the given pool will always be added to this load balancer. This agent pool must have at least one node and minCount&gt;=1 for autoscaling operations. An agent pool can only be the primary pool for a single load balancer.
        /// </summary>
        public readonly string PrimaryAgentPoolName;
        /// <summary>
        /// The current provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Only services that must match this selector can be placed on this load balancer.
        /// </summary>
        public readonly Outputs.LabelSelectorResponse? ServiceLabelSelector;
        /// <summary>
        /// Services created in namespaces that match the selector can be placed on this load balancer.
        /// </summary>
        public readonly Outputs.LabelSelectorResponse? ServiceNamespaceSelector;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLoadBalancerResult(
            bool? allowServicePlacement,

            string azureApiVersion,

            string id,

            string name,

            Outputs.LabelSelectorResponse? nodeSelector,

            string primaryAgentPoolName,

            string provisioningState,

            Outputs.LabelSelectorResponse? serviceLabelSelector,

            Outputs.LabelSelectorResponse? serviceNamespaceSelector,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AllowServicePlacement = allowServicePlacement;
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            NodeSelector = nodeSelector;
            PrimaryAgentPoolName = primaryAgentPoolName;
            ProvisioningState = provisioningState;
            ServiceLabelSelector = serviceLabelSelector;
            ServiceNamespaceSelector = serviceNamespaceSelector;
            SystemData = systemData;
            Type = type;
        }
    }
}
