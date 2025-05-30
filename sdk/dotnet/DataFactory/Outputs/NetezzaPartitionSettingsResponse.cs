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
    /// The settings that will be leveraged for Netezza source partitioning.
    /// </summary>
    [OutputType]
    public sealed class NetezzaPartitionSettingsResponse
    {
        /// <summary>
        /// The name of the column in integer type that will be used for proceeding range partitioning. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionColumnName;
        /// <summary>
        /// The minimum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionLowerBound;
        /// <summary>
        /// The maximum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionUpperBound;

        [OutputConstructor]
        private NetezzaPartitionSettingsResponse(
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
