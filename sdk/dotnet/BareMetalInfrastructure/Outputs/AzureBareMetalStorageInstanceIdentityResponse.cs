// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BareMetalInfrastructure.Outputs
{

    /// <summary>
    /// Identity for Azure Bare Metal Storage Instance.
    /// </summary>
    [OutputType]
    public sealed class AzureBareMetalStorageInstanceIdentityResponse
    {
        /// <summary>
        /// The principal ID of Azure Bare Metal Storage Instance identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID associated with the Azure Bare Metal Storage Instance. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the Azure Bare Metal Storage Instance. The type 'SystemAssigned' refers to an implicitly created identity. The type 'None' will remove any identities from the Azure Bare Metal Storage Instance.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private AzureBareMetalStorageInstanceIdentityResponse(
            string principalId,

            string tenantId,

            string? type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
