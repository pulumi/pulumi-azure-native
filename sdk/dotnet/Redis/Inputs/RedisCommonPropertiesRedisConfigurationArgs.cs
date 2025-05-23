// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Redis.Inputs
{

    /// <summary>
    /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta, maxmemory-policy,notify-keyspace-events, aof-backup-enabled, aof-storage-connection-string-0, aof-storage-connection-string-1 etc.
    /// </summary>
    public sealed class RedisCommonPropertiesRedisConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether AAD based authentication has been enabled or disabled for the cache
        /// </summary>
        [Input("aadEnabled")]
        public Input<string>? AadEnabled { get; set; }

        /// <summary>
        /// Specifies whether the aof backup is enabled
        /// </summary>
        [Input("aofBackupEnabled")]
        public Input<string>? AofBackupEnabled { get; set; }

        /// <summary>
        /// First storage account connection string
        /// </summary>
        [Input("aofStorageConnectionString0")]
        public Input<string>? AofStorageConnectionString0 { get; set; }

        /// <summary>
        /// Second storage account connection string
        /// </summary>
        [Input("aofStorageConnectionString1")]
        public Input<string>? AofStorageConnectionString1 { get; set; }

        /// <summary>
        /// Specifies whether the authentication is disabled. Setting this property is highly discouraged from security point of view.
        /// </summary>
        [Input("authnotrequired")]
        public Input<string>? Authnotrequired { get; set; }

        /// <summary>
        /// Value in megabytes reserved for fragmentation per shard
        /// </summary>
        [Input("maxfragmentationmemoryReserved")]
        public Input<string>? MaxfragmentationmemoryReserved { get; set; }

        /// <summary>
        /// Value in megabytes reserved for non-cache usage per shard e.g. failover.
        /// </summary>
        [Input("maxmemoryDelta")]
        public Input<string>? MaxmemoryDelta { get; set; }

        /// <summary>
        /// The eviction strategy used when your data won't fit within its memory limit.
        /// </summary>
        [Input("maxmemoryPolicy")]
        public Input<string>? MaxmemoryPolicy { get; set; }

        /// <summary>
        /// Value in megabytes reserved for non-cache usage per shard e.g. failover.
        /// </summary>
        [Input("maxmemoryReserved")]
        public Input<string>? MaxmemoryReserved { get; set; }

        /// <summary>
        /// The keyspace events which should be monitored.
        /// </summary>
        [Input("notifyKeyspaceEvents")]
        public Input<string>? NotifyKeyspaceEvents { get; set; }

        /// <summary>
        /// Preferred auth method to communicate to storage account used for data persistence, specify SAS or ManagedIdentity, default value is SAS
        /// </summary>
        [Input("preferredDataPersistenceAuthMethod")]
        public Input<string>? PreferredDataPersistenceAuthMethod { get; set; }

        /// <summary>
        /// Specifies whether the rdb backup is enabled
        /// </summary>
        [Input("rdbBackupEnabled")]
        public Input<string>? RdbBackupEnabled { get; set; }

        /// <summary>
        /// Specifies the frequency for creating rdb backup in minutes. Valid values: (15, 30, 60, 360, 720, 1440)
        /// </summary>
        [Input("rdbBackupFrequency")]
        public Input<string>? RdbBackupFrequency { get; set; }

        /// <summary>
        /// Specifies the maximum number of snapshots for rdb backup
        /// </summary>
        [Input("rdbBackupMaxSnapshotCount")]
        public Input<string>? RdbBackupMaxSnapshotCount { get; set; }

        /// <summary>
        /// The storage account connection string for storing rdb file
        /// </summary>
        [Input("rdbStorageConnectionString")]
        public Input<string>? RdbStorageConnectionString { get; set; }

        /// <summary>
        /// SubscriptionId of the storage account for persistence (aof/rdb) using ManagedIdentity.
        /// </summary>
        [Input("storageSubscriptionId")]
        public Input<string>? StorageSubscriptionId { get; set; }

        public RedisCommonPropertiesRedisConfigurationArgs()
        {
        }
        public static new RedisCommonPropertiesRedisConfigurationArgs Empty => new RedisCommonPropertiesRedisConfigurationArgs();
    }
}
