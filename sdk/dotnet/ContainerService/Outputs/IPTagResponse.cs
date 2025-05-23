// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Contains the IPTag associated with the object.
    /// </summary>
    [OutputType]
    public sealed class IPTagResponse
    {
        /// <summary>
        /// The IP tag type. Example: RoutingPreference.
        /// </summary>
        public readonly string? IpTagType;
        /// <summary>
        /// The value of the IP tag associated with the public IP. Example: Internet.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private IPTagResponse(
            string? ipTagType,

            string? tag)
        {
            IpTagType = ipTagType;
            Tag = tag;
        }
    }
}
