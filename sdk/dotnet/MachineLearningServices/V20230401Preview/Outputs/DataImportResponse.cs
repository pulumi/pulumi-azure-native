// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230401Preview.Outputs
{

    [OutputType]
    public sealed class DataImportResponse
    {
        /// <summary>
        /// Name of the asset for data import job to create
        /// </summary>
        public readonly string? AssetName;
        /// <summary>
        /// Specifies the lifecycle setting of managed data asset.
        /// </summary>
        public readonly Outputs.AutoDeleteSettingResponse? AutoDeleteSetting;
        /// <summary>
        /// Enum to determine the type of data.
        /// Expected value is 'uri_folder'.
        /// </summary>
        public readonly string DataType;
        /// <summary>
        /// [Required] Uri of the data. Example: https://go.microsoft.com/fwlink/?linkid=2202330
        /// </summary>
        public readonly string DataUri;
        /// <summary>
        /// The asset description text.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Intellectual Property details. Used if data is an Intellectual Property.
        /// </summary>
        public readonly Outputs.IntellectualPropertyResponse? IntellectualProperty;
        /// <summary>
        /// If the name version are system generated (anonymous registration). For types where Stage is defined, when Stage is provided it will be used to populate IsAnonymous
        /// </summary>
        public readonly bool? IsAnonymous;
        /// <summary>
        /// Is the asset archived? For types where Stage is defined, when Stage is provided it will be used to populate IsArchived
        /// </summary>
        public readonly bool? IsArchived;
        /// <summary>
        /// The asset property dictionary.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// Source data of the asset to import from
        /// </summary>
        public readonly Union<Outputs.DatabaseSourceResponse, Outputs.FileSystemSourceResponse>? Source;
        /// <summary>
        /// Stage in the data lifecycle assigned to this data asset
        /// </summary>
        public readonly string? Stage;
        /// <summary>
        /// Tag dictionary. Tags can be added, removed, and updated.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private DataImportResponse(
            string? assetName,

            Outputs.AutoDeleteSettingResponse? autoDeleteSetting,

            string dataType,

            string dataUri,

            string? description,

            Outputs.IntellectualPropertyResponse? intellectualProperty,

            bool? isAnonymous,

            bool? isArchived,

            ImmutableDictionary<string, string>? properties,

            Union<Outputs.DatabaseSourceResponse, Outputs.FileSystemSourceResponse>? source,

            string? stage,

            ImmutableDictionary<string, string>? tags)
        {
            AssetName = assetName;
            AutoDeleteSetting = autoDeleteSetting;
            DataType = dataType;
            DataUri = dataUri;
            Description = description;
            IntellectualProperty = intellectualProperty;
            IsAnonymous = isAnonymous;
            IsArchived = isArchived;
            Properties = properties;
            Source = source;
            Stage = stage;
            Tags = tags;
        }
    }
}
