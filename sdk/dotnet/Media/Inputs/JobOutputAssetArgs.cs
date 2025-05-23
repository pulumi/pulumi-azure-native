// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Inputs
{

    /// <summary>
    /// Represents an Asset used as a JobOutput.
    /// </summary>
    public sealed class JobOutputAssetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the output Asset.
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        /// <summary>
        /// A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.JobOutputAsset'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// A preset used to override the preset in the corresponding transform output.
        /// </summary>
        [Input("presetOverride")]
        public object? PresetOverride { get; set; }

        public JobOutputAssetArgs()
        {
        }
        public static new JobOutputAssetArgs Empty => new JobOutputAssetArgs();
    }
}
