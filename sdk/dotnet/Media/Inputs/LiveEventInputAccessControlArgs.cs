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
    /// The IP access control for live event input.
    /// </summary>
    public sealed class LiveEventInputAccessControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP access control properties.
        /// </summary>
        [Input("ip")]
        public Input<Inputs.IPAccessControlArgs>? Ip { get; set; }

        public LiveEventInputAccessControlArgs()
        {
        }
        public static new LiveEventInputAccessControlArgs Empty => new LiveEventInputAccessControlArgs();
    }
}
