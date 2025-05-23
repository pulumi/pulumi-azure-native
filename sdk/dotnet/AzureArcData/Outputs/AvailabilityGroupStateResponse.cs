// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// The specifications of the availability group state
    /// </summary>
    [OutputType]
    public sealed class AvailabilityGroupStateResponse
    {
        /// <summary>
        /// Current Always On availability groups role of the availability group replica.
        /// </summary>
        public readonly string AvailabilityGroupReplicaRole;
        /// <summary>
        /// Whether a secondary replica is currently connected to the primary replica.
        /// </summary>
        public readonly string ConnectedStateDescription;
        /// <summary>
        /// Text description of the last connection error of the availability group replica.
        /// </summary>
        public readonly string LastConnectErrorDescription;
        /// <summary>
        /// Date and time timestamp indicating when the last connect error occurred.
        /// </summary>
        public readonly string LastConnectErrorTimestamp;
        /// <summary>
        /// Current operational state of the availability group replica
        /// </summary>
        public readonly string OperationalStateDescription;
        /// <summary>
        /// Recovery health of the availability group replica.
        /// </summary>
        public readonly string RecoveryHealthDescription;
        /// <summary>
        /// Reflects a rollup of the database synchronization state (synchronization_state) of all joined availability databases (also known as replicas) and the availability mode of the replica (synchronous-commit or asynchronous-commit mode). The rollup will reflect the least healthy accumulated state the databases on the replica.
        /// </summary>
        public readonly string SynchronizationHealthDescription;

        [OutputConstructor]
        private AvailabilityGroupStateResponse(
            string availabilityGroupReplicaRole,

            string connectedStateDescription,

            string lastConnectErrorDescription,

            string lastConnectErrorTimestamp,

            string operationalStateDescription,

            string recoveryHealthDescription,

            string synchronizationHealthDescription)
        {
            AvailabilityGroupReplicaRole = availabilityGroupReplicaRole;
            ConnectedStateDescription = connectedStateDescription;
            LastConnectErrorDescription = lastConnectErrorDescription;
            LastConnectErrorTimestamp = lastConnectErrorTimestamp;
            OperationalStateDescription = operationalStateDescription;
            RecoveryHealthDescription = recoveryHealthDescription;
            SynchronizationHealthDescription = synchronizationHealthDescription;
        }
    }
}
