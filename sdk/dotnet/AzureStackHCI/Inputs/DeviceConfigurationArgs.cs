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
    /// The device Configuration for edge device.
    /// </summary>
    public sealed class DeviceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device metadata details.
        /// </summary>
        [Input("deviceMetadata")]
        public Input<string>? DeviceMetadata { get; set; }

        [Input("nicDetails")]
        private InputList<Inputs.NicDetailArgs>? _nicDetails;

        /// <summary>
        /// NIC Details of device
        /// </summary>
        public InputList<Inputs.NicDetailArgs> NicDetails
        {
            get => _nicDetails ?? (_nicDetails = new InputList<Inputs.NicDetailArgs>());
            set => _nicDetails = value;
        }

        public DeviceConfigurationArgs()
        {
        }
        public static new DeviceConfigurationArgs Empty => new DeviceConfigurationArgs();
    }
}
