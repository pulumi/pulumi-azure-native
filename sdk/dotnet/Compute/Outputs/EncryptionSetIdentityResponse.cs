// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
    /// </summary>
    [OutputType]
    public sealed class EncryptionSetIdentityResponse
    {
        /// <summary>
        /// The object id of the Managed Identity Resource. This will be sent to the RP from ARM via the x-ms-identity-principal-id header in the PUT request if the resource has a systemAssigned(implicit) identity
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id of the Managed Identity Resource. This will be sent to the RP from ARM via the x-ms-client-tenant-id header in the PUT request if the resource has a systemAssigned(implicit) identity
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported for new creations. Disk Encryption Sets can be updated with Identity type None during migration of subscription to a new Azure Active Directory tenant; it will cause the encrypted resources to lose access to the keys.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The list of user identities associated with the disk encryption set. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.UserAssignedIdentitiesValueResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private EncryptionSetIdentityResponse(
            string principalId,

            string tenantId,

            string? type,

            ImmutableDictionary<string, Outputs.UserAssignedIdentitiesValueResponse>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
