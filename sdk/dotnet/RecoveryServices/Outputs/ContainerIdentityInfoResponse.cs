// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Container identity information
    /// </summary>
    [OutputType]
    public sealed class ContainerIdentityInfoResponse
    {
        /// <summary>
        /// Protection container identity - AAD Tenant
        /// </summary>
        public readonly string? AadTenantId;
        /// <summary>
        /// Protection container identity - Audience
        /// </summary>
        public readonly string? Audience;
        /// <summary>
        /// Protection container identity - AAD Service Principal
        /// </summary>
        public readonly string? ServicePrincipalClientId;
        /// <summary>
        /// Unique name of the container
        /// </summary>
        public readonly string? UniqueName;

        [OutputConstructor]
        private ContainerIdentityInfoResponse(
            string? aadTenantId,

            string? audience,

            string? servicePrincipalClientId,

            string? uniqueName)
        {
            AadTenantId = aadTenantId;
            Audience = audience;
            ServicePrincipalClientId = servicePrincipalClientId;
            UniqueName = uniqueName;
        }
    }
}
