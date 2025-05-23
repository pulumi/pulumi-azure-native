// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Gets the application HTTP endpoints.
    /// </summary>
    [OutputType]
    public sealed class ApplicationGetHttpsEndpointResponse
    {
        /// <summary>
        /// The list of access modes for the application.
        /// </summary>
        public readonly ImmutableArray<string> AccessModes;
        /// <summary>
        /// The destination port to connect to.
        /// </summary>
        public readonly int? DestinationPort;
        /// <summary>
        /// The value indicates whether to disable GatewayAuth.
        /// </summary>
        public readonly bool? DisableGatewayAuth;
        /// <summary>
        /// The location of the endpoint.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The private ip address of the endpoint.
        /// </summary>
        public readonly string? PrivateIPAddress;
        /// <summary>
        /// The public port to connect to.
        /// </summary>
        public readonly int PublicPort;

        [OutputConstructor]
        private ApplicationGetHttpsEndpointResponse(
            ImmutableArray<string> accessModes,

            int? destinationPort,

            bool? disableGatewayAuth,

            string location,

            string? privateIPAddress,

            int publicPort)
        {
            AccessModes = accessModes;
            DestinationPort = destinationPort;
            DisableGatewayAuth = disableGatewayAuth;
            Location = location;
            PrivateIPAddress = privateIPAddress;
            PublicPort = publicPort;
        }
    }
}
