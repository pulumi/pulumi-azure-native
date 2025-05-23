// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.Inputs
{

    /// <summary>
    /// Stateless profile meaning that the machines will be cleaned up after running a job.
    /// </summary>
    public sealed class StatelessAgentProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Discriminator property for AgentProfile.
        /// Expected value is 'Stateless'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Defines pool buffer/stand-by agents.
        /// </summary>
        [Input("resourcePredictions")]
        public Input<object>? ResourcePredictions { get; set; }

        /// <summary>
        /// Defines how the pool buffer/stand-by agents is provided.
        /// </summary>
        [Input("resourcePredictionsProfile")]
        public InputUnion<Inputs.AutomaticResourcePredictionsProfileArgs, Inputs.ManualResourcePredictionsProfileArgs>? ResourcePredictionsProfile { get; set; }

        public StatelessAgentProfileArgs()
        {
        }
        public static new StatelessAgentProfileArgs Empty => new StatelessAgentProfileArgs();
    }
}
