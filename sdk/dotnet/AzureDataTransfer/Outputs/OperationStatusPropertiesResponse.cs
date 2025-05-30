// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer.Outputs
{

    /// <summary>
    /// Operation status associated with the last patch request
    /// </summary>
    [OutputType]
    public sealed class OperationStatusPropertiesResponse
    {
        /// <summary>
        /// Operation status ID of the last patch request for this connection.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Message for the operation for the last patch request for this connection.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Operation status for the last patch request for this connection.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private OperationStatusPropertiesResponse(
            string id,

            string message,

            string status)
        {
            Id = id;
            Message = message;
            Status = status;
        }
    }
}
