// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Gets or sets the central server configuration.
    /// </summary>
    [OutputType]
    public sealed class CentralServerConfigurationResponse
    {
        /// <summary>
        /// The number of central server VMs.
        /// </summary>
        public readonly double InstanceCount;
        /// <summary>
        /// The subnet id.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Gets or sets the virtual machine configuration.
        /// </summary>
        public readonly Outputs.VirtualMachineConfigurationResponse VirtualMachineConfiguration;

        [OutputConstructor]
        private CentralServerConfigurationResponse(
            double instanceCount,

            string subnetId,

            Outputs.VirtualMachineConfigurationResponse virtualMachineConfiguration)
        {
            InstanceCount = instanceCount;
            SubnetId = subnetId;
            VirtualMachineConfiguration = virtualMachineConfiguration;
        }
    }
}
