// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    [OutputType]
    public sealed class FinetuningDetailsResponse
    {
        /// <summary>
        /// Finetuning Hyperparameters
        /// </summary>
        public readonly ImmutableDictionary<string, string>? HyperParameters;
        /// <summary>
        /// [Required] Student model for fine tuning.
        /// </summary>
        public readonly object StudentModel;

        [OutputConstructor]
        private FinetuningDetailsResponse(
            ImmutableDictionary<string, string>? hyperParameters,

            object studentModel)
        {
            HyperParameters = hyperParameters;
            StudentModel = studentModel;
        }
    }
}
