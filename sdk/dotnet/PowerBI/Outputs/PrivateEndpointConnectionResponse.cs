// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerBI.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionResponse
    {
        /// <summary>
        /// Specifies the id of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the private endpoint.
        /// </summary>
        public readonly Outputs.PrivateEndpointResponse? PrivateEndpoint;
        /// <summary>
        /// Specifies the connection state.
        /// </summary>
        public readonly Outputs.ConnectionStateResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Specifies the type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointConnectionResponse(
            string id,

            string name,

            Outputs.PrivateEndpointResponse? privateEndpoint,

            Outputs.ConnectionStateResponse? privateLinkServiceConnectionState,

            string? provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
