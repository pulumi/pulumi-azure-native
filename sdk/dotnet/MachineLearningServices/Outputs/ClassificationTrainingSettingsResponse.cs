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
    /// Classification Training related configuration.
    /// </summary>
    [OutputType]
    public sealed class ClassificationTrainingSettingsResponse
    {
        /// <summary>
        /// Allowed models for classification task.
        /// </summary>
        public readonly ImmutableArray<string> AllowedTrainingAlgorithms;
        /// <summary>
        /// Blocked models for classification task.
        /// </summary>
        public readonly ImmutableArray<string> BlockedTrainingAlgorithms;
        /// <summary>
        /// Enable recommendation of DNN models.
        /// </summary>
        public readonly bool? EnableDnnTraining;
        /// <summary>
        /// Flag to turn on explainability on best model.
        /// </summary>
        public readonly bool? EnableModelExplainability;
        /// <summary>
        /// Flag for enabling onnx compatible models.
        /// </summary>
        public readonly bool? EnableOnnxCompatibleModels;
        /// <summary>
        /// Enable stack ensemble run.
        /// </summary>
        public readonly bool? EnableStackEnsemble;
        /// <summary>
        /// Enable voting ensemble run.
        /// </summary>
        public readonly bool? EnableVoteEnsemble;
        /// <summary>
        /// During VotingEnsemble and StackEnsemble model generation, multiple fitted models from the previous child runs are downloaded.
        /// Configure this parameter with a higher value than 300 secs, if more time is needed.
        /// </summary>
        public readonly string? EnsembleModelDownloadTimeout;
        /// <summary>
        /// Stack ensemble settings for stack ensemble run.
        /// </summary>
        public readonly Outputs.StackEnsembleSettingsResponse? StackEnsembleSettings;

        [OutputConstructor]
        private ClassificationTrainingSettingsResponse(
            ImmutableArray<string> allowedTrainingAlgorithms,

            ImmutableArray<string> blockedTrainingAlgorithms,

            bool? enableDnnTraining,

            bool? enableModelExplainability,

            bool? enableOnnxCompatibleModels,

            bool? enableStackEnsemble,

            bool? enableVoteEnsemble,

            string? ensembleModelDownloadTimeout,

            Outputs.StackEnsembleSettingsResponse? stackEnsembleSettings)
        {
            AllowedTrainingAlgorithms = allowedTrainingAlgorithms;
            BlockedTrainingAlgorithms = blockedTrainingAlgorithms;
            EnableDnnTraining = enableDnnTraining;
            EnableModelExplainability = enableModelExplainability;
            EnableOnnxCompatibleModels = enableOnnxCompatibleModels;
            EnableStackEnsemble = enableStackEnsemble;
            EnableVoteEnsemble = enableVoteEnsemble;
            EnsembleModelDownloadTimeout = ensembleModelDownloadTimeout;
            StackEnsembleSettings = stackEnsembleSettings;
        }
    }
}
