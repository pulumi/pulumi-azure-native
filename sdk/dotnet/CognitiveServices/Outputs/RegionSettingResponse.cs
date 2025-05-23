// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.Outputs
{

    /// <summary>
    /// The call rate limit Cognitive Services account.
    /// </summary>
    [OutputType]
    public sealed class RegionSettingResponse
    {
        /// <summary>
        /// Maps the region to the regional custom subdomain.
        /// </summary>
        public readonly string? Customsubdomain;
        /// <summary>
        /// Name of the region.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A value for priority or weighted routing methods.
        /// </summary>
        public readonly double? Value;

        [OutputConstructor]
        private RegionSettingResponse(
            string? customsubdomain,

            string? name,

            double? value)
        {
            Customsubdomain = customsubdomain;
            Name = name;
            Value = value;
        }
    }
}
