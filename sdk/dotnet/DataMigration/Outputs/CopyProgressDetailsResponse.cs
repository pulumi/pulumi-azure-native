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
    /// Details on progress of ADF copy activity
    /// </summary>
    [OutputType]
    public sealed class CopyProgressDetailsResponse
    {
        /// <summary>
        /// Copy Duration in seconds
        /// </summary>
        public readonly int CopyDuration;
        /// <summary>
        /// Copy Start
        /// </summary>
        public readonly string CopyStart;
        /// <summary>
        /// Copy throughput in KBps
        /// </summary>
        public readonly double CopyThroughput;
        /// <summary>
        /// Bytes read
        /// </summary>
        public readonly double DataRead;
        /// <summary>
        /// Bytes written
        /// </summary>
        public readonly double DataWritten;
        /// <summary>
        /// Type of parallel copy (Dynamic range, Physical partition, none).
        /// </summary>
        public readonly string ParallelCopyType;
        /// <summary>
        /// Rows Copied
        /// </summary>
        public readonly double RowsCopied;
        /// <summary>
        /// Rows read
        /// </summary>
        public readonly double RowsRead;
        /// <summary>
        /// Status of the Copy activity (InProgress, Succeeded, Failed, Canceled).
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Table Name
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// The degree of parallelization.
        /// </summary>
        public readonly int UsedParallelCopies;

        [OutputConstructor]
        private CopyProgressDetailsResponse(
            int copyDuration,

            string copyStart,

            double copyThroughput,

            double dataRead,

            double dataWritten,

            string parallelCopyType,

            double rowsCopied,

            double rowsRead,

            string status,

            string tableName,

            int usedParallelCopies)
        {
            CopyDuration = copyDuration;
            CopyStart = copyStart;
            CopyThroughput = copyThroughput;
            DataRead = dataRead;
            DataWritten = dataWritten;
            ParallelCopyType = parallelCopyType;
            RowsCopied = rowsCopied;
            RowsRead = rowsRead;
            Status = status;
            TableName = tableName;
            UsedParallelCopies = usedParallelCopies;
        }
    }
}
