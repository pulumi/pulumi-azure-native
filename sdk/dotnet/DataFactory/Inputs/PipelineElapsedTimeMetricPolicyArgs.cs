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
    /// Pipeline ElapsedTime Metric Policy.
    /// </summary>
    public sealed class PipelineElapsedTimeMetricPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TimeSpan value, after which an Azure Monitoring Metric is fired.
        /// </summary>
        [Input("duration")]
        public Input<object>? Duration { get; set; }

        public PipelineElapsedTimeMetricPolicyArgs()
        {
        }
        public static new PipelineElapsedTimeMetricPolicyArgs Empty => new PipelineElapsedTimeMetricPolicyArgs();
    }
}
