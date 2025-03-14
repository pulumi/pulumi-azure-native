// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230601Preview.Outputs
{

    [OutputType]
    public sealed class TableParameterSubspaceResponse
    {
        /// <summary>
        /// Specify the boosting type, e.g gbdt for XGBoost.
        /// </summary>
        public readonly string? Booster;
        /// <summary>
        /// Specify the boosting type, e.g gbdt for LightGBM.
        /// </summary>
        public readonly string? BoostingType;
        /// <summary>
        /// Specify the grow policy, which controls the way new nodes are added to the tree.
        /// </summary>
        public readonly string? GrowPolicy;
        /// <summary>
        /// The learning rate for the training procedure.
        /// </summary>
        public readonly string? LearningRate;
        /// <summary>
        /// Specify the Maximum number of discrete bins to bucket continuous features .
        /// </summary>
        public readonly string? MaxBin;
        /// <summary>
        /// Specify the max depth to limit the tree depth explicitly.
        /// </summary>
        public readonly string? MaxDepth;
        /// <summary>
        /// Specify the max leaves to limit the tree leaves explicitly.
        /// </summary>
        public readonly string? MaxLeaves;
        /// <summary>
        /// The minimum number of data per leaf.
        /// </summary>
        public readonly string? MinDataInLeaf;
        /// <summary>
        /// Minimum loss reduction required to make a further partition on a leaf node of the tree.
        /// </summary>
        public readonly string? MinSplitGain;
        /// <summary>
        /// The name of the model to train.
        /// </summary>
        public readonly string? ModelName;
        /// <summary>
        /// Specify the number of trees (or rounds) in an model.
        /// </summary>
        public readonly string? NEstimators;
        /// <summary>
        /// Specify the number of leaves.
        /// </summary>
        public readonly string? NumLeaves;
        /// <summary>
        /// The name of the preprocessor to use.
        /// </summary>
        public readonly string? PreprocessorName;
        /// <summary>
        /// L1 regularization term on weights.
        /// </summary>
        public readonly string? RegAlpha;
        /// <summary>
        /// L2 regularization term on weights.
        /// </summary>
        public readonly string? RegLambda;
        /// <summary>
        /// Subsample ratio of the training instance.
        /// </summary>
        public readonly string? Subsample;
        /// <summary>
        /// Frequency of subsample
        /// </summary>
        public readonly string? SubsampleFreq;
        /// <summary>
        /// Specify the tree method.
        /// </summary>
        public readonly string? TreeMethod;
        /// <summary>
        /// If true, center before scaling the data with StandardScalar.
        /// </summary>
        public readonly string? WithMean;
        /// <summary>
        /// If true, scaling the data with Unit Variance with StandardScalar.
        /// </summary>
        public readonly string? WithStd;

        [OutputConstructor]
        private TableParameterSubspaceResponse(
            string? booster,

            string? boostingType,

            string? growPolicy,

            string? learningRate,

            string? maxBin,

            string? maxDepth,

            string? maxLeaves,

            string? minDataInLeaf,

            string? minSplitGain,

            string? modelName,

            string? nEstimators,

            string? numLeaves,

            string? preprocessorName,

            string? regAlpha,

            string? regLambda,

            string? subsample,

            string? subsampleFreq,

            string? treeMethod,

            string? withMean,

            string? withStd)
        {
            Booster = booster;
            BoostingType = boostingType;
            GrowPolicy = growPolicy;
            LearningRate = learningRate;
            MaxBin = maxBin;
            MaxDepth = maxDepth;
            MaxLeaves = maxLeaves;
            MinDataInLeaf = minDataInLeaf;
            MinSplitGain = minSplitGain;
            ModelName = modelName;
            NEstimators = nEstimators;
            NumLeaves = numLeaves;
            PreprocessorName = preprocessorName;
            RegAlpha = regAlpha;
            RegLambda = regLambda;
            Subsample = subsample;
            SubsampleFreq = subsampleFreq;
            TreeMethod = treeMethod;
            WithMean = withMean;
            WithStd = withStd;
        }
    }
}
