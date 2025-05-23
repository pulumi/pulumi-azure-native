// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// Management policy action for snapshot.
    /// </summary>
    [OutputType]
    public sealed class ManagementPolicySnapShotResponse
    {
        /// <summary>
        /// The function to delete the blob snapshot
        /// </summary>
        public readonly Outputs.DateAfterCreationResponse? Delete;
        /// <summary>
        /// The function to tier blob snapshot to archive storage.
        /// </summary>
        public readonly Outputs.DateAfterCreationResponse? TierToArchive;
        /// <summary>
        /// The function to tier blobs to cold storage.
        /// </summary>
        public readonly Outputs.DateAfterCreationResponse? TierToCold;
        /// <summary>
        /// The function to tier blob snapshot to cool storage.
        /// </summary>
        public readonly Outputs.DateAfterCreationResponse? TierToCool;
        /// <summary>
        /// The function to tier blobs to hot storage. This action can only be used with Premium Block Blob Storage Accounts
        /// </summary>
        public readonly Outputs.DateAfterCreationResponse? TierToHot;

        [OutputConstructor]
        private ManagementPolicySnapShotResponse(
            Outputs.DateAfterCreationResponse? delete,

            Outputs.DateAfterCreationResponse? tierToArchive,

            Outputs.DateAfterCreationResponse? tierToCold,

            Outputs.DateAfterCreationResponse? tierToCool,

            Outputs.DateAfterCreationResponse? tierToHot)
        {
            Delete = delete;
            TierToArchive = tierToArchive;
            TierToCold = tierToCold;
            TierToCool = tierToCool;
            TierToHot = tierToHot;
        }
    }
}
