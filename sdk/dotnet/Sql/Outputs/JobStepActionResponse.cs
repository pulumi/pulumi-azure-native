// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// The action to be executed by a job step.
    /// </summary>
    [OutputType]
    public sealed class JobStepActionResponse
    {
        /// <summary>
        /// The source of the action to execute.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// Type of action being executed by the job step.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The action value, for example the text of the T-SQL script to execute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private JobStepActionResponse(
            string? source,

            string? type,

            string value)
        {
            Source = source;
            Type = type;
            Value = value;
        }
    }
}
