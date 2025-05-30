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
    /// Defines an early termination policy based on slack criteria, and a frequency and delay interval for evaluation
    /// </summary>
    public sealed class BanditPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of intervals by which to delay the first evaluation.
        /// </summary>
        [Input("delayEvaluation")]
        public Input<int>? DelayEvaluation { get; set; }

        /// <summary>
        /// Interval (number of runs) between policy evaluations.
        /// </summary>
        [Input("evaluationInterval")]
        public Input<int>? EvaluationInterval { get; set; }

        /// <summary>
        /// 
        /// Expected value is 'Bandit'.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<string> PolicyType { get; set; } = null!;

        /// <summary>
        /// Absolute distance allowed from the best performing run.
        /// </summary>
        [Input("slackAmount")]
        public Input<double>? SlackAmount { get; set; }

        /// <summary>
        /// Ratio of the allowed distance from the best performing run.
        /// </summary>
        [Input("slackFactor")]
        public Input<double>? SlackFactor { get; set; }

        public BanditPolicyArgs()
        {
            DelayEvaluation = 0;
            EvaluationInterval = 0;
            SlackAmount = 0;
            SlackFactor = 0;
        }
        public static new BanditPolicyArgs Empty => new BanditPolicyArgs();
    }
}
