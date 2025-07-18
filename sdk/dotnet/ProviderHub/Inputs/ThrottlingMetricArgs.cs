// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class ThrottlingMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The interval.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// The limit.
        /// </summary>
        [Input("limit", required: true)]
        public Input<double> Limit { get; set; } = null!;

        /// <summary>
        /// The throttling metric type
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ProviderHub.ThrottlingMetricType> Type { get; set; } = null!;

        public ThrottlingMetricArgs()
        {
        }
        public static new ThrottlingMetricArgs Empty => new ThrottlingMetricArgs();
    }
}
