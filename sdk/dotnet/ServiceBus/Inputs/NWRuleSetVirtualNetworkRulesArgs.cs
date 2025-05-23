// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceBus.Inputs
{

    /// <summary>
    /// Description of VirtualNetworkRules - NetworkRules resource.
    /// </summary>
    public sealed class NWRuleSetVirtualNetworkRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value that indicates whether to ignore missing VNet Service Endpoint
        /// </summary>
        [Input("ignoreMissingVnetServiceEndpoint")]
        public Input<bool>? IgnoreMissingVnetServiceEndpoint { get; set; }

        /// <summary>
        /// Subnet properties
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubnetArgs>? Subnet { get; set; }

        public NWRuleSetVirtualNetworkRulesArgs()
        {
        }
        public static new NWRuleSetVirtualNetworkRulesArgs Empty => new NWRuleSetVirtualNetworkRulesArgs();
    }
}
