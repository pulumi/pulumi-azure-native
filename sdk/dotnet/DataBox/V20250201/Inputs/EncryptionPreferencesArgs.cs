// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.V20250201.Inputs
{

    /// <summary>
    /// Preferences related to the Encryption.
    /// </summary>
    public sealed class EncryptionPreferencesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines secondary layer of software-based encryption enablement.
        /// </summary>
        [Input("doubleEncryption")]
        public InputUnion<string, Pulumi.AzureNative.DataBox.V20250201.DoubleEncryption>? DoubleEncryption { get; set; }

        /// <summary>
        /// Defines Hardware level encryption (Only for disk)
        /// </summary>
        [Input("hardwareEncryption")]
        public InputUnion<string, Pulumi.AzureNative.DataBox.V20250201.HardwareEncryption>? HardwareEncryption { get; set; }

        public EncryptionPreferencesArgs()
        {
            DoubleEncryption = "Disabled";
        }
        public static new EncryptionPreferencesArgs Empty => new EncryptionPreferencesArgs();
    }
}
