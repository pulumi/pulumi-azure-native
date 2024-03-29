// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20240101.Outputs
{

    /// <summary>
    /// SSH configuration for VMs of the provisioned cluster.
    /// </summary>
    [OutputType]
    public sealed class LinuxProfilePropertiesResponseSsh
    {
        /// <summary>
        /// The list of SSH public keys used to authenticate with VMs. A maximum of 1 key may be specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.LinuxProfilePropertiesResponsePublicKeys> PublicKeys;

        [OutputConstructor]
        private LinuxProfilePropertiesResponseSsh(ImmutableArray<Outputs.LinuxProfilePropertiesResponsePublicKeys> publicKeys)
        {
            PublicKeys = publicKeys;
        }
    }
}
