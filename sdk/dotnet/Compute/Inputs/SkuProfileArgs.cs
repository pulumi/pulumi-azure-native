// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Specifies the sku profile for the virtual machine scale set. With this property the customer is able to specify a list of VM sizes and an allocation strategy.
    /// </summary>
    public sealed class SkuProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the allocation strategy for the virtual machine scale set based on which the VMs will be allocated.
        /// </summary>
        [Input("allocationStrategy")]
        public InputUnion<string, Pulumi.AzureNative.Compute.AllocationStrategy>? AllocationStrategy { get; set; }

        [Input("vmSizes")]
        private InputList<Inputs.SkuProfileVMSizeArgs>? _vmSizes;

        /// <summary>
        /// Specifies the VM sizes for the virtual machine scale set.
        /// </summary>
        public InputList<Inputs.SkuProfileVMSizeArgs> VmSizes
        {
            get => _vmSizes ?? (_vmSizes = new InputList<Inputs.SkuProfileVMSizeArgs>());
            set => _vmSizes = value;
        }

        public SkuProfileArgs()
        {
        }
        public static new SkuProfileArgs Empty => new SkuProfileArgs();
    }
}
