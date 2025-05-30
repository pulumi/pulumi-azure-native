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
    /// Delimited text read settings.
    /// </summary>
    [OutputType]
    public sealed class DelimitedTextReadSettingsResponse
    {
        /// <summary>
        /// Compression settings.
        /// </summary>
        public readonly object? CompressionProperties;
        /// <summary>
        /// Indicates the number of non-empty rows to skip when reading data from input files. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? SkipLineCount;
        /// <summary>
        /// The read setting type.
        /// Expected value is 'DelimitedTextReadSettings'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DelimitedTextReadSettingsResponse(
            object? compressionProperties,

            object? skipLineCount,

            string type)
        {
            CompressionProperties = compressionProperties;
            SkipLineCount = skipLineCount;
            Type = type;
        }
    }
}
