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
    /// The properties of a stateful service resource.
    /// </summary>
    [OutputType]
    public sealed class StatefulServicePropertiesResponse
    {
        /// <summary>
        /// A list that describes the correlation of the service with other services.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceCorrelationResponse> CorrelationScheme;
        /// <summary>
        /// Specifies the move cost for the service.
        /// </summary>
        public readonly string? DefaultMoveCost;
        /// <summary>
        /// A flag indicating whether this is a persistent service which stores states on the local disk. If it is then the value of this property is true, if not it is false.
        /// </summary>
        public readonly bool? HasPersistedState;
        /// <summary>
        /// The minimum replica set size as a number.
        /// </summary>
        public readonly int? MinReplicaSetSize;
        /// <summary>
        /// Describes how the service is partitioned.
        /// </summary>
        public readonly object PartitionDescription;
        /// <summary>
        /// The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
        /// </summary>
        public readonly string? PlacementConstraints;
        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The maximum duration for which a partition is allowed to be in a state of quorum loss, represented in ISO 8601 format "hh:mm:ss".
        /// </summary>
        public readonly string? QuorumLossWaitDuration;
        /// <summary>
        /// The duration between when a replica goes down and when a new replica is created, represented in ISO 8601 format "hh:mm:ss".
        /// </summary>
        public readonly string? ReplicaRestartWaitDuration;
        /// <summary>
        /// Scaling policies for this service.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingPolicyResponse> ScalingPolicies;
        /// <summary>
        /// Dns name used for the service. If this is specified, then the DNS name can be used to return the IP addresses of service endpoints for application layer protocols (e.g., HTTP).
        /// When updating serviceDnsName, old name may be temporarily resolvable. However, rely on new name.
        /// When removing serviceDnsName, removed name may temporarily be resolvable. Do not rely on the name being unresolvable.
        /// </summary>
        public readonly string? ServiceDnsName;
        /// <summary>
        /// The kind of service (Stateless or Stateful).
        /// Expected value is 'Stateful'.
        /// </summary>
        public readonly string ServiceKind;
        /// <summary>
        /// The service load metrics is given as an array of ServiceLoadMetric objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceLoadMetricResponse> ServiceLoadMetrics;
        /// <summary>
        /// The activation Mode of the service package
        /// </summary>
        public readonly string? ServicePackageActivationMode;
        /// <summary>
        /// A list that describes the correlation of the service with other services.
        /// </summary>
        public readonly ImmutableArray<object> ServicePlacementPolicies;
        /// <summary>
        /// The duration for which replicas can stay InBuild before reporting that build is stuck, represented in ISO 8601 format "hh:mm:ss".
        /// </summary>
        public readonly string? ServicePlacementTimeLimit;
        /// <summary>
        /// The name of the service type
        /// </summary>
        public readonly string ServiceTypeName;
        /// <summary>
        /// The definition on how long StandBy replicas should be maintained before being removed, represented in ISO 8601 format "hh:mm:ss".
        /// </summary>
        public readonly string? StandByReplicaKeepDuration;
        /// <summary>
        /// The target replica set size as a number.
        /// </summary>
        public readonly int? TargetReplicaSetSize;

        [OutputConstructor]
        private StatefulServicePropertiesResponse(
            ImmutableArray<Outputs.ServiceCorrelationResponse> correlationScheme,

            string? defaultMoveCost,

            bool? hasPersistedState,

            int? minReplicaSetSize,

            object partitionDescription,

            string? placementConstraints,

            string provisioningState,

            string? quorumLossWaitDuration,

            string? replicaRestartWaitDuration,

            ImmutableArray<Outputs.ScalingPolicyResponse> scalingPolicies,

            string? serviceDnsName,

            string serviceKind,

            ImmutableArray<Outputs.ServiceLoadMetricResponse> serviceLoadMetrics,

            string? servicePackageActivationMode,

            ImmutableArray<object> servicePlacementPolicies,

            string? servicePlacementTimeLimit,

            string serviceTypeName,

            string? standByReplicaKeepDuration,

            int? targetReplicaSetSize)
        {
            CorrelationScheme = correlationScheme;
            DefaultMoveCost = defaultMoveCost;
            HasPersistedState = hasPersistedState;
            MinReplicaSetSize = minReplicaSetSize;
            PartitionDescription = partitionDescription;
            PlacementConstraints = placementConstraints;
            ProvisioningState = provisioningState;
            QuorumLossWaitDuration = quorumLossWaitDuration;
            ReplicaRestartWaitDuration = replicaRestartWaitDuration;
            ScalingPolicies = scalingPolicies;
            ServiceDnsName = serviceDnsName;
            ServiceKind = serviceKind;
            ServiceLoadMetrics = serviceLoadMetrics;
            ServicePackageActivationMode = servicePackageActivationMode;
            ServicePlacementPolicies = servicePlacementPolicies;
            ServicePlacementTimeLimit = servicePlacementTimeLimit;
            ServiceTypeName = serviceTypeName;
            StandByReplicaKeepDuration = standByReplicaKeepDuration;
            TargetReplicaSetSize = targetReplicaSetSize;
        }
    }
}
