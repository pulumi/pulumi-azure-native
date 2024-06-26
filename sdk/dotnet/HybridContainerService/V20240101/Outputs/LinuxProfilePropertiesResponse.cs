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
    /// SSH profile for control plane and nodepool VMs of the provisioned cluster.
    /// </summary>
    [OutputType]
    public sealed class LinuxProfilePropertiesResponse
    {
        /// <summary>
        /// SSH configuration for VMs of the provisioned cluster.
        /// </summary>
        public readonly Outputs.LinuxProfilePropertiesResponseSsh? Ssh;

        [OutputConstructor]
        private LinuxProfilePropertiesResponse(Outputs.LinuxProfilePropertiesResponseSsh? ssh)
        {
            Ssh = ssh;
        }
    }
}
