// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    [OutputType]
    public sealed class SharedPrivateLinkResourceResponse
    {
        /// <summary>
        /// group id of the private link
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// Unique name of the private link
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// the resource id that private link links to
        /// </summary>
        public readonly string? PrivateLinkResourceId;
        /// <summary>
        /// Request message
        /// </summary>
        public readonly string? RequestMessage;
        /// <summary>
        /// Connection status of the service consumer with the service provider
        /// Possible state transitions
        /// Pending -&gt; Approved (Service provider approves the connection request)
        /// Pending -&gt; Rejected (Service provider rejects the connection request)
        /// Pending -&gt; Disconnected (Service provider deletes the connection)
        /// Approved -&gt; Rejected (Service provider rejects the approved connection)
        /// Approved -&gt; Disconnected (Service provider deletes the connection)
        /// Rejected -&gt; Pending (Service consumer re-initiates the connection request that was rejected)
        /// Rejected -&gt; Disconnected (Service provider deletes the connection)
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SharedPrivateLinkResourceResponse(
            string? groupId,

            string? name,

            string? privateLinkResourceId,

            string? requestMessage,

            string? status)
        {
            GroupId = groupId;
            Name = name;
            PrivateLinkResourceId = privateLinkResourceId;
            RequestMessage = requestMessage;
            Status = status;
        }
    }
}
