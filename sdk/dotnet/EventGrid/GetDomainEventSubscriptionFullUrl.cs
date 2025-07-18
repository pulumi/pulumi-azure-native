// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class GetDomainEventSubscriptionFullUrl
    {
        /// <summary>
        /// Get the full endpoint URL for an event subscription for domain.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDomainEventSubscriptionFullUrlResult> InvokeAsync(GetDomainEventSubscriptionFullUrlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainEventSubscriptionFullUrlResult>("azure-native:eventgrid:getDomainEventSubscriptionFullUrl", args ?? new GetDomainEventSubscriptionFullUrlArgs(), options.WithDefaults());

        /// <summary>
        /// Get the full endpoint URL for an event subscription for domain.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDomainEventSubscriptionFullUrlResult> Invoke(GetDomainEventSubscriptionFullUrlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainEventSubscriptionFullUrlResult>("azure-native:eventgrid:getDomainEventSubscriptionFullUrl", args ?? new GetDomainEventSubscriptionFullUrlInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the full endpoint URL for an event subscription for domain.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDomainEventSubscriptionFullUrlResult> Invoke(GetDomainEventSubscriptionFullUrlInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainEventSubscriptionFullUrlResult>("azure-native:eventgrid:getDomainEventSubscriptionFullUrl", args ?? new GetDomainEventSubscriptionFullUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainEventSubscriptionFullUrlArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain topic.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

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

        public GetDomainEventSubscriptionFullUrlArgs()
        {
        }
        public static new GetDomainEventSubscriptionFullUrlArgs Empty => new GetDomainEventSubscriptionFullUrlArgs();
    }

    public sealed class GetDomainEventSubscriptionFullUrlInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain topic.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

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

        public GetDomainEventSubscriptionFullUrlInvokeArgs()
        {
        }
        public static new GetDomainEventSubscriptionFullUrlInvokeArgs Empty => new GetDomainEventSubscriptionFullUrlInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainEventSubscriptionFullUrlResult
    {
        /// <summary>
        /// The URL that represents the endpoint of the destination of an event subscription.
        /// </summary>
        public readonly string? EndpointUrl;

        [OutputConstructor]
        private GetDomainEventSubscriptionFullUrlResult(string? endpointUrl)
        {
            EndpointUrl = endpointUrl;
        }
    }
}
