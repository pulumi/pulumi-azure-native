// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20250401Preview.Inputs
{

    /// <summary>
    /// PlacementProfile - Specifies the placement related settings for the virtual machine.
    /// </summary>
    public sealed class VirtualMachineInstancePropertiesPlacementProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether VM can only failover strictly within the zone it was placed in
        /// </summary>
        [Input("strictPlacementPolicy")]
        public Input<bool>? StrictPlacementPolicy { get; set; }

        /// <summary>
        /// The zone in which the VM should be placed in.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VirtualMachineInstancePropertiesPlacementProfileArgs()
        {
        }
        public static new VirtualMachineInstancePropertiesPlacementProfileArgs Empty => new VirtualMachineInstancePropertiesPlacementProfileArgs();
    }
}
