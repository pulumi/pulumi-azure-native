// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Outputs
{

    /// <summary>
    /// SSH - SSH configuration for Linux-based VMs running on Azure.
    /// </summary>
    [OutputType]
    public sealed class LinuxProfilePropertiesResponseSsh
    {
        /// <summary>
        /// PublicKeys - The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.LinuxProfilePropertiesResponsePublicKeys> PublicKeys;

        [OutputConstructor]
        private LinuxProfilePropertiesResponseSsh(ImmutableArray<Outputs.LinuxProfilePropertiesResponsePublicKeys> publicKeys)
        {
            PublicKeys = publicKeys;
        }
    }
}
