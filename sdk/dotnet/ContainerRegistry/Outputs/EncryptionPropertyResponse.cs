// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    [OutputType]
    public sealed class EncryptionPropertyResponse
    {
        /// <summary>
        /// Key vault properties.
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponse? KeyVaultProperties;
        /// <summary>
        /// Indicates whether or not the encryption is enabled for container registry.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private EncryptionPropertyResponse(
            Outputs.KeyVaultPropertiesResponse? keyVaultProperties,

            string? status)
        {
            KeyVaultProperties = keyVaultProperties;
            Status = status;
        }
    }
}
