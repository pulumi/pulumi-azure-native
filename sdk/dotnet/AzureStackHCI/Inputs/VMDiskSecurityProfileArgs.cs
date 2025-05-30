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
    /// Specifies the security profile settings for the managed disk. NOTE: It can only be set for Confidential VMs
    /// </summary>
    public sealed class VMDiskSecurityProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the EncryptionType of the managed disk. It is set to NonPersistedTPM for not persisting firmware state in the VMGuestState blob. NOTE: It can be set for only Confidential VMs.
        /// </summary>
        [Input("securityEncryptionType")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.SecurityEncryptionType>? SecurityEncryptionType { get; set; }

        public VMDiskSecurityProfileArgs()
        {
        }
        public static new VMDiskSecurityProfileArgs Empty => new VMDiskSecurityProfileArgs();
    }
}
