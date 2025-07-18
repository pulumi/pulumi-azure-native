// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class GetPartnerTopicEventSubscriptionDeliveryAttributes
    {
        /// <summary>
        /// Get all delivery attributes for an event subscription of a partner topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetPartnerTopicEventSubscriptionDeliveryAttributesResult> InvokeAsync(GetPartnerTopicEventSubscriptionDeliveryAttributesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPartnerTopicEventSubscriptionDeliveryAttributesResult>("azure-native:eventgrid:getPartnerTopicEventSubscriptionDeliveryAttributes", args ?? new GetPartnerTopicEventSubscriptionDeliveryAttributesArgs(), options.WithDefaults());

        /// <summary>
        /// Get all delivery attributes for an event subscription of a partner topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPartnerTopicEventSubscriptionDeliveryAttributesResult> Invoke(GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPartnerTopicEventSubscriptionDeliveryAttributesResult>("azure-native:eventgrid:getPartnerTopicEventSubscriptionDeliveryAttributes", args ?? new GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get all delivery attributes for an event subscription of a partner topic.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPartnerTopicEventSubscriptionDeliveryAttributesResult> Invoke(GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPartnerTopicEventSubscriptionDeliveryAttributesResult>("azure-native:eventgrid:getPartnerTopicEventSubscriptionDeliveryAttributes", args ?? new GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPartnerTopicEventSubscriptionDeliveryAttributesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the event subscription.
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public string EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// Name of the partner topic.
        /// </summary>
        [Input("partnerTopicName", required: true)]
        public string PartnerTopicName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPartnerTopicEventSubscriptionDeliveryAttributesArgs()
        {
        }
        public static new GetPartnerTopicEventSubscriptionDeliveryAttributesArgs Empty => new GetPartnerTopicEventSubscriptionDeliveryAttributesArgs();
    }

    public sealed class GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the event subscription.
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public Input<string> EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// Name of the partner topic.
        /// </summary>
        [Input("partnerTopicName", required: true)]
        public Input<string> PartnerTopicName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs()
        {
        }
        public static new GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs Empty => new GetPartnerTopicEventSubscriptionDeliveryAttributesInvokeArgs();
    }


    [OutputType]
    public sealed class GetPartnerTopicEventSubscriptionDeliveryAttributesResult
    {
        /// <summary>
        /// A collection of DeliveryAttributeMapping
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.DynamicDeliveryAttributeMappingResponse, Outputs.StaticDeliveryAttributeMappingResponse>> Value;

        [OutputConstructor]
        private GetPartnerTopicEventSubscriptionDeliveryAttributesResult(ImmutableArray<Union<Outputs.DynamicDeliveryAttributeMappingResponse, Outputs.StaticDeliveryAttributeMappingResponse>> value)
        {
            Value = value;
        }
    }
}
