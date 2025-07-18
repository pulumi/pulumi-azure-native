// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedCache
{
    public static class GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistory
    {
        /// <summary>
        /// This api gets ispCacheNode resource issues details histrory information
        /// 
        /// Uses Azure REST API version 2024-11-30-preview.
        /// </summary>
        public static Task<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult> InvokeAsync(GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult>("azure-native:connectedcache:getIspCacheNodesOperationCacheNodeMccIssueDetailsHistory", args ?? new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs(), options.WithDefaults());

        /// <summary>
        /// This api gets ispCacheNode resource issues details histrory information
        /// 
        /// Uses Azure REST API version 2024-11-30-preview.
        /// </summary>
        public static Output<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult> Invoke(GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult>("azure-native:connectedcache:getIspCacheNodesOperationCacheNodeMccIssueDetailsHistory", args ?? new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This api gets ispCacheNode resource issues details histrory information
        /// 
        /// Uses Azure REST API version 2024-11-30-preview.
        /// </summary>
        public static Output<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult> Invoke(GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult>("azure-native:connectedcache:getIspCacheNodesOperationCacheNodeMccIssueDetailsHistory", args ?? new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ConnectedCache resource
        /// </summary>
        [Input("cacheNodeResourceName", required: true)]
        public string CacheNodeResourceName { get; set; } = null!;

        /// <summary>
        /// Name of the Customer resource
        /// </summary>
        [Input("customerResourceName", required: true)]
        public string CustomerResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs()
        {
        }
        public static new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs Empty => new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryArgs();
    }

    public sealed class GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ConnectedCache resource
        /// </summary>
        [Input("cacheNodeResourceName", required: true)]
        public Input<string> CacheNodeResourceName { get; set; } = null!;

        /// <summary>
        /// Name of the Customer resource
        /// </summary>
        [Input("customerResourceName", required: true)]
        public Input<string> CustomerResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs()
        {
        }
        public static new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs Empty => new GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
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
        /// Mcc cache node resource issue history properties.
        /// </summary>
        public readonly Outputs.MccCacheNodeIssueHistoryPropertiesResponse Properties;
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
        private GetIspCacheNodesOperationCacheNodeMccIssueDetailsHistoryResult(
            string id,

            string location,

            string name,

            Outputs.MccCacheNodeIssueHistoryPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
