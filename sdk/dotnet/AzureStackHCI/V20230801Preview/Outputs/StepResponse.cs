// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20230801Preview.Outputs
{

    /// <summary>
    /// Progress representation of the update run steps.
    /// </summary>
    [OutputType]
    public sealed class StepResponse
    {
        /// <summary>
        /// More detailed description of the step.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// When the step reached a terminal state.
        /// </summary>
        public readonly string? EndTimeUtc;
        /// <summary>
        /// Error message, specified if the step is in a failed state.
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// List of exceptions in AzureStackHCI Cluster Deployment.
        /// </summary>
        public readonly ImmutableArray<string> Exception;
        /// <summary>
        /// FullStepIndex of step.
        /// </summary>
        public readonly string? FullStepIndex;
        /// <summary>
        /// Completion time of this step or the last completed sub-step.
        /// </summary>
        public readonly string? LastUpdatedTimeUtc;
        /// <summary>
        /// Name of the step.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// When the step started, or empty if it has not started executing.
        /// </summary>
        public readonly string? StartTimeUtc;
        /// <summary>
        /// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Recursive model for child steps of this step.
        /// </summary>
        public readonly ImmutableArray<Outputs.StepResponse> Steps;

        [OutputConstructor]
        private StepResponse(
            string? description,

            string? endTimeUtc,

            string? errorMessage,

            ImmutableArray<string> exception,

            string? fullStepIndex,

            string? lastUpdatedTimeUtc,

            string? name,

            string? startTimeUtc,

            string? status,

            ImmutableArray<Outputs.StepResponse> steps)
        {
            Description = description;
            EndTimeUtc = endTimeUtc;
            ErrorMessage = errorMessage;
            Exception = exception;
            FullStepIndex = fullStepIndex;
            LastUpdatedTimeUtc = lastUpdatedTimeUtc;
            Name = name;
            StartTimeUtc = startTimeUtc;
            Status = status;
            Steps = steps;
        }
    }
}
