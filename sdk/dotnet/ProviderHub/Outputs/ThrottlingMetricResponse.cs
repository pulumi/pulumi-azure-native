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
    public sealed class ThrottlingMetricResponse
    {
        /// <summary>
        /// The interval.
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// The limit.
        /// </summary>
        public readonly double Limit;
        /// <summary>
        /// The throttling metric type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ThrottlingMetricResponse(
            string? interval,

            double limit,

            string type)
        {
            Interval = interval;
            Limit = limit;
            Type = type;
        }
    }
}
