// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesRuntime.Inputs
{

    /// <summary>
    /// The properties of NFS StorageClass
    /// </summary>
    public sealed class NfsStorageClassTypePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mounted folder permissions. Default is 0. If set as non-zero, driver will perform `chmod` after mount
        /// </summary>
        [Input("mountPermissions")]
        public Input<string>? MountPermissions { get; set; }

        /// <summary>
        /// The action to take when a NFS volume is deleted. Default is Delete
        /// </summary>
        [Input("onDelete")]
        public InputUnion<string, Pulumi.AzureNative.KubernetesRuntime.NfsDirectoryActionOnVolumeDeletion>? OnDelete { get; set; }

        /// <summary>
        /// NFS Server
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// NFS share
        /// </summary>
        [Input("share", required: true)]
        public Input<string> Share { get; set; } = null!;

        /// <summary>
        /// Sub directory under share. If the sub directory doesn't exist, driver will create it
        /// </summary>
        [Input("subDir")]
        public Input<string>? SubDir { get; set; }

        /// <summary>
        /// Type of a storage class
        /// Expected value is 'NFS'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NfsStorageClassTypePropertiesArgs()
        {
        }
        public static new NfsStorageClassTypePropertiesArgs Empty => new NfsStorageClassTypePropertiesArgs();
    }
}
