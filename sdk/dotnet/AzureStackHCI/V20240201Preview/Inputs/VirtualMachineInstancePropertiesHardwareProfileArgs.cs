// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240201Preview.Inputs
{

    /// <summary>
    /// HardwareProfile - Specifies the hardware settings for the virtual machine instance.
    /// </summary>
    public sealed class VirtualMachineInstancePropertiesHardwareProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dynamic memory config
        /// </summary>
        [Input("dynamicMemoryConfig")]
        public Input<Inputs.VirtualMachineInstancePropertiesHardwareProfileDynamicMemoryConfigArgs>? DynamicMemoryConfig { get; set; }

        /// <summary>
        /// RAM in MB for the virtual machine instance
        /// </summary>
        [Input("memoryMB")]
        public Input<double>? MemoryMB { get; set; }

        /// <summary>
        /// number of processors for the virtual machine instance
        /// </summary>
        [Input("processors")]
        public Input<int>? Processors { get; set; }

        /// <summary>
        /// Enum of VM Sizes
        /// </summary>
        [Input("vmSize")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.V20240201Preview.VmSizeEnum>? VmSize { get; set; }

        public VirtualMachineInstancePropertiesHardwareProfileArgs()
        {
            VmSize = "Default";
        }
        public static new VirtualMachineInstancePropertiesHardwareProfileArgs Empty => new VirtualMachineInstancePropertiesHardwareProfileArgs();
    }
}
