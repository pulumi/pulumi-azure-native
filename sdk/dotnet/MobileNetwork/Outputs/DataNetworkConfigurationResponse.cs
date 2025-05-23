// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Outputs
{

    /// <summary>
    /// Settings controlling data network use
    /// </summary>
    [OutputType]
    public sealed class DataNetworkConfigurationResponse
    {
        /// <summary>
        /// Allowed session types in addition to the default session type. Must not duplicate the default session type.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalAllowedSessionTypes;
        /// <summary>
        /// Default QoS Flow allocation and retention priority (ARP) level. Flows with higher priority preempt flows with lower priority, if the settings of `preemptionCapability` and `preemptionVulnerability` allow it. 1 is the highest level of priority. If this field is not specified then `5qi` is used to derive the ARP value. See 3GPP TS23.501 section 5.7.2.2 for a full description of the ARP parameters.
        /// </summary>
        public readonly int? AllocationAndRetentionPriorityLevel;
        /// <summary>
        /// List of services that can be used as part of this SIM policy. The list must not contain duplicate items and must contain at least one item. The services must be in the same location as the SIM policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceResourceIdResponse> AllowedServices;
        /// <summary>
        /// A reference to the data network that these settings apply to. The data network must be in the same location as the SIM policy.
        /// </summary>
        public readonly Outputs.DataNetworkResourceIdResponse DataNetwork;
        /// <summary>
        /// The default PDU session type, which is used if the UE does not request a specific session type.
        /// </summary>
        public readonly string? DefaultSessionType;
        /// <summary>
        /// Default 5G QoS Flow Indicator value. The 5QI identifies a specific QoS forwarding treatment to be provided to a flow. See 3GPP TS23.501 section 5.7.2.1 for a full description of the 5QI parameter, and table 5.7.4-1 for the definition the 5QI values.
        /// </summary>
        public readonly int? FiveQi;
        /// <summary>
        /// The maximum number of downlink packets to buffer at the user plane for High Latency Communication - Extended Buffering. See 3GPP TS29.272 v15.10.0 section 7.3.188 for a full description. This maximum is not guaranteed because there is a internal limit on buffered packets across all PDU sessions.
        /// </summary>
        public readonly int? MaximumNumberOfBufferedPackets;
        /// <summary>
        /// Default QoS Flow preemption capability. The preemption capability of a QoS Flow controls whether it can preempt another QoS Flow with a lower priority level. See 3GPP TS23.501 section 5.7.2.2 for a full description of the ARP parameters.
        /// </summary>
        public readonly string? PreemptionCapability;
        /// <summary>
        /// Default QoS Flow preemption vulnerability. The preemption vulnerability of a QoS Flow controls whether it can be preempted by a QoS Flow with a higher priority level. See 3GPP TS23.501 section 5.7.2.2 for a full description of the ARP parameters.
        /// </summary>
        public readonly string? PreemptionVulnerability;
        /// <summary>
        /// Aggregate maximum bit rate across all non-GBR QoS flows of a given PDU session. See 3GPP TS23.501 section 5.7.2.6 for a full description of the Session-AMBR.
        /// </summary>
        public readonly Outputs.AmbrResponse SessionAmbr;

        [OutputConstructor]
        private DataNetworkConfigurationResponse(
            ImmutableArray<string> additionalAllowedSessionTypes,

            int? allocationAndRetentionPriorityLevel,

            ImmutableArray<Outputs.ServiceResourceIdResponse> allowedServices,

            Outputs.DataNetworkResourceIdResponse dataNetwork,

            string? defaultSessionType,

            int? fiveQi,

            int? maximumNumberOfBufferedPackets,

            string? preemptionCapability,

            string? preemptionVulnerability,

            Outputs.AmbrResponse sessionAmbr)
        {
            AdditionalAllowedSessionTypes = additionalAllowedSessionTypes;
            AllocationAndRetentionPriorityLevel = allocationAndRetentionPriorityLevel;
            AllowedServices = allowedServices;
            DataNetwork = dataNetwork;
            DefaultSessionType = defaultSessionType;
            FiveQi = fiveQi;
            MaximumNumberOfBufferedPackets = maximumNumberOfBufferedPackets;
            PreemptionCapability = preemptionCapability;
            PreemptionVulnerability = preemptionVulnerability;
            SessionAmbr = sessionAmbr;
        }
    }
}
