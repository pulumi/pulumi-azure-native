// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins.Outputs
{

    /// <summary>
    /// The information about the user assigned identity.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedIdentityResponse
    {
        /// <summary>
        /// The client id of the User Assigned Identity Resource.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The object id of the User Assigned Identity Resource.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private UserAssignedIdentityResponse(
            string clientId,

            string principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
