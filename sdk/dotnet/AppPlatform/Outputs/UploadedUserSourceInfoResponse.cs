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
    /// Source with uploaded location
    /// </summary>
    [OutputType]
    public sealed class UploadedUserSourceInfoResponse
    {
        /// <summary>
        /// Relative path of the storage which stores the source
        /// </summary>
        public readonly string? RelativePath;
        /// <summary>
        /// Type of the source uploaded
        /// Expected value is 'UploadedUserSourceInfo'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the source
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private UploadedUserSourceInfoResponse(
            string? relativePath,

            string type,

            string? version)
        {
            RelativePath = relativePath;
            Type = type;
            Version = version;
        }
    }
}
