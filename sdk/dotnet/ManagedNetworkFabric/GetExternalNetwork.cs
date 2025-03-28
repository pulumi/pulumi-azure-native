// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    public static class GetExternalNetwork
    {
        /// <summary>
        /// Implements ExternalNetworks GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Task<GetExternalNetworkResult> InvokeAsync(GetExternalNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Implements ExternalNetworks GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Output<GetExternalNetworkResult> Invoke(GetExternalNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements ExternalNetworks GET method.
        /// 
        /// Uses Azure REST API version 2023-02-01-preview.
        /// 
        /// Other available API versions: 2023-06-15.
        /// </summary>
        public static Output<GetExternalNetworkResult> Invoke(GetExternalNetworkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExternalNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ExternalNetwork
        /// </summary>
        [Input("externalNetworkName", required: true)]
        public string ExternalNetworkName { get; set; } = null!;

        /// <summary>
        /// Name of the L3IsolationDomain
        /// </summary>
        [Input("l3IsolationDomainName", required: true)]
        public string L3IsolationDomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExternalNetworkArgs()
        {
        }
        public static new GetExternalNetworkArgs Empty => new GetExternalNetworkArgs();
    }

    public sealed class GetExternalNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ExternalNetwork
        /// </summary>
        [Input("externalNetworkName", required: true)]
        public Input<string> ExternalNetworkName { get; set; } = null!;

        /// <summary>
        /// Name of the L3IsolationDomain
        /// </summary>
        [Input("l3IsolationDomainName", required: true)]
        public Input<string> L3IsolationDomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetExternalNetworkInvokeArgs()
        {
        }
        public static new GetExternalNetworkInvokeArgs Empty => new GetExternalNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetExternalNetworkResult
    {
        /// <summary>
        /// AdministrativeState of the externalNetwork. Example: Enabled | Disabled.
        /// </summary>
        public readonly string AdministrativeState;
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        public readonly string? Annotation;
        /// <summary>
        /// List of resources the externalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
        /// </summary>
        public readonly ImmutableArray<string> DisabledOnResources;
        /// <summary>
        /// ARM resource ID of exportRoutePolicy.
        /// </summary>
        public readonly string? ExportRoutePolicyId;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ARM resource ID of importRoutePolicy.
        /// </summary>
        public readonly string? ImportRoutePolicyId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets the networkToNetworkInterconnectId of the resource.
        /// </summary>
        public readonly string NetworkToNetworkInterconnectId;
        /// <summary>
        /// option A properties object
        /// </summary>
        public readonly Outputs.ExternalNetworkPropertiesResponseOptionAProperties? OptionAProperties;
        /// <summary>
        /// option B properties object
        /// </summary>
        public readonly Outputs.OptionBPropertiesResponse? OptionBProperties;
        /// <summary>
        /// Peering option list.
        /// </summary>
        public readonly string PeeringOption;
        /// <summary>
        /// Gets the provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetExternalNetworkResult(
            string administrativeState,

            string? annotation,

            ImmutableArray<string> disabledOnResources,

            string? exportRoutePolicyId,

            string id,

            string? importRoutePolicyId,

            string name,

            string networkToNetworkInterconnectId,

            Outputs.ExternalNetworkPropertiesResponseOptionAProperties? optionAProperties,

            Outputs.OptionBPropertiesResponse? optionBProperties,

            string peeringOption,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AdministrativeState = administrativeState;
            Annotation = annotation;
            DisabledOnResources = disabledOnResources;
            ExportRoutePolicyId = exportRoutePolicyId;
            Id = id;
            ImportRoutePolicyId = importRoutePolicyId;
            Name = name;
            NetworkToNetworkInterconnectId = networkToNetworkInterconnectId;
            OptionAProperties = optionAProperties;
            OptionBProperties = optionBProperties;
            PeeringOption = peeringOption;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
