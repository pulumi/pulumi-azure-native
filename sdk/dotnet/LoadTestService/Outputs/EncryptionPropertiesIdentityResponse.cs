// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LoadTestService.Outputs
{

    /// <summary>
    /// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
    /// </summary>
    [OutputType]
    public sealed class EncryptionPropertiesIdentityResponse
    {
        /// <summary>
        /// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/a0a0a0a0-bbbb-cccd-dddd-e1e1e1e1e1e1/resourceGroups/&lt;resource group&gt;/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Managed identity type to use for accessing encryption key Url.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private EncryptionPropertiesIdentityResponse(
            string? resourceId,

            string? type)
        {
            ResourceId = resourceId;
            Type = type;
        }
    }
}
