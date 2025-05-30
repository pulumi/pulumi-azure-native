// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent.Outputs
{

    /// <summary>
    /// The connector service type is AzureBlobStorageSourceConnector
    /// </summary>
    [OutputType]
    public sealed class AzureBlobStorageSourceConnectorServiceInfoResponse
    {
        /// <summary>
        /// The connector service type.
        /// Expected value is 'AzureBlobStorageSourceConnector'.
        /// </summary>
        public readonly string ConnectorServiceType;
        /// <summary>
        /// Azure Blob Storage Account Key
        /// </summary>
        public readonly string? StorageAccountKey;
        /// <summary>
        /// Azure Blob Storage Account Name
        /// </summary>
        public readonly string? StorageAccountName;
        /// <summary>
        /// Azure Blob Storage Account Container Name
        /// </summary>
        public readonly string? StorageContainerName;

        [OutputConstructor]
        private AzureBlobStorageSourceConnectorServiceInfoResponse(
            string connectorServiceType,

            string? storageAccountKey,

            string? storageAccountName,

            string? storageContainerName)
        {
            ConnectorServiceType = connectorServiceType;
            StorageAccountKey = storageAccountKey;
            StorageAccountName = storageAccountName;
            StorageContainerName = storageContainerName;
        }
    }
}
