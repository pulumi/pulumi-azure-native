// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTHub.Outputs
{

    /// <summary>
    /// The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed across all endpoint types for free hubs.
    /// </summary>
    [OutputType]
    public sealed class RoutingEndpointsResponse
    {
        /// <summary>
        /// The list of Cosmos DB container endpoints that IoT hub routes messages to, based on the routing rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingCosmosDBSqlApiPropertiesResponse> CosmosDBSqlContainers;
        /// <summary>
        /// The list of Event Hubs endpoints that IoT hub routes messages to, based on the routing rules. This list does not include the built-in Event Hubs endpoint.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingEventHubPropertiesResponse> EventHubs;
        /// <summary>
        /// The list of Service Bus queue endpoints that IoT hub routes the messages to, based on the routing rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingServiceBusQueueEndpointPropertiesResponse> ServiceBusQueues;
        /// <summary>
        /// The list of Service Bus topic endpoints that the IoT hub routes the messages to, based on the routing rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingServiceBusTopicEndpointPropertiesResponse> ServiceBusTopics;
        /// <summary>
        /// The list of storage container endpoints that IoT hub routes messages to, based on the routing rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutingStorageContainerPropertiesResponse> StorageContainers;

        [OutputConstructor]
        private RoutingEndpointsResponse(
            ImmutableArray<Outputs.RoutingCosmosDBSqlApiPropertiesResponse> cosmosDBSqlContainers,

            ImmutableArray<Outputs.RoutingEventHubPropertiesResponse> eventHubs,

            ImmutableArray<Outputs.RoutingServiceBusQueueEndpointPropertiesResponse> serviceBusQueues,

            ImmutableArray<Outputs.RoutingServiceBusTopicEndpointPropertiesResponse> serviceBusTopics,

            ImmutableArray<Outputs.RoutingStorageContainerPropertiesResponse> storageContainers)
        {
            CosmosDBSqlContainers = cosmosDBSqlContainers;
            EventHubs = eventHubs;
            ServiceBusQueues = serviceBusQueues;
            ServiceBusTopics = serviceBusTopics;
            StorageContainers = storageContainers;
        }
    }
}
