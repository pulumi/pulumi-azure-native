// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Specifies the sku profile for the virtual machine scale set. With this property the customer is able to specify a list of VM sizes and an allocation strategy.
    /// </summary>
    [OutputType]
    public sealed class SkuProfileResponse
    {
        /// <summary>
        /// Specifies the allocation strategy for the virtual machine scale set based on which the VMs will be allocated.
        /// </summary>
        public readonly string? AllocationStrategy;
        /// <summary>
        /// Specifies the VM sizes for the virtual machine scale set.
        /// </summary>
        public readonly ImmutableArray<Outputs.SkuProfileVMSizeResponse> VmSizes;

        [OutputConstructor]
        private SkuProfileResponse(
            string? allocationStrategy,

            ImmutableArray<Outputs.SkuProfileVMSizeResponse> vmSizes)
        {
            AllocationStrategy = allocationStrategy;
            VmSizes = vmSizes;
        }
    }
}
