// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20250401Preview.Outputs
{

    /// <summary>
    /// PlacementProfile - Specifies the placement related settings for the virtual machine.
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineInstancePropertiesPlacementProfileResponse
    {
        /// <summary>
        /// Specifies whether VM can only failover strictly within the zone it was placed in
        /// </summary>
        public readonly bool? StrictPlacementPolicy;
        /// <summary>
        /// The zone in which the VM should be placed in.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private VirtualMachineInstancePropertiesPlacementProfileResponse(
            bool? strictPlacementPolicy,

            string? zone)
        {
            StrictPlacementPolicy = strictPlacementPolicy;
            Zone = zone;
        }
    }
}
