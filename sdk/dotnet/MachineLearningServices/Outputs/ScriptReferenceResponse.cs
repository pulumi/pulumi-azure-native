// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Script reference
    /// </summary>
    [OutputType]
    public sealed class ScriptReferenceResponse
    {
        /// <summary>
        /// Optional command line arguments passed to the script to run.
        /// </summary>
        public readonly string? ScriptArguments;
        /// <summary>
        /// The location of scripts in the mounted volume.
        /// </summary>
        public readonly string? ScriptData;
        /// <summary>
        /// The storage source of the script: workspace.
        /// </summary>
        public readonly string? ScriptSource;
        /// <summary>
        /// Optional time period passed to timeout command.
        /// </summary>
        public readonly string? Timeout;

        [OutputConstructor]
        private ScriptReferenceResponse(
            string? scriptArguments,

            string? scriptData,

            string? scriptSource,

            string? timeout)
        {
            ScriptArguments = scriptArguments;
            ScriptData = scriptData;
            ScriptSource = scriptSource;
            Timeout = timeout;
        }
    }
}
