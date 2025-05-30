// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// Azure Active Directory identity configuration for a resource.
    /// </summary>
    [OutputType]
    public sealed class ResourceIdentityResponse
    {
        /// <summary>
        /// The Azure Active Directory principal id.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The Azure Active Directory tenant id.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The resource ids of the user assigned identities to use
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.UserIdentityResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private ResourceIdentityResponse(
            string principalId,

            string tenantId,

            string? type,

            ImmutableDictionary<string, Outputs.UserIdentityResponse>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
