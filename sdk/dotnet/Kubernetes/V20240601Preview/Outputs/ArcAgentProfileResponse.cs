// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kubernetes.V20240601Preview.Outputs
{

    /// <summary>
    /// Defines the Arc Agent properties for the clusters.
    /// </summary>
    [OutputType]
    public sealed class ArcAgentProfileResponse
    {
        /// <summary>
        /// Indicates whether the Arc agents on the be upgraded automatically to the latest version. Defaults to Enabled.
        /// </summary>
        public readonly string? AgentAutoUpgrade;
        /// <summary>
        /// List of arc agentry and system components errors on the cluster resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentErrorResponse> AgentErrors;
        /// <summary>
        /// Represents the current state of the Arc agentry and its dependent components.
        /// </summary>
        public readonly string AgentState;
        /// <summary>
        /// Version of the Arc agents to be installed on the cluster resource
        /// </summary>
        public readonly string? DesiredAgentVersion;
        /// <summary>
        /// List of system extensions that are installed on the cluster resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SystemComponentResponse> SystemComponents;

        [OutputConstructor]
        private ArcAgentProfileResponse(
            string? agentAutoUpgrade,

            ImmutableArray<Outputs.AgentErrorResponse> agentErrors,

            string agentState,

            string? desiredAgentVersion,

            ImmutableArray<Outputs.SystemComponentResponse> systemComponents)
        {
            AgentAutoUpgrade = agentAutoUpgrade;
            AgentErrors = agentErrors;
            AgentState = agentState;
            DesiredAgentVersion = desiredAgentVersion;
            SystemComponents = systemComponents;
        }
    }
}
