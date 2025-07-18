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
    /// Compute information of a flexible server.
    /// </summary>
    [OutputType]
    public sealed class ServerSkuResponse
    {
        /// <summary>
        /// Compute tier and size of the database server. This object is empty for an Azure Database for PostgreSQL single server.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tier of the compute assigned to a flexible server.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private ServerSkuResponse(
            string? name,

            string? tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
