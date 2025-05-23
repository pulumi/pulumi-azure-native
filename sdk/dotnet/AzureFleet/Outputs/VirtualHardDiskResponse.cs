// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Describes the uri of a disk.
    /// </summary>
    [OutputType]
    public sealed class VirtualHardDiskResponse
    {
        /// <summary>
        /// Specifies the virtual hard disk's uri.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private VirtualHardDiskResponse(string? uri)
        {
            Uri = uri;
        }
    }
}
