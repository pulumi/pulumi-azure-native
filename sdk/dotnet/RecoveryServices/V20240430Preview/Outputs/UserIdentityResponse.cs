// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.V20240430Preview.Outputs
{

    /// <summary>
    /// A resource identity that is managed by the user of the service.
    /// </summary>
    [OutputType]
    public sealed class UserIdentityResponse
    {
        /// <summary>
        /// The client ID of the user-assigned identity.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The principal ID of the user-assigned identity.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private UserIdentityResponse(
            string clientId,

            string principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
