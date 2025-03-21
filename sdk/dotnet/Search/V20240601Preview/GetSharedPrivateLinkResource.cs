// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Search.V20240601Preview
{
    public static class GetSharedPrivateLinkResource
    {
        /// <summary>
        /// Gets the details of the shared private link resource managed by the search service in the given resource group.
        /// </summary>
        public static Task<GetSharedPrivateLinkResourceResult> InvokeAsync(GetSharedPrivateLinkResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSharedPrivateLinkResourceResult>("azure-native:search/v20240601preview:getSharedPrivateLinkResource", args ?? new GetSharedPrivateLinkResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the shared private link resource managed by the search service in the given resource group.
        /// </summary>
        public static Output<GetSharedPrivateLinkResourceResult> Invoke(GetSharedPrivateLinkResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSharedPrivateLinkResourceResult>("azure-native:search/v20240601preview:getSharedPrivateLinkResource", args ?? new GetSharedPrivateLinkResourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the shared private link resource managed by the search service in the given resource group.
        /// </summary>
        public static Output<GetSharedPrivateLinkResourceResult> Invoke(GetSharedPrivateLinkResourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSharedPrivateLinkResourceResult>("azure-native:search/v20240601preview:getSharedPrivateLinkResource", args ?? new GetSharedPrivateLinkResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSharedPrivateLinkResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure AI Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public string SearchServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the shared private link resource managed by the Azure AI Search service within the specified resource group.
        /// </summary>
        [Input("sharedPrivateLinkResourceName", required: true)]
        public string SharedPrivateLinkResourceName { get; set; } = null!;

        public GetSharedPrivateLinkResourceArgs()
        {
        }
        public static new GetSharedPrivateLinkResourceArgs Empty => new GetSharedPrivateLinkResourceArgs();
    }

    public sealed class GetSharedPrivateLinkResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure AI Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public Input<string> SearchServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the shared private link resource managed by the Azure AI Search service within the specified resource group.
        /// </summary>
        [Input("sharedPrivateLinkResourceName", required: true)]
        public Input<string> SharedPrivateLinkResourceName { get; set; } = null!;

        public GetSharedPrivateLinkResourceInvokeArgs()
        {
        }
        public static new GetSharedPrivateLinkResourceInvokeArgs Empty => new GetSharedPrivateLinkResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSharedPrivateLinkResourceResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes the properties of a shared private link resource managed by the Azure AI Search service.
        /// </summary>
        public readonly Outputs.SharedPrivateLinkResourcePropertiesResponse Properties;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSharedPrivateLinkResourceResult(
            string id,

            string name,

            Outputs.SharedPrivateLinkResourcePropertiesResponse properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
