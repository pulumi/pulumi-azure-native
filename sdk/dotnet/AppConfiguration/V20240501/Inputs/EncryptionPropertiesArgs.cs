// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppConfiguration.V20240501.Inputs
{

    /// <summary>
    /// The encryption settings for a configuration store.
    /// </summary>
    public sealed class EncryptionPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key vault properties.
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        public EncryptionPropertiesArgs()
        {
        }
        public static new EncryptionPropertiesArgs Empty => new EncryptionPropertiesArgs();
    }
}
