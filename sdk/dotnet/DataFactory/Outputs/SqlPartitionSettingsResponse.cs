// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// The settings that will be leveraged for Sql source partitioning.
    /// </summary>
    [OutputType]
    public sealed class SqlPartitionSettingsResponse
    {
        /// <summary>
        /// The name of the column in integer or datetime type that will be used for proceeding partitioning. If not specified, the primary key of the table is auto-detected and used as the partition column. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionColumnName;
        /// <summary>
        /// The minimum value of the partition column for partition range splitting. This value is used to decide the partition stride, not for filtering the rows in table. All rows in the table or query result will be partitioned and copied. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionLowerBound;
        /// <summary>
        /// The maximum value of the partition column for partition range splitting. This value is used to decide the partition stride, not for filtering the rows in table. All rows in the table or query result will be partitioned and copied. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionUpperBound;

        [OutputConstructor]
        private SqlPartitionSettingsResponse(
            object? partitionColumnName,

            object? partitionLowerBound,

            object? partitionUpperBound)
        {
            PartitionColumnName = partitionColumnName;
            PartitionLowerBound = partitionLowerBound;
            PartitionUpperBound = partitionUpperBound;
        }
    }
}
