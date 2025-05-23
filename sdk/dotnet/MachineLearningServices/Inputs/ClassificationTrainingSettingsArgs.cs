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
    /// Classification Training related configuration.
    /// </summary>
    public sealed class ClassificationTrainingSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedTrainingAlgorithms")]
        private InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>>? _allowedTrainingAlgorithms;

        /// <summary>
        /// Allowed models for classification task.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>> AllowedTrainingAlgorithms
        {
            get => _allowedTrainingAlgorithms ?? (_allowedTrainingAlgorithms = new InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>>());
            set => _allowedTrainingAlgorithms = value;
        }

        [Input("blockedTrainingAlgorithms")]
        private InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>>? _blockedTrainingAlgorithms;

        /// <summary>
        /// Blocked models for classification task.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>> BlockedTrainingAlgorithms
        {
            get => _blockedTrainingAlgorithms ?? (_blockedTrainingAlgorithms = new InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.ClassificationModels>>());
            set => _blockedTrainingAlgorithms = value;
        }

        /// <summary>
        /// Enable recommendation of DNN models.
        /// </summary>
        [Input("enableDnnTraining")]
        public Input<bool>? EnableDnnTraining { get; set; }

        /// <summary>
        /// Flag to turn on explainability on best model.
        /// </summary>
        [Input("enableModelExplainability")]
        public Input<bool>? EnableModelExplainability { get; set; }

        /// <summary>
        /// Flag for enabling onnx compatible models.
        /// </summary>
        [Input("enableOnnxCompatibleModels")]
        public Input<bool>? EnableOnnxCompatibleModels { get; set; }

        /// <summary>
        /// Enable stack ensemble run.
        /// </summary>
        [Input("enableStackEnsemble")]
        public Input<bool>? EnableStackEnsemble { get; set; }

        /// <summary>
        /// Enable voting ensemble run.
        /// </summary>
        [Input("enableVoteEnsemble")]
        public Input<bool>? EnableVoteEnsemble { get; set; }

        /// <summary>
        /// During VotingEnsemble and StackEnsemble model generation, multiple fitted models from the previous child runs are downloaded.
        /// Configure this parameter with a higher value than 300 secs, if more time is needed.
        /// </summary>
        [Input("ensembleModelDownloadTimeout")]
        public Input<string>? EnsembleModelDownloadTimeout { get; set; }

        /// <summary>
        /// Stack ensemble settings for stack ensemble run.
        /// </summary>
        [Input("stackEnsembleSettings")]
        public Input<Inputs.StackEnsembleSettingsArgs>? StackEnsembleSettings { get; set; }

        public ClassificationTrainingSettingsArgs()
        {
            EnableDnnTraining = false;
            EnableModelExplainability = true;
            EnableOnnxCompatibleModels = false;
            EnableStackEnsemble = true;
            EnableVoteEnsemble = true;
            EnsembleModelDownloadTimeout = "PT5M";
        }
        public static new ClassificationTrainingSettingsArgs Empty => new ClassificationTrainingSettingsArgs();
    }
}
