// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The properties of the archive package source.
    /// </summary>
    [OutputType]
    public sealed class ArchivePackageSourcePropertiesResponse
    {
        /// <summary>
        /// The type of package source for a archive.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The external repository url.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private ArchivePackageSourcePropertiesResponse(
            string? type,

            string? url)
        {
            Type = type;
            Url = url;
        }
    }
}
