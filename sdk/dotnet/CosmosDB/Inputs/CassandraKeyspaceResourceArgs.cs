// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Inputs
{

    /// <summary>
    /// Cosmos DB Cassandra keyspace id object
    /// </summary>
    public sealed class CassandraKeyspaceResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Cosmos DB Cassandra keyspace
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public CassandraKeyspaceResourceArgs()
        {
        }
        public static new CassandraKeyspaceResourceArgs Empty => new CassandraKeyspaceResourceArgs();
    }
}
