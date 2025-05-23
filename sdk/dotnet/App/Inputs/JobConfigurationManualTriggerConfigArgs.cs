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
    /// Manual trigger configuration for a single execution job. Properties replicaCompletionCount and parallelism would be set to 1 by default
    /// </summary>
    public sealed class JobConfigurationManualTriggerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of parallel replicas of a job that can run at a given time.
        /// </summary>
        [Input("parallelism")]
        public Input<int>? Parallelism { get; set; }

        /// <summary>
        /// Minimum number of successful replica completions before overall job completion.
        /// </summary>
        [Input("replicaCompletionCount")]
        public Input<int>? ReplicaCompletionCount { get; set; }

        public JobConfigurationManualTriggerConfigArgs()
        {
        }
        public static new JobConfigurationManualTriggerConfigArgs Empty => new JobConfigurationManualTriggerConfigArgs();
    }
}
