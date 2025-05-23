// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.Inputs
{

    /// <summary>
    /// Properties to EncryptionScope
    /// </summary>
    public sealed class EncryptionScopePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enumerates the possible value of keySource for Encryption
        /// </summary>
        [Input("keySource")]
        public InputUnion<string, Pulumi.AzureNative.CognitiveServices.KeySource>? KeySource { get; set; }

        /// <summary>
        /// Properties of KeyVault
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        /// <summary>
        /// The encryptionScope state.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.CognitiveServices.EncryptionScopeState>? State { get; set; }

        public EncryptionScopePropertiesArgs()
        {
            KeySource = "Microsoft.KeyVault";
        }
        public static new EncryptionScopePropertiesArgs Empty => new EncryptionScopePropertiesArgs();
    }
}
