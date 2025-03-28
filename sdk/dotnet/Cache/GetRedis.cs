// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cache
{
    public static class GetRedis
    {
        /// <summary>
        /// Gets a Redis cache (resource description).
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2020-06-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview, 2024-11-01.
        /// </summary>
        public static Task<GetRedisResult> InvokeAsync(GetRedisArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRedisResult>("azure-native:cache:getRedis", args ?? new GetRedisArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Redis cache (resource description).
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2020-06-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview, 2024-11-01.
        /// </summary>
        public static Output<GetRedisResult> Invoke(GetRedisInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRedisResult>("azure-native:cache:getRedis", args ?? new GetRedisInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Redis cache (resource description).
        /// 
        /// Uses Azure REST API version 2023-04-01.
        /// 
        /// Other available API versions: 2020-06-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview, 2024-11-01.
        /// </summary>
        public static Output<GetRedisResult> Invoke(GetRedisInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRedisResult>("azure-native:cache:getRedis", args ?? new GetRedisInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRedisArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRedisArgs()
        {
        }
        public static new GetRedisArgs Empty => new GetRedisArgs();
    }

    public sealed class GetRedisInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetRedisInvokeArgs()
        {
        }
        public static new GetRedisInvokeArgs Empty => new GetRedisInvokeArgs();
    }


    [OutputType]
    public sealed class GetRedisResult
    {
        /// <summary>
        /// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
        /// </summary>
        public readonly Outputs.RedisAccessKeysResponse AccessKeys;
        /// <summary>
        /// Specifies whether the non-ssl Redis server port (6379) is enabled.
        /// </summary>
        public readonly bool? EnableNonSslPort;
        /// <summary>
        /// Redis host name.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// List of the Redis instances associated with the cache
        /// </summary>
        public readonly ImmutableArray<Outputs.RedisInstanceDetailsResponse> Instances;
        /// <summary>
        /// List of the linked servers associated with the cache
        /// </summary>
        public readonly ImmutableArray<Outputs.RedisLinkedServerResponse> LinkedServers;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
        /// </summary>
        public readonly string? MinimumTlsVersion;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Redis non-SSL port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// List of private endpoint connection associated with the specified redis cache
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Redis instance provisioning status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Whether or not public endpoint access is allowed for this cache.  Value is optional, but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is 'Enabled'. Note: This setting is important for caches with private endpoints. It has *no effect* on caches that are joined to, or injected into, a virtual network subnet.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
        /// </summary>
        public readonly Outputs.RedisCommonPropertiesResponseRedisConfiguration? RedisConfiguration;
        /// <summary>
        /// Redis version. This should be in the form 'major[.minor]' (only 'major' is required) or the value 'latest' which refers to the latest stable Redis version that is available. Supported versions: 4.0, 6.0 (latest). Default value is 'latest'.
        /// </summary>
        public readonly string? RedisVersion;
        /// <summary>
        /// The number of replicas to be created per primary.
        /// </summary>
        public readonly int? ReplicasPerMaster;
        /// <summary>
        /// The number of replicas to be created per primary.
        /// </summary>
        public readonly int? ReplicasPerPrimary;
        /// <summary>
        /// The number of shards to be created on a Premium Cluster Cache.
        /// </summary>
        public readonly int? ShardCount;
        /// <summary>
        /// The SKU of the Redis cache to deploy.
        /// </summary>
        public readonly Outputs.SkuResponse Sku;
        /// <summary>
        /// Redis SSL port.
        /// </summary>
        public readonly int SslPort;
        /// <summary>
        /// Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual Network; auto assigned by default.
        /// </summary>
        public readonly string? StaticIP;
        /// <summary>
        /// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// A dictionary of tenant settings
        /// </summary>
        public readonly ImmutableDictionary<string, string>? TenantSettings;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetRedisResult(
            Outputs.RedisAccessKeysResponse accessKeys,

            bool? enableNonSslPort,

            string hostName,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            ImmutableArray<Outputs.RedisInstanceDetailsResponse> instances,

            ImmutableArray<Outputs.RedisLinkedServerResponse> linkedServers,

            string location,

            string? minimumTlsVersion,

            string name,

            int port,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            Outputs.RedisCommonPropertiesResponseRedisConfiguration? redisConfiguration,

            string? redisVersion,

            int? replicasPerMaster,

            int? replicasPerPrimary,

            int? shardCount,

            Outputs.SkuResponse sku,

            int sslPort,

            string? staticIP,

            string? subnetId,

            ImmutableDictionary<string, string>? tags,

            ImmutableDictionary<string, string>? tenantSettings,

            string type,

            ImmutableArray<string> zones)
        {
            AccessKeys = accessKeys;
            EnableNonSslPort = enableNonSslPort;
            HostName = hostName;
            Id = id;
            Identity = identity;
            Instances = instances;
            LinkedServers = linkedServers;
            Location = location;
            MinimumTlsVersion = minimumTlsVersion;
            Name = name;
            Port = port;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            RedisConfiguration = redisConfiguration;
            RedisVersion = redisVersion;
            ReplicasPerMaster = replicasPerMaster;
            ReplicasPerPrimary = replicasPerPrimary;
            ShardCount = shardCount;
            Sku = sku;
            SslPort = sslPort;
            StaticIP = staticIP;
            SubnetId = subnetId;
            Tags = tags;
            TenantSettings = tenantSettings;
            Type = type;
            Zones = zones;
        }
    }
}
