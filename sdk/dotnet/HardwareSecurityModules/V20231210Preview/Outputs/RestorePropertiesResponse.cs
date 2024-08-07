// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HardwareSecurityModules.V20231210Preview.Outputs
{

    /// <summary>
    /// Cloud Hsm Cluster restore information
    /// </summary>
    [OutputType]
    public sealed class RestorePropertiesResponse
    {
        /// <summary>
        /// Azure Blob storage container Uri
        /// </summary>
        public readonly string? AzureStorageResourceUri;
        /// <summary>
        /// Directory name in Azure Storage Blob where the backup is stored
        /// </summary>
        public readonly string? Foldername;

        [OutputConstructor]
        private RestorePropertiesResponse(
            string? azureStorageResourceUri,

            string? foldername)
        {
            AzureStorageResourceUri = azureStorageResourceUri;
            Foldername = foldername;
        }
    }
}
