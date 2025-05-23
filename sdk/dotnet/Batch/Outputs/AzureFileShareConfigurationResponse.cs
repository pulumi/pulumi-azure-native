// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    [OutputType]
    public sealed class AzureFileShareConfigurationResponse
    {
        public readonly string AccountKey;
        public readonly string AccountName;
        /// <summary>
        /// This is of the form 'https://{account}.file.core.windows.net/'.
        /// </summary>
        public readonly string AzureFileUrl;
        /// <summary>
        /// These are 'net use' options in Windows and 'mount' options in Linux.
        /// </summary>
        public readonly string? MountOptions;
        /// <summary>
        /// All file systems are mounted relative to the Batch mounts directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment variable.
        /// </summary>
        public readonly string RelativeMountPath;

        [OutputConstructor]
        private AzureFileShareConfigurationResponse(
            string accountKey,

            string accountName,

            string azureFileUrl,

            string? mountOptions,

            string relativeMountPath)
        {
            AccountKey = accountKey;
            AccountName = accountName;
            AzureFileUrl = azureFileUrl;
            MountOptions = mountOptions;
            RelativeMountPath = relativeMountPath;
        }
    }
}
