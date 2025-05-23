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
    /// Static input data definition.
    /// </summary>
    public sealed class StaticInputDataArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputMap<string>? _columns;

        /// <summary>
        /// Mapping of column names to special uses.
        /// </summary>
        public InputMap<string> Columns
        {
            get => _columns ?? (_columns = new InputMap<string>());
            set => _columns = value;
        }

        /// <summary>
        /// The context metadata of the data source.
        /// </summary>
        [Input("dataContext")]
        public Input<string>? DataContext { get; set; }

        /// <summary>
        /// Monitoring input data type enum.
        /// Expected value is 'Static'.
        /// </summary>
        [Input("inputDataType", required: true)]
        public Input<string> InputDataType { get; set; } = null!;

        /// <summary>
        /// [Required] Specifies the type of job.
        /// </summary>
        [Input("jobInputType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.JobInputType> JobInputType { get; set; } = null!;

        /// <summary>
        /// Reference to the component asset used to preprocess the data.
        /// </summary>
        [Input("preprocessingComponentId")]
        public Input<string>? PreprocessingComponentId { get; set; }

        /// <summary>
        /// [Required] Input Asset URI.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        /// <summary>
        /// [Required] The end date of the data window.
        /// </summary>
        [Input("windowEnd", required: true)]
        public Input<string> WindowEnd { get; set; } = null!;

        /// <summary>
        /// [Required] The start date of the data window.
        /// </summary>
        [Input("windowStart", required: true)]
        public Input<string> WindowStart { get; set; } = null!;

        public StaticInputDataArgs()
        {
        }
        public static new StaticInputDataArgs Empty => new StaticInputDataArgs();
    }
}
