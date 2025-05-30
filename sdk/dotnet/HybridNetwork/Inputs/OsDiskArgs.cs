// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Specifies information about the operating system disk used by the virtual machine. &lt;br&gt;&lt;br&gt; For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
    /// </summary>
    public sealed class OsDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the size of os disk in gigabytes. This is the fully expanded disk size needed of the VHD image on the ASE. This disk size should be greater than the size of the VHD provided in vhdUri.
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// The VHD name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OS type.
        /// </summary>
        [Input("osType")]
        public InputUnion<string, Pulumi.AzureNative.HybridNetwork.OperatingSystemTypes>? OsType { get; set; }

        /// <summary>
        /// The virtual hard disk.
        /// </summary>
        [Input("vhd")]
        public Input<Inputs.VirtualHardDiskArgs>? Vhd { get; set; }

        public OsDiskArgs()
        {
        }
        public static new OsDiskArgs Empty => new OsDiskArgs();
    }
}
