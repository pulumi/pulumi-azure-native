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
    /// Workload profile to scope container app execution.
    /// </summary>
    public sealed class WorkloadProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum capacity.
        /// </summary>
        [Input("maximumCount")]
        public Input<int>? MaximumCount { get; set; }

        /// <summary>
        /// The minimum capacity.
        /// </summary>
        [Input("minimumCount")]
        public Input<int>? MinimumCount { get; set; }

        /// <summary>
        /// Workload profile type for the workloads to run on.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Workload profile type for the workloads to run on.
        /// </summary>
        [Input("workloadProfileType", required: true)]
        public Input<string> WorkloadProfileType { get; set; } = null!;

        public WorkloadProfileArgs()
        {
        }
        public static new WorkloadProfileArgs Empty => new WorkloadProfileArgs();
    }
}
