// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.Inputs
{

    /// <summary>
    /// Describes a DocumentDB output data source.
    /// </summary>
    public sealed class DocumentDbOutputDataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DocumentDB account name or ID. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The account key for the DocumentDB account. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("accountKey")]
        public Input<string>? AccountKey { get; set; }

        /// <summary>
        /// The collection name pattern for the collections to be used. The collection name format can be constructed using the optional {partition} token, where partitions start from 0. See the DocumentDB section of https://docs.microsoft.com/en-us/rest/api/streamanalytics/stream-analytics-output for more information. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("collectionNamePattern")]
        public Input<string>? CollectionNamePattern { get; set; }

        /// <summary>
        /// The name of the DocumentDB database. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The name of the field in output events used to specify the primary key which insert or update operations are based on.
        /// </summary>
        [Input("documentId")]
        public Input<string>? DocumentId { get; set; }

        /// <summary>
        /// The name of the field in output events used to specify the key for partitioning output across collections. If 'collectionNamePattern' contains the {partition} token, this property is required to be specified.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        /// <summary>
        /// Indicates the type of data source output will be written to. Required on PUT (CreateOrReplace) requests.
        /// Expected value is 'Microsoft.Storage/DocumentDB'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DocumentDbOutputDataSourceArgs()
        {
        }
        public static new DocumentDbOutputDataSourceArgs Empty => new DocumentDbOutputDataSourceArgs();
    }
}
