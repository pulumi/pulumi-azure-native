// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearning.V20160501Preview.Inputs
{

    /// <summary>
    /// Defines a global parameter in the graph.
    /// </summary>
    public sealed class GraphParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of this graph parameter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("links", required: true)]
        private InputList<Inputs.GraphParameterLinkArgs>? _links;

        /// <summary>
        /// Association links for this parameter to nodes in the graph.
        /// </summary>
        public InputList<Inputs.GraphParameterLinkArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.GraphParameterLinkArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Graph parameter's type.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.MachineLearning.V20160501Preview.ParameterType> Type { get; set; } = null!;

        public GraphParameterArgs()
        {
        }
        public static new GraphParameterArgs Empty => new GraphParameterArgs();
    }
}
