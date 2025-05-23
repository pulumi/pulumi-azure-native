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
    public sealed class JobAgentIdentityResponse
    {
        /// <summary>
        /// The job agent identity tenant id
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The job agent identity type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The resource ids of the user assigned identities to use
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.JobAgentUserAssignedIdentityResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private JobAgentIdentityResponse(
            string? tenantId,

            string type,

            ImmutableDictionary<string, Outputs.JobAgentUserAssignedIdentityResponse>? userAssignedIdentities)
        {
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
