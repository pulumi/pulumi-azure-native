// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Outputs
{

    /// <summary>
    /// The properties describing private cloud availability zone distribution
    /// </summary>
    [OutputType]
    public sealed class AvailabilityPropertiesResponse
    {
        /// <summary>
        /// The secondary availability zone for the private cloud
        /// </summary>
        public readonly int? SecondaryZone;
        /// <summary>
        /// The availability strategy for the private cloud
        /// </summary>
        public readonly string? Strategy;
        /// <summary>
        /// The primary availability zone for the private cloud
        /// </summary>
        public readonly int? Zone;

        [OutputConstructor]
        private AvailabilityPropertiesResponse(
            int? secondaryZone,

            string? strategy,

            int? zone)
        {
            SecondaryZone = secondaryZone;
            Strategy = strategy;
            Zone = zone;
        }
    }
}
