// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Outputs
{

    /// <summary>
    /// Cosmos DB Cassandra table column
    /// </summary>
    [OutputType]
    public sealed class ColumnResponse
    {
        /// <summary>
        /// Name of the Cosmos DB Cassandra table column
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Type of the Cosmos DB Cassandra table column
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ColumnResponse(
            string? name,

            string? type)
        {
            Name = name;
            Type = type;
        }
    }
}
