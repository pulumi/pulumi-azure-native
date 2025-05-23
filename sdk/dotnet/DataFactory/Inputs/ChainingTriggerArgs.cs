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
    /// Trigger that allows the referenced pipeline to depend on other pipeline runs based on runDimension Name/Value pairs. Upstream pipelines should declare the same runDimension Name and their runs should have the values for those runDimensions. The referenced pipeline run would be triggered if the values for the runDimension match for all upstream pipeline runs.
    /// </summary>
    public sealed class ChainingTriggerArgs : global::Pulumi.ResourceArgs
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

        [Input("dependsOn", required: true)]
        private InputList<Inputs.PipelineReferenceArgs>? _dependsOn;

        /// <summary>
        /// Upstream Pipelines.
        /// </summary>
        public InputList<Inputs.PipelineReferenceArgs> DependsOn
        {
            get => _dependsOn ?? (_dependsOn = new InputList<Inputs.PipelineReferenceArgs>());
            set => _dependsOn = value;
        }

        /// <summary>
        /// Trigger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Pipeline for which runs are created when all upstream pipelines complete successfully.
        /// </summary>
        [Input("pipeline", required: true)]
        public Input<Inputs.TriggerPipelineReferenceArgs> Pipeline { get; set; } = null!;

        /// <summary>
        /// Run Dimension property that needs to be emitted by upstream pipelines.
        /// </summary>
        [Input("runDimension", required: true)]
        public Input<string> RunDimension { get; set; } = null!;

        /// <summary>
        /// Trigger type.
        /// Expected value is 'ChainingTrigger'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ChainingTriggerArgs()
        {
        }
        public static new ChainingTriggerArgs Empty => new ChainingTriggerArgs();
    }
}
