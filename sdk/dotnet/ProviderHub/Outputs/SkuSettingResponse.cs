// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    [OutputType]
    public sealed class SkuSettingResponse
    {
        public readonly ImmutableArray<Outputs.SkuCapabilityResponse> Capabilities;
        public readonly Outputs.SkuSettingResponseCapacity? Capacity;
        public readonly ImmutableArray<Outputs.SkuCostResponse> Costs;
        public readonly string? Family;
        public readonly string? Kind;
        public readonly ImmutableArray<Outputs.SkuLocationInfoResponse> LocationInfo;
        public readonly ImmutableArray<string> Locations;
        public readonly string Name;
        public readonly ImmutableArray<string> RequiredFeatures;
        public readonly ImmutableArray<string> RequiredQuotaIds;
        public readonly string? Size;
        public readonly string? Tier;

        [OutputConstructor]
        private SkuSettingResponse(
            ImmutableArray<Outputs.SkuCapabilityResponse> capabilities,

            Outputs.SkuSettingResponseCapacity? capacity,

            ImmutableArray<Outputs.SkuCostResponse> costs,

            string? family,

            string? kind,

            ImmutableArray<Outputs.SkuLocationInfoResponse> locationInfo,

            ImmutableArray<string> locations,

            string name,

            ImmutableArray<string> requiredFeatures,

            ImmutableArray<string> requiredQuotaIds,

            string? size,

            string? tier)
        {
            Capabilities = capabilities;
            Capacity = capacity;
            Costs = costs;
            Family = family;
            Kind = kind;
            LocationInfo = locationInfo;
            Locations = locations;
            Name = name;
            RequiredFeatures = requiredFeatures;
            RequiredQuotaIds = requiredQuotaIds;
            Size = size;
            Tier = tier;
        }
    }
}
