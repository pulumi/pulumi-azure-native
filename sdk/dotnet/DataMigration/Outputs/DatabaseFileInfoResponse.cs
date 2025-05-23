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
    /// Database file specific information
    /// </summary>
    [OutputType]
    public sealed class DatabaseFileInfoResponse
    {
        /// <summary>
        /// Name of the database
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// Database file type
        /// </summary>
        public readonly string? FileType;
        /// <summary>
        /// Unique identifier for database file
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Logical name of the file
        /// </summary>
        public readonly string? LogicalName;
        /// <summary>
        /// Operating-system full path of the file
        /// </summary>
        public readonly string? PhysicalFullName;
        /// <summary>
        /// Suggested full path of the file for restoring
        /// </summary>
        public readonly string? RestoreFullName;
        /// <summary>
        /// Size of the file in megabytes
        /// </summary>
        public readonly double? SizeMB;

        [OutputConstructor]
        private DatabaseFileInfoResponse(
            string? databaseName,

            string? fileType,

            string? id,

            string? logicalName,

            string? physicalFullName,

            string? restoreFullName,

            double? sizeMB)
        {
            DatabaseName = databaseName;
            FileType = fileType;
            Id = id;
            LogicalName = logicalName;
            PhysicalFullName = physicalFullName;
            RestoreFullName = restoreFullName;
            SizeMB = sizeMB;
        }
    }
}
