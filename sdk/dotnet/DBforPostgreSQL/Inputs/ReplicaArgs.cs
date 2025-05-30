// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Inputs
{

    /// <summary>
    /// Replica properties of a server
    /// </summary>
    public sealed class ReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets the promote mode for a replica server. This is a write only property.
        /// </summary>
        [Input("promoteMode")]
        public InputUnion<string, Pulumi.AzureNative.DBforPostgreSQL.ReadReplicaPromoteMode>? PromoteMode { get; set; }

        /// <summary>
        /// Sets the promote options for a replica server. This is a write only property.
        /// </summary>
        [Input("promoteOption")]
        public InputUnion<string, Pulumi.AzureNative.DBforPostgreSQL.ReplicationPromoteOption>? PromoteOption { get; set; }

        /// <summary>
        /// Used to indicate role of the server in replication set.
        /// </summary>
        [Input("role")]
        public InputUnion<string, Pulumi.AzureNative.DBforPostgreSQL.ReplicationRole>? Role { get; set; }

        public ReplicaArgs()
        {
        }
        public static new ReplicaArgs Empty => new ReplicaArgs();
    }
}
