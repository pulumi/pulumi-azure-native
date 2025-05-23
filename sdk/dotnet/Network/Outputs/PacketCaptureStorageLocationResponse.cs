// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// The storage location for a packet capture session.
    /// </summary>
    [OutputType]
    public sealed class PacketCaptureStorageLocationResponse
    {
        /// <summary>
        /// This path is invalid if 'Continuous Capture' is provided with 'true' or 'false'. A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with /var/captures. Required if no storage ID is provided, otherwise optional.
        /// </summary>
        public readonly string? FilePath;
        /// <summary>
        /// This path is valid if 'Continuous Capture' is provided with 'true' or 'false' and required if no storage ID is provided, otherwise optional. Must include the name of the capture file (*.cap). For linux virtual machine it must start with /var/captures.
        /// </summary>
        public readonly string? LocalPath;
        /// <summary>
        /// The ID of the storage account to save the packet capture session. Required if no localPath or filePath is provided.
        /// </summary>
        public readonly string? StorageId;
        /// <summary>
        /// The URI of the storage path to save the packet capture. Must be a well-formed URI describing the location to save the packet capture.
        /// </summary>
        public readonly string? StoragePath;

        [OutputConstructor]
        private PacketCaptureStorageLocationResponse(
            string? filePath,

            string? localPath,

            string? storageId,

            string? storagePath)
        {
            FilePath = filePath;
            LocalPath = localPath;
            StorageId = storageId;
            StoragePath = storagePath;
        }
    }
}
