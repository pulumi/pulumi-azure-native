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
    /// The properties of Arc Sql availability group replica resource
    /// </summary>
    [OutputType]
    public sealed class SqlAvailabilityGroupReplicaResourcePropertiesResponse
    {
        /// <summary>
        /// null
        /// </summary>
        public readonly Outputs.AvailabilityGroupConfigureResponse? Configure;
        /// <summary>
        /// ID GUID of the availability group.
        /// </summary>
        public readonly string ReplicaId;
        /// <summary>
        /// the replica name.
        /// </summary>
        public readonly string? ReplicaName;
        /// <summary>
        /// null
        /// </summary>
        public readonly Outputs.AvailabilityGroupStateResponse? State;

        [OutputConstructor]
        private SqlAvailabilityGroupReplicaResourcePropertiesResponse(
            Outputs.AvailabilityGroupConfigureResponse? configure,

            string replicaId,

            string? replicaName,

            Outputs.AvailabilityGroupStateResponse? state)
        {
            Configure = configure;
            ReplicaId = replicaId;
            ReplicaName = replicaName;
            State = state;
        }
    }
}
