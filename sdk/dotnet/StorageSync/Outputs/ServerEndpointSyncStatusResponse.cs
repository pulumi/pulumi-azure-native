// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageSync.Outputs
{

    /// <summary>
    /// Server Endpoint sync status
    /// </summary>
    [OutputType]
    public sealed class ServerEndpointSyncStatusResponse
    {
        /// <summary>
        /// Background data download activity
        /// </summary>
        public readonly Outputs.ServerEndpointBackgroundDataDownloadActivityResponse BackgroundDataDownloadActivity;
        /// <summary>
        /// Combined Health Status.
        /// </summary>
        public readonly string CombinedHealth;
        /// <summary>
        /// Download sync activity
        /// </summary>
        public readonly Outputs.ServerEndpointSyncActivityStatusResponse DownloadActivity;
        /// <summary>
        /// Download Health Status.
        /// </summary>
        public readonly string DownloadHealth;
        /// <summary>
        /// Download Status
        /// </summary>
        public readonly Outputs.ServerEndpointSyncSessionStatusResponse DownloadStatus;
        /// <summary>
        /// Last Updated Timestamp
        /// </summary>
        public readonly string LastUpdatedTimestamp;
        /// <summary>
        /// Offline Data Transfer State
        /// </summary>
        public readonly string OfflineDataTransferStatus;
        /// <summary>
        /// Sync activity
        /// </summary>
        public readonly string SyncActivity;
        /// <summary>
        /// Total count of persistent files not syncing (combined upload + download).
        /// </summary>
        public readonly double TotalPersistentFilesNotSyncingCount;
        /// <summary>
        /// Upload sync activity
        /// </summary>
        public readonly Outputs.ServerEndpointSyncActivityStatusResponse UploadActivity;
        /// <summary>
        /// Upload Health Status.
        /// </summary>
        public readonly string UploadHealth;
        /// <summary>
        /// Upload Status
        /// </summary>
        public readonly Outputs.ServerEndpointSyncSessionStatusResponse UploadStatus;

        [OutputConstructor]
        private ServerEndpointSyncStatusResponse(
            Outputs.ServerEndpointBackgroundDataDownloadActivityResponse backgroundDataDownloadActivity,

            string combinedHealth,

            Outputs.ServerEndpointSyncActivityStatusResponse downloadActivity,

            string downloadHealth,

            Outputs.ServerEndpointSyncSessionStatusResponse downloadStatus,

            string lastUpdatedTimestamp,

            string offlineDataTransferStatus,

            string syncActivity,

            double totalPersistentFilesNotSyncingCount,

            Outputs.ServerEndpointSyncActivityStatusResponse uploadActivity,

            string uploadHealth,

            Outputs.ServerEndpointSyncSessionStatusResponse uploadStatus)
        {
            BackgroundDataDownloadActivity = backgroundDataDownloadActivity;
            CombinedHealth = combinedHealth;
            DownloadActivity = downloadActivity;
            DownloadHealth = downloadHealth;
            DownloadStatus = downloadStatus;
            LastUpdatedTimestamp = lastUpdatedTimestamp;
            OfflineDataTransferStatus = offlineDataTransferStatus;
            SyncActivity = syncActivity;
            TotalPersistentFilesNotSyncingCount = totalPersistentFilesNotSyncingCount;
            UploadActivity = uploadActivity;
            UploadHealth = uploadHealth;
            UploadStatus = uploadStatus;
        }
    }
}
