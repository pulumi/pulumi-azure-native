// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The properties of a task step.
    /// </summary>
    [OutputType]
    public sealed class FileTaskStepResponse
    {
        /// <summary>
        /// List of base image dependencies for a step.
        /// </summary>
        public readonly ImmutableArray<Outputs.BaseImageDependencyResponse> BaseImageDependencies;
        /// <summary>
        /// The token (git PAT or SAS token of storage account blob) associated with the context for a step.
        /// </summary>
        public readonly string? ContextAccessToken;
        /// <summary>
        /// The URL(absolute or relative) of the source context for the task step.
        /// </summary>
        public readonly string? ContextPath;
        /// <summary>
        /// The task template/definition file path relative to the source context.
        /// </summary>
        public readonly string TaskFilePath;
        /// <summary>
        /// The type of the step.
        /// Expected value is 'FileTask'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The collection of overridable values that can be passed when running a task.
        /// </summary>
        public readonly ImmutableArray<Outputs.SetValueResponse> Values;
        /// <summary>
        /// The task values/parameters file path relative to the source context.
        /// </summary>
        public readonly string? ValuesFilePath;

        [OutputConstructor]
        private FileTaskStepResponse(
            ImmutableArray<Outputs.BaseImageDependencyResponse> baseImageDependencies,

            string? contextAccessToken,

            string? contextPath,

            string taskFilePath,

            string type,

            ImmutableArray<Outputs.SetValueResponse> values,

            string? valuesFilePath)
        {
            BaseImageDependencies = baseImageDependencies;
            ContextAccessToken = contextAccessToken;
            ContextPath = contextPath;
            TaskFilePath = taskFilePath;
            Type = type;
            Values = values;
            ValuesFilePath = valuesFilePath;
        }
    }
}
