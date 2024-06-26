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
    /// IP Address of the Kubernetes API server
    /// </summary>
    [OutputType]
    public sealed class ControlPlaneProfileResponseControlPlaneEndpoint
    {
        /// <summary>
        /// IP address of the Kubernetes API server
        /// </summary>
        public readonly string? HostIP;

        [OutputConstructor]
        private ControlPlaneProfileResponseControlPlaneEndpoint(string? hostIP)
        {
            HostIP = hostIP;
        }
    }
}
