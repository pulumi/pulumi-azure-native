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
    /// Data flow flowlet
    /// </summary>
    public sealed class FlowletArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<object>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the data flow.
        /// </summary>
        public InputList<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// The description of the data flow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The folder that this data flow is in. If not specified, Data flow will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<Inputs.DataFlowFolderArgs>? Folder { get; set; }

        /// <summary>
        /// Flowlet script.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        [Input("scriptLines")]
        private InputList<string>? _scriptLines;

        /// <summary>
        /// Flowlet script lines.
        /// </summary>
        public InputList<string> ScriptLines
        {
            get => _scriptLines ?? (_scriptLines = new InputList<string>());
            set => _scriptLines = value;
        }

        [Input("sinks")]
        private InputList<Inputs.DataFlowSinkArgs>? _sinks;

        /// <summary>
        /// List of sinks in Flowlet.
        /// </summary>
        public InputList<Inputs.DataFlowSinkArgs> Sinks
        {
            get => _sinks ?? (_sinks = new InputList<Inputs.DataFlowSinkArgs>());
            set => _sinks = value;
        }

        [Input("sources")]
        private InputList<Inputs.DataFlowSourceArgs>? _sources;

        /// <summary>
        /// List of sources in Flowlet.
        /// </summary>
        public InputList<Inputs.DataFlowSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DataFlowSourceArgs>());
            set => _sources = value;
        }

        [Input("transformations")]
        private InputList<Inputs.TransformationArgs>? _transformations;

        /// <summary>
        /// List of transformations in Flowlet.
        /// </summary>
        public InputList<Inputs.TransformationArgs> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<Inputs.TransformationArgs>());
            set => _transformations = value;
        }

        /// <summary>
        /// Type of data flow.
        /// Expected value is 'Flowlet'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FlowletArgs()
        {
        }
        public static new FlowletArgs Empty => new FlowletArgs();
    }
}
