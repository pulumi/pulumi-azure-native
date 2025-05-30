// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Inputs
{

    /// <summary>
    /// Kafka Token KeyVault properties.
    /// </summary>
    public sealed class KafkaTokenKeyVaultPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Username to connect with.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// KeyVault properties.
        /// </summary>
        [Input("vault", required: true)]
        public Input<Inputs.KeyVaultConnectionPropertiesArgs> Vault { get; set; } = null!;

        /// <summary>
        /// KeyVault secret details.
        /// </summary>
        [Input("vaultSecret", required: true)]
        public Input<Inputs.KeyVaultSecretObjectArgs> VaultSecret { get; set; } = null!;

        public KafkaTokenKeyVaultPropertiesArgs()
        {
        }
        public static new KafkaTokenKeyVaultPropertiesArgs Empty => new KafkaTokenKeyVaultPropertiesArgs();
    }
}
