// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// Configuration to be used to setup the management network.
    /// </summary>
    public sealed class ManagementNetworkConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPN Configuration properties.
        /// </summary>
        [Input("infrastructureVpnConfiguration", required: true)]
        public Input<Inputs.VpnConfigurationPropertiesArgs> InfrastructureVpnConfiguration { get; set; } = null!;

        /// <summary>
        /// VPN Configuration properties.
        /// </summary>
        [Input("workloadVpnConfiguration", required: true)]
        public Input<Inputs.VpnConfigurationPropertiesArgs> WorkloadVpnConfiguration { get; set; } = null!;

        public ManagementNetworkConfigurationPropertiesArgs()
        {
        }
        public static new ManagementNetworkConfigurationPropertiesArgs Empty => new ManagementNetworkConfigurationPropertiesArgs();
    }
}
