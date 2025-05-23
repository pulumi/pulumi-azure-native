// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTCentral.Outputs
{

    /// <summary>
    /// Network Rule Set Properties of this IoT Central application.
    /// </summary>
    [OutputType]
    public sealed class NetworkRuleSetsResponse
    {
        /// <summary>
        /// Whether these rules apply for device connectivity to IoT Hub and Device Provisioning service associated with this application.
        /// </summary>
        public readonly bool? ApplyToDevices;
        /// <summary>
        /// Whether these rules apply for connectivity via IoT Central web portal and APIs.
        /// </summary>
        public readonly bool? ApplyToIoTCentral;
        /// <summary>
        /// The default network action to apply.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// List of IP rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkRuleSetIpRuleResponse> IpRules;

        [OutputConstructor]
        private NetworkRuleSetsResponse(
            bool? applyToDevices,

            bool? applyToIoTCentral,

            string? defaultAction,

            ImmutableArray<Outputs.NetworkRuleSetIpRuleResponse> ipRules)
        {
            ApplyToDevices = applyToDevices;
            ApplyToIoTCentral = applyToIoTCentral;
            DefaultAction = defaultAction;
            IpRules = ipRules;
        }
    }
}
