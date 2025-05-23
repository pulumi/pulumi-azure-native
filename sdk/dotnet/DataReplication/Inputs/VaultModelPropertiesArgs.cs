// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Inputs
{

    /// <summary>
    /// Vault properties.
    /// </summary>
    public sealed class VaultModelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the type of vault.
        /// </summary>
        [Input("vaultType")]
        public InputUnion<string, Pulumi.AzureNative.DataReplication.ReplicationVaultType>? VaultType { get; set; }

        public VaultModelPropertiesArgs()
        {
        }
        public static new VaultModelPropertiesArgs Empty => new VaultModelPropertiesArgs();
    }
}
