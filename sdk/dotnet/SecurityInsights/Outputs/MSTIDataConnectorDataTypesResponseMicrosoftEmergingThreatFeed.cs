// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// Data type for Microsoft Threat Intelligence data connector.
    /// </summary>
    [OutputType]
    public sealed class MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed
    {
        /// <summary>
        /// The lookback period for the feed to be imported. The date-time to begin importing the feed from, for example: 2024-01-01T00:00:00.000Z.
        /// </summary>
        public readonly string LookbackPeriod;
        /// <summary>
        /// Describe whether this data type connection is enabled or not.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed(
            string lookbackPeriod,

            string? state)
        {
            LookbackPeriod = lookbackPeriod;
            State = state;
        }
    }
}
