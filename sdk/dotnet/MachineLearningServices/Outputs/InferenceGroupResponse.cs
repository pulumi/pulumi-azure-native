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
    /// Inference group configuration
    /// </summary>
    [OutputType]
    public sealed class InferenceGroupResponse
    {
        /// <summary>
        /// Description of the resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets or sets environment configuration for the inference group. Used if PoolType=ScaleUnit.
        /// </summary>
        public readonly Outputs.GroupEnvironmentConfigurationResponse? EnvironmentConfiguration;
        /// <summary>
        /// Gets or sets model configuration for the inference group. Used if PoolType=ScaleUnit.
        /// </summary>
        public readonly Outputs.GroupModelConfigurationResponse? ModelConfiguration;
        /// <summary>
        /// Gets or sets compute instance type.
        /// </summary>
        public readonly string? NodeSkuType;
        /// <summary>
        /// Property dictionary. Properties can be added, but not removed or altered.
        /// </summary>
        public readonly ImmutableArray<Outputs.StringStringKeyValuePairResponse> Properties;
        /// <summary>
        /// Provisioning state for the inference group.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Gets or sets Scale Unit size.
        /// </summary>
        public readonly int? ScaleUnitSize;

        [OutputConstructor]
        private InferenceGroupResponse(
            string? description,

            Outputs.GroupEnvironmentConfigurationResponse? environmentConfiguration,

            Outputs.GroupModelConfigurationResponse? modelConfiguration,

            string? nodeSkuType,

            ImmutableArray<Outputs.StringStringKeyValuePairResponse> properties,

            string provisioningState,

            int? scaleUnitSize)
        {
            Description = description;
            EnvironmentConfiguration = environmentConfiguration;
            ModelConfiguration = modelConfiguration;
            NodeSkuType = nodeSkuType;
            Properties = properties;
            ProvisioningState = provisioningState;
            ScaleUnitSize = scaleUnitSize;
        }
    }
}
