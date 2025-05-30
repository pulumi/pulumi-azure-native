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
    /// Optimization objective.
    /// </summary>
    [OutputType]
    public sealed class ObjectiveResponse
    {
        /// <summary>
        /// [Required] Defines supported metric goals for hyperparameter tuning
        /// </summary>
        public readonly string Goal;
        /// <summary>
        /// [Required] Name of the metric to optimize.
        /// </summary>
        public readonly string PrimaryMetric;

        [OutputConstructor]
        private ObjectiveResponse(
            string goal,

            string primaryMetric)
        {
            Goal = goal;
            PrimaryMetric = primaryMetric;
        }
    }
}
