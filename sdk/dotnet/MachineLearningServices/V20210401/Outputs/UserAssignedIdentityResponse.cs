// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210401.Outputs
{

    /// <summary>
    /// User Assigned Identity
    /// </summary>
    [OutputType]
    public sealed class UserAssignedIdentityResponse
    {
        /// <summary>
        /// The clientId(aka appId) of the user assigned identity.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The principal ID of the user assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID of the user assigned identity.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private UserAssignedIdentityResponse(
            string clientId,

            string principalId,

            string tenantId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
            TenantId = tenantId;
        }
    }
}
