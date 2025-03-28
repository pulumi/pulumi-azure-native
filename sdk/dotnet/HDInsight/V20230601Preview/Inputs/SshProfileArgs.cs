// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.V20230601Preview.Inputs
{

    /// <summary>
    /// Ssh profile for the cluster.
    /// </summary>
    public sealed class SshProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of ssh pods per cluster.
        /// </summary>
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public SshProfileArgs()
        {
        }
        public static new SshProfileArgs Empty => new SshProfileArgs();
    }
}
