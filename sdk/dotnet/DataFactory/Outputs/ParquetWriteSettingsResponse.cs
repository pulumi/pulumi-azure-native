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
    /// Parquet write settings.
    /// </summary>
    [OutputType]
    public sealed class ParquetWriteSettingsResponse
    {
        /// <summary>
        /// Specifies the file name pattern &lt;fileNamePrefix&gt;_&lt;fileIndex&gt;.&lt;fileExtension&gt; when copy from non-file based store without partitionOptions. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? FileNamePrefix;
        /// <summary>
        /// Limit the written file's row count to be smaller than or equal to the specified count. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxRowsPerFile;
        /// <summary>
        /// The write setting type.
        /// Expected value is 'ParquetWriteSettings'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ParquetWriteSettingsResponse(
            object? fileNamePrefix,

            object? maxRowsPerFile,

            string type)
        {
            FileNamePrefix = fileNamePrefix;
            MaxRowsPerFile = maxRowsPerFile;
            Type = type;
        }
    }
}
