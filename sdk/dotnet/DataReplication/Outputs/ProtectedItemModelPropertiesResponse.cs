// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Outputs
{

    /// <summary>
    /// Protected item model properties.
    /// </summary>
    [OutputType]
    public sealed class ProtectedItemModelPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the allowed scenarios on the protected item.
        /// </summary>
        public readonly ImmutableArray<string> AllowedJobs;
        /// <summary>
        /// Gets or sets the protected item correlation Id.
        /// </summary>
        public readonly string CorrelationId;
        public readonly Outputs.ProtectedItemModelPropertiesResponseCurrentJob CurrentJob;
        /// <summary>
        /// Protected item model custom properties.
        /// </summary>
        public readonly Union<Outputs.HyperVToAzStackHCIProtectedItemModelCustomPropertiesResponse, Outputs.VMwareToAzStackHCIProtectedItemModelCustomPropertiesResponse> CustomProperties;
        /// <summary>
        /// Gets or sets the DRA Id.
        /// </summary>
        public readonly string DraId;
        /// <summary>
        /// Gets or sets the fabric Id.
        /// </summary>
        public readonly string FabricId;
        /// <summary>
        /// Gets or sets the fabric object Id.
        /// </summary>
        public readonly string FabricObjectId;
        /// <summary>
        /// Gets or sets the fabric object name.
        /// </summary>
        public readonly string FabricObjectName;
        /// <summary>
        /// Gets or sets the list of health errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.HealthErrorModelResponse> HealthErrors;
        public readonly Outputs.ProtectedItemModelPropertiesResponseLastFailedEnableProtectionJob LastFailedEnableProtectionJob;
        public readonly Outputs.ProtectedItemModelPropertiesResponseLastFailedPlannedFailoverJob LastFailedPlannedFailoverJob;
        /// <summary>
        /// Gets or sets the Last successful planned failover time.
        /// </summary>
        public readonly string LastSuccessfulPlannedFailoverTime;
        /// <summary>
        /// Gets or sets the Last successful test failover time.
        /// </summary>
        public readonly string LastSuccessfulTestFailoverTime;
        /// <summary>
        /// Gets or sets the Last successful unplanned failover time.
        /// </summary>
        public readonly string LastSuccessfulUnplannedFailoverTime;
        public readonly Outputs.ProtectedItemModelPropertiesResponseLastTestFailoverJob LastTestFailoverJob;
        /// <summary>
        /// Gets or sets the policy name.
        /// </summary>
        public readonly string PolicyName;
        /// <summary>
        /// Gets or sets the protection state.
        /// </summary>
        public readonly string ProtectionState;
        /// <summary>
        /// Gets or sets the protection state description.
        /// </summary>
        public readonly string ProtectionStateDescription;
        /// <summary>
        /// Gets or sets the provisioning state of the Dra.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Gets or sets the replication extension name.
        /// </summary>
        public readonly string ReplicationExtensionName;
        /// <summary>
        /// Gets or sets protected item replication health.
        /// </summary>
        public readonly string ReplicationHealth;
        /// <summary>
        /// Gets or sets a value indicating whether resynchronization is required or not.
        /// </summary>
        public readonly bool ResyncRequired;
        /// <summary>
        /// Gets or sets the resynchronization state.
        /// </summary>
        public readonly string ResynchronizationState;
        /// <summary>
        /// Gets or sets the source fabric provider Id.
        /// </summary>
        public readonly string SourceFabricProviderId;
        /// <summary>
        /// Gets or sets the target DRA Id.
        /// </summary>
        public readonly string TargetDraId;
        /// <summary>
        /// Gets or sets the target fabric Id.
        /// </summary>
        public readonly string TargetFabricId;
        /// <summary>
        /// Gets or sets the target fabric provider Id.
        /// </summary>
        public readonly string TargetFabricProviderId;
        /// <summary>
        /// Gets or sets the test failover state.
        /// </summary>
        public readonly string TestFailoverState;
        /// <summary>
        /// Gets or sets the Test failover state description.
        /// </summary>
        public readonly string TestFailoverStateDescription;

        [OutputConstructor]
        private ProtectedItemModelPropertiesResponse(
            ImmutableArray<string> allowedJobs,

            string correlationId,

            Outputs.ProtectedItemModelPropertiesResponseCurrentJob currentJob,

            Union<Outputs.HyperVToAzStackHCIProtectedItemModelCustomPropertiesResponse, Outputs.VMwareToAzStackHCIProtectedItemModelCustomPropertiesResponse> customProperties,

            string draId,

            string fabricId,

            string fabricObjectId,

            string fabricObjectName,

            ImmutableArray<Outputs.HealthErrorModelResponse> healthErrors,

            Outputs.ProtectedItemModelPropertiesResponseLastFailedEnableProtectionJob lastFailedEnableProtectionJob,

            Outputs.ProtectedItemModelPropertiesResponseLastFailedPlannedFailoverJob lastFailedPlannedFailoverJob,

            string lastSuccessfulPlannedFailoverTime,

            string lastSuccessfulTestFailoverTime,

            string lastSuccessfulUnplannedFailoverTime,

            Outputs.ProtectedItemModelPropertiesResponseLastTestFailoverJob lastTestFailoverJob,

            string policyName,

            string protectionState,

            string protectionStateDescription,

            string provisioningState,

            string replicationExtensionName,

            string replicationHealth,

            bool resyncRequired,

            string resynchronizationState,

            string sourceFabricProviderId,

            string targetDraId,

            string targetFabricId,

            string targetFabricProviderId,

            string testFailoverState,

            string testFailoverStateDescription)
        {
            AllowedJobs = allowedJobs;
            CorrelationId = correlationId;
            CurrentJob = currentJob;
            CustomProperties = customProperties;
            DraId = draId;
            FabricId = fabricId;
            FabricObjectId = fabricObjectId;
            FabricObjectName = fabricObjectName;
            HealthErrors = healthErrors;
            LastFailedEnableProtectionJob = lastFailedEnableProtectionJob;
            LastFailedPlannedFailoverJob = lastFailedPlannedFailoverJob;
            LastSuccessfulPlannedFailoverTime = lastSuccessfulPlannedFailoverTime;
            LastSuccessfulTestFailoverTime = lastSuccessfulTestFailoverTime;
            LastSuccessfulUnplannedFailoverTime = lastSuccessfulUnplannedFailoverTime;
            LastTestFailoverJob = lastTestFailoverJob;
            PolicyName = policyName;
            ProtectionState = protectionState;
            ProtectionStateDescription = protectionStateDescription;
            ProvisioningState = provisioningState;
            ReplicationExtensionName = replicationExtensionName;
            ReplicationHealth = replicationHealth;
            ResyncRequired = resyncRequired;
            ResynchronizationState = resynchronizationState;
            SourceFabricProviderId = sourceFabricProviderId;
            TargetDraId = targetDraId;
            TargetFabricId = targetFabricId;
            TargetFabricProviderId = targetFabricProviderId;
            TestFailoverState = testFailoverState;
            TestFailoverStateDescription = testFailoverStateDescription;
        }
    }
}
