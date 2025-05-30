// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Outputs
{

    /// <summary>
    /// An Azure NetApp Files volume from Microsoft.NetApp provider
    /// </summary>
    [OutputType]
    public sealed class NetAppVolumeResponse
    {
        /// <summary>
        /// Azure resource ID of the NetApp volume
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private NetAppVolumeResponse(string id)
        {
            Id = id;
        }
    }
}
