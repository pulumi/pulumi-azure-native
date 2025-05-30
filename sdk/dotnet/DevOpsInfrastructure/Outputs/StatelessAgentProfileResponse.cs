// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.Outputs
{

    /// <summary>
    /// Stateless profile meaning that the machines will be cleaned up after running a job.
    /// </summary>
    [OutputType]
    public sealed class StatelessAgentProfileResponse
    {
        /// <summary>
        /// Discriminator property for AgentProfile.
        /// Expected value is 'Stateless'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Defines pool buffer/stand-by agents.
        /// </summary>
        public readonly object? ResourcePredictions;
        /// <summary>
        /// Defines how the pool buffer/stand-by agents is provided.
        /// </summary>
        public readonly Union<Outputs.AutomaticResourcePredictionsProfileResponse, Outputs.ManualResourcePredictionsProfileResponse>? ResourcePredictionsProfile;

        [OutputConstructor]
        private StatelessAgentProfileResponse(
            string kind,

            object? resourcePredictions,

            Union<Outputs.AutomaticResourcePredictionsProfileResponse, Outputs.ManualResourcePredictionsProfileResponse>? resourcePredictionsProfile)
        {
            Kind = kind;
            ResourcePredictions = resourcePredictions;
            ResourcePredictionsProfile = resourcePredictionsProfile;
        }
    }
}
