// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    public static class GetNetworkInterface
    {
        /// <summary>
        /// Get the Network Interface resource details.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNetworkInterfaceResult> InvokeAsync(GetNetworkInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkInterfaceResult>("azure-native:managednetworkfabric:getNetworkInterface", args ?? new GetNetworkInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Network Interface resource details.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkInterfaceResult> Invoke(GetNetworkInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkInterfaceResult>("azure-native:managednetworkfabric:getNetworkInterface", args ?? new GetNetworkInterfaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Network Interface resource details.
        /// 
        /// Uses Azure REST API version 2023-06-15.
        /// 
        /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkInterfaceResult> Invoke(GetNetworkInterfaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkInterfaceResult>("azure-native:managednetworkfabric:getNetworkInterface", args ?? new GetNetworkInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Network Device.
        /// </summary>
        [Input("networkDeviceName", required: true)]
        public string NetworkDeviceName { get; set; } = null!;

        /// <summary>
        /// Name of the Network Interface.
        /// </summary>
        [Input("networkInterfaceName", required: true)]
        public string NetworkInterfaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkInterfaceArgs()
        {
        }
        public static new GetNetworkInterfaceArgs Empty => new GetNetworkInterfaceArgs();
    }

    public sealed class GetNetworkInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Network Device.
        /// </summary>
        [Input("networkDeviceName", required: true)]
        public Input<string> NetworkDeviceName { get; set; } = null!;

        /// <summary>
        /// Name of the Network Interface.
        /// </summary>
        [Input("networkInterfaceName", required: true)]
        public Input<string> NetworkInterfaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNetworkInterfaceInvokeArgs()
        {
        }
        public static new GetNetworkInterfaceInvokeArgs Empty => new GetNetworkInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkInterfaceResult
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
        /// The ARM resource id of the interface or compute server its connected to.
        /// </summary>
        public readonly string ConnectedTo;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Interface Type. Example: Management/Data
        /// </summary>
        public readonly string InterfaceType;
        /// <summary>
        /// IPv4Address of the interface.
        /// </summary>
        public readonly string Ipv4Address;
        /// <summary>
        /// IPv6Address of the interface.
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Physical Identifier of the network interface.
        /// </summary>
        public readonly string PhysicalIdentifier;
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
        private GetNetworkInterfaceResult(
            string administrativeState,

            string? annotation,

            string azureApiVersion,

            string connectedTo,

            string id,

            string interfaceType,

            string ipv4Address,

            string ipv6Address,

            string name,

            string physicalIdentifier,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AdministrativeState = administrativeState;
            Annotation = annotation;
            AzureApiVersion = azureApiVersion;
            ConnectedTo = connectedTo;
            Id = id;
            InterfaceType = interfaceType;
            Ipv4Address = ipv4Address;
            Ipv6Address = ipv6Address;
            Name = name;
            PhysicalIdentifier = physicalIdentifier;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
