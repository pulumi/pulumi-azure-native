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
    /// Encryption settings for disk or snapshot
    /// </summary>
    [OutputType]
    public sealed class EncryptionSettingsCollectionResponse
    {
        /// <summary>
        /// Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// A collection of encryption settings, one for each disk volume.
        /// </summary>
        public readonly ImmutableArray<Outputs.EncryptionSettingsElementResponse> EncryptionSettings;
        /// <summary>
        /// Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
        /// </summary>
        public readonly string? EncryptionSettingsVersion;

        [OutputConstructor]
        private EncryptionSettingsCollectionResponse(
            bool enabled,

            ImmutableArray<Outputs.EncryptionSettingsElementResponse> encryptionSettings,

            string? encryptionSettingsVersion)
        {
            Enabled = enabled;
            EncryptionSettings = encryptionSettings;
            EncryptionSettingsVersion = encryptionSettingsVersion;
        }
    }
}
