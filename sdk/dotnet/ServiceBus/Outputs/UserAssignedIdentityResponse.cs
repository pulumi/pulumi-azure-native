// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceBus.Outputs
{

    /// <summary>
    /// Recognized Dictionary value.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedIdentityResponse
    {
        /// <summary>
        /// Client Id of user assigned identity
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Principal Id of user assigned identity
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
