// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Hub Item.
    /// </summary>
    public sealed class HubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Id.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Resource Type.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public HubArgs()
        {
        }
        public static new HubArgs Empty => new HubArgs();
    }
}
