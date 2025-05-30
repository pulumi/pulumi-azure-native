// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Specifies policy and settings for SSH access.
    /// </summary>
    public sealed class ComputeInstanceSshSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH key pairs.
        /// </summary>
        [Input("adminPublicKey")]
        public Input<string>? AdminPublicKey { get; set; }

        /// <summary>
        /// State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is closed on this instance. Enabled - Indicates that the public ssh port is open and accessible according to the VNet/subnet policy if applicable.
        /// </summary>
        [Input("sshPublicAccess")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.SshPublicAccess>? SshPublicAccess { get; set; }

        public ComputeInstanceSshSettingsArgs()
        {
            SshPublicAccess = "Disabled";
        }
        public static new ComputeInstanceSshSettingsArgs Empty => new ComputeInstanceSshSettingsArgs();
    }
}
