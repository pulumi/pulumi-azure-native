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
    /// Distribution expressions to sweep over values of model settings.
    /// &lt;example&gt;
    /// Some examples are:
    /// ```
    /// ModelName = "choice('seresnext', 'resnest50')";
    /// LearningRate = "uniform(0.001, 0.01)";
    /// LayersToFreeze = "choice(0, 2)";
    /// ```&lt;/example&gt;
    /// For more details on how to compose distribution expressions please check the documentation:
    /// https://docs.microsoft.com/en-us/azure/machine-learning/how-to-tune-hyperparameters
    /// For more information on the available settings please visit the official documentation:
    /// https://docs.microsoft.com/en-us/azure/machine-learning/how-to-auto-train-image-models.
    /// </summary>
    [OutputType]
    public sealed class ImageModelDistributionSettingsObjectDetectionResponse
    {
        /// <summary>
        /// Enable AMSGrad when optimizer is 'adam' or 'adamw'.
        /// </summary>
        public readonly string? AmsGradient;
        /// <summary>
        /// Settings for using Augmentations.
        /// </summary>
        public readonly string? Augmentations;
        /// <summary>
        /// Value of 'beta1' when optimizer is 'adam' or 'adamw'. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? Beta1;
        /// <summary>
        /// Value of 'beta2' when optimizer is 'adam' or 'adamw'. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? Beta2;
        /// <summary>
        /// Maximum number of detections per image, for all classes. Must be a positive integer.
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? BoxDetectionsPerImage;
        /// <summary>
        /// During inference, only return proposals with a classification score greater than
        /// BoxScoreThreshold. Must be a float in the range[0, 1].
        /// </summary>
        public readonly string? BoxScoreThreshold;
        /// <summary>
        /// Whether to use distributer training.
        /// </summary>
        public readonly string? Distributed;
        /// <summary>
        /// Enable early stopping logic during training.
        /// </summary>
        public readonly string? EarlyStopping;
        /// <summary>
        /// Minimum number of epochs or validation evaluations to wait before primary metric improvement
        /// is tracked for early stopping. Must be a positive integer.
        /// </summary>
        public readonly string? EarlyStoppingDelay;
        /// <summary>
        /// Minimum number of epochs or validation evaluations with no primary metric improvement before
        /// the run is stopped. Must be a positive integer.
        /// </summary>
        public readonly string? EarlyStoppingPatience;
        /// <summary>
        /// Enable normalization when exporting ONNX model.
        /// </summary>
        public readonly string? EnableOnnxNormalization;
        /// <summary>
        /// Frequency to evaluate validation dataset to get metric scores. Must be a positive integer.
        /// </summary>
        public readonly string? EvaluationFrequency;
        /// <summary>
        /// Gradient accumulation means running a configured number of "GradAccumulationStep" steps without
        /// updating the model weights while accumulating the gradients of those steps, and then using
        /// the accumulated gradients to compute the weight updates. Must be a positive integer.
        /// </summary>
        public readonly string? GradientAccumulationStep;
        /// <summary>
        /// Image size for train and validation. Must be a positive integer.
        /// Note: The training run may get into CUDA OOM if the size is too big.
        /// Note: This settings is only supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? ImageSize;
        /// <summary>
        /// Number of layers to freeze for the model. Must be a positive integer.
        /// For instance, passing 2 as value for 'seresnext' means
        /// freezing layer0 and layer1. For a full list of models supported and details on layer freeze, please
        /// see: https://docs.microsoft.com/en-us/azure/machine-learning/how-to-auto-train-image-models.
        /// </summary>
        public readonly string? LayersToFreeze;
        /// <summary>
        /// Initial learning rate. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? LearningRate;
        /// <summary>
        /// Type of learning rate scheduler. Must be 'warmup_cosine' or 'step'.
        /// </summary>
        public readonly string? LearningRateScheduler;
        /// <summary>
        /// Maximum size of the image to be rescaled before feeding it to the backbone.
        /// Must be a positive integer. Note: training run may get into CUDA OOM if the size is too big.
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? MaxSize;
        /// <summary>
        /// Minimum size of the image to be rescaled before feeding it to the backbone.
        /// Must be a positive integer. Note: training run may get into CUDA OOM if the size is too big.
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? MinSize;
        /// <summary>
        /// Name of the model to use for training.
        /// For more information on the available models please visit the official documentation:
        /// https://docs.microsoft.com/en-us/azure/machine-learning/how-to-auto-train-image-models.
        /// </summary>
        public readonly string? ModelName;
        /// <summary>
        /// Model size. Must be 'small', 'medium', 'large', or 'xlarge'.
        /// Note: training run may get into CUDA OOM if the model size is too big.
        /// Note: This settings is only supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? ModelSize;
        /// <summary>
        /// Value of momentum when optimizer is 'sgd'. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? Momentum;
        /// <summary>
        /// Enable multi-scale image by varying image size by +/- 50%.
        /// Note: training run may get into CUDA OOM if no sufficient GPU memory.
        /// Note: This settings is only supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? MultiScale;
        /// <summary>
        /// Enable nesterov when optimizer is 'sgd'.
        /// </summary>
        public readonly string? Nesterov;
        /// <summary>
        /// IOU threshold used during inference in NMS post processing. Must be float in the range [0, 1].
        /// </summary>
        public readonly string? NmsIouThreshold;
        /// <summary>
        /// Number of training epochs. Must be a positive integer.
        /// </summary>
        public readonly string? NumberOfEpochs;
        /// <summary>
        /// Number of data loader workers. Must be a non-negative integer.
        /// </summary>
        public readonly string? NumberOfWorkers;
        /// <summary>
        /// Type of optimizer. Must be either 'sgd', 'adam', or 'adamw'.
        /// </summary>
        public readonly string? Optimizer;
        /// <summary>
        /// Random seed to be used when using deterministic training.
        /// </summary>
        public readonly string? RandomSeed;
        /// <summary>
        /// Value of gamma when learning rate scheduler is 'step'. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? StepLRGamma;
        /// <summary>
        /// Value of step size when learning rate scheduler is 'step'. Must be a positive integer.
        /// </summary>
        public readonly string? StepLRStepSize;
        /// <summary>
        /// The grid size to use for tiling each image. Note: TileGridSize must not be
        /// None to enable small object detection logic. A string containing two integers in mxn format.
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? TileGridSize;
        /// <summary>
        /// Overlap ratio between adjacent tiles in each dimension. Must be float in the range [0, 1).
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// </summary>
        public readonly string? TileOverlapRatio;
        /// <summary>
        /// The IOU threshold to use to perform NMS while merging predictions from tiles and image.
        /// Used in validation/ inference. Must be float in the range [0, 1].
        /// Note: This settings is not supported for the 'yolov5' algorithm.
        /// NMS: Non-maximum suppression
        /// </summary>
        public readonly string? TilePredictionsNmsThreshold;
        /// <summary>
        /// Training batch size. Must be a positive integer.
        /// </summary>
        public readonly string? TrainingBatchSize;
        /// <summary>
        /// Validation batch size. Must be a positive integer.
        /// </summary>
        public readonly string? ValidationBatchSize;
        /// <summary>
        /// IOU threshold to use when computing validation metric. Must be float in the range [0, 1].
        /// </summary>
        public readonly string? ValidationIouThreshold;
        /// <summary>
        /// Metric computation method to use for validation metrics. Must be 'none', 'coco', 'voc', or 'coco_voc'.
        /// </summary>
        public readonly string? ValidationMetricType;
        /// <summary>
        /// Value of cosine cycle when learning rate scheduler is 'warmup_cosine'. Must be a float in the range [0, 1].
        /// </summary>
        public readonly string? WarmupCosineLRCycles;
        /// <summary>
        /// Value of warmup epochs when learning rate scheduler is 'warmup_cosine'. Must be a positive integer.
        /// </summary>
        public readonly string? WarmupCosineLRWarmupEpochs;
        /// <summary>
        /// Value of weight decay when optimizer is 'sgd', 'adam', or 'adamw'. Must be a float in the range[0, 1].
        /// </summary>
        public readonly string? WeightDecay;

        [OutputConstructor]
        private ImageModelDistributionSettingsObjectDetectionResponse(
            string? amsGradient,

            string? augmentations,

            string? beta1,

            string? beta2,

            string? boxDetectionsPerImage,

            string? boxScoreThreshold,

            string? distributed,

            string? earlyStopping,

            string? earlyStoppingDelay,

            string? earlyStoppingPatience,

            string? enableOnnxNormalization,

            string? evaluationFrequency,

            string? gradientAccumulationStep,

            string? imageSize,

            string? layersToFreeze,

            string? learningRate,

            string? learningRateScheduler,

            string? maxSize,

            string? minSize,

            string? modelName,

            string? modelSize,

            string? momentum,

            string? multiScale,

            string? nesterov,

            string? nmsIouThreshold,

            string? numberOfEpochs,

            string? numberOfWorkers,

            string? optimizer,

            string? randomSeed,

            string? stepLRGamma,

            string? stepLRStepSize,

            string? tileGridSize,

            string? tileOverlapRatio,

            string? tilePredictionsNmsThreshold,

            string? trainingBatchSize,

            string? validationBatchSize,

            string? validationIouThreshold,

            string? validationMetricType,

            string? warmupCosineLRCycles,

            string? warmupCosineLRWarmupEpochs,

            string? weightDecay)
        {
            AmsGradient = amsGradient;
            Augmentations = augmentations;
            Beta1 = beta1;
            Beta2 = beta2;
            BoxDetectionsPerImage = boxDetectionsPerImage;
            BoxScoreThreshold = boxScoreThreshold;
            Distributed = distributed;
            EarlyStopping = earlyStopping;
            EarlyStoppingDelay = earlyStoppingDelay;
            EarlyStoppingPatience = earlyStoppingPatience;
            EnableOnnxNormalization = enableOnnxNormalization;
            EvaluationFrequency = evaluationFrequency;
            GradientAccumulationStep = gradientAccumulationStep;
            ImageSize = imageSize;
            LayersToFreeze = layersToFreeze;
            LearningRate = learningRate;
            LearningRateScheduler = learningRateScheduler;
            MaxSize = maxSize;
            MinSize = minSize;
            ModelName = modelName;
            ModelSize = modelSize;
            Momentum = momentum;
            MultiScale = multiScale;
            Nesterov = nesterov;
            NmsIouThreshold = nmsIouThreshold;
            NumberOfEpochs = numberOfEpochs;
            NumberOfWorkers = numberOfWorkers;
            Optimizer = optimizer;
            RandomSeed = randomSeed;
            StepLRGamma = stepLRGamma;
            StepLRStepSize = stepLRStepSize;
            TileGridSize = tileGridSize;
            TileOverlapRatio = tileOverlapRatio;
            TilePredictionsNmsThreshold = tilePredictionsNmsThreshold;
            TrainingBatchSize = trainingBatchSize;
            ValidationBatchSize = validationBatchSize;
            ValidationIouThreshold = validationIouThreshold;
            ValidationMetricType = validationMetricType;
            WarmupCosineLRCycles = warmupCosineLRCycles;
            WarmupCosineLRWarmupEpochs = warmupCosineLRWarmupEpochs;
            WeightDecay = weightDecay;
        }
    }
}
