// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs.Outputs
{

    /// <summary>
    /// Private Endpoint Connection properties.
    /// </summary>
    [OutputType]
    public sealed class PrivateEndpointConnectionPropertiesResponse
    {
        /// <summary>
        /// List of group ids. For Notification Hubs, it always contains a single "namespace" element.
        /// </summary>
        public readonly ImmutableArray<string> GroupIds;
        /// <summary>
        /// Represents a Private Endpoint that is connected to Notification Hubs namespace using Private Endpoint Connection.
        /// </summary>
        public readonly Outputs.RemotePrivateEndpointConnectionResponse? PrivateEndpoint;
        /// <summary>
        /// State of the Private Link Service connection.
        /// </summary>
        public readonly Outputs.RemotePrivateLinkServiceConnectionStateResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// State of Private Endpoint Connection.
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private PrivateEndpointConnectionPropertiesResponse(
            ImmutableArray<string> groupIds,

            Outputs.RemotePrivateEndpointConnectionResponse? privateEndpoint,

            Outputs.RemotePrivateLinkServiceConnectionStateResponse? privateLinkServiceConnectionState,

            string? provisioningState)
        {
            GroupIds = groupIds;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
        }
    }
}
