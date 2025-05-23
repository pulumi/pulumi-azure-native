// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Inputs
{

    /// <summary>
    /// The metric counter
    /// </summary>
    public sealed class MetricCounterArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalDimensions")]
        private InputList<Inputs.MetricDimensionArgs>? _additionalDimensions;

        /// <summary>
        /// The additional dimensions to be added to metric.
        /// </summary>
        public InputList<Inputs.MetricDimensionArgs> AdditionalDimensions
        {
            get => _additionalDimensions ?? (_additionalDimensions = new InputList<Inputs.MetricDimensionArgs>());
            set => _additionalDimensions = value;
        }

        [Input("dimensionFilter")]
        private InputList<Inputs.MetricDimensionArgs>? _dimensionFilter;

        /// <summary>
        /// The dimension filter.
        /// </summary>
        public InputList<Inputs.MetricDimensionArgs> DimensionFilter
        {
            get => _dimensionFilter ?? (_dimensionFilter = new InputList<Inputs.MetricDimensionArgs>());
            set => _dimensionFilter = value;
        }

        /// <summary>
        /// The instance from which counter should be collected.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// The counter name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public MetricCounterArgs()
        {
        }
        public static new MetricCounterArgs Empty => new MetricCounterArgs();
    }
}
