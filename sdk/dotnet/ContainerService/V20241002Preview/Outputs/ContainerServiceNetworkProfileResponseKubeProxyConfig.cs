// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Outputs
{

    /// <summary>
    /// Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy defaulting behavior. See https://v&lt;version&gt;.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/ where &lt;version&gt; is represented by a &lt;major version&gt;-&lt;minor version&gt; string. Kubernetes version 1.23 would be '1-23'.
    /// </summary>
    [OutputType]
    public sealed class ContainerServiceNetworkProfileResponseKubeProxyConfig
    {
        /// <summary>
        /// Whether to enable on kube-proxy on the cluster (if no 'kubeProxyConfig' exists, kube-proxy is enabled in AKS by default without these customizations).
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Holds configuration customizations for IPVS. May only be specified if 'mode' is set to 'IPVS'.
        /// </summary>
        public readonly Outputs.ContainerServiceNetworkProfileResponseIpvsConfig? IpvsConfig;
        /// <summary>
        /// Specify which proxy mode to use ('IPTABLES' or 'IPVS')
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private ContainerServiceNetworkProfileResponseKubeProxyConfig(
            bool? enabled,

            Outputs.ContainerServiceNetworkProfileResponseIpvsConfig? ipvsConfig,

            string? mode)
        {
            Enabled = enabled;
            IpvsConfig = ipvsConfig;
            Mode = mode;
        }
    }
}
