// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.V20240201.Inputs
{

    /// <summary>
    /// Signaling configuration for the packet core.
    /// </summary>
    public sealed class SignalingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration enabling 4G NAS reroute.
        /// </summary>
        [Input("nasReroute")]
        public Input<Inputs.NASRerouteConfigurationArgs>? NasReroute { get; set; }

        public SignalingConfigurationArgs()
        {
        }
        public static new SignalingConfigurationArgs Empty => new SignalingConfigurationArgs();
    }
}
