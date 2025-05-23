// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Details of the appliance resource.
    /// </summary>
    [OutputType]
    public sealed class DataStoreUtilizationDetailsResponse
    {
        /// <summary>
        /// The datastore name.
        /// </summary>
        public readonly string DataStoreName;
        /// <summary>
        /// The total snapshots created for server migration in the datastore.
        /// </summary>
        public readonly double TotalSnapshotsCreated;
        /// <summary>
        /// The total count of snapshots supported by the datastore.
        /// </summary>
        public readonly double TotalSnapshotsSupported;

        [OutputConstructor]
        private DataStoreUtilizationDetailsResponse(
            string dataStoreName,

            double totalSnapshotsCreated,

            double totalSnapshotsSupported)
        {
            DataStoreName = dataStoreName;
            TotalSnapshotsCreated = totalSnapshotsCreated;
            TotalSnapshotsSupported = totalSnapshotsSupported;
        }
    }
}
