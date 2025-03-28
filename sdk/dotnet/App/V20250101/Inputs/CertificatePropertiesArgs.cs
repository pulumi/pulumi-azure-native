// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Inputs
{

    /// <summary>
    /// Certificate resource specific properties
    /// </summary>
    public sealed class CertificatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Properties for a certificate stored in a Key Vault.
        /// </summary>
        [Input("certificateKeyVaultProperties")]
        public Input<Inputs.CertificateKeyVaultPropertiesArgs>? CertificateKeyVaultProperties { get; set; }

        /// <summary>
        /// Certificate password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// PFX or PEM blob
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public CertificatePropertiesArgs()
        {
        }
        public static new CertificatePropertiesArgs Empty => new CertificatePropertiesArgs();
    }
}
