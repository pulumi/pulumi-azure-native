// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Inputs
{

    /// <summary>
    /// High availability properties of a flexible server.
    /// </summary>
    public sealed class HighAvailabilityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// High availability mode for a flexible server.
        /// </summary>
        [Input("mode")]
        public InputUnion<string, Pulumi.AzureNative.DBforPostgreSQL.HighAvailabilityMode>? Mode { get; set; }

        /// <summary>
        /// Availability zone associated to the standby server created when high availability is set to SameZone or ZoneRedundant.
        /// </summary>
        [Input("standbyAvailabilityZone")]
        public Input<string>? StandbyAvailabilityZone { get; set; }

        public HighAvailabilityArgs()
        {
            Mode = "Disabled";
            StandbyAvailabilityZone = "";
        }
        public static new HighAvailabilityArgs Empty => new HighAvailabilityArgs();
    }
}
