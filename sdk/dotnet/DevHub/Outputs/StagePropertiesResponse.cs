// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevHub.Outputs
{

    /// <summary>
    /// Properties of a Stage.
    /// </summary>
    [OutputType]
    public sealed class StagePropertiesResponse
    {
        public readonly ImmutableArray<string> Dependencies;
        public readonly string? GitEnvironment;
        /// <summary>
        /// Stage Name
        /// </summary>
        public readonly string? StageName;

        [OutputConstructor]
        private StagePropertiesResponse(
            ImmutableArray<string> dependencies,

            string? gitEnvironment,

            string? stageName)
        {
            Dependencies = dependencies;
            GitEnvironment = gitEnvironment;
            StageName = stageName;
        }
    }
}
