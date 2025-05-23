// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LoadTestService.Inputs
{

    /// <summary>
    /// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
    /// </summary>
    public sealed class EncryptionPropertiesIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/a0a0a0a0-bbbb-cccd-dddd-e1e1e1e1e1e1/resourceGroups/&lt;resource group&gt;/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Managed identity type to use for accessing encryption key Url.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.LoadTestService.Type>? Type { get; set; }

        public EncryptionPropertiesIdentityArgs()
        {
        }
        public static new EncryptionPropertiesIdentityArgs Empty => new EncryptionPropertiesIdentityArgs();
    }
}
