// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer
{
    public static class GetVideo
    {
        /// <summary>
        /// Retrieves an existing video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Task<GetVideoResult> InvokeAsync(GetVideoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVideoResult>("azure-native:videoanalyzer:getVideo", args ?? new GetVideoArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an existing video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<GetVideoResult> Invoke(GetVideoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVideoResult>("azure-native:videoanalyzer:getVideo", args ?? new GetVideoInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an existing video resource with the given name.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<GetVideoResult> Invoke(GetVideoInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVideoResult>("azure-native:videoanalyzer:getVideo", args ?? new GetVideoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVideoArgs : global::Pulumi.InvokeArgs
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

        public GetVideoArgs()
        {
        }
        public static new GetVideoArgs Empty => new GetVideoArgs();
    }

    public sealed class GetVideoInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetVideoInvokeArgs()
        {
        }
        public static new GetVideoInvokeArgs Empty => new GetVideoInvokeArgs();
    }


    [OutputType]
    public sealed class GetVideoResult
    {
        /// <summary>
        /// Video archival properties.
        /// </summary>
        public readonly Outputs.VideoArchivalResponse? Archival;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Set of URLs to the video content.
        /// </summary>
        public readonly Outputs.VideoContentUrlsResponse ContentUrls;
        /// <summary>
        /// Optional video description provided by the user. Value can be up to 2048 characters long.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Video flags contain information about the available video actions and its dynamic properties based on the current video state.
        /// </summary>
        public readonly Outputs.VideoFlagsResponse Flags;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Contains information about the video and audio content.
        /// </summary>
        public readonly Outputs.VideoMediaInfoResponse? MediaInfo;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Optional video title provided by the user. Value can be up to 256 characters long.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVideoResult(
            Outputs.VideoArchivalResponse? archival,

            string azureApiVersion,

            Outputs.VideoContentUrlsResponse contentUrls,

            string? description,

            Outputs.VideoFlagsResponse flags,

            string id,

            Outputs.VideoMediaInfoResponse? mediaInfo,

            string name,

            Outputs.SystemDataResponse systemData,

            string? title,

            string type)
        {
            Archival = archival;
            AzureApiVersion = azureApiVersion;
            ContentUrls = contentUrls;
            Description = description;
            Flags = flags;
            Id = id;
            MediaInfo = mediaInfo;
            Name = name;
            SystemData = systemData;
            Title = title;
            Type = type;
        }
    }
}
