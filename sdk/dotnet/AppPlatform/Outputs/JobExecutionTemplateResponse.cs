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
    /// Job's execution template, containing configuration for an execution
    /// </summary>
    [OutputType]
    public sealed class JobExecutionTemplateResponse
    {
        /// <summary>
        /// Arguments for the Job execution.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Environment variables of Job execution
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvVarResponse> EnvironmentVariables;
        /// <summary>
        /// The requested resource quantity for required CPU and Memory.
        /// </summary>
        public readonly Outputs.JobResourceRequestsResponse? ResourceRequests;

        [OutputConstructor]
        private JobExecutionTemplateResponse(
            ImmutableArray<string> args,

            ImmutableArray<Outputs.EnvVarResponse> environmentVariables,

            Outputs.JobResourceRequestsResponse? resourceRequests)
        {
            Args = args;
            EnvironmentVariables = environmentVariables;
            ResourceRequests = resourceRequests;
        }
    }
}
