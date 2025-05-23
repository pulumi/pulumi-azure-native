// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// HardwareProfile - Specifies the hardware settings for the virtual machine.
    /// </summary>
    [OutputType]
    public sealed class VirtualMachinePropertiesResponseHardwareProfile
    {
        public readonly Outputs.VirtualMachinePropertiesResponseDynamicMemoryConfig? DynamicMemoryConfig;
        /// <summary>
        /// RAM in MB for the virtual machine
        /// </summary>
        public readonly double? MemoryMB;
        /// <summary>
        /// number of processors for the virtual machine
        /// </summary>
        public readonly int? Processors;
        public readonly string? VmSize;

        [OutputConstructor]
        private VirtualMachinePropertiesResponseHardwareProfile(
            Outputs.VirtualMachinePropertiesResponseDynamicMemoryConfig? dynamicMemoryConfig,

            double? memoryMB,

            int? processors,

            string? vmSize)
        {
            DynamicMemoryConfig = dynamicMemoryConfig;
            MemoryMB = memoryMB;
            Processors = processors;
            VmSize = vmSize;
        }
    }
}
