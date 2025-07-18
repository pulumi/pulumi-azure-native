// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry
{
    public static class GetWebhook
    {
        /// <summary>
        /// Gets the properties of the specified webhook.
        /// 
        /// Uses Azure REST API version 2024-11-01-preview.
        /// 
        /// Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-09-01, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebhookResult> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhookResult>("azure-native:containerregistry:getWebhook", args ?? new GetWebhookArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified webhook.
        /// 
        /// Uses Azure REST API version 2024-11-01-preview.
        /// 
        /// Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-09-01, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebhookResult> Invoke(GetWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhookResult>("azure-native:containerregistry:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified webhook.
        /// 
        /// Uses Azure REST API version 2024-11-01-preview.
        /// 
        /// Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-09-01, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebhookResult> Invoke(GetWebhookInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhookResult>("azure-native:containerregistry:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the webhook.
        /// </summary>
        [Input("webhookName", required: true)]
        public string WebhookName { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
        public static new GetWebhookArgs Empty => new GetWebhookArgs();
    }

    public sealed class GetWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the webhook.
        /// </summary>
        [Input("webhookName", required: true)]
        public Input<string> WebhookName { get; set; } = null!;

        public GetWebhookInvokeArgs()
        {
        }
        public static new GetWebhookInvokeArgs Empty => new GetWebhookInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhookResult
    {
        /// <summary>
        /// The list of actions that trigger the webhook to post notifications.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the webhook at the time the operation was called.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// The status of the webhook at the time the operation was called.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebhookResult(
            ImmutableArray<string> actions,

            string azureApiVersion,

            string id,

            string location,

            string name,

            string provisioningState,

            string? scope,

            string? status,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Actions = actions;
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Scope = scope;
            Status = status;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
