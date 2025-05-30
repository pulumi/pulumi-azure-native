// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesRuntime.Outputs
{

    /// <summary>
    /// The properties of NFS StorageClass
    /// </summary>
    [OutputType]
    public sealed class NfsStorageClassTypePropertiesResponse
    {
        /// <summary>
        /// Mounted folder permissions. Default is 0. If set as non-zero, driver will perform `chmod` after mount
        /// </summary>
        public readonly string? MountPermissions;
        /// <summary>
        /// The action to take when a NFS volume is deleted. Default is Delete
        /// </summary>
        public readonly string? OnDelete;
        /// <summary>
        /// NFS Server
        /// </summary>
        public readonly string Server;
        /// <summary>
        /// NFS share
        /// </summary>
        public readonly string Share;
        /// <summary>
        /// Sub directory under share. If the sub directory doesn't exist, driver will create it
        /// </summary>
        public readonly string? SubDir;
        /// <summary>
        /// Type of a storage class
        /// Expected value is 'NFS'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NfsStorageClassTypePropertiesResponse(
            string? mountPermissions,

            string? onDelete,

            string server,

            string share,

            string? subDir,

            string type)
        {
            MountPermissions = mountPermissions;
            OnDelete = onDelete;
            Server = server;
            Share = share;
            SubDir = subDir;
            Type = type;
        }
    }
}
