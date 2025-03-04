// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20220201Preview.Inputs
{

    /// <summary>
    /// Text Classification Multilabel task in AutoML NLP vertical.
    /// NLP - Natural Language Processing.
    /// </summary>
    public sealed class TextClassificationMultilabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data inputs for AutoMLJob.
        /// </summary>
        [Input("dataSettings")]
        public Input<Inputs.NlpVerticalDataSettingsArgs>? DataSettings { get; set; }

        /// <summary>
        /// Featurization inputs needed for AutoML job.
        /// </summary>
        [Input("featurizationSettings")]
        public Input<Inputs.NlpVerticalFeaturizationSettingsArgs>? FeaturizationSettings { get; set; }

        /// <summary>
        /// Execution constraints for AutoMLJob.
        /// </summary>
        [Input("limitSettings")]
        public Input<Inputs.NlpVerticalLimitSettingsArgs>? LimitSettings { get; set; }

        /// <summary>
        /// Log verbosity for the job.
        /// </summary>
        [Input("logVerbosity")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20220201Preview.LogVerbosity>? LogVerbosity { get; set; }

        /// <summary>
        /// AutoMLJob Task type.
        /// Expected value is 'TextClassificationMultilabel'.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public TextClassificationMultilabelArgs()
        {
            LogVerbosity = "Info";
        }
        public static new TextClassificationMultilabelArgs Empty => new TextClassificationMultilabelArgs();
    }
}
