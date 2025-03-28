// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Search.V20240601Preview.Inputs
{

    /// <summary>
    /// Network specific rules that determine how the Azure AI Search service may be reached.
    /// </summary>
    public sealed class NetworkRuleSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Possible origins of inbound traffic that can bypass the rules defined in the 'ipRules' section.
        /// </summary>
        [Input("bypass")]
        public InputUnion<string, Pulumi.AzureNative.Search.V20240601Preview.SearchBypass>? Bypass { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.IpRuleArgs>? _ipRules;

        /// <summary>
        /// A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint. At the meantime, all other public IP networks are blocked by the firewall. These restriction rules are applied only when the 'publicNetworkAccess' of the search service is 'enabled'; otherwise, traffic over public interface is not allowed even with any public IP rules, and private endpoint connections would be the exclusive access method.
        /// </summary>
        public InputList<Inputs.IpRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.IpRuleArgs>());
            set => _ipRules = value;
        }

        public NetworkRuleSetArgs()
        {
        }
        public static new NetworkRuleSetArgs Empty => new NetworkRuleSetArgs();
    }
}
