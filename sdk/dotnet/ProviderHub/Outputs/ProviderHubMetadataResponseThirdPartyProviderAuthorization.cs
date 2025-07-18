// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    /// <summary>
    /// The third party provider authorization.
    /// </summary>
    [OutputType]
    public sealed class ProviderHubMetadataResponseThirdPartyProviderAuthorization
    {
        /// <summary>
        /// The authorizations.
        /// </summary>
        public readonly ImmutableArray<Outputs.LightHouseAuthorizationResponse> Authorizations;
        /// <summary>
        /// The managed by tenant id.
        /// </summary>
        public readonly string? ManagedByTenantId;

        [OutputConstructor]
        private ProviderHubMetadataResponseThirdPartyProviderAuthorization(
            ImmutableArray<Outputs.LightHouseAuthorizationResponse> authorizations,

            string? managedByTenantId)
        {
            Authorizations = authorizations;
            ManagedByTenantId = managedByTenantId;
        }
    }
}
