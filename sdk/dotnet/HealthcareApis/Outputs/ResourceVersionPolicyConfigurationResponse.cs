// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Outputs
{

    /// <summary>
    /// The settings for history tracking for FHIR resources.
    /// </summary>
    [OutputType]
    public sealed class ResourceVersionPolicyConfigurationResponse
    {
        /// <summary>
        /// The default value for tracking history across all resources.
        /// </summary>
        public readonly string? Default;
        /// <summary>
        /// A list of FHIR Resources and their version policy overrides.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ResourceTypeOverrides;

        [OutputConstructor]
        private ResourceVersionPolicyConfigurationResponse(
            string? @default,

            ImmutableDictionary<string, string>? resourceTypeOverrides)
        {
            Default = @default;
            ResourceTypeOverrides = resourceTypeOverrides;
        }
    }
}
