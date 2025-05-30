// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.Outputs
{

    [OutputType]
    public sealed class ThrottlingRuleResponse
    {
        public readonly double? Count;
        public readonly bool? DynamicThrottlingEnabled;
        public readonly string? Key;
        public readonly ImmutableArray<Outputs.RequestMatchPatternResponse> MatchPatterns;
        public readonly double? MinCount;
        public readonly double? RenewalPeriod;

        [OutputConstructor]
        private ThrottlingRuleResponse(
            double? count,

            bool? dynamicThrottlingEnabled,

            string? key,

            ImmutableArray<Outputs.RequestMatchPatternResponse> matchPatterns,

            double? minCount,

            double? renewalPeriod)
        {
            Count = count;
            DynamicThrottlingEnabled = dynamicThrottlingEnabled;
            Key = key;
            MatchPatterns = matchPatterns;
            MinCount = minCount;
            RenewalPeriod = renewalPeriod;
        }
    }
}
