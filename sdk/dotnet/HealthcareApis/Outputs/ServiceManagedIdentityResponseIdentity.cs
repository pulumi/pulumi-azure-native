// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Outputs
{

    /// <summary>
    /// Setting indicating whether the service has a managed identity associated with it.
    /// </summary>
    [OutputType]
    public sealed class ServiceManagedIdentityResponseIdentity
    {
        /// <summary>
        /// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Type of identity being specified, currently SystemAssigned and None are allowed.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.UserAssignedIdentityResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private ServiceManagedIdentityResponseIdentity(
            string principalId,

            string tenantId,

            string type,

            ImmutableDictionary<string, Outputs.UserAssignedIdentityResponse>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
