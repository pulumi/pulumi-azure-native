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
    /// The properties of an operating system release.
    /// </summary>
    [OutputType]
    public sealed class ReleasePropertiesResponse
    {
        /// <summary>
        /// The build number of the OS release.
        /// </summary>
        public readonly string? BuildNumber;
        /// <summary>
        /// The build revision of the OS release.
        /// </summary>
        public readonly string? BuildRevision;
        /// <summary>
        /// The name of the OS release.
        /// </summary>
        public readonly string? ReleaseName;
        /// <summary>
        /// The release version date of the OS release.
        /// </summary>
        public readonly string? ReleaseVersionDate;

        [OutputConstructor]
        private ReleasePropertiesResponse(
            string? buildNumber,

            string? buildRevision,

            string? releaseName,

            string? releaseVersionDate)
        {
            BuildNumber = buildNumber;
            BuildRevision = buildRevision;
            ReleaseName = releaseName;
            ReleaseVersionDate = releaseVersionDate;
        }
    }
}
