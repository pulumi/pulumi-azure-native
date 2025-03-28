// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240701Preview.Outputs
{

    [OutputType]
    public sealed class AzureOpenAiFineTuningResponse
    {
        /// <summary>
        /// HyperParameters for fine tuning Azure Open AI model.
        /// </summary>
        public readonly Outputs.AzureOpenAiHyperParametersResponse? HyperParameters;
        /// <summary>
        /// [Required] Input model for fine tuning.
        /// </summary>
        public readonly Outputs.MLFlowModelJobInputResponse Model;
        /// <summary>
        /// Enum to determine the type of fine tuning.
        /// Expected value is 'AzureOpenAI'.
        /// </summary>
        public readonly string ModelProvider;
        /// <summary>
        /// [Required] Fine tuning task type.
        /// </summary>
        public readonly string TaskType;
        /// <summary>
        /// [Required] Training data for fine tuning.
        /// </summary>
        public readonly object TrainingData;
        /// <summary>
        /// Validation data for fine tuning.
        /// </summary>
        public readonly object? ValidationData;

        [OutputConstructor]
        private AzureOpenAiFineTuningResponse(
            Outputs.AzureOpenAiHyperParametersResponse? hyperParameters,

            Outputs.MLFlowModelJobInputResponse model,

            string modelProvider,

            string taskType,

            object trainingData,

            object? validationData)
        {
            HyperParameters = hyperParameters;
            Model = model;
            ModelProvider = modelProvider;
            TaskType = taskType;
            TrainingData = trainingData;
            ValidationData = validationData;
        }
    }
}
