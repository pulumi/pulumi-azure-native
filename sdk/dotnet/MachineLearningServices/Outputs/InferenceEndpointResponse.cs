// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// InferenceEndpoint configuration
    /// </summary>
    [OutputType]
    public sealed class InferenceEndpointResponse
    {
        /// <summary>
        /// [Required] Authentication mode for the endpoint.
        /// </summary>
        public readonly string AuthMode;
        /// <summary>
        /// Description of the resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Endpoint URI for the inference endpoint.
        /// </summary>
        public readonly string EndpointUri;
        /// <summary>
        /// [Required] Group within the same pool with which this endpoint needs to be associated with.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Property dictionary. Properties can be added, but not removed or altered.
        /// </summary>
        public readonly ImmutableArray<Outputs.StringStringKeyValuePairResponse> Properties;
        /// <summary>
        /// Provisioning state for the endpoint.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// RequestConfiguration for endpoint.
        /// </summary>
        public readonly Outputs.RequestConfigurationResponse? RequestConfiguration;

        [OutputConstructor]
        private InferenceEndpointResponse(
            string authMode,

            string? description,

            string endpointUri,

            string groupName,

            ImmutableArray<Outputs.StringStringKeyValuePairResponse> properties,

            string provisioningState,

            Outputs.RequestConfigurationResponse? requestConfiguration)
        {
            AuthMode = authMode;
            Description = description;
            EndpointUri = endpointUri;
            GroupName = groupName;
            Properties = properties;
            ProvisioningState = provisioningState;
            RequestConfiguration = requestConfiguration;
        }
    }
}
