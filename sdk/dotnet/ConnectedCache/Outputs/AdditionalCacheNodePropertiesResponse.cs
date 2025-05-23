// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedCache.Outputs
{

    /// <summary>
    /// Model representing cache node for connected cache resource
    /// </summary>
    [OutputType]
    public sealed class AdditionalCacheNodePropertiesResponse
    {
        /// <summary>
        /// Cache node resource aggregated status code.
        /// </summary>
        public readonly int AggregatedStatusCode;
        /// <summary>
        /// Cache node resource aggregated status details.
        /// </summary>
        public readonly string AggregatedStatusDetails;
        /// <summary>
        /// Cache node resource aggregated status text.
        /// </summary>
        public readonly string AggregatedStatusText;
        /// <summary>
        /// Auto update version that is the applied to update on mcc cache node
        /// </summary>
        public readonly string AutoUpdateAppliedVersion;
        /// <summary>
        /// Auto update last applied date time of mcc install
        /// </summary>
        public readonly string AutoUpdateLastAppliedDateTime;
        /// <summary>
        /// Auto Update status details from the backend after applying the new version details
        /// </summary>
        public readonly string AutoUpdateLastAppliedDetails;
        /// <summary>
        /// Last applied auto update state for mcc install of auto update cycle
        /// </summary>
        public readonly string AutoUpdateLastAppliedState;
        /// <summary>
        /// Auto update last triggered date time of mcc install
        /// </summary>
        public readonly string AutoUpdateLastTriggeredDateTime;
        /// <summary>
        /// Auto update last applied date time of mcc install
        /// </summary>
        public readonly string AutoUpdateNextAvailableDateTime;
        /// <summary>
        /// Auto update version that is the Next available version to update on mcc cache node
        /// </summary>
        public readonly string AutoUpdateNextAvailableVersion;
        /// <summary>
        /// Auto update or fast update version
        /// </summary>
        public readonly string? AutoUpdateVersion;
        /// <summary>
        /// Cache node resource Bgp configuration.
        /// </summary>
        public readonly Outputs.BgpConfigurationResponse? BgpConfiguration;
        /// <summary>
        /// issues list to return the issues as part of the additional cache node properties
        /// </summary>
        public readonly ImmutableArray<string> CacheNodePropertiesDetailsIssuesList;
        /// <summary>
        /// Cache node resource state as integer.
        /// </summary>
        public readonly int CacheNodeState;
        /// <summary>
        /// Cache node resource detailed state text.
        /// </summary>
        public readonly string CacheNodeStateDetailedText;
        /// <summary>
        /// Cache node resource short state text.
        /// </summary>
        public readonly string CacheNodeStateShortText;
        /// <summary>
        /// Cache node resource drive configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.CacheNodeDriveConfigurationResponse> DriveConfiguration;
        /// <summary>
        /// Cache node resource flag indicating if cache node has been physically installed or provisioned on their physical lab.
        /// </summary>
        public readonly bool IsProvisioned;
        /// <summary>
        /// Cache node resource requires a proxy
        /// </summary>
        public readonly string? IsProxyRequired;
        /// <summary>
        /// Optional property #1 of Mcc response object
        /// </summary>
        public readonly string? OptionalProperty1;
        /// <summary>
        /// Optional property #2 of Mcc response object
        /// </summary>
        public readonly string? OptionalProperty2;
        /// <summary>
        /// Optional property #3 of Mcc response object
        /// </summary>
        public readonly string? OptionalProperty3;
        /// <summary>
        /// Optional property #4 of Mcc response object
        /// </summary>
        public readonly string? OptionalProperty4;
        /// <summary>
        /// Optional property #5 of Mcc response object
        /// </summary>
        public readonly string? OptionalProperty5;
        /// <summary>
        /// Operating system of the cache node
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// Cache node resource Mcc product version.
        /// </summary>
        public readonly string ProductVersion;
        /// <summary>
        /// Cache node resource Mcc proxy Url
        /// </summary>
        public readonly string? ProxyUrl;
        /// <summary>
        /// proxyUrl configuration of the cache node
        /// </summary>
        public readonly Outputs.ProxyUrlConfigurationResponse? ProxyUrlConfiguration;
        /// <summary>
        /// Update Cycle Type
        /// </summary>
        public readonly string? UpdateCycleType;
        /// <summary>
        /// Update related information details
        /// </summary>
        public readonly string? UpdateInfoDetails;
        /// <summary>
        /// customer requested date time for mcc install of update cycle
        /// </summary>
        public readonly string? UpdateRequestedDateTime;

        [OutputConstructor]
        private AdditionalCacheNodePropertiesResponse(
            int aggregatedStatusCode,

            string aggregatedStatusDetails,

            string aggregatedStatusText,

            string autoUpdateAppliedVersion,

            string autoUpdateLastAppliedDateTime,

            string autoUpdateLastAppliedDetails,

            string autoUpdateLastAppliedState,

            string autoUpdateLastTriggeredDateTime,

            string autoUpdateNextAvailableDateTime,

            string autoUpdateNextAvailableVersion,

            string? autoUpdateVersion,

            Outputs.BgpConfigurationResponse? bgpConfiguration,

            ImmutableArray<string> cacheNodePropertiesDetailsIssuesList,

            int cacheNodeState,

            string cacheNodeStateDetailedText,

            string cacheNodeStateShortText,

            ImmutableArray<Outputs.CacheNodeDriveConfigurationResponse> driveConfiguration,

            bool isProvisioned,

            string? isProxyRequired,

            string? optionalProperty1,

            string? optionalProperty2,

            string? optionalProperty3,

            string? optionalProperty4,

            string? optionalProperty5,

            string? osType,

            string productVersion,

            string? proxyUrl,

            Outputs.ProxyUrlConfigurationResponse? proxyUrlConfiguration,

            string? updateCycleType,

            string? updateInfoDetails,

            string? updateRequestedDateTime)
        {
            AggregatedStatusCode = aggregatedStatusCode;
            AggregatedStatusDetails = aggregatedStatusDetails;
            AggregatedStatusText = aggregatedStatusText;
            AutoUpdateAppliedVersion = autoUpdateAppliedVersion;
            AutoUpdateLastAppliedDateTime = autoUpdateLastAppliedDateTime;
            AutoUpdateLastAppliedDetails = autoUpdateLastAppliedDetails;
            AutoUpdateLastAppliedState = autoUpdateLastAppliedState;
            AutoUpdateLastTriggeredDateTime = autoUpdateLastTriggeredDateTime;
            AutoUpdateNextAvailableDateTime = autoUpdateNextAvailableDateTime;
            AutoUpdateNextAvailableVersion = autoUpdateNextAvailableVersion;
            AutoUpdateVersion = autoUpdateVersion;
            BgpConfiguration = bgpConfiguration;
            CacheNodePropertiesDetailsIssuesList = cacheNodePropertiesDetailsIssuesList;
            CacheNodeState = cacheNodeState;
            CacheNodeStateDetailedText = cacheNodeStateDetailedText;
            CacheNodeStateShortText = cacheNodeStateShortText;
            DriveConfiguration = driveConfiguration;
            IsProvisioned = isProvisioned;
            IsProxyRequired = isProxyRequired;
            OptionalProperty1 = optionalProperty1;
            OptionalProperty2 = optionalProperty2;
            OptionalProperty3 = optionalProperty3;
            OptionalProperty4 = optionalProperty4;
            OptionalProperty5 = optionalProperty5;
            OsType = osType;
            ProductVersion = productVersion;
            ProxyUrl = proxyUrl;
            ProxyUrlConfiguration = proxyUrlConfiguration;
            UpdateCycleType = updateCycleType;
            UpdateInfoDetails = updateInfoDetails;
            UpdateRequestedDateTime = updateRequestedDateTime;
        }
    }
}
