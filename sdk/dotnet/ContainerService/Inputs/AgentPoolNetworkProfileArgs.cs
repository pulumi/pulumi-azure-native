// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Network settings of an agent pool.
    /// </summary>
    public sealed class AgentPoolNetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedHostPorts")]
        private InputList<Inputs.PortRangeArgs>? _allowedHostPorts;

        /// <summary>
        /// The port ranges that are allowed to access. The specified ranges are allowed to overlap.
        /// </summary>
        public InputList<Inputs.PortRangeArgs> AllowedHostPorts
        {
            get => _allowedHostPorts ?? (_allowedHostPorts = new InputList<Inputs.PortRangeArgs>());
            set => _allowedHostPorts = value;
        }

        [Input("applicationSecurityGroups")]
        private InputList<string>? _applicationSecurityGroups;

        /// <summary>
        /// The IDs of the application security groups which agent pool will associate when created.
        /// </summary>
        public InputList<string> ApplicationSecurityGroups
        {
            get => _applicationSecurityGroups ?? (_applicationSecurityGroups = new InputList<string>());
            set => _applicationSecurityGroups = value;
        }

        [Input("nodePublicIPTags")]
        private InputList<Inputs.IPTagArgs>? _nodePublicIPTags;

        /// <summary>
        /// IPTags of instance-level public IPs.
        /// </summary>
        public InputList<Inputs.IPTagArgs> NodePublicIPTags
        {
            get => _nodePublicIPTags ?? (_nodePublicIPTags = new InputList<Inputs.IPTagArgs>());
            set => _nodePublicIPTags = value;
        }

        public AgentPoolNetworkProfileArgs()
        {
        }
        public static new AgentPoolNetworkProfileArgs Empty => new AgentPoolNetworkProfileArgs();
    }
}
