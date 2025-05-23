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
    /// Agent details.
    /// </summary>
    [OutputType]
    public sealed class AgentDetailsResponse
    {
        /// <summary>
        /// The Id of the agent running on the server.
        /// </summary>
        public readonly string AgentId;
        /// <summary>
        /// The machine BIOS Id.
        /// </summary>
        public readonly string BiosId;
        /// <summary>
        /// The disks.
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentDiskDetailsResponse> Disks;
        /// <summary>
        /// The machine FQDN.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The Id of the machine to which the agent is registered.
        /// </summary>
        public readonly string MachineId;

        [OutputConstructor]
        private AgentDetailsResponse(
            string agentId,

            string biosId,

            ImmutableArray<Outputs.AgentDiskDetailsResponse> disks,

            string fqdn,

            string machineId)
        {
            AgentId = agentId;
            BiosId = biosId;
            Disks = disks;
            Fqdn = fqdn;
            MachineId = machineId;
        }
    }
}
