// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20201201Preview.Inputs
{

    /// <summary>
    /// The storage account information for storing Service Fabric diagnostic logs.
    /// </summary>
    public sealed class DiagnosticsStorageAccountConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The blob endpoint of the azure storage account.
        /// </summary>
        [Input("blobEndpoint", required: true)]
        public Input<string> BlobEndpoint { get; set; } = null!;

        /// <summary>
        /// The protected diagnostics storage key name.
        /// </summary>
        [Input("protectedAccountKeyName", required: true)]
        public Input<string> ProtectedAccountKeyName { get; set; } = null!;

        /// <summary>
        /// The secondary protected diagnostics storage key name. If one of the storage account keys is rotated the cluster will fallback to using the other.
        /// </summary>
        [Input("protectedAccountKeyName2")]
        public Input<string>? ProtectedAccountKeyName2 { get; set; }

        /// <summary>
        /// The queue endpoint of the azure storage account.
        /// </summary>
        [Input("queueEndpoint", required: true)]
        public Input<string> QueueEndpoint { get; set; } = null!;

        /// <summary>
        /// The Azure storage account name.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// The table endpoint of the azure storage account.
        /// </summary>
        [Input("tableEndpoint", required: true)]
        public Input<string> TableEndpoint { get; set; } = null!;

        public DiagnosticsStorageAccountConfigArgs()
        {
        }
        public static new DiagnosticsStorageAccountConfigArgs Empty => new DiagnosticsStorageAccountConfigArgs();
    }
}
