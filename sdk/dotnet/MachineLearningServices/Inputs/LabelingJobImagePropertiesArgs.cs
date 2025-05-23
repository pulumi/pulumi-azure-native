// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Properties of a labeling job for image data
    /// </summary>
    public sealed class LabelingJobImagePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Annotation type of image labeling job.
        /// </summary>
        [Input("annotationType")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.ImageAnnotationType>? AnnotationType { get; set; }

        /// <summary>
        /// Media type of data asset.
        /// Expected value is 'Image'.
        /// </summary>
        [Input("mediaType", required: true)]
        public Input<string> MediaType { get; set; } = null!;

        public LabelingJobImagePropertiesArgs()
        {
            AnnotationType = "Classification";
        }
        public static new LabelingJobImagePropertiesArgs Empty => new LabelingJobImagePropertiesArgs();
    }
}
