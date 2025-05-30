// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Inputs
{

    /// <summary>
    /// Represents the profile for a single additional location in the Fleet. The location and the virtualMachineProfileOverride (optional).
    /// </summary>
    public sealed class LocationProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARM location name of the additional region. If LocationProfile is specified, then location is required.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// An override for computeProfile.baseVirtualMachineProfile specific to this region. 
        /// This override is merged with the base virtual machine profile to define the final virtual machine profile for the resources deployed in this location.
        /// </summary>
        [Input("virtualMachineProfileOverride")]
        public Input<Inputs.BaseVirtualMachineProfileArgs>? VirtualMachineProfileOverride { get; set; }

        public LocationProfileArgs()
        {
        }
        public static new LocationProfileArgs Empty => new LocationProfileArgs();
    }
}
