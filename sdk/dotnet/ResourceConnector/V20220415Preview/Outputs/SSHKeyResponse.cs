// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ResourceConnector.V20220415Preview.Outputs
{

    /// <summary>
    /// Appliance SSHKey definition.
    /// </summary>
    [OutputType]
    public sealed class SSHKeyResponse
    {
        /// <summary>
        /// User Private Key.
        /// </summary>
        public readonly string? PrivateKey;
        /// <summary>
        /// User Public Key.
        /// </summary>
        public readonly string? PublicKey;

        [OutputConstructor]
        private SSHKeyResponse(
            string? privateKey,

            string? publicKey)
        {
            PrivateKey = privateKey;
            PublicKey = publicKey;
        }
    }
}
