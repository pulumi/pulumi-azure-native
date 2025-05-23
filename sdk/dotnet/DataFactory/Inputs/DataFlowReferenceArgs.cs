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
    /// Data flow reference type.
    /// </summary>
    public sealed class DataFlowReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reference data flow parameters from dataset.
        /// </summary>
        [Input("datasetParameters")]
        public Input<object>? DatasetParameters { get; set; }

        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// Data flow parameters
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        /// <summary>
        /// Reference data flow name.
        /// </summary>
        [Input("referenceName", required: true)]
        public Input<string> ReferenceName { get; set; } = null!;

        /// <summary>
        /// Data flow reference type.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.DataFlowReferenceType> Type { get; set; } = null!;

        public DataFlowReferenceArgs()
        {
        }
        public static new DataFlowReferenceArgs Empty => new DataFlowReferenceArgs();
    }
}
