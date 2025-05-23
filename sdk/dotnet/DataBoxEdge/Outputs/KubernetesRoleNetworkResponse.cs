// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Outputs
{

    /// <summary>
    /// Kubernetes role network resource
    /// </summary>
    [OutputType]
    public sealed class KubernetesRoleNetworkResponse
    {
        /// <summary>
        /// Cni configuration
        /// </summary>
        public readonly Outputs.CniConfigResponse CniConfig;
        /// <summary>
        /// Load balancer configuration
        /// </summary>
        public readonly Outputs.LoadBalancerConfigResponse LoadBalancerConfig;

        [OutputConstructor]
        private KubernetesRoleNetworkResponse(
            Outputs.CniConfigResponse cniConfig,

            Outputs.LoadBalancerConfigResponse loadBalancerConfig)
        {
            CniConfig = cniConfig;
            LoadBalancerConfig = loadBalancerConfig;
        }
    }
}
