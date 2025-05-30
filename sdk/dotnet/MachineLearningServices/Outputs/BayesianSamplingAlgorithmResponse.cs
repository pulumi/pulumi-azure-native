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
    /// Defines a Sampling Algorithm that generates values based on previous values
    /// </summary>
    [OutputType]
    public sealed class BayesianSamplingAlgorithmResponse
    {
        /// <summary>
        /// 
        /// Expected value is 'Bayesian'.
        /// </summary>
        public readonly string SamplingAlgorithmType;

        [OutputConstructor]
        private BayesianSamplingAlgorithmResponse(string samplingAlgorithmType)
        {
            SamplingAlgorithmType = samplingAlgorithmType;
        }
    }
}
