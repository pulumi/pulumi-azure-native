// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class ListNamespaceTopicSharedAccessKeys
    {
        /// <summary>
        /// List the two keys used to publish to a namespace topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListNamespaceTopicSharedAccessKeysResult> InvokeAsync(ListNamespaceTopicSharedAccessKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListNamespaceTopicSharedAccessKeysResult>("azure-native:eventgrid:listNamespaceTopicSharedAccessKeys", args ?? new ListNamespaceTopicSharedAccessKeysArgs(), options.WithDefaults());

        /// <summary>
        /// List the two keys used to publish to a namespace topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListNamespaceTopicSharedAccessKeysResult> Invoke(ListNamespaceTopicSharedAccessKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListNamespaceTopicSharedAccessKeysResult>("azure-native:eventgrid:listNamespaceTopicSharedAccessKeys", args ?? new ListNamespaceTopicSharedAccessKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List the two keys used to publish to a namespace topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListNamespaceTopicSharedAccessKeysResult> Invoke(ListNamespaceTopicSharedAccessKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListNamespaceTopicSharedAccessKeysResult>("azure-native:eventgrid:listNamespaceTopicSharedAccessKeys", args ?? new ListNamespaceTopicSharedAccessKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListNamespaceTopicSharedAccessKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the topic.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public ListNamespaceTopicSharedAccessKeysArgs()
        {
        }
        public static new ListNamespaceTopicSharedAccessKeysArgs Empty => new ListNamespaceTopicSharedAccessKeysArgs();
    }

    public sealed class ListNamespaceTopicSharedAccessKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the topic.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public ListNamespaceTopicSharedAccessKeysInvokeArgs()
        {
        }
        public static new ListNamespaceTopicSharedAccessKeysInvokeArgs Empty => new ListNamespaceTopicSharedAccessKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListNamespaceTopicSharedAccessKeysResult
    {
        /// <summary>
        /// Shared access key1 for the topic.
        /// </summary>
        public readonly string? Key1;
        /// <summary>
        /// Shared access key2 for the topic.
        /// </summary>
        public readonly string? Key2;

        [OutputConstructor]
        private ListNamespaceTopicSharedAccessKeysResult(
            string? key1,

            string? key2)
        {
            Key1 = key1;
            Key2 = key2;
        }
    }
}
