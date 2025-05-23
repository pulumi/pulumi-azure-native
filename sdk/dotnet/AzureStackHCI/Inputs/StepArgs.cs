// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// Progress representation of the update run steps.
    /// </summary>
    public sealed class StepArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// More detailed description of the step.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When the step reached a terminal state.
        /// </summary>
        [Input("endTimeUtc")]
        public Input<string>? EndTimeUtc { get; set; }

        /// <summary>
        /// Error message, specified if the step is in a failed state.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
        /// </summary>
        [Input("expectedExecutionTime")]
        public Input<string>? ExpectedExecutionTime { get; set; }

        /// <summary>
        /// Completion time of this step or the last completed sub-step.
        /// </summary>
        [Input("lastUpdatedTimeUtc")]
        public Input<string>? LastUpdatedTimeUtc { get; set; }

        /// <summary>
        /// Name of the step.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When the step started, or empty if it has not started executing.
        /// </summary>
        [Input("startTimeUtc")]
        public Input<string>? StartTimeUtc { get; set; }

        /// <summary>
        /// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("steps")]
        private InputList<Inputs.StepArgs>? _steps;

        /// <summary>
        /// Recursive model for child steps of this step.
        /// </summary>
        public InputList<Inputs.StepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.StepArgs>());
            set => _steps = value;
        }

        public StepArgs()
        {
        }
        public static new StepArgs Empty => new StepArgs();
    }
}
