// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Trigger based on total requests.
    /// </summary>
    [OutputType]
    public sealed class RequestsBasedTriggerResponse
    {
        /// <summary>
        /// Request Count.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// Time interval.
        /// </summary>
        public readonly string? TimeInterval;

        [OutputConstructor]
        private RequestsBasedTriggerResponse(
            int? count,

            string? timeInterval)
        {
            Count = count;
            TimeInterval = timeInterval;
        }
    }
}
