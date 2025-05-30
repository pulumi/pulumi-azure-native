// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Entity Uptime.
    /// </summary>
    public sealed class EntityUptimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets the days per month.
        /// </summary>
        [Input("daysPerMonth")]
        public Input<int>? DaysPerMonth { get; set; }

        /// <summary>
        /// Gets the hours per day.
        /// </summary>
        [Input("hoursPerDay")]
        public Input<int>? HoursPerDay { get; set; }

        public EntityUptimeArgs()
        {
        }
        public static new EntityUptimeArgs Empty => new EntityUptimeArgs();
    }
}
