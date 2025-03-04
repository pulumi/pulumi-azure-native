// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230401Preview.Outputs
{

    /// <summary>
    /// Details of the Registry
    /// </summary>
    [OutputType]
    public sealed class RegistryResponse
    {
        /// <summary>
        /// Discovery URL for the Registry
        /// </summary>
        public readonly string? DiscoveryUrl;
        /// <summary>
        /// IntellectualPropertyPublisher for the registry
        /// </summary>
        public readonly string? IntellectualPropertyPublisher;
        /// <summary>
        /// ResourceId of the managed RG if the registry has system created resources
        /// </summary>
        public readonly Outputs.ArmResourceIdResponse? ManagedResourceGroup;
        /// <summary>
        /// MLFlow Registry URI for the Registry
        /// </summary>
        public readonly string? MlFlowRegistryUri;
        /// <summary>
        /// Private endpoint connections info used for pending connections in private link portal
        /// </summary>
        public readonly ImmutableArray<Outputs.RegistryPrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Is the Registry accessible from the internet?
        /// Possible values: "Enabled" or "Disabled"
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// Details of each region the registry is in
        /// </summary>
        public readonly ImmutableArray<Outputs.RegistryRegionArmDetailsResponse> RegionDetails;

        [OutputConstructor]
        private RegistryResponse(
            string? discoveryUrl,

            string? intellectualPropertyPublisher,

            Outputs.ArmResourceIdResponse? managedResourceGroup,

            string? mlFlowRegistryUri,

            ImmutableArray<Outputs.RegistryPrivateEndpointConnectionResponse> privateEndpointConnections,

            string? publicNetworkAccess,

            ImmutableArray<Outputs.RegistryRegionArmDetailsResponse> regionDetails)
        {
            DiscoveryUrl = discoveryUrl;
            IntellectualPropertyPublisher = intellectualPropertyPublisher;
            ManagedResourceGroup = managedResourceGroup;
            MlFlowRegistryUri = mlFlowRegistryUri;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            RegionDetails = regionDetails;
        }
    }
}
