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
    /// Concurrent publishing configuration.
    /// </summary>
    public sealed class ConcurrencyConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the queue for log batches.
        /// </summary>
        [Input("batchQueueSize")]
        public Input<int>? BatchQueueSize { get; set; }

        /// <summary>
        /// Number of parallel workers processing the log queues.
        /// </summary>
        [Input("workerCount")]
        public Input<int>? WorkerCount { get; set; }

        public ConcurrencyConfigurationArgs()
        {
            BatchQueueSize = 100;
            WorkerCount = 4;
        }
        public static new ConcurrencyConfigurationArgs Empty => new ConcurrencyConfigurationArgs();
    }
}
