// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Inputs
{

    public sealed class AzureFileShareConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountKey", required: true)]
        public Input<string> AccountKey { get; set; } = null!;

        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// This is of the form 'https://{account}.file.core.windows.net/'.
        /// </summary>
        [Input("azureFileUrl", required: true)]
        public Input<string> AzureFileUrl { get; set; } = null!;

        /// <summary>
        /// These are 'net use' options in Windows and 'mount' options in Linux.
        /// </summary>
        [Input("mountOptions")]
        public Input<string>? MountOptions { get; set; }

        /// <summary>
        /// All file systems are mounted relative to the Batch mounts directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment variable.
        /// </summary>
        [Input("relativeMountPath", required: true)]
        public Input<string> RelativeMountPath { get; set; } = null!;

        public AzureFileShareConfigurationArgs()
        {
        }
        public static new AzureFileShareConfigurationArgs Empty => new AzureFileShareConfigurationArgs();
    }
}
