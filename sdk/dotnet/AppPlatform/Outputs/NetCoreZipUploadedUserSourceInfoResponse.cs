// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// Uploaded Jar binary for a deployment
    /// </summary>
    [OutputType]
    public sealed class NetCoreZipUploadedUserSourceInfoResponse
    {
        /// <summary>
        /// The path to the .NET executable relative to zip root
        /// </summary>
        public readonly string? NetCoreMainEntryPath;
        /// <summary>
        /// Relative path of the storage which stores the source
        /// </summary>
        public readonly string? RelativePath;
        /// <summary>
        /// Runtime version of the .Net file
        /// </summary>
        public readonly string? RuntimeVersion;
        /// <summary>
        /// Type of the source uploaded
        /// Expected value is 'NetCoreZip'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the source
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private NetCoreZipUploadedUserSourceInfoResponse(
            string? netCoreMainEntryPath,

            string? relativePath,

            string? runtimeVersion,

            string type,

            string? version)
        {
            NetCoreMainEntryPath = netCoreMainEntryPath;
            RelativePath = relativePath;
            RuntimeVersion = runtimeVersion;
            Type = type;
            Version = version;
        }
    }
}
