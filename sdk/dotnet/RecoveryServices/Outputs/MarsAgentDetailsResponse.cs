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
    /// Mars agent details.
    /// </summary>
    [OutputType]
    public sealed class MarsAgentDetailsResponse
    {
        /// <summary>
        /// The Mars agent Bios Id.
        /// </summary>
        public readonly string BiosId;
        /// <summary>
        /// The fabric object Id.
        /// </summary>
        public readonly string FabricObjectId;
        /// <summary>
        /// The Mars agent Fqdn.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The health of the Mars agent.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// The health errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.HealthErrorResponse> HealthErrors;
        /// <summary>
        /// The Mars agent Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last heartbeat received from the Mars agent.
        /// </summary>
        public readonly string LastHeartbeatUtc;
        /// <summary>
        /// The Mars agent name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private MarsAgentDetailsResponse(
            string biosId,

            string fabricObjectId,

            string fqdn,

            string health,

            ImmutableArray<Outputs.HealthErrorResponse> healthErrors,

            string id,

            string lastHeartbeatUtc,

            string name,

            string version)
        {
            BiosId = biosId;
            FabricObjectId = fabricObjectId;
            Fqdn = fqdn;
            Health = health;
            HealthErrors = healthErrors;
            Id = id;
            LastHeartbeatUtc = lastHeartbeatUtc;
            Name = name;
            Version = version;
        }
    }
}
