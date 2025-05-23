// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Machine Learning dataset object.
    /// </summary>
    [OutputType]
    public sealed class DatasetResponse
    {
        /// <summary>
        /// The dataset creation time (UTC).
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// Unique Dataset identifier.
        /// </summary>
        public readonly string DatasetId;
        /// <summary>
        /// Dataset state
        /// </summary>
        public readonly Outputs.DatasetStateResponse? DatasetState;
        /// <summary>
        /// Dataset Type.
        /// </summary>
        public readonly string DatasetType;
        /// <summary>
        /// Name of the default compute to be used for any Dataset actions (such as Profile, Write).
        /// </summary>
        public readonly string DefaultCompute;
        /// <summary>
        /// Description about this dataset version.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// eTag description
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Flag to hide Dataset in UI
        /// </summary>
        public readonly bool IsVisible;
        /// <summary>
        /// Last created Dataset definition.
        /// </summary>
        public readonly Outputs.DatasetResponseLatest? Latest;
        /// <summary>
        /// The dataset last modified time (UTC).
        /// </summary>
        public readonly string ModifiedTime;
        /// <summary>
        /// Unique dataset name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Tags for this dataset version.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private DatasetResponse(
            string createdTime,

            string datasetId,

            Outputs.DatasetStateResponse? datasetState,

            string datasetType,

            string defaultCompute,

            string description,

            string etag,

            bool isVisible,

            Outputs.DatasetResponseLatest? latest,

            string modifiedTime,

            string name,

            ImmutableDictionary<string, string> tags)
        {
            CreatedTime = createdTime;
            DatasetId = datasetId;
            DatasetState = datasetState;
            DatasetType = datasetType;
            DefaultCompute = defaultCompute;
            Description = description;
            Etag = etag;
            IsVisible = isVisible;
            Latest = latest;
            ModifiedTime = modifiedTime;
            Name = name;
            Tags = tags;
        }
    }
}
