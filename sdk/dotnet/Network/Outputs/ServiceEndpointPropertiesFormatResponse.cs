// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// The service endpoint properties.
    /// </summary>
    [OutputType]
    public sealed class ServiceEndpointPropertiesFormatResponse
    {
        /// <summary>
        /// A list of locations.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// SubResource as network identifier.
        /// </summary>
        public readonly Outputs.SubResourceResponse? NetworkIdentifier;
        /// <summary>
        /// The provisioning state of the service endpoint resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the endpoint service.
        /// </summary>
        public readonly string? Service;

        [OutputConstructor]
        private ServiceEndpointPropertiesFormatResponse(
            ImmutableArray<string> locations,

            Outputs.SubResourceResponse? networkIdentifier,

            string provisioningState,

            string? service)
        {
            Locations = locations;
            NetworkIdentifier = networkIdentifier;
            ProvisioningState = provisioningState;
            Service = service;
        }
    }
}
