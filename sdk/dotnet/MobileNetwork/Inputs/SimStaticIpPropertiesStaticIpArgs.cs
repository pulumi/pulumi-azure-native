// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// The static IP configuration for the SIM to use at the defined network scope.
    /// </summary>
    public sealed class SimStaticIpPropertiesStaticIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv4 address assigned to the SIM at this network scope. This address must be in the userEquipmentStaticAddressPoolPrefix defined in the attached data network.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        public SimStaticIpPropertiesStaticIpArgs()
        {
        }
        public static new SimStaticIpPropertiesStaticIpArgs Empty => new SimStaticIpPropertiesStaticIpArgs();
    }
}
