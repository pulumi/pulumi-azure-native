// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Inputs
{

    /// <summary>
    /// Class to specify which protocols are enabled
    /// </summary>
    public sealed class EnabledProtocolsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable DASH protocol or not
        /// </summary>
        [Input("dash", required: true)]
        public Input<bool> Dash { get; set; } = null!;

        /// <summary>
        /// Enable Download protocol or not
        /// </summary>
        [Input("download", required: true)]
        public Input<bool> Download { get; set; } = null!;

        /// <summary>
        /// Enable HLS protocol or not
        /// </summary>
        [Input("hls", required: true)]
        public Input<bool> Hls { get; set; } = null!;

        /// <summary>
        /// Enable SmoothStreaming protocol or not
        /// </summary>
        [Input("smoothStreaming", required: true)]
        public Input<bool> SmoothStreaming { get; set; } = null!;

        public EnabledProtocolsArgs()
        {
        }
        public static new EnabledProtocolsArgs Empty => new EnabledProtocolsArgs();
    }
}
