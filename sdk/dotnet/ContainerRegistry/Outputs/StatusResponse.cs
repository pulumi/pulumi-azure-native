// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The status of an Azure resource at the time the operation was called.
    /// </summary>
    [OutputType]
    public sealed class StatusResponse
    {
        /// <summary>
        /// The short label for the status.
        /// </summary>
        public readonly string DisplayStatus;
        /// <summary>
        /// The detailed message for the status, including alerts and error messages.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The timestamp when the status was changed to the current value.
        /// </summary>
        public readonly string Timestamp;

        [OutputConstructor]
        private StatusResponse(
            string displayStatus,

            string message,

            string timestamp)
        {
            DisplayStatus = displayStatus;
            Message = message;
            Timestamp = timestamp;
        }
    }
}
