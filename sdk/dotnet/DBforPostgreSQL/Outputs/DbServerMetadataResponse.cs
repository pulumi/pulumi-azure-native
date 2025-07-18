// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Outputs
{

    /// <summary>
    /// Database server metadata.
    /// </summary>
    [OutputType]
    public sealed class DbServerMetadataResponse
    {
        /// <summary>
        /// Location of database server.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Compute tier and size of the database server. This object is empty for an Azure Database for PostgreSQL single server.
        /// </summary>
        public readonly Outputs.ServerSkuResponse? Sku;
        /// <summary>
        /// Storage size (in MB) for database server.
        /// </summary>
        public readonly int? StorageMb;
        /// <summary>
        /// Major version of PostgreSQL database engine.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private DbServerMetadataResponse(
            string location,

            Outputs.ServerSkuResponse? sku,

            int? storageMb,

            string? version)
        {
            Location = location;
            Sku = sku;
            StorageMb = storageMb;
            Version = version;
        }
    }
}
