// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.Outputs
{

    /// <summary>
    /// A class represent managed identities used for request and response
    /// </summary>
    [OutputType]
    public sealed class ManagedIdentityResponse
    {
        /// <summary>
        /// Get the principal id for the system assigned identity.
        /// Only be used in response.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// Get the tenant id for the system assigned identity.
        /// Only be used in response
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Represents the identity type: systemAssigned, userAssigned, None
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Get or set the user assigned identities
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.UserAssignedIdentityPropertyResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private ManagedIdentityResponse(
            string principalId,

            string tenantId,

            string? type,

            ImmutableDictionary<string, Outputs.UserAssignedIdentityPropertyResponse>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
