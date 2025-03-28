// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Inputs
{

    /// <summary>
    /// Settings for upgrading an agentpool
    /// </summary>
    public sealed class AgentPoolUpgradeSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time (in minutes) to wait on eviction of pods and graceful termination per node. This eviction wait time honors waiting on pod disruption budgets. If this time is exceeded, the upgrade fails. If not specified, the default is 30 minutes.
        /// </summary>
        [Input("drainTimeoutInMinutes")]
        public Input<int>? DrainTimeoutInMinutes { get; set; }

        /// <summary>
        /// This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 10%. For more information, including best practices, see: https://learn.microsoft.com/en-us/azure/aks/upgrade-cluster
        /// </summary>
        [Input("maxSurge")]
        public Input<string>? MaxSurge { get; set; }

        /// <summary>
        /// This can either be set to an integer (e.g. '1') or a percentage (e.g. '5%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 0. For more information, including best practices, see: https://learn.microsoft.com/en-us/azure/aks/upgrade-cluster
        /// </summary>
        [Input("maxUnavailable")]
        public Input<string>? MaxUnavailable { get; set; }

        /// <summary>
        /// The amount of time (in minutes) to wait after draining a node and before reimaging it and moving on to next node. If not specified, the default is 0 minutes.
        /// </summary>
        [Input("nodeSoakDurationInMinutes")]
        public Input<int>? NodeSoakDurationInMinutes { get; set; }

        /// <summary>
        /// Defines the behavior for undrainable nodes during upgrade. The most common cause of undrainable nodes is Pod Disruption Budgets (PDBs), but other issues, such as pod termination grace period is exceeding the remaining per-node drain timeout or pod is still being in a running state, can also cause undrainable nodes.
        /// </summary>
        [Input("undrainableNodeBehavior")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20241002Preview.UndrainableNodeBehavior>? UndrainableNodeBehavior { get; set; }

        public AgentPoolUpgradeSettingsArgs()
        {
        }
        public static new AgentPoolUpgradeSettingsArgs Empty => new AgentPoolUpgradeSettingsArgs();
    }
}
