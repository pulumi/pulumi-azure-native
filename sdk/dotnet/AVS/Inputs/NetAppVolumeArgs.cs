// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Inputs
{

    /// <summary>
    /// An Azure NetApp Files volume from Microsoft.NetApp provider
    /// </summary>
    public sealed class NetAppVolumeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure resource ID of the NetApp volume
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public NetAppVolumeArgs()
        {
        }
        public static new NetAppVolumeArgs Empty => new NetAppVolumeArgs();
    }
}
