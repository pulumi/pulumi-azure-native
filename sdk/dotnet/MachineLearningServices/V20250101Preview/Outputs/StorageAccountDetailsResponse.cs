// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    /// <summary>
    /// Details of storage account to be used for the Registry
    /// </summary>
    [OutputType]
    public sealed class StorageAccountDetailsResponse
    {
        /// <summary>
        /// Details of system created storage account to be used for the registry
        /// </summary>
        public readonly Outputs.SystemCreatedStorageAccountResponse? SystemCreatedStorageAccount;
        /// <summary>
        /// Details of user created storage account to be used for the registry
        /// </summary>
        public readonly Outputs.UserCreatedStorageAccountResponse? UserCreatedStorageAccount;

        [OutputConstructor]
        private StorageAccountDetailsResponse(
            Outputs.SystemCreatedStorageAccountResponse? systemCreatedStorageAccount,

            Outputs.UserCreatedStorageAccountResponse? userCreatedStorageAccount)
        {
            SystemCreatedStorageAccount = systemCreatedStorageAccount;
            UserCreatedStorageAccount = userCreatedStorageAccount;
        }
    }
}
