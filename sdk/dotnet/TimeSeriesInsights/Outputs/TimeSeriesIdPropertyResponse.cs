// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TimeSeriesInsights.Outputs
{

    /// <summary>
    /// The structure of the property that a time series id can have. An environment can have multiple such properties.
    /// </summary>
    [OutputType]
    public sealed class TimeSeriesIdPropertyResponse
    {
        /// <summary>
        /// The name of the property.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of the property.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private TimeSeriesIdPropertyResponse(
            string? name,

            string? type)
        {
            Name = name;
            Type = type;
        }
    }
}
