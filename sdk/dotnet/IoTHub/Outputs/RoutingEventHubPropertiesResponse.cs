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
    /// The properties related to an event hub endpoint.
    /// </summary>
    [OutputType]
    public sealed class RoutingEventHubPropertiesResponse
    {
        /// <summary>
        /// Method used to authenticate against the event hub endpoint
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// The connection string of the event hub endpoint. 
        /// </summary>
        public readonly string? ConnectionString;
        /// <summary>
        /// The url of the event hub endpoint. It must include the protocol sb://
        /// </summary>
        public readonly string? EndpointUri;
        /// <summary>
        /// Event hub name on the event hub namespace
        /// </summary>
        public readonly string? EntityPath;
        /// <summary>
        /// Id of the event hub endpoint
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Managed identity properties of routing event hub endpoint.
        /// </summary>
        public readonly Outputs.ManagedIdentityResponse? Identity;
        /// <summary>
        /// The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the resource group of the event hub endpoint.
        /// </summary>
        public readonly string? ResourceGroup;
        /// <summary>
        /// The subscription identifier of the event hub endpoint.
        /// </summary>
        public readonly string? SubscriptionId;

        [OutputConstructor]
        private RoutingEventHubPropertiesResponse(
            string? authenticationType,

            string? connectionString,

            string? endpointUri,

            string? entityPath,

            string? id,

            Outputs.ManagedIdentityResponse? identity,

            string name,

            string? resourceGroup,

            string? subscriptionId)
        {
            AuthenticationType = authenticationType;
            ConnectionString = connectionString;
            EndpointUri = endpointUri;
            EntityPath = entityPath;
            Id = id;
            Identity = identity;
            Name = name;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
        }
    }
}
