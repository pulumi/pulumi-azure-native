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
    /// Connection string for the Cosmos DB account
    /// </summary>
    [OutputType]
    public sealed class DatabaseAccountConnectionStringResponse
    {
        /// <summary>
        /// Value of the connection string
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// Description of the connection string
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Kind of the connection string key
        /// </summary>
        public readonly string KeyKind;
        /// <summary>
        /// Type of the connection string
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DatabaseAccountConnectionStringResponse(
            string connectionString,

            string description,

            string keyKind,

            string type)
        {
            ConnectionString = connectionString;
            Description = description;
            KeyKind = keyKind;
            Type = type;
        }
    }
}
