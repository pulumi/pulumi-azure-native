// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    [OutputType]
    public sealed class StorageTableDestinationResponse
    {
        /// <summary>
        /// A friendly name for the destination. 
        /// This name should be unique across all destinations (regardless of type) within the data collection rule.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource ID of the storage account.
        /// </summary>
        public readonly string? StorageAccountResourceId;
        /// <summary>
        /// The name of the Storage Table.
        /// </summary>
        public readonly string? TableName;

        [OutputConstructor]
        private StorageTableDestinationResponse(
            string? name,

            string? storageAccountResourceId,

            string? tableName)
        {
            Name = name;
            StorageAccountResourceId = storageAccountResourceId;
            TableName = tableName;
        }
    }
}
