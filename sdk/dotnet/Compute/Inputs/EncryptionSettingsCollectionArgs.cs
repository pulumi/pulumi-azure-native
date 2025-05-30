// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Encryption settings for disk or snapshot
    /// </summary>
    public sealed class EncryptionSettingsCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("encryptionSettings")]
        private InputList<Inputs.EncryptionSettingsElementArgs>? _encryptionSettings;

        /// <summary>
        /// A collection of encryption settings, one for each disk volume.
        /// </summary>
        public InputList<Inputs.EncryptionSettingsElementArgs> EncryptionSettings
        {
            get => _encryptionSettings ?? (_encryptionSettings = new InputList<Inputs.EncryptionSettingsElementArgs>());
            set => _encryptionSettings = value;
        }

        /// <summary>
        /// Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
        /// </summary>
        [Input("encryptionSettingsVersion")]
        public Input<string>? EncryptionSettingsVersion { get; set; }

        public EncryptionSettingsCollectionArgs()
        {
        }
        public static new EncryptionSettingsCollectionArgs Empty => new EncryptionSettingsCollectionArgs();
    }
}
