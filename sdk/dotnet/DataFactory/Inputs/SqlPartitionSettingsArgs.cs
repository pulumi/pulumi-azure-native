// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// The settings that will be leveraged for Sql source partitioning.
    /// </summary>
    public sealed class SqlPartitionSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the column in integer or datetime type that will be used for proceeding partitioning. If not specified, the primary key of the table is auto-detected and used as the partition column. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("partitionColumnName")]
        public Input<object>? PartitionColumnName { get; set; }

        /// <summary>
        /// The minimum value of the partition column for partition range splitting. This value is used to decide the partition stride, not for filtering the rows in table. All rows in the table or query result will be partitioned and copied. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("partitionLowerBound")]
        public Input<object>? PartitionLowerBound { get; set; }

        /// <summary>
        /// The maximum value of the partition column for partition range splitting. This value is used to decide the partition stride, not for filtering the rows in table. All rows in the table or query result will be partitioned and copied. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("partitionUpperBound")]
        public Input<object>? PartitionUpperBound { get; set; }

        public SqlPartitionSettingsArgs()
        {
        }
        public static new SqlPartitionSettingsArgs Empty => new SqlPartitionSettingsArgs();
    }
}
