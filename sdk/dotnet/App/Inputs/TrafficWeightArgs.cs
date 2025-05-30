// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Traffic weight assigned to a revision
    /// </summary>
    public sealed class TrafficWeightArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Associates a traffic label with a revision
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Indicates that the traffic weight belongs to a latest stable revision
        /// </summary>
        [Input("latestRevision")]
        public Input<bool>? LatestRevision { get; set; }

        /// <summary>
        /// Name of a revision
        /// </summary>
        [Input("revisionName")]
        public Input<string>? RevisionName { get; set; }

        /// <summary>
        /// Traffic weight assigned to a revision
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TrafficWeightArgs()
        {
            LatestRevision = false;
        }
        public static new TrafficWeightArgs Empty => new TrafficWeightArgs();
    }
}
