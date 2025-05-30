// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    public sealed class ContainerAppsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of a subnet for control plane infrastructure components. This subnet must be in the same VNET as the subnet defined in appSubnetResourceId. Must not overlap with the IP range defined in platformReservedCidr, if defined.
        /// </summary>
        [Input("appSubnetResourceId")]
        public Input<string>? AppSubnetResourceId { get; set; }

        /// <summary>
        /// Resource ID of a subnet for control plane infrastructure components. This subnet must be in the same VNET as the subnet defined in appSubnetResourceId. Must not overlap with the IP range defined in platformReservedCidr, if defined.
        /// </summary>
        [Input("controlPlaneSubnetResourceId")]
        public Input<string>? ControlPlaneSubnetResourceId { get; set; }

        /// <summary>
        /// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
        /// </summary>
        [Input("daprAIInstrumentationKey")]
        public Input<string>? DaprAIInstrumentationKey { get; set; }

        /// <summary>
        /// CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the IP range defined in platformReservedCidr, if defined.
        /// </summary>
        [Input("dockerBridgeCidr")]
        public Input<string>? DockerBridgeCidr { get; set; }

        /// <summary>
        /// IP range in CIDR notation that can be reserved for environment infrastructure IP addresses. It must not overlap with any other Subnet IP ranges.
        /// </summary>
        [Input("platformReservedCidr")]
        public Input<string>? PlatformReservedCidr { get; set; }

        /// <summary>
        /// An IP address from the IP range defined by platformReservedCidr that will be reserved for the internal DNS server
        /// </summary>
        [Input("platformReservedDnsIP")]
        public Input<string>? PlatformReservedDnsIP { get; set; }

        public ContainerAppsConfigurationArgs()
        {
        }
        public static new ContainerAppsConfigurationArgs Empty => new ContainerAppsConfigurationArgs();
    }
}
