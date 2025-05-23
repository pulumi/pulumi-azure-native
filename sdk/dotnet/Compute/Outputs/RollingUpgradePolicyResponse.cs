// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// The configuration parameters used while performing a rolling upgrade.
    /// </summary>
    [OutputType]
    public sealed class RollingUpgradePolicyResponse
    {
        /// <summary>
        /// Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration the Update Domain and maxBatchInstancePercent to determine the batch size.
        /// </summary>
        public readonly bool? EnableCrossZoneUpgrade;
        /// <summary>
        /// The maximum percent of total virtual machine instances that will be upgraded simultaneously by the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.
        /// </summary>
        public readonly int? MaxBatchInstancePercent;
        /// <summary>
        /// Create new virtual machines to upgrade the scale set, rather than updating the existing virtual machines. Existing virtual machines will be deleted once the new virtual machines are created for each batch.
        /// </summary>
        public readonly bool? MaxSurge;
        /// <summary>
        /// The maximum percentage of the total virtual machine instances in the scale set that can be simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch. The default value for this parameter is 20%.
        /// </summary>
        public readonly int? MaxUnhealthyInstancePercent;
        /// <summary>
        /// The maximum percentage of upgraded virtual machine instances that can be found to be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the rolling update aborts. The default value for this parameter is 20%.
        /// </summary>
        public readonly int? MaxUnhealthyUpgradedInstancePercent;
        /// <summary>
        /// The wait time between completing the update for all virtual machines in one batch and starting the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).
        /// </summary>
        public readonly string? PauseTimeBetweenBatches;
        /// <summary>
        /// Upgrade all unhealthy instances in a scale set before any healthy instances.
        /// </summary>
        public readonly bool? PrioritizeUnhealthyInstances;
        /// <summary>
        /// Rollback failed instances to previous model if the Rolling Upgrade policy is violated.
        /// </summary>
        public readonly bool? RollbackFailedInstancesOnPolicyBreach;

        [OutputConstructor]
        private RollingUpgradePolicyResponse(
            bool? enableCrossZoneUpgrade,

            int? maxBatchInstancePercent,

            bool? maxSurge,

            int? maxUnhealthyInstancePercent,

            int? maxUnhealthyUpgradedInstancePercent,

            string? pauseTimeBetweenBatches,

            bool? prioritizeUnhealthyInstances,

            bool? rollbackFailedInstancesOnPolicyBreach)
        {
            EnableCrossZoneUpgrade = enableCrossZoneUpgrade;
            MaxBatchInstancePercent = maxBatchInstancePercent;
            MaxSurge = maxSurge;
            MaxUnhealthyInstancePercent = maxUnhealthyInstancePercent;
            MaxUnhealthyUpgradedInstancePercent = maxUnhealthyUpgradedInstancePercent;
            PauseTimeBetweenBatches = pauseTimeBetweenBatches;
            PrioritizeUnhealthyInstances = prioritizeUnhealthyInstances;
            RollbackFailedInstancesOnPolicyBreach = rollbackFailedInstancesOnPolicyBreach;
        }
    }
}
