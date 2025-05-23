// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Outputs
{

    /// <summary>
    /// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
    /// </summary>
    [OutputType]
    public sealed class AgentPoolProvisioningStatusResponseStatus
    {
        /// <summary>
        /// ErrorMessage - Error messages during creation of cluster
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// Contains Provisioning errors
        /// </summary>
        public readonly Outputs.AgentPoolProvisioningStatusResponseProvisioningStatus? ProvisioningStatus;
        /// <summary>
        /// Total number of ready machines targeted by this deployment.
        /// </summary>
        public readonly int? ReadyReplicas;
        /// <summary>
        /// Total number of non-terminated machines targeted by this deployment
        /// </summary>
        public readonly int? Replicas;

        [OutputConstructor]
        private AgentPoolProvisioningStatusResponseStatus(
            string? errorMessage,

            Outputs.AgentPoolProvisioningStatusResponseProvisioningStatus? provisioningStatus,

            int? readyReplicas,

            int? replicas)
        {
            ErrorMessage = errorMessage;
            ProvisioningStatus = provisioningStatus;
            ReadyReplicas = readyReplicas;
            Replicas = replicas;
        }
    }
}
