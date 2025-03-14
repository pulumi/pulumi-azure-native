// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20240202Preview.Inputs
{

    /// <summary>
    /// Header of otlp configuration
    /// </summary>
    public sealed class HeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of otlp configuration header
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value of otlp configuration header
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public HeaderArgs()
        {
        }
        public static new HeaderArgs Empty => new HeaderArgs();
    }
}
