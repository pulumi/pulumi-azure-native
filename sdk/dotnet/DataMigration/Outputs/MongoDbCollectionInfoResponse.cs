// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    /// <summary>
    /// Describes a supported collection within a MongoDB database
    /// </summary>
    [OutputType]
    public sealed class MongoDbCollectionInfoResponse
    {
        /// <summary>
        /// The average document size, or -1 if the average size is unknown
        /// </summary>
        public readonly double AverageDocumentSize;
        /// <summary>
        /// The estimated total data size, in bytes, or -1 if the size is unknown.
        /// </summary>
        public readonly double DataSize;
        /// <summary>
        /// The name of the database containing the collection
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The estimated total number of documents, or -1 if the document count is unknown
        /// </summary>
        public readonly double DocumentCount;
        /// <summary>
        /// Whether the collection is a capped collection (i.e. whether it has a fixed size and acts like a circular buffer)
        /// </summary>
        public readonly bool IsCapped;
        /// <summary>
        /// Whether the collection is system collection
        /// </summary>
        public readonly bool IsSystemCollection;
        /// <summary>
        /// Whether the collection is a view of another collection
        /// </summary>
        public readonly bool IsView;
        /// <summary>
        /// The unqualified name of the database or collection
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The qualified name of the database or collection. For a collection, this is the database-qualified name.
        /// </summary>
        public readonly string QualifiedName;
        /// <summary>
        /// The shard key on the collection, or null if the collection is not sharded
        /// </summary>
        public readonly Outputs.MongoDbShardKeyInfoResponse? ShardKey;
        /// <summary>
        /// Whether the database has sharding enabled. Note that the migration task will enable sharding on the target if necessary.
        /// </summary>
        public readonly bool SupportsSharding;
        /// <summary>
        /// The name of the collection that this is a view of, if IsView is true
        /// </summary>
        public readonly string? ViewOf;

        [OutputConstructor]
        private MongoDbCollectionInfoResponse(
            double averageDocumentSize,

            double dataSize,

            string databaseName,

            double documentCount,

            bool isCapped,

            bool isSystemCollection,

            bool isView,

            string name,

            string qualifiedName,

            Outputs.MongoDbShardKeyInfoResponse? shardKey,

            bool supportsSharding,

            string? viewOf)
        {
            AverageDocumentSize = averageDocumentSize;
            DataSize = dataSize;
            DatabaseName = databaseName;
            DocumentCount = documentCount;
            IsCapped = isCapped;
            IsSystemCollection = isSystemCollection;
            IsView = isView;
            Name = name;
            QualifiedName = qualifiedName;
            ShardKey = shardKey;
            SupportsSharding = supportsSharding;
            ViewOf = viewOf;
        }
    }
}
