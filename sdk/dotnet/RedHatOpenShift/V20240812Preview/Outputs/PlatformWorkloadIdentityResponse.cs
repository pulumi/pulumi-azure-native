// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.V20240812Preview.Outputs
{

    /// <summary>
    /// PlatformWorkloadIdentity stores information representing a single workload identity.
    /// </summary>
    [OutputType]
    public sealed class PlatformWorkloadIdentityResponse
    {
        /// <summary>
        /// The ClientID of the PlatformWorkloadIdentity resource
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The ObjectID of the PlatformWorkloadIdentity resource
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The resource ID of the PlatformWorkloadIdentity resource
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private PlatformWorkloadIdentityResponse(
            string clientId,

            string objectId,

            string? resourceId)
        {
            ClientId = clientId;
            ObjectId = objectId;
            ResourceId = resourceId;
        }
    }
}
