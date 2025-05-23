// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// Single-network slice selection assistance information (S-NSSAI).
    /// </summary>
    public sealed class SnssaiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Slice differentiator (SD).
        /// </summary>
        [Input("sd")]
        public Input<string>? Sd { get; set; }

        /// <summary>
        /// Slice/service type (SST).
        /// </summary>
        [Input("sst", required: true)]
        public Input<int> Sst { get; set; } = null!;

        public SnssaiArgs()
        {
        }
        public static new SnssaiArgs Empty => new SnssaiArgs();
    }
}
