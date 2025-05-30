// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Base class for all triggers that support one to many model for trigger to pipeline.
    /// </summary>
    public sealed class MultiplePipelineTriggerArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<object>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the trigger.
        /// </summary>
        public InputList<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// Trigger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("pipelines")]
        private InputList<Inputs.TriggerPipelineReferenceArgs>? _pipelines;

        /// <summary>
        /// Pipelines that need to be started.
        /// </summary>
        public InputList<Inputs.TriggerPipelineReferenceArgs> Pipelines
        {
            get => _pipelines ?? (_pipelines = new InputList<Inputs.TriggerPipelineReferenceArgs>());
            set => _pipelines = value;
        }

        /// <summary>
        /// Trigger type.
        /// Expected value is 'MultiplePipelineTrigger'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MultiplePipelineTriggerArgs()
        {
        }
        public static new MultiplePipelineTriggerArgs Empty => new MultiplePipelineTriggerArgs();
    }
}
