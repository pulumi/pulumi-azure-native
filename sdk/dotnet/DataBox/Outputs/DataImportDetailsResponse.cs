// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Outputs
{

    /// <summary>
    /// Details of the data to be used for importing data to azure.
    /// </summary>
    [OutputType]
    public sealed class DataImportDetailsResponse
    {
        /// <summary>
        /// Account details of the data to be transferred
        /// </summary>
        public readonly Union<Outputs.ManagedDiskDetailsResponse, Outputs.StorageAccountDetailsResponse> AccountDetails;
        /// <summary>
        /// Level of the logs to be collected.
        /// </summary>
        public readonly string? LogCollectionLevel;

        [OutputConstructor]
        private DataImportDetailsResponse(
            Union<Outputs.ManagedDiskDetailsResponse, Outputs.StorageAccountDetailsResponse> accountDetails,

            string? logCollectionLevel)
        {
            AccountDetails = accountDetails;
            LogCollectionLevel = logCollectionLevel;
        }
    }
}
