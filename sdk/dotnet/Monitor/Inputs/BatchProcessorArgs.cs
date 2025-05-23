// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// Batch processor.
    /// </summary>
    public sealed class BatchProcessorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the batch.
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        /// <summary>
        /// Timeout in milliseconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public BatchProcessorArgs()
        {
            BatchSize = 8192;
            Timeout = 200;
        }
        public static new BatchProcessorArgs Empty => new BatchProcessorArgs();
    }
}
