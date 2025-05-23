// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Advance configuration for AKS networking
    /// </summary>
    public sealed class AksNetworkingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
        /// </summary>
        [Input("dnsServiceIP")]
        public Input<string>? DnsServiceIP { get; set; }

        /// <summary>
        /// A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
        /// </summary>
        [Input("dockerBridgeCidr")]
        public Input<string>? DockerBridgeCidr { get; set; }

        /// <summary>
        /// A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        /// <summary>
        /// Virtual network subnet resource ID the compute nodes belong to
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public AksNetworkingConfigurationArgs()
        {
        }
        public static new AksNetworkingConfigurationArgs Empty => new AksNetworkingConfigurationArgs();
    }
}
