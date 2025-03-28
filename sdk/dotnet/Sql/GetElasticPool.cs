// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    public static class GetElasticPool
    {
        /// <summary>
        /// Gets an elastic pool.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2014-04-01, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Task<GetElasticPoolResult> InvokeAsync(GetElasticPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetElasticPoolResult>("azure-native:sql:getElasticPool", args ?? new GetElasticPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an elastic pool.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2014-04-01, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetElasticPoolResult> Invoke(GetElasticPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetElasticPoolResult>("azure-native:sql:getElasticPool", args ?? new GetElasticPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an elastic pool.
        /// 
        /// Uses Azure REST API version 2021-11-01.
        /// 
        /// Other available API versions: 2014-04-01, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetElasticPoolResult> Invoke(GetElasticPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetElasticPoolResult>("azure-native:sql:getElasticPool", args ?? new GetElasticPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetElasticPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the elastic pool.
        /// </summary>
        [Input("elasticPoolName", required: true)]
        public string ElasticPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetElasticPoolArgs()
        {
        }
        public static new GetElasticPoolArgs Empty => new GetElasticPoolArgs();
    }

    public sealed class GetElasticPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the elastic pool.
        /// </summary>
        [Input("elasticPoolName", required: true)]
        public Input<string> ElasticPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public GetElasticPoolInvokeArgs()
        {
        }
        public static new GetElasticPoolInvokeArgs Empty => new GetElasticPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetElasticPoolResult
    {
        /// <summary>
        /// The creation date of the elastic pool (ISO8601 format).
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
        /// </summary>
        public readonly int? HighAvailabilityReplicaCount;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of elastic pool. This is metadata used for the Azure portal experience.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The license type to apply for this elastic pool.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
        /// </summary>
        public readonly string? MaintenanceConfigurationId;
        /// <summary>
        /// The storage limit for the database elastic pool in bytes.
        /// </summary>
        public readonly double? MaxSizeBytes;
        /// <summary>
        /// Minimal capacity that serverless pool will not shrink below, if not paused
        /// </summary>
        public readonly double? MinCapacity;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The per database settings for the elastic pool.
        /// </summary>
        public readonly Outputs.ElasticPoolPerDatabaseSettingsResponse? PerDatabaseSettings;
        /// <summary>
        /// The elastic pool SKU.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
        /// 
        /// ```azurecli
        /// az sql elastic-pool list-editions -l &lt;location&gt; -o table
        /// ````
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The state of the elastic pool.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
        /// </summary>
        public readonly bool? ZoneRedundant;

        [OutputConstructor]
        private GetElasticPoolResult(
            string creationDate,

            int? highAvailabilityReplicaCount,

            string id,

            string kind,

            string? licenseType,

            string location,

            string? maintenanceConfigurationId,

            double? maxSizeBytes,

            double? minCapacity,

            string name,

            Outputs.ElasticPoolPerDatabaseSettingsResponse? perDatabaseSettings,

            Outputs.SkuResponse? sku,

            string state,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? zoneRedundant)
        {
            CreationDate = creationDate;
            HighAvailabilityReplicaCount = highAvailabilityReplicaCount;
            Id = id;
            Kind = kind;
            LicenseType = licenseType;
            Location = location;
            MaintenanceConfigurationId = maintenanceConfigurationId;
            MaxSizeBytes = maxSizeBytes;
            MinCapacity = minCapacity;
            Name = name;
            PerDatabaseSettings = perDatabaseSettings;
            Sku = sku;
            State = state;
            Tags = tags;
            Type = type;
            ZoneRedundant = zoneRedundant;
        }
    }
}
