// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer
{
    public static class ListVideoContentToken
    {
        /// <summary>
        /// Generates a streaming token which can be used for accessing content from video content URLs, for a video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Task<ListVideoContentTokenResult> InvokeAsync(ListVideoContentTokenArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListVideoContentTokenResult>("azure-native:videoanalyzer:listVideoContentToken", args ?? new ListVideoContentTokenArgs(), options.WithDefaults());

        /// <summary>
        /// Generates a streaming token which can be used for accessing content from video content URLs, for a video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<ListVideoContentTokenResult> Invoke(ListVideoContentTokenInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListVideoContentTokenResult>("azure-native:videoanalyzer:listVideoContentToken", args ?? new ListVideoContentTokenInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Generates a streaming token which can be used for accessing content from video content URLs, for a video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<ListVideoContentTokenResult> Invoke(ListVideoContentTokenInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListVideoContentTokenResult>("azure-native:videoanalyzer:listVideoContentToken", args ?? new ListVideoContentTokenInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVideoContentTokenArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Video Analyzer account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Video name.
        /// </summary>
        [Input("videoName", required: true)]
        public string VideoName { get; set; } = null!;

        public ListVideoContentTokenArgs()
        {
        }
        public static new ListVideoContentTokenArgs Empty => new ListVideoContentTokenArgs();
    }

    public sealed class ListVideoContentTokenInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Video Analyzer account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Video name.
        /// </summary>
        [Input("videoName", required: true)]
        public Input<string> VideoName { get; set; } = null!;

        public ListVideoContentTokenInvokeArgs()
        {
        }
        public static new ListVideoContentTokenInvokeArgs Empty => new ListVideoContentTokenInvokeArgs();
    }


    [OutputType]
    public sealed class ListVideoContentTokenResult
    {
        /// <summary>
        /// The content token expiration date in ISO8601 format (eg. 2021-01-01T00:00:00Z).
        /// </summary>
        public readonly string ExpirationDate;
        /// <summary>
        /// The content token value to be added to the video content URL as the value for the "token" query string parameter. The token is specific to a single video.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private ListVideoContentTokenResult(
            string expirationDate,

            string token)
        {
            ExpirationDate = expirationDate;
            Token = token;
        }
    }
}
