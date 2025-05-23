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
    /// The unique key on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
    /// </summary>
    public sealed class UniqueKeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// List of paths must be unique for each document in the Azure Cosmos DB service
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        public UniqueKeyArgs()
        {
        }
        public static new UniqueKeyArgs Empty => new UniqueKeyArgs();
    }
}
