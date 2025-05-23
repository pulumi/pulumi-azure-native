// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of ImageConfig
    /// </summary>
    [OutputType]
    public sealed class ImageConfigResponse
    {
        /// <summary>
        /// Specifies parameters that you want to pass in with ENTRYPOINT. You can specify a maximum of 1,500 parameters in the list.
        /// </summary>
        public readonly ImmutableArray<string> Command;
        /// <summary>
        /// Specifies the entry point to their application, which is typically the location of the runtime executable. You can specify a maximum of 1,500 string entries in the list.
        /// </summary>
        public readonly ImmutableArray<string> EntryPoint;
        /// <summary>
        /// Specifies the working directory. The length of the directory string cannot exceed 1,000 characters.
        /// </summary>
        public readonly string? WorkingDirectory;

        [OutputConstructor]
        private ImageConfigResponse(
            ImmutableArray<string> command,

            ImmutableArray<string> entryPoint,

            string? workingDirectory)
        {
            Command = command;
            EntryPoint = entryPoint;
            WorkingDirectory = workingDirectory;
        }
    }
}
