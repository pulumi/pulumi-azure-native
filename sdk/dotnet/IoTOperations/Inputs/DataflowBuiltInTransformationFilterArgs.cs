// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Inputs
{

    /// <summary>
    /// Dataflow BuiltIn Transformation filter properties
    /// </summary>
    public sealed class DataflowBuiltInTransformationFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user provided optional description of the filter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Condition to filter data. Can reference input fields with {n} where n is the index of the input field starting from 1. Example: $1 &lt; 0 || $1 &gt; $2 (Assuming inputs section $1 and $2 are provided)
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("inputs", required: true)]
        private InputList<string>? _inputs;

        /// <summary>
        /// List of fields for filtering in JSON path expression.
        /// </summary>
        public InputList<string> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<string>());
            set => _inputs = value;
        }

        /// <summary>
        /// The type of dataflow operation.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.IoTOperations.FilterType>? Type { get; set; }

        public DataflowBuiltInTransformationFilterArgs()
        {
            Type = "Filter";
        }
        public static new DataflowBuiltInTransformationFilterArgs Empty => new DataflowBuiltInTransformationFilterArgs();
    }
}
