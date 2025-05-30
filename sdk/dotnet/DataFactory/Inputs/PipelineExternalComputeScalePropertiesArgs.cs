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
    /// PipelineExternalComputeScale properties for managed integration runtime.
    /// </summary>
    public sealed class PipelineExternalComputeScalePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of the the external nodes, which should be greater than 0 and less than 11.
        /// </summary>
        [Input("numberOfExternalNodes")]
        public Input<int>? NumberOfExternalNodes { get; set; }

        /// <summary>
        /// Number of the pipeline nodes, which should be greater than 0 and less than 11.
        /// </summary>
        [Input("numberOfPipelineNodes")]
        public Input<int>? NumberOfPipelineNodes { get; set; }

        /// <summary>
        /// Time to live (in minutes) setting of integration runtime which will execute pipeline and external activity.
        /// </summary>
        [Input("timeToLive")]
        public Input<int>? TimeToLive { get; set; }

        public PipelineExternalComputeScalePropertiesArgs()
        {
        }
        public static new PipelineExternalComputeScalePropertiesArgs Empty => new PipelineExternalComputeScalePropertiesArgs();
    }
}
