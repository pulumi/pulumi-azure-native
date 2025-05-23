// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Reprotect agent details.
    /// </summary>
    [OutputType]
    public sealed class ReprotectAgentDetailsResponse
    {
        /// <summary>
        /// The list of accessible datastores fetched from discovery.
        /// </summary>
        public readonly ImmutableArray<string> AccessibleDatastores;
        /// <summary>
        /// The reprotect agent Bios Id.
        /// </summary>
        public readonly string BiosId;
        /// <summary>
        /// The fabric object Id.
        /// </summary>
        public readonly string FabricObjectId;
        /// <summary>
        /// The reprotect agent Fqdn.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The health of the reprotect agent.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// The health errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.HealthErrorResponse> HealthErrors;
        /// <summary>
        /// The reprotect agent Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last time when SDS information discovered in SRS.
        /// </summary>
        public readonly string LastDiscoveryInUtc;
        /// <summary>
        /// The last heartbeat received from the reprotect agent.
        /// </summary>
        public readonly string LastHeartbeatUtc;
        /// <summary>
        /// The reprotect agent name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The protected item count.
        /// </summary>
        public readonly int ProtectedItemCount;
        /// <summary>
        /// The Vcenter Id.
        /// </summary>
        public readonly string VcenterId;
        /// <summary>
        /// The version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ReprotectAgentDetailsResponse(
            ImmutableArray<string> accessibleDatastores,

            string biosId,

            string fabricObjectId,

            string fqdn,

            string health,

            ImmutableArray<Outputs.HealthErrorResponse> healthErrors,

            string id,

            string lastDiscoveryInUtc,

            string lastHeartbeatUtc,

            string name,

            int protectedItemCount,

            string vcenterId,

            string version)
        {
            AccessibleDatastores = accessibleDatastores;
            BiosId = biosId;
            FabricObjectId = fabricObjectId;
            Fqdn = fqdn;
            Health = health;
            HealthErrors = healthErrors;
            Id = id;
            LastDiscoveryInUtc = lastDiscoveryInUtc;
            LastHeartbeatUtc = lastHeartbeatUtc;
            Name = name;
            ProtectedItemCount = protectedItemCount;
            VcenterId = vcenterId;
            Version = version;
        }
    }
}
