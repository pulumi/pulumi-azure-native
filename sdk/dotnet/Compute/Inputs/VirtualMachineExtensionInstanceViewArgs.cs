// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// The instance view of a virtual machine extension.
    /// </summary>
    public sealed class VirtualMachineExtensionInstanceViewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The virtual machine extension name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("statuses")]
        private InputList<Inputs.InstanceViewStatusArgs>? _statuses;

        /// <summary>
        /// The resource status information.
        /// </summary>
        public InputList<Inputs.InstanceViewStatusArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.InstanceViewStatusArgs>());
            set => _statuses = value;
        }

        [Input("substatuses")]
        private InputList<Inputs.InstanceViewStatusArgs>? _substatuses;

        /// <summary>
        /// The resource status information.
        /// </summary>
        public InputList<Inputs.InstanceViewStatusArgs> Substatuses
        {
            get => _substatuses ?? (_substatuses = new InputList<Inputs.InstanceViewStatusArgs>());
            set => _substatuses = value;
        }

        /// <summary>
        /// Specifies the type of the extension; an example is "CustomScriptExtension".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        [Input("typeHandlerVersion")]
        public Input<string>? TypeHandlerVersion { get; set; }

        public VirtualMachineExtensionInstanceViewArgs()
        {
        }
        public static new VirtualMachineExtensionInstanceViewArgs Empty => new VirtualMachineExtensionInstanceViewArgs();
    }
}
