// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CloudHealth.Outputs
{

    /// <summary>
    /// HealthModel properties
    /// </summary>
    [OutputType]
    public sealed class HealthModelPropertiesResponse
    {
        /// <summary>
        /// The data plane endpoint for interacting with health data
        /// </summary>
        public readonly string DataplaneEndpoint;
        /// <summary>
        /// Configure to automatically discover entities from a given scope, such as a Service Group. The discovered entities will be linked to the root entity of the health model.
        /// </summary>
        public readonly Outputs.ModelDiscoverySettingsResponse? Discovery;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private HealthModelPropertiesResponse(
            string dataplaneEndpoint,

            Outputs.ModelDiscoverySettingsResponse? discovery,

            string provisioningState)
        {
            DataplaneEndpoint = dataplaneEndpoint;
            Discovery = discovery;
            ProvisioningState = provisioningState;
        }
    }
}
