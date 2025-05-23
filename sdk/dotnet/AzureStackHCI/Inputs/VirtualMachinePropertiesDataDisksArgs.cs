// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    public sealed class VirtualMachinePropertiesDataDisksArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of the data disk
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public VirtualMachinePropertiesDataDisksArgs()
        {
        }
        public static new VirtualMachinePropertiesDataDisksArgs Empty => new VirtualMachinePropertiesDataDisksArgs();
    }
}
