// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240701Preview.Outputs
{

    /// <summary>
    /// A collection of information about the state of the connection between service consumer and provider.
    /// </summary>
    [OutputType]
    public sealed class PrivateLinkServiceConnectionStateResponse
    {
        /// <summary>
        /// Some RP chose "None". Other RPs use this for region expansion.
        /// </summary>
        public readonly string? ActionsRequired;
        /// <summary>
        /// User-defined message that, per NRP doc, may be used for approval-related message.
        /// </summary>
        public readonly string? Description;
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
        private PrivateLinkServiceConnectionStateResponse(
            string? actionsRequired,

            string? description,

            string? status)
        {
            ActionsRequired = actionsRequired;
            Description = description;
            Status = status;
        }
    }
}
