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
    /// KeyVault secret object properties
    /// </summary>
    public sealed class KeyVaultSecretObjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// KeyVault secret name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// KeyVault secret version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KeyVaultSecretObjectArgs()
        {
        }
        public static new KeyVaultSecretObjectArgs Empty => new KeyVaultSecretObjectArgs();
    }
}
