// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20250101.Inputs
{

    /// <summary>
    /// GPU settings for the Agent Pool.
    /// </summary>
    public sealed class GPUProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to install GPU drivers. When it's not specified, default is Install.
        /// </summary>
        [Input("driver")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20250101.GPUDriver>? Driver { get; set; }

        public GPUProfileArgs()
        {
        }
        public static new GPUProfileArgs Empty => new GPUProfileArgs();
    }
}
