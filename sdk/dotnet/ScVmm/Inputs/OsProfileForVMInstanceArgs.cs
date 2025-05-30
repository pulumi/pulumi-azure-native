// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm.Inputs
{

    /// <summary>
    /// Defines the resource properties.
    /// </summary>
    public sealed class OsProfileForVMInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin password of the virtual machine.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        /// <summary>
        /// Gets or sets computer name.
        /// </summary>
        [Input("computerName")]
        public Input<string>? ComputerName { get; set; }

        public OsProfileForVMInstanceArgs()
        {
        }
        public static new OsProfileForVMInstanceArgs Empty => new OsProfileForVMInstanceArgs();
    }
}
