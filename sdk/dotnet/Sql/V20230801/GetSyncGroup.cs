// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20230801
{
    public static class GetSyncGroup
    {
        /// <summary>
        /// Gets a sync group.
        /// </summary>
        public static Task<GetSyncGroupResult> InvokeAsync(GetSyncGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSyncGroupResult>("azure-native:sql/v20230801:getSyncGroup", args ?? new GetSyncGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a sync group.
        /// </summary>
        public static Output<GetSyncGroupResult> Invoke(GetSyncGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncGroupResult>("azure-native:sql/v20230801:getSyncGroup", args ?? new GetSyncGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a sync group.
        /// </summary>
        public static Output<GetSyncGroupResult> Invoke(GetSyncGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncGroupResult>("azure-native:sql/v20230801:getSyncGroup", args ?? new GetSyncGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSyncGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database on which the sync group is hosted.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

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

        /// <summary>
        /// The name of the sync group.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public string SyncGroupName { get; set; } = null!;

        public GetSyncGroupArgs()
        {
        }
        public static new GetSyncGroupArgs Empty => new GetSyncGroupArgs();
    }

    public sealed class GetSyncGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database on which the sync group is hosted.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

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

        /// <summary>
        /// The name of the sync group.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public Input<string> SyncGroupName { get; set; } = null!;

        public GetSyncGroupInvokeArgs()
        {
        }
        public static new GetSyncGroupInvokeArgs Empty => new GetSyncGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSyncGroupResult
    {
        /// <summary>
        /// Conflict logging retention period.
        /// </summary>
        public readonly int? ConflictLoggingRetentionInDays;
        /// <summary>
        /// Conflict resolution policy of the sync group.
        /// </summary>
        public readonly string? ConflictResolutionPolicy;
        /// <summary>
        /// If conflict logging is enabled.
        /// </summary>
        public readonly bool? EnableConflictLogging;
        /// <summary>
        /// User name for the sync group hub database credential.
        /// </summary>
        public readonly string? HubDatabaseUserName;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Sync interval of the sync group.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// Last sync time of the sync group.
        /// </summary>
        public readonly string LastSyncTime;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint name of the sync group if use private link connection is enabled.
        /// </summary>
        public readonly string PrivateEndpointName;
        /// <summary>
        /// Sync schema of the sync group.
        /// </summary>
        public readonly Outputs.SyncGroupSchemaResponse? Schema;
        /// <summary>
        /// The name and capacity of the SKU.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// ARM resource id of the sync database in the sync group.
        /// </summary>
        public readonly string? SyncDatabaseId;
        /// <summary>
        /// Sync state of the sync group.
        /// </summary>
        public readonly string SyncState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// If use private link connection is enabled.
        /// </summary>
        public readonly bool? UsePrivateLinkConnection;

        [OutputConstructor]
        private GetSyncGroupResult(
            int? conflictLoggingRetentionInDays,

            string? conflictResolutionPolicy,

            bool? enableConflictLogging,

            string? hubDatabaseUserName,

            string id,

            int? interval,

            string lastSyncTime,

            string name,

            string privateEndpointName,

            Outputs.SyncGroupSchemaResponse? schema,

            Outputs.SkuResponse? sku,

            string? syncDatabaseId,

            string syncState,

            string type,

            bool? usePrivateLinkConnection)
        {
            ConflictLoggingRetentionInDays = conflictLoggingRetentionInDays;
            ConflictResolutionPolicy = conflictResolutionPolicy;
            EnableConflictLogging = enableConflictLogging;
            HubDatabaseUserName = hubDatabaseUserName;
            Id = id;
            Interval = interval;
            LastSyncTime = lastSyncTime;
            Name = name;
            PrivateEndpointName = privateEndpointName;
            Schema = schema;
            Sku = sku;
            SyncDatabaseId = syncDatabaseId;
            SyncState = syncState;
            Type = type;
            UsePrivateLinkConnection = usePrivateLinkConnection;
        }
    }
}
