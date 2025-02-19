// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.V20240901.Outputs
{

    /// <summary>
    /// Storage details of all the Storage accounts attached to the VM. For e.g. NFS on AFS Shared Storage.
    /// </summary>
    [OutputType]
    public sealed class StorageInformationResponse
    {
        /// <summary>
        /// Fully qualified resource ID for the storage account.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private StorageInformationResponse(string id)
        {
            Id = id;
        }
    }
}
