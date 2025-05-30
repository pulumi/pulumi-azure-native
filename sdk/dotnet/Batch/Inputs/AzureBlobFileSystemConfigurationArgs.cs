// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Inputs
{

    public sealed class AzureBlobFileSystemConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This property is mutually exclusive with both sasKey and identity; exactly one must be specified.
        /// </summary>
        [Input("accountKey")]
        public Input<string>? AccountKey { get; set; }

        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// These are 'net use' options in Windows and 'mount' options in Linux.
        /// </summary>
        [Input("blobfuseOptions")]
        public Input<string>? BlobfuseOptions { get; set; }

        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// This property is mutually exclusive with both accountKey and sasKey; exactly one must be specified.
        /// </summary>
        [Input("identityReference")]
        public Input<Inputs.ComputeNodeIdentityReferenceArgs>? IdentityReference { get; set; }

        /// <summary>
        /// All file systems are mounted relative to the Batch mounts directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment variable.
        /// </summary>
        [Input("relativeMountPath", required: true)]
        public Input<string> RelativeMountPath { get; set; } = null!;

        /// <summary>
        /// This property is mutually exclusive with both accountKey and identity; exactly one must be specified.
        /// </summary>
        [Input("sasKey")]
        public Input<string>? SasKey { get; set; }

        public AzureBlobFileSystemConfigurationArgs()
        {
        }
        public static new AzureBlobFileSystemConfigurationArgs Empty => new AzureBlobFileSystemConfigurationArgs();
    }
}
