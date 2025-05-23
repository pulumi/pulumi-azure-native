// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ResourceConnector.Outputs
{

    /// <summary>
    /// Appliance SSHKey definition.
    /// </summary>
    [OutputType]
    public sealed class SSHKeyResponse
    {
        /// <summary>
        /// Certificate associated with the public key if the key is signed.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Certificate creation timestamp (Unix).
        /// </summary>
        public readonly double CreationTimeStamp;
        /// <summary>
        /// Certificate expiration timestamp (Unix).
        /// </summary>
        public readonly double ExpirationTimeStamp;
        /// <summary>
        /// Private Key.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// Public Key.
        /// </summary>
        public readonly string PublicKey;

        [OutputConstructor]
        private SSHKeyResponse(
            string certificate,

            double creationTimeStamp,

            double expirationTimeStamp,

            string privateKey,

            string publicKey)
        {
            Certificate = certificate;
            CreationTimeStamp = creationTimeStamp;
            ExpirationTimeStamp = expirationTimeStamp;
            PrivateKey = privateKey;
            PublicKey = publicKey;
        }
    }
}
