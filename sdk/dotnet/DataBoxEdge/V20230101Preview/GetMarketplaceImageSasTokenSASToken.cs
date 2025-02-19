// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.V20230101Preview
{
    public static class GetMarketplaceImageSasTokenSASToken
    {
        public static Task<GetMarketplaceImageSasTokenSASTokenResult> InvokeAsync(GetMarketplaceImageSasTokenSASTokenArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMarketplaceImageSasTokenSASTokenResult>("azure-native:databoxedge/v20230101preview:getMarketplaceImageSasTokenSASToken", args ?? new GetMarketplaceImageSasTokenSASTokenArgs(), options.WithDefaults());

        public static Output<GetMarketplaceImageSasTokenSASTokenResult> Invoke(GetMarketplaceImageSasTokenSASTokenInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMarketplaceImageSasTokenSASTokenResult>("azure-native:databoxedge/v20230101preview:getMarketplaceImageSasTokenSASToken", args ?? new GetMarketplaceImageSasTokenSASTokenInvokeArgs(), options.WithDefaults());

        public static Output<GetMarketplaceImageSasTokenSASTokenResult> Invoke(GetMarketplaceImageSasTokenSASTokenInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMarketplaceImageSasTokenSASTokenResult>("azure-native:databoxedge/v20230101preview:getMarketplaceImageSasTokenSASToken", args ?? new GetMarketplaceImageSasTokenSASTokenInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMarketplaceImageSasTokenSASTokenArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        [Input("offerName", required: true)]
        public string OfferName { get; set; } = null!;

        [Input("publisherName", required: true)]
        public string PublisherName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("skuName", required: true)]
        public string SkuName { get; set; } = null!;

        [Input("versionName", required: true)]
        public string VersionName { get; set; } = null!;

        public GetMarketplaceImageSasTokenSASTokenArgs()
        {
        }
        public static new GetMarketplaceImageSasTokenSASTokenArgs Empty => new GetMarketplaceImageSasTokenSASTokenArgs();
    }

    public sealed class GetMarketplaceImageSasTokenSASTokenInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("offerName", required: true)]
        public Input<string> OfferName { get; set; } = null!;

        [Input("publisherName", required: true)]
        public Input<string> PublisherName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("versionName", required: true)]
        public Input<string> VersionName { get; set; } = null!;

        public GetMarketplaceImageSasTokenSASTokenInvokeArgs()
        {
        }
        public static new GetMarketplaceImageSasTokenSASTokenInvokeArgs Empty => new GetMarketplaceImageSasTokenSASTokenInvokeArgs();
    }


    [OutputType]
    public sealed class GetMarketplaceImageSasTokenSASTokenResult
    {
        public readonly string? SasUri;
        public readonly string? Status;

        [OutputConstructor]
        private GetMarketplaceImageSasTokenSASTokenResult(
            string? sasUri,

            string? status)
        {
            SasUri = sasUri;
            Status = status;
        }
    }
}
