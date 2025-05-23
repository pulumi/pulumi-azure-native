// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Outputs
{

    /// <summary>
    /// Identity for the Provisioned cluster.
    /// </summary>
    [OutputType]
    public sealed class ProvisionedClusterIdentityResponse
    {
        /// <summary>
        /// The principal id of provisioned cluster identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id associated with the provisioned cluster. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the provisioned cluster. The type SystemAssigned, includes a system created identity. The type None means no identity is assigned to the provisioned cluster.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ProvisionedClusterIdentityResponse(
            string principalId,

            string tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
