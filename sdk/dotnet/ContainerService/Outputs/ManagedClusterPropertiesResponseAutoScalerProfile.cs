// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Parameters to be applied to the cluster-autoscaler when enabled
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterPropertiesResponseAutoScalerProfile
    {
        /// <summary>
        /// Valid values are 'true' and 'false'
        /// </summary>
        public readonly string? BalanceSimilarNodeGroups;
        /// <summary>
        /// If set to true, all daemonset pods on empty nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.
        /// </summary>
        public readonly bool? DaemonsetEvictionForEmptyNodes;
        /// <summary>
        /// If set to true, all daemonset pods on occupied nodes will be evicted before deletion of the node. If the daemonset pod cannot be evicted another node will be chosen for scaling. If set to false, the node will be deleted without ensuring that daemonset pods are deleted or evicted.
        /// </summary>
        public readonly bool? DaemonsetEvictionForOccupiedNodes;
        /// <summary>
        /// If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.
        /// </summary>
        public readonly string? Expander;
        /// <summary>
        /// If set to true, the resources used by daemonset will be taken into account when making scaling down decisions.
        /// </summary>
        public readonly bool? IgnoreDaemonsetsUtilization;
        /// <summary>
        /// The default is 10.
        /// </summary>
        public readonly string? MaxEmptyBulkDelete;
        /// <summary>
        /// The default is 600.
        /// </summary>
        public readonly string? MaxGracefulTerminationSec;
        /// <summary>
        /// The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? MaxNodeProvisionTime;
        /// <summary>
        /// The default is 45. The maximum is 100 and the minimum is 0.
        /// </summary>
        public readonly string? MaxTotalUnreadyPercentage;
        /// <summary>
        /// For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they're a certain age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours, etc).
        /// </summary>
        public readonly string? NewPodScaleUpDelay;
        /// <summary>
        /// This must be an integer. The default is 3.
        /// </summary>
        public readonly string? OkTotalUnreadyCount;
        /// <summary>
        /// The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? ScaleDownDelayAfterAdd;
        /// <summary>
        /// The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? ScaleDownDelayAfterDelete;
        /// <summary>
        /// The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? ScaleDownDelayAfterFailure;
        /// <summary>
        /// The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? ScaleDownUnneededTime;
        /// <summary>
        /// The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
        /// </summary>
        public readonly string? ScaleDownUnreadyTime;
        /// <summary>
        /// The default is '0.5'.
        /// </summary>
        public readonly string? ScaleDownUtilizationThreshold;
        /// <summary>
        /// The default is '10'. Values must be an integer number of seconds.
        /// </summary>
        public readonly string? ScanInterval;
        /// <summary>
        /// The default is true.
        /// </summary>
        public readonly string? SkipNodesWithLocalStorage;
        /// <summary>
        /// The default is true.
        /// </summary>
        public readonly string? SkipNodesWithSystemPods;

        [OutputConstructor]
        private ManagedClusterPropertiesResponseAutoScalerProfile(
            string? balanceSimilarNodeGroups,

            bool? daemonsetEvictionForEmptyNodes,

            bool? daemonsetEvictionForOccupiedNodes,

            string? expander,

            bool? ignoreDaemonsetsUtilization,

            string? maxEmptyBulkDelete,

            string? maxGracefulTerminationSec,

            string? maxNodeProvisionTime,

            string? maxTotalUnreadyPercentage,

            string? newPodScaleUpDelay,

            string? okTotalUnreadyCount,

            string? scaleDownDelayAfterAdd,

            string? scaleDownDelayAfterDelete,

            string? scaleDownDelayAfterFailure,

            string? scaleDownUnneededTime,

            string? scaleDownUnreadyTime,

            string? scaleDownUtilizationThreshold,

            string? scanInterval,

            string? skipNodesWithLocalStorage,

            string? skipNodesWithSystemPods)
        {
            BalanceSimilarNodeGroups = balanceSimilarNodeGroups;
            DaemonsetEvictionForEmptyNodes = daemonsetEvictionForEmptyNodes;
            DaemonsetEvictionForOccupiedNodes = daemonsetEvictionForOccupiedNodes;
            Expander = expander;
            IgnoreDaemonsetsUtilization = ignoreDaemonsetsUtilization;
            MaxEmptyBulkDelete = maxEmptyBulkDelete;
            MaxGracefulTerminationSec = maxGracefulTerminationSec;
            MaxNodeProvisionTime = maxNodeProvisionTime;
            MaxTotalUnreadyPercentage = maxTotalUnreadyPercentage;
            NewPodScaleUpDelay = newPodScaleUpDelay;
            OkTotalUnreadyCount = okTotalUnreadyCount;
            ScaleDownDelayAfterAdd = scaleDownDelayAfterAdd;
            ScaleDownDelayAfterDelete = scaleDownDelayAfterDelete;
            ScaleDownDelayAfterFailure = scaleDownDelayAfterFailure;
            ScaleDownUnneededTime = scaleDownUnneededTime;
            ScaleDownUnreadyTime = scaleDownUnreadyTime;
            ScaleDownUtilizationThreshold = scaleDownUtilizationThreshold;
            ScanInterval = scanInterval;
            SkipNodesWithLocalStorage = skipNodesWithLocalStorage;
            SkipNodesWithSystemPods = skipNodesWithSystemPods;
        }
    }
}
