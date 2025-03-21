// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240101Preview.Inputs
{

    /// <summary>
    /// Text Classification task in AutoML NLP vertical.
    /// NLP - Natural Language Processing.
    /// </summary>
    public sealed class TextClassificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Featurization inputs needed for AutoML job.
        /// </summary>
        [Input("featurizationSettings")]
        public Input<Inputs.NlpVerticalFeaturizationSettingsArgs>? FeaturizationSettings { get; set; }

        /// <summary>
        /// Model/training parameters that will remain constant throughout training.
        /// </summary>
        [Input("fixedParameters")]
        public Input<Inputs.NlpFixedParametersArgs>? FixedParameters { get; set; }

        /// <summary>
        /// Execution constraints for AutoMLJob.
        /// </summary>
        [Input("limitSettings")]
        public Input<Inputs.NlpVerticalLimitSettingsArgs>? LimitSettings { get; set; }

        /// <summary>
        /// Log verbosity for the job.
        /// </summary>
        [Input("logVerbosity")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20240101Preview.LogVerbosity>? LogVerbosity { get; set; }

        /// <summary>
        /// Primary metric for Text-Classification task.
        /// </summary>
        [Input("primaryMetric")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20240101Preview.ClassificationPrimaryMetrics>? PrimaryMetric { get; set; }

        [Input("searchSpace")]
        private InputList<Inputs.NlpParameterSubspaceArgs>? _searchSpace;

        /// <summary>
        /// Search space for sampling different combinations of models and their hyperparameters.
        /// </summary>
        public InputList<Inputs.NlpParameterSubspaceArgs> SearchSpace
        {
            get => _searchSpace ?? (_searchSpace = new InputList<Inputs.NlpParameterSubspaceArgs>());
            set => _searchSpace = value;
        }

        /// <summary>
        /// Settings for model sweeping and hyperparameter tuning.
        /// </summary>
        [Input("sweepSettings")]
        public Input<Inputs.NlpSweepSettingsArgs>? SweepSettings { get; set; }

        /// <summary>
        /// Target column name: This is prediction values column.
        /// Also known as label column name in context of classification tasks.
        /// </summary>
        [Input("targetColumnName")]
        public Input<string>? TargetColumnName { get; set; }

        /// <summary>
        /// AutoMLJob Task type.
        /// Expected value is 'TextClassification'.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        /// <summary>
        /// [Required] Training data input.
        /// </summary>
        [Input("trainingData", required: true)]
        public Input<Inputs.MLTableJobInputArgs> TrainingData { get; set; } = null!;

        /// <summary>
        /// Validation data inputs.
        /// </summary>
        [Input("validationData")]
        public Input<Inputs.MLTableJobInputArgs>? ValidationData { get; set; }

        public TextClassificationArgs()
        {
            LogVerbosity = "Info";
            PrimaryMetric = "Accuracy";
        }
        public static new TextClassificationArgs Empty => new TextClassificationArgs();
    }
}
