// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Inputs
{

    /// <summary>
    /// Specifies that the content key ID is specified in the PlayReady configuration.
    /// </summary>
    public sealed class ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content key ID.
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        public ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs()
        {
        }
        public static new ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs Empty => new ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs();
    }
}
