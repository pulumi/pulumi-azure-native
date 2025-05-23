// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of GlobalSecondaryIndex
    /// </summary>
    public sealed class GlobalSecondaryIndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings used to enable or disable CloudWatch Contributor Insights for the specified global secondary index. The settings used to enable or disable CloudWatch Contributor Insights.
        /// </summary>
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.ContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        /// <summary>
        /// The name of the global secondary index. The name must be unique among all other indexes on this table.
        /// </summary>
        [Input("indexName")]
        public Input<string>? IndexName { get; set; }

        [Input("keySchema")]
        private InputList<Inputs.KeySchemaArgs>? _keySchema;

        /// <summary>
        /// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  +   ``HASH`` - partition key  +   ``RANGE`` - sort key    The partition key of an item is also known as its *hash attribute*. The term 'hash attribute' derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values. The sort key of an item is also known as its *range attribute*. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
        /// </summary>
        public InputList<Inputs.KeySchemaArgs> KeySchema
        {
            get => _keySchema ?? (_keySchema = new InputList<Inputs.KeySchemaArgs>());
            set => _keySchema = value;
        }

        /// <summary>
        /// Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
        /// </summary>
        [Input("projection")]
        public Input<Inputs.ProjectionArgs>? Projection { get; set; }

        /// <summary>
        /// Represents the provisioned throughput settings for the specified global secondary index. For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide*. Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html).
        /// </summary>
        [Input("provisionedThroughput")]
        public Input<Inputs.ProvisionedThroughputArgs>? ProvisionedThroughput { get; set; }

        public GlobalSecondaryIndexArgs()
        {
        }
        public static new GlobalSecondaryIndexArgs Empty => new GlobalSecondaryIndexArgs();
    }
}
