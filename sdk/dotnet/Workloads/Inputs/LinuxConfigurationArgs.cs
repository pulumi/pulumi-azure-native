// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Inputs
{

    /// <summary>
    /// Specifies the Linux operating system settings on the virtual machine. For a list of supported Linux distributions, see [Linux on Azure-Endorsed Distributions](https://learn.microsoft.com/azure/virtual-machines/linux/endorsed-distros).
    /// </summary>
    public sealed class LinuxConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether password authentication should be disabled.
        /// </summary>
        [Input("disablePasswordAuthentication")]
        public Input<bool>? DisablePasswordAuthentication { get; set; }

        /// <summary>
        /// The OS Type
        /// Expected value is 'Linux'.
        /// </summary>
        [Input("osType", required: true)]
        public Input<string> OsType { get; set; } = null!;

        /// <summary>
        /// Specifies the ssh key configuration for a Linux OS. (This property is deprecated, please use 'sshKeyPair' instead)
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.SshConfigurationArgs>? Ssh { get; set; }

        /// <summary>
        /// The SSH Key-pair used to authenticate with the VM's.
        /// </summary>
        [Input("sshKeyPair")]
        public Input<Inputs.SshKeyPairArgs>? SshKeyPair { get; set; }

        public LinuxConfigurationArgs()
        {
        }
        public static new LinuxConfigurationArgs Empty => new LinuxConfigurationArgs();
    }
}
