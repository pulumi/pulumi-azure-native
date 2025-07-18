// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class GetTopicEventSubscriptionFullUrl
    {
        /// <summary>
        /// Get the full endpoint URL for an event subscription for topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTopicEventSubscriptionFullUrlResult> InvokeAsync(GetTopicEventSubscriptionFullUrlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTopicEventSubscriptionFullUrlResult>("azure-native:eventgrid:getTopicEventSubscriptionFullUrl", args ?? new GetTopicEventSubscriptionFullUrlArgs(), options.WithDefaults());

        /// <summary>
        /// Get the full endpoint URL for an event subscription for topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTopicEventSubscriptionFullUrlResult> Invoke(GetTopicEventSubscriptionFullUrlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicEventSubscriptionFullUrlResult>("azure-native:eventgrid:getTopicEventSubscriptionFullUrl", args ?? new GetTopicEventSubscriptionFullUrlInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the full endpoint URL for an event subscription for topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTopicEventSubscriptionFullUrlResult> Invoke(GetTopicEventSubscriptionFullUrlInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicEventSubscriptionFullUrlResult>("azure-native:eventgrid:getTopicEventSubscriptionFullUrl", args ?? new GetTopicEventSubscriptionFullUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicEventSubscriptionFullUrlArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the event subscription.
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public string EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the domain topic.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetTopicEventSubscriptionFullUrlArgs()
        {
        }
        public static new GetTopicEventSubscriptionFullUrlArgs Empty => new GetTopicEventSubscriptionFullUrlArgs();
    }

    public sealed class GetTopicEventSubscriptionFullUrlInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the event subscription.
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public Input<string> EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the domain topic.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetTopicEventSubscriptionFullUrlInvokeArgs()
        {
        }
        public static new GetTopicEventSubscriptionFullUrlInvokeArgs Empty => new GetTopicEventSubscriptionFullUrlInvokeArgs();
    }


    [OutputType]
    public sealed class GetTopicEventSubscriptionFullUrlResult
    {
        /// <summary>
        /// The URL that represents the endpoint of the destination of an event subscription.
        /// </summary>
        public readonly string? EndpointUrl;

        [OutputConstructor]
        private GetTopicEventSubscriptionFullUrlResult(string? endpointUrl)
        {
            EndpointUrl = endpointUrl;
        }
    }
}
