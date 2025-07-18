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
    /// Describes a Virtual Machine Extension.
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineExtensionResponse
    {
        /// <summary>
        /// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
        /// </summary>
        public readonly bool? AutoUpgradeMinorVersion;
        /// <summary>
        /// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
        /// </summary>
        public readonly bool? EnableAutomaticUpgrade;
        /// <summary>
        /// How the extension handler should be forced to update even if the extension configuration has not changed.
        /// </summary>
        public readonly string? ForceUpdateTag;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The virtual machine extension instance view.
        /// </summary>
        public readonly Outputs.VirtualMachineExtensionInstanceViewResponse? InstanceView;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
        /// </summary>
        public readonly object? ProtectedSettings;
        /// <summary>
        /// The extensions protected settings that are passed by reference, and consumed from key vault
        /// </summary>
        public readonly Outputs.KeyVaultSecretReferenceResponse? ProtectedSettingsFromKeyVault;
        /// <summary>
        /// Collection of extension names after which this extension needs to be provisioned.
        /// </summary>
        public readonly ImmutableArray<string> ProvisionAfterExtensions;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The name of the extension handler publisher.
        /// </summary>
        public readonly string? Publisher;
        /// <summary>
        /// Json formatted public settings for the extension.
        /// </summary>
        public readonly object? Settings;
        /// <summary>
        /// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
        /// </summary>
        public readonly bool? SuppressFailures;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        public readonly string? TypeHandlerVersion;

        [OutputConstructor]
        private VirtualMachineExtensionResponse(
            bool? autoUpgradeMinorVersion,

            bool? enableAutomaticUpgrade,

            string? forceUpdateTag,

            string id,

            Outputs.VirtualMachineExtensionInstanceViewResponse? instanceView,

            string location,

            string name,

            object? protectedSettings,

            Outputs.KeyVaultSecretReferenceResponse? protectedSettingsFromKeyVault,

            ImmutableArray<string> provisionAfterExtensions,

            string provisioningState,

            string? publisher,

            object? settings,

            bool? suppressFailures,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? typeHandlerVersion)
        {
            AutoUpgradeMinorVersion = autoUpgradeMinorVersion;
            EnableAutomaticUpgrade = enableAutomaticUpgrade;
            ForceUpdateTag = forceUpdateTag;
            Id = id;
            InstanceView = instanceView;
            Location = location;
            Name = name;
            ProtectedSettings = protectedSettings;
            ProtectedSettingsFromKeyVault = protectedSettingsFromKeyVault;
            ProvisionAfterExtensions = provisionAfterExtensions;
            ProvisioningState = provisioningState;
            Publisher = publisher;
            Settings = settings;
            SuppressFailures = suppressFailures;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            TypeHandlerVersion = typeHandlerVersion;
        }
    }
}
