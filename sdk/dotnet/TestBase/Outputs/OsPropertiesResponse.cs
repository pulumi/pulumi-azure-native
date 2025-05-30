// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase.Outputs
{

    /// <summary>
    /// The properties of an operating system.
    /// </summary>
    [OutputType]
    public sealed class OsPropertiesResponse
    {
        /// <summary>
        /// The name of the custom image resource.
        /// </summary>
        public readonly string CustomImageDisplayName;
        /// <summary>
        /// Specify the referenced Test Base Custom Image Id if available.
        /// </summary>
        public readonly string? CustomImageId;
        /// <summary>
        /// The name of the OS.
        /// </summary>
        public readonly string? OsName;
        /// <summary>
        /// The properties of the OS release.
        /// </summary>
        public readonly Outputs.ReleasePropertiesResponse? ReleaseProperties;

        [OutputConstructor]
        private OsPropertiesResponse(
            string customImageDisplayName,

            string? customImageId,

            string? osName,

            Outputs.ReleasePropertiesResponse? releaseProperties)
        {
            CustomImageDisplayName = customImageDisplayName;
            CustomImageId = customImageId;
            OsName = osName;
            ReleaseProperties = releaseProperties;
        }
    }
}
