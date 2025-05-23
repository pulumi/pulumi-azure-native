// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard.Outputs
{

    /// <summary>
    /// The state of managed private endpoint connection.
    /// </summary>
    [OutputType]
    public sealed class ManagedPrivateEndpointConnectionStateResponse
    {
        /// <summary>
        /// Gets or sets the reason for approval/rejection of the connection.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The approval/rejection status of managed private endpoint connection.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private ManagedPrivateEndpointConnectionStateResponse(
            string description,

            string status)
        {
            Description = description;
            Status = status;
        }
    }
}
