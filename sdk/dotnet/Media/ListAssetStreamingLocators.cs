// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media
{
    public static class ListAssetStreamingLocators
    {
        /// <summary>
        /// Lists Streaming Locators which are associated with this asset.
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// </summary>
        public static Task<ListAssetStreamingLocatorsResult> InvokeAsync(ListAssetStreamingLocatorsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListAssetStreamingLocatorsResult>("azure-native:media:listAssetStreamingLocators", args ?? new ListAssetStreamingLocatorsArgs(), options.WithDefaults());

        /// <summary>
        /// Lists Streaming Locators which are associated with this asset.
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// </summary>
        public static Output<ListAssetStreamingLocatorsResult> Invoke(ListAssetStreamingLocatorsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListAssetStreamingLocatorsResult>("azure-native:media:listAssetStreamingLocators", args ?? new ListAssetStreamingLocatorsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists Streaming Locators which are associated with this asset.
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// </summary>
        public static Output<ListAssetStreamingLocatorsResult> Invoke(ListAssetStreamingLocatorsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListAssetStreamingLocatorsResult>("azure-native:media:listAssetStreamingLocators", args ?? new ListAssetStreamingLocatorsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListAssetStreamingLocatorsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The Asset name.
        /// </summary>
        [Input("assetName", required: true)]
        public string AssetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListAssetStreamingLocatorsArgs()
        {
        }
        public static new ListAssetStreamingLocatorsArgs Empty => new ListAssetStreamingLocatorsArgs();
    }

    public sealed class ListAssetStreamingLocatorsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The Asset name.
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListAssetStreamingLocatorsInvokeArgs()
        {
        }
        public static new ListAssetStreamingLocatorsInvokeArgs Empty => new ListAssetStreamingLocatorsInvokeArgs();
    }


    [OutputType]
    public sealed class ListAssetStreamingLocatorsResult
    {
        /// <summary>
        /// The list of Streaming Locators.
        /// </summary>
        public readonly ImmutableArray<Outputs.AssetStreamingLocatorResponse> StreamingLocators;

        [OutputConstructor]
        private ListAssetStreamingLocatorsResult(ImmutableArray<Outputs.AssetStreamingLocatorResponse> streamingLocators)
        {
            StreamingLocators = streamingLocators;
        }
    }
}
