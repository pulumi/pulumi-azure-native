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
    /// The parameters of a capacity reservation Profile.
    /// </summary>
    public sealed class CapacityReservationProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the capacity reservation group resource id that should be used for
        /// allocating the virtual machine or scaleset vm instances provided enough
        /// capacity has been reserved. Please refer to https://aka.ms/CapacityReservation
        /// for more details.
        /// </summary>
        [Input("capacityReservationGroup")]
        public Input<Inputs.SubResourceArgs>? CapacityReservationGroup { get; set; }

        public CapacityReservationProfileArgs()
        {
        }
        public static new CapacityReservationProfileArgs Empty => new CapacityReservationProfileArgs();
    }
}
