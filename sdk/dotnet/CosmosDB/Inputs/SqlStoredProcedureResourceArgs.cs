// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Inputs
{

    /// <summary>
    /// Cosmos DB SQL storedProcedure resource object
    /// </summary>
    public sealed class SqlStoredProcedureResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Body of the Stored Procedure
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Name of the Cosmos DB SQL storedProcedure
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public SqlStoredProcedureResourceArgs()
        {
        }
        public static new SqlStoredProcedureResourceArgs Empty => new SqlStoredProcedureResourceArgs();
    }
}
