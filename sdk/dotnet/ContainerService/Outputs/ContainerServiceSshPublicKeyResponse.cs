// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Contains information about SSH certificate public key data.
    /// </summary>
    [OutputType]
    public sealed class ContainerServiceSshPublicKeyResponse
    {
        /// <summary>
        /// Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        /// </summary>
        public readonly string KeyData;

        [OutputConstructor]
        private ContainerServiceSshPublicKeyResponse(string keyData)
        {
            KeyData = keyData;
        }
    }
}
