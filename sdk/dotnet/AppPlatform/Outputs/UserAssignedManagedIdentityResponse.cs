// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// The details of the user-assigned managed identity assigned to an App.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedManagedIdentityResponse
    {
        /// <summary>
        /// Client Id of user-assigned managed identity.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Principal Id of user-assigned managed identity.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private UserAssignedManagedIdentityResponse(
            string clientId,

            string principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
