// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetExternalNetworkResult> InvokeAsync(GetExternalNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Implements ExternalNetworks GET method.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExternalNetworkResult> Invoke(GetExternalNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements ExternalNetworks GET method.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExternalNetworkResult> Invoke(GetExternalNetworkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkResult>("azure-native:managednetworkfabric:getExternalNetwork", args ?? new GetExternalNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExternalNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the External Network.
        /// </summary>
        [Input("externalNetworkName", required: true)]
        public string ExternalNetworkName { get; set; } = null!;

        /// <summary>
        /// Name of the L3 Isolation Domain.
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
        /// Name of the External Network.
        /// </summary>
        [Input("externalNetworkName", required: true)]
        public Input<string> ExternalNetworkName { get; set; } = null!;

        /// <summary>
        /// Name of the L3 Isolation Domain.
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
        /// Administrative state of the resource.
        /// </summary>
        public readonly string AdministrativeState;
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        public readonly string? Annotation;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Configuration state of the resource.
        /// </summary>
        public readonly string ConfigurationState;
        /// <summary>
        /// Export Route Policy either IPv4 or IPv6.
        /// </summary>
        public readonly Outputs.ExportRoutePolicyResponse? ExportRoutePolicy;
        /// <summary>
        /// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        /// </summary>
        public readonly string? ExportRoutePolicyId;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Import Route Policy either IPv4 or IPv6.
        /// </summary>
        public readonly Outputs.ImportRoutePolicyResponse? ImportRoutePolicy;
        /// <summary>
        /// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        /// </summary>
        public readonly string? ImportRoutePolicyId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ARM Resource ID of the networkToNetworkInterconnectId of the ExternalNetwork resource.
        /// </summary>
        public readonly string? NetworkToNetworkInterconnectId;
        /// <summary>
        /// option A properties object
        /// </summary>
        public readonly Outputs.ExternalNetworkPropertiesResponseOptionAProperties? OptionAProperties;
        /// <summary>
        /// option B properties object
        /// </summary>
        public readonly Outputs.L3OptionBPropertiesResponse? OptionBProperties;
        /// <summary>
        /// Peering option list.
        /// </summary>
        public readonly string PeeringOption;
        /// <summary>
        /// Provisioning state of the resource.
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

            string azureApiVersion,

            string configurationState,

            Outputs.ExportRoutePolicyResponse? exportRoutePolicy,

            string? exportRoutePolicyId,

            string id,

            Outputs.ImportRoutePolicyResponse? importRoutePolicy,

            string? importRoutePolicyId,

            string name,

            string? networkToNetworkInterconnectId,

            Outputs.ExternalNetworkPropertiesResponseOptionAProperties? optionAProperties,

            Outputs.L3OptionBPropertiesResponse? optionBProperties,

            string peeringOption,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AdministrativeState = administrativeState;
            Annotation = annotation;
            AzureApiVersion = azureApiVersion;
            ConfigurationState = configurationState;
            ExportRoutePolicy = exportRoutePolicy;
            ExportRoutePolicyId = exportRoutePolicyId;
            Id = id;
            ImportRoutePolicy = importRoutePolicy;
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
