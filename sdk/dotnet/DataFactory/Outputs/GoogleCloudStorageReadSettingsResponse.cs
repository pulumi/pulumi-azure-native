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
    /// Google Cloud Storage read settings.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudStorageReadSettingsResponse
    {
        /// <summary>
        /// Indicates whether the source files need to be deleted after copy completion. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DeleteFilesAfterCompletion;
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DisableMetricsCollection;
        /// <summary>
        /// Indicates whether to enable partition discovery. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? EnablePartitionDiscovery;
        /// <summary>
        /// Point to a text file that lists each file (relative path to the path configured in the dataset) that you want to copy. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? FileListPath;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// The end of file's modified datetime. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ModifiedDatetimeEnd;
        /// <summary>
        /// The start of file's modified datetime. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ModifiedDatetimeStart;
        /// <summary>
        /// Specify the root path where partition discovery starts from. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PartitionRootPath;
        /// <summary>
        /// The prefix filter for the Google Cloud Storage object name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Prefix;
        /// <summary>
        /// If true, files under the folder path will be read recursively. Default is true. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? Recursive;
        /// <summary>
        /// The read setting type.
        /// Expected value is 'GoogleCloudStorageReadSettings'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Google Cloud Storage wildcardFileName. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? WildcardFileName;
        /// <summary>
        /// Google Cloud Storage wildcardFolderPath. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? WildcardFolderPath;

        [OutputConstructor]
        private GoogleCloudStorageReadSettingsResponse(
            object? deleteFilesAfterCompletion,

            object? disableMetricsCollection,

            object? enablePartitionDiscovery,

            object? fileListPath,

            object? maxConcurrentConnections,

            object? modifiedDatetimeEnd,

            object? modifiedDatetimeStart,

            object? partitionRootPath,

            object? prefix,

            object? recursive,

            string type,

            object? wildcardFileName,

            object? wildcardFolderPath)
        {
            DeleteFilesAfterCompletion = deleteFilesAfterCompletion;
            DisableMetricsCollection = disableMetricsCollection;
            EnablePartitionDiscovery = enablePartitionDiscovery;
            FileListPath = fileListPath;
            MaxConcurrentConnections = maxConcurrentConnections;
            ModifiedDatetimeEnd = modifiedDatetimeEnd;
            ModifiedDatetimeStart = modifiedDatetimeStart;
            PartitionRootPath = partitionRootPath;
            Prefix = prefix;
            Recursive = recursive;
            Type = type;
            WildcardFileName = wildcardFileName;
            WildcardFolderPath = wildcardFolderPath;
        }
    }
}
