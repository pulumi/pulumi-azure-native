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
    /// Environment configuration options.
    /// </summary>
    [OutputType]
    public sealed class GroupEnvironmentConfigurationResponse
    {
        /// <summary>
        /// ARM resource ID of the environment specification for the inference pool.
        /// </summary>
        public readonly string? EnvironmentId;
        /// <summary>
        /// Environment variables configuration for the inference pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.StringStringKeyValuePairResponse> EnvironmentVariables;
        /// <summary>
        /// Liveness probe monitors the health of the container regularly.
        /// </summary>
        public readonly Outputs.ProbeSettingsResponse? LivenessProbe;
        /// <summary>
        /// Readiness probe validates if the container is ready to serve traffic. The properties and defaults are the same as liveness probe.
        /// </summary>
        public readonly Outputs.ProbeSettingsResponse? ReadinessProbe;
        /// <summary>
        /// This verifies whether the application within a container is started. Startup probes run before any other probe, and, unless it finishes successfully, disables other probes.
        /// </summary>
        public readonly Outputs.ProbeSettingsResponse? StartupProbe;

        [OutputConstructor]
        private GroupEnvironmentConfigurationResponse(
            string? environmentId,

            ImmutableArray<Outputs.StringStringKeyValuePairResponse> environmentVariables,

            Outputs.ProbeSettingsResponse? livenessProbe,

            Outputs.ProbeSettingsResponse? readinessProbe,

            Outputs.ProbeSettingsResponse? startupProbe)
        {
            EnvironmentId = environmentId;
            EnvironmentVariables = environmentVariables;
            LivenessProbe = livenessProbe;
            ReadinessProbe = readinessProbe;
            StartupProbe = startupProbe;
        }
    }
}
