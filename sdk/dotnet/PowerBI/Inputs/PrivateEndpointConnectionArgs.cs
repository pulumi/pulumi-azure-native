// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerBI.Inputs
{

    public sealed class PrivateEndpointConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the private endpoint.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// Specifies the connection state.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.ConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNative.PowerBI.ResourceProvisioningState>? ProvisioningState { get; set; }

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
