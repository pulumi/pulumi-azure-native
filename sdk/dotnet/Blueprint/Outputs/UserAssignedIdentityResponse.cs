// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint.Outputs
{

    /// <summary>
    /// User-assigned managed identity.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedIdentityResponse
    {
        /// <summary>
        /// Client App Id associated with this identity.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Azure Active Directory principal ID associated with this Identity.
        /// </summary>
        public readonly string? PrincipalId;

        [OutputConstructor]
        private UserAssignedIdentityResponse(
            string? clientId,

            string? principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
