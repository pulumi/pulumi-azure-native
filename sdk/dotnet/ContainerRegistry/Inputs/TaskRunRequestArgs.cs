// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Inputs
{

    /// <summary>
    /// The parameters for a task run request.
    /// </summary>
    public sealed class TaskRunRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dedicated agent pool for the run.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// The value that indicates whether archiving is enabled for the run or not.
        /// </summary>
        [Input("isArchiveEnabled")]
        public Input<bool>? IsArchiveEnabled { get; set; }

        /// <summary>
        /// The template that describes the repository and tag information for run log artifact.
        /// </summary>
        [Input("logTemplate")]
        public Input<string>? LogTemplate { get; set; }

        /// <summary>
        /// Set of overridable parameters that can be passed when running a Task.
        /// </summary>
        [Input("overrideTaskStepProperties")]
        public Input<Inputs.OverrideTaskStepPropertiesArgs>? OverrideTaskStepProperties { get; set; }

        /// <summary>
        /// The resource ID of task against which run has to be queued.
        /// </summary>
        [Input("taskId", required: true)]
        public Input<string> TaskId { get; set; } = null!;

        /// <summary>
        /// The type of the run request.
        /// Expected value is 'TaskRunRequest'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TaskRunRequestArgs()
        {
            IsArchiveEnabled = false;
        }
        public static new TaskRunRequestArgs Empty => new TaskRunRequestArgs();
    }
}
