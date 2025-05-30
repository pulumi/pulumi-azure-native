// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Outputs
{

    /// <summary>
    /// Managed Virtual Network Settings
    /// </summary>
    [OutputType]
    public sealed class ManagedVirtualNetworkSettingsResponse
    {
        /// <summary>
        /// Allowed Aad Tenant Ids For Linking
        /// </summary>
        public readonly ImmutableArray<string> AllowedAadTenantIdsForLinking;
        /// <summary>
        /// Linked Access Check On Target Resource
        /// </summary>
        public readonly bool? LinkedAccessCheckOnTargetResource;
        /// <summary>
        /// Prevent Data Exfiltration
        /// </summary>
        public readonly bool? PreventDataExfiltration;

        [OutputConstructor]
        private ManagedVirtualNetworkSettingsResponse(
            ImmutableArray<string> allowedAadTenantIdsForLinking,

            bool? linkedAccessCheckOnTargetResource,

            bool? preventDataExfiltration)
        {
            AllowedAadTenantIdsForLinking = allowedAadTenantIdsForLinking;
            LinkedAccessCheckOnTargetResource = linkedAccessCheckOnTargetResource;
            PreventDataExfiltration = preventDataExfiltration;
        }
    }
}
