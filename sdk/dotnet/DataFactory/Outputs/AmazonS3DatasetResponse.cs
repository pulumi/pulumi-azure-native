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
    /// A single Amazon Simple Storage Service (S3) object or a set of S3 objects.
    /// </summary>
    [OutputType]
    public sealed class AmazonS3DatasetResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the Dataset.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The name of the Amazon S3 bucket. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object BucketName;
        /// <summary>
        /// The data compression method used for the Amazon S3 object.
        /// </summary>
        public readonly Outputs.DatasetCompressionResponse? Compression;
        /// <summary>
        /// Dataset description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        /// </summary>
        public readonly Outputs.DatasetResponseFolder? Folder;
        /// <summary>
        /// The format of files.
        /// </summary>
        public readonly object? Format;
        /// <summary>
        /// The key of the Amazon S3 object. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Key;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse LinkedServiceName;
        /// <summary>
        /// The end of S3 object's modified datetime. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ModifiedDatetimeEnd;
        /// <summary>
        /// The start of S3 object's modified datetime. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ModifiedDatetimeStart;
        /// <summary>
        /// Parameters for dataset.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The prefix filter for the S3 object name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Prefix;
        /// <summary>
        /// Columns that define the physical type schema of the dataset. Type: array (or Expression with resultType array), itemType: DatasetSchemaDataElement.
        /// </summary>
        public readonly object? Schema;
        /// <summary>
        /// Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        /// </summary>
        public readonly object? Structure;
        /// <summary>
        /// Type of dataset.
        /// Expected value is 'AmazonS3Object'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The version for the S3 object. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Version;

        [OutputConstructor]
        private AmazonS3DatasetResponse(
            ImmutableArray<object> annotations,

            object bucketName,

            Outputs.DatasetCompressionResponse? compression,

            string? description,

            Outputs.DatasetResponseFolder? folder,

            object? format,

            object? key,

            Outputs.LinkedServiceReferenceResponse linkedServiceName,

            object? modifiedDatetimeEnd,

            object? modifiedDatetimeStart,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            object? prefix,

            object? schema,

            object? structure,

            string type,

            object? version)
        {
            Annotations = annotations;
            BucketName = bucketName;
            Compression = compression;
            Description = description;
            Folder = folder;
            Format = format;
            Key = key;
            LinkedServiceName = linkedServiceName;
            ModifiedDatetimeEnd = modifiedDatetimeEnd;
            ModifiedDatetimeStart = modifiedDatetimeStart;
            Parameters = parameters;
            Prefix = prefix;
            Schema = schema;
            Structure = structure;
            Type = type;
            Version = version;
        }
    }
}
