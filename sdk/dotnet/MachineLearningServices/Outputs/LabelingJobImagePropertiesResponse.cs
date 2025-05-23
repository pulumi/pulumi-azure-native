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
    /// Properties of a labeling job for image data
    /// </summary>
    [OutputType]
    public sealed class LabelingJobImagePropertiesResponse
    {
        /// <summary>
        /// Annotation type of image labeling job.
        /// </summary>
        public readonly string? AnnotationType;
        /// <summary>
        /// Media type of data asset.
        /// Expected value is 'Image'.
        /// </summary>
        public readonly string MediaType;

        [OutputConstructor]
        private LabelingJobImagePropertiesResponse(
            string? annotationType,

            string mediaType)
        {
            AnnotationType = annotationType;
            MediaType = mediaType;
        }
    }
}
