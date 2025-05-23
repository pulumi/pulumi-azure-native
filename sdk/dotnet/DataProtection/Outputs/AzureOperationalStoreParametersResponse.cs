// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Outputs
{

    /// <summary>
    /// Parameters for Operational-Tier DataStore
    /// </summary>
    [OutputType]
    public sealed class AzureOperationalStoreParametersResponse
    {
        /// <summary>
        /// type of datastore; Operational/Vault/Archive
        /// </summary>
        public readonly string DataStoreType;
        /// <summary>
        /// Type of the specific object - used for deserializing
        /// Expected value is 'AzureOperationalStoreParameters'.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// Gets or sets the Snapshot Resource Group Uri.
        /// </summary>
        public readonly string? ResourceGroupId;

        [OutputConstructor]
        private AzureOperationalStoreParametersResponse(
            string dataStoreType,

            string objectType,

            string? resourceGroupId)
        {
            DataStoreType = dataStoreType;
            ObjectType = objectType;
            ResourceGroupId = resourceGroupId;
        }
    }
}
