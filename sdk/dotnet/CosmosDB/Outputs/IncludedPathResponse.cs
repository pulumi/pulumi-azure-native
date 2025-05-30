// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Outputs
{

    /// <summary>
    /// The paths that are included in indexing
    /// </summary>
    [OutputType]
    public sealed class IncludedPathResponse
    {
        /// <summary>
        /// List of indexes for this path
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexesResponse> Indexes;
        /// <summary>
        /// The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private IncludedPathResponse(
            ImmutableArray<Outputs.IndexesResponse> indexes,

            string? path)
        {
            Indexes = indexes;
            Path = path;
        }
    }
}
