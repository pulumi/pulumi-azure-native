// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetNetworkVirtualApplianceConnection
    {
        /// <summary>
        /// Retrieves the details of specified NVA connection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNetworkVirtualApplianceConnectionResult> InvokeAsync(GetNetworkVirtualApplianceConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkVirtualApplianceConnectionResult>("azure-native:network:getNetworkVirtualApplianceConnection", args ?? new GetNetworkVirtualApplianceConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of specified NVA connection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkVirtualApplianceConnectionResult> Invoke(GetNetworkVirtualApplianceConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkVirtualApplianceConnectionResult>("azure-native:network:getNetworkVirtualApplianceConnection", args ?? new GetNetworkVirtualApplianceConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of specified NVA connection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkVirtualApplianceConnectionResult> Invoke(GetNetworkVirtualApplianceConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkVirtualApplianceConnectionResult>("azure-native:network:getNetworkVirtualApplianceConnection", args ?? new GetNetworkVirtualApplianceConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkVirtualApplianceConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NVA connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the Network Virtual Appliance.
        /// </summary>
        [Input("networkVirtualApplianceName", required: true)]
        public string NetworkVirtualApplianceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkVirtualApplianceConnectionArgs()
        {
        }
        public static new GetNetworkVirtualApplianceConnectionArgs Empty => new GetNetworkVirtualApplianceConnectionArgs();
    }

    public sealed class GetNetworkVirtualApplianceConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NVA connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the Network Virtual Appliance.
        /// </summary>
        [Input("networkVirtualApplianceName", required: true)]
        public Input<string> NetworkVirtualApplianceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNetworkVirtualApplianceConnectionInvokeArgs()
        {
        }
        public static new GetNetworkVirtualApplianceConnectionInvokeArgs Empty => new GetNetworkVirtualApplianceConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkVirtualApplianceConnectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of the express route connection.
        /// </summary>
        public readonly Outputs.NetworkVirtualApplianceConnectionPropertiesResponse Properties;

        [OutputConstructor]
        private GetNetworkVirtualApplianceConnectionResult(
            string azureApiVersion,

            string? id,

            string? name,

            Outputs.NetworkVirtualApplianceConnectionPropertiesResponse properties)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
        }
    }
}
