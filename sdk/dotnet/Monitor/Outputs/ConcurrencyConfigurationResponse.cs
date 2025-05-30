// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Concurrent publishing configuration.
    /// </summary>
    [OutputType]
    public sealed class ConcurrencyConfigurationResponse
    {
        /// <summary>
        /// Size of the queue for log batches.
        /// </summary>
        public readonly int? BatchQueueSize;
        /// <summary>
        /// Number of parallel workers processing the log queues.
        /// </summary>
        public readonly int? WorkerCount;

        [OutputConstructor]
        private ConcurrencyConfigurationResponse(
            int? batchQueueSize,

            int? workerCount)
        {
            BatchQueueSize = batchQueueSize;
            WorkerCount = workerCount;
        }
    }
}
