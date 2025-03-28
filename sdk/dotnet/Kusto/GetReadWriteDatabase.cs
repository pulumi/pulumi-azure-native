// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kusto
{
    public static class GetReadWriteDatabase
    {
        /// <summary>
        /// Returns a database.
        /// 
        /// Uses Azure REST API version 2022-12-29.
        /// </summary>
        public static Task<GetReadWriteDatabaseResult> InvokeAsync(GetReadWriteDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReadWriteDatabaseResult>("azure-native:kusto:getReadWriteDatabase", args ?? new GetReadWriteDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a database.
        /// 
        /// Uses Azure REST API version 2022-12-29.
        /// </summary>
        public static Output<GetReadWriteDatabaseResult> Invoke(GetReadWriteDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReadWriteDatabaseResult>("azure-native:kusto:getReadWriteDatabase", args ?? new GetReadWriteDatabaseInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a database.
        /// 
        /// Uses Azure REST API version 2022-12-29.
        /// </summary>
        public static Output<GetReadWriteDatabaseResult> Invoke(GetReadWriteDatabaseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetReadWriteDatabaseResult>("azure-native:kusto:getReadWriteDatabase", args ?? new GetReadWriteDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReadWriteDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto cluster.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetReadWriteDatabaseArgs()
        {
        }
        public static new GetReadWriteDatabaseArgs Empty => new GetReadWriteDatabaseArgs();
    }

    public sealed class GetReadWriteDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto cluster.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetReadWriteDatabaseInvokeArgs()
        {
        }
        public static new GetReadWriteDatabaseInvokeArgs Empty => new GetReadWriteDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetReadWriteDatabaseResult
    {
        /// <summary>
        /// The time the data should be kept in cache for fast queries in TimeSpan.
        /// </summary>
        public readonly string? HotCachePeriod;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether the database is followed.
        /// </summary>
        public readonly bool IsFollowed;
        /// <summary>
        /// Kind of the database
        /// Expected value is 'ReadWrite'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The time the data should be kept before it stops being accessible to queries in TimeSpan.
        /// </summary>
        public readonly string? SoftDeletePeriod;
        /// <summary>
        /// The statistics of the database.
        /// </summary>
        public readonly Outputs.DatabaseStatisticsResponse Statistics;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetReadWriteDatabaseResult(
            string? hotCachePeriod,

            string id,

            bool isFollowed,

            string kind,

            string? location,

            string name,

            string provisioningState,

            string? softDeletePeriod,

            Outputs.DatabaseStatisticsResponse statistics,

            string type)
        {
            HotCachePeriod = hotCachePeriod;
            Id = id;
            IsFollowed = isFollowed;
            Kind = kind;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SoftDeletePeriod = softDeletePeriod;
            Statistics = statistics;
            Type = type;
        }
    }
}
