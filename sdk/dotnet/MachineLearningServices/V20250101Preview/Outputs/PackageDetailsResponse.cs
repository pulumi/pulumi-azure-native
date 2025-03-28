// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    [OutputType]
    public sealed class PackageDetailsResponse
    {
        /// <summary>
        /// Install path.
        /// </summary>
        public readonly string? InstallPath;
        /// <summary>
        /// Installed version.
        /// </summary>
        public readonly string? InstalledVersion;
        /// <summary>
        /// Package or dependency name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Patched version.
        /// </summary>
        public readonly string? PatchedVersion;

        [OutputConstructor]
        private PackageDetailsResponse(
            string? installPath,

            string? installedVersion,

            string? name,

            string? patchedVersion)
        {
            InstallPath = installPath;
            InstalledVersion = installedVersion;
            Name = name;
            PatchedVersion = patchedVersion;
        }
    }
}
