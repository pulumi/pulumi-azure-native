// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// The configuration settings of the storage of the tokens if a file system is used.
    /// </summary>
    [OutputType]
    public sealed class FileSystemTokenStoreResponse
    {
        /// <summary>
        /// The directory in which the tokens will be stored.
        /// </summary>
        public readonly string? Directory;

        [OutputConstructor]
        private FileSystemTokenStoreResponse(string? directory)
        {
            Directory = directory;
        }
    }
}
