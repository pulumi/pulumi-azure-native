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
    /// Details on the total up-time for the VM.
    /// </summary>
    public sealed class VmUptimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days in a month for VM uptime.
        /// </summary>
        [Input("daysPerMonth")]
        public Input<double>? DaysPerMonth { get; set; }

        /// <summary>
        /// Number of hours per day for VM uptime.
        /// </summary>
        [Input("hoursPerDay")]
        public Input<double>? HoursPerDay { get; set; }

        public VmUptimeArgs()
        {
        }
        public static new VmUptimeArgs Empty => new VmUptimeArgs();
    }
}
