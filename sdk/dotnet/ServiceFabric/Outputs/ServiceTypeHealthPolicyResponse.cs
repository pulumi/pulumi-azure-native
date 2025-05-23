// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.Outputs
{

    /// <summary>
    /// Represents the health policy used to evaluate the health of services belonging to a service type.
    /// </summary>
    [OutputType]
    public sealed class ServiceTypeHealthPolicyResponse
    {
        /// <summary>
        /// The maximum allowed percentage of unhealthy partitions per service.
        /// 
        /// The percentage represents the maximum tolerated percentage of partitions that can be unhealthy before the service is considered in error.
        /// If the percentage is respected but there is at least one unhealthy partition, the health is evaluated as Warning.
        /// The percentage is calculated by dividing the number of unhealthy partitions over the total number of partitions in the service.
        /// The computation rounds up to tolerate one failure on small numbers of partitions.
        /// </summary>
        public readonly int MaxPercentUnhealthyPartitionsPerService;
        /// <summary>
        /// The maximum allowed percentage of unhealthy replicas per partition.
        /// 
        /// The percentage represents the maximum tolerated percentage of replicas that can be unhealthy before the partition is considered in error.
        /// If the percentage is respected but there is at least one unhealthy replica, the health is evaluated as Warning.
        /// The percentage is calculated by dividing the number of unhealthy replicas over the total number of replicas in the partition.
        /// The computation rounds up to tolerate one failure on small numbers of replicas.
        /// </summary>
        public readonly int MaxPercentUnhealthyReplicasPerPartition;
        /// <summary>
        /// The maximum allowed percentage of unhealthy services.
        /// 
        /// The percentage represents the maximum tolerated percentage of services that can be unhealthy before the application is considered in error.
        /// If the percentage is respected but there is at least one unhealthy service, the health is evaluated as Warning.
        /// This is calculated by dividing the number of unhealthy services of the specific service type over the total number of services of the specific service type.
        /// The computation rounds up to tolerate one failure on small numbers of services.
        /// </summary>
        public readonly int MaxPercentUnhealthyServices;

        [OutputConstructor]
        private ServiceTypeHealthPolicyResponse(
            int maxPercentUnhealthyPartitionsPerService,

            int maxPercentUnhealthyReplicasPerPartition,

            int maxPercentUnhealthyServices)
        {
            MaxPercentUnhealthyPartitionsPerService = maxPercentUnhealthyPartitionsPerService;
            MaxPercentUnhealthyReplicasPerPartition = maxPercentUnhealthyReplicasPerPartition;
            MaxPercentUnhealthyServices = maxPercentUnhealthyServices;
        }
    }
}
