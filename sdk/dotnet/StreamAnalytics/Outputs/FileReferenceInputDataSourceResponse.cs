// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.Outputs
{

    /// <summary>
    /// Describes a file input data source that contains reference data.
    /// </summary>
    [OutputType]
    public sealed class FileReferenceInputDataSourceResponse
    {
        /// <summary>
        /// The path of the file.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Indicates the type of input data source containing reference data. Required on PUT (CreateOrReplace) requests.
        /// Expected value is 'File'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FileReferenceInputDataSourceResponse(
            string? path,

            string type)
        {
            Path = path;
            Type = type;
        }
    }
}
