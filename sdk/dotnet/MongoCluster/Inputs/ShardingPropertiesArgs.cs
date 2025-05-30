// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MongoCluster.Inputs
{

    /// <summary>
    /// The sharding properties of the cluster. This includes the shard count and scaling options for the cluster.
    /// </summary>
    public sealed class ShardingPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of shards to provision on the cluster.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        public ShardingPropertiesArgs()
        {
        }
        public static new ShardingPropertiesArgs Empty => new ShardingPropertiesArgs();
    }
}
