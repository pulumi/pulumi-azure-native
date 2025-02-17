// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview.Inputs
{

    /// <summary>
    /// Infra network profile for VMware platform
    /// </summary>
    public sealed class VirtualNetworkPropertiesVmwareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the network segment in VSphere
        /// </summary>
        [Input("segmentName")]
        public Input<string>? SegmentName { get; set; }

        public VirtualNetworkPropertiesVmwareArgs()
        {
        }
        public static new VirtualNetworkPropertiesVmwareArgs Empty => new VirtualNetworkPropertiesVmwareArgs();
    }
}
