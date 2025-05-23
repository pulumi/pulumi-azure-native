// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// LinuxConfiguration - linux specific configuration values for the virtual machine instance
    /// </summary>
    public sealed class VirtualMachineInstancePropertiesOsProfileLinuxConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DisablePasswordAuthentication - whether password authentication should be disabled
        /// </summary>
        [Input("disablePasswordAuthentication")]
        public Input<bool>? DisablePasswordAuthentication { get; set; }

        /// <summary>
        /// Used to indicate whether Arc for Servers agent onboarding should be triggered during the virtual machine instance creation process.
        /// </summary>
        [Input("provisionVMAgent")]
        public Input<bool>? ProvisionVMAgent { get; set; }

        /// <summary>
        /// Used to indicate whether the VM Config Agent should be installed during the virtual machine creation process.
        /// </summary>
        [Input("provisionVMConfigAgent")]
        public Input<bool>? ProvisionVMConfigAgent { get; set; }

        /// <summary>
        /// Specifies the ssh key configuration for a Linux OS.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.SshConfigurationArgs>? Ssh { get; set; }

        public VirtualMachineInstancePropertiesOsProfileLinuxConfigurationArgs()
        {
            ProvisionVMAgent = true;
            ProvisionVMConfigAgent = true;
        }
        public static new VirtualMachineInstancePropertiesOsProfileLinuxConfigurationArgs Empty => new VirtualMachineInstancePropertiesOsProfileLinuxConfigurationArgs();
    }
}
