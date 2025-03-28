// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.V20230101.Outputs
{

    /// <summary>
    /// Properties of a private endpoint connection.
    /// </summary>
    [OutputType]
    public sealed class PrivateEndpointConnectionPropertiesResponse
    {
        /// <summary>
        /// Defines resource ID of a private endpoint connection.
        /// </summary>
        public readonly Outputs.ResourceIdResponse PrivateEndpoint;
        /// <summary>
        /// Gets the properties of the object.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private PrivateEndpointConnectionPropertiesResponse(
            Outputs.ResourceIdResponse privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponse? privateLinkServiceConnectionState,

            string provisioningState)
        {
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
        }
    }
}
