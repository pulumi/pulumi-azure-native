// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TimeSeriesInsights.Inputs
{

    /// <summary>
    /// The structure of the property that a time series id can have. An environment can have multiple such properties.
    /// </summary>
    public sealed class TimeSeriesIdPropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the property.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the property.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.TimeSeriesInsights.PropertyType>? Type { get; set; }

        public TimeSeriesIdPropertyArgs()
        {
        }
        public static new TimeSeriesIdPropertyArgs Empty => new TimeSeriesIdPropertyArgs();
    }
}
