// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork
{
    public static class GetNetworkFunction
    {
        /// <summary>
        /// Gets information about the specified network function resource.
        /// 
        /// Uses Azure REST API version 2024-04-15.
        /// 
        /// Other available API versions: 2022-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridnetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNetworkFunctionResult> InvokeAsync(GetNetworkFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkFunctionResult>("azure-native:hybridnetwork:getNetworkFunction", args ?? new GetNetworkFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified network function resource.
        /// 
        /// Uses Azure REST API version 2024-04-15.
        /// 
        /// Other available API versions: 2022-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridnetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkFunctionResult> Invoke(GetNetworkFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkFunctionResult>("azure-native:hybridnetwork:getNetworkFunction", args ?? new GetNetworkFunctionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified network function resource.
        /// 
        /// Uses Azure REST API version 2024-04-15.
        /// 
        /// Other available API versions: 2022-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridnetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNetworkFunctionResult> Invoke(GetNetworkFunctionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkFunctionResult>("azure-native:hybridnetwork:getNetworkFunction", args ?? new GetNetworkFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkFunctionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network function resource.
        /// </summary>
        [Input("networkFunctionName", required: true)]
        public string NetworkFunctionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkFunctionArgs()
        {
        }
        public static new GetNetworkFunctionArgs Empty => new GetNetworkFunctionArgs();
    }

    public sealed class GetNetworkFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network function resource.
        /// </summary>
        [Input("networkFunctionName", required: true)]
        public Input<string> NetworkFunctionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNetworkFunctionInvokeArgs()
        {
        }
        public static new GetNetworkFunctionInvokeArgs Empty => new GetNetworkFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkFunctionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed identity of the network function.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network function properties.
        /// </summary>
        public readonly Union<Outputs.NetworkFunctionValueWithSecretsResponse, Outputs.NetworkFunctionValueWithoutSecretsResponse> Properties;
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

        [OutputConstructor]
        private GetNetworkFunctionResult(
            string azureApiVersion,

            string? etag,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Union<Outputs.NetworkFunctionValueWithSecretsResponse, Outputs.NetworkFunctionValueWithoutSecretsResponse> properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
