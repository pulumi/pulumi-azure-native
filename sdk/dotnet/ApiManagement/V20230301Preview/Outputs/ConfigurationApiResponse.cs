// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20230301Preview.Outputs
{

    /// <summary>
    /// Information regarding the Configuration API of the API Management service.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationApiResponse
    {
        /// <summary>
        /// Indication whether or not the legacy Configuration API (v1) should be exposed on the API Management service. Value is optional but must be 'Enabled' or 'Disabled'. If 'Disabled', legacy Configuration API (v1) will not be available for self-hosted gateways. Default value is 'Enabled'
        /// </summary>
        public readonly string? LegacyApi;

        [OutputConstructor]
        private ConfigurationApiResponse(string? legacyApi)
        {
            LegacyApi = legacyApi;
        }
    }
}
