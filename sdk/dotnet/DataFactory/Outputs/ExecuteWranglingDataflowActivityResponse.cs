// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Execute power query activity.
    /// </summary>
    [OutputType]
    public sealed class ExecuteWranglingDataflowActivityResponse
    {
        /// <summary>
        /// Compute properties for data flow activity.
        /// </summary>
        public readonly Outputs.ExecuteDataFlowActivityTypePropertiesResponseCompute? Compute;
        /// <summary>
        /// Continuation settings for execute data flow activity.
        /// </summary>
        public readonly Outputs.ContinuationSettingsReferenceResponse? ContinuationSettings;
        /// <summary>
        /// Continue on error setting used for data flow execution. Enables processing to continue if a sink fails. Type: boolean (or Expression with resultType boolean)
        /// </summary>
        public readonly object? ContinueOnError;
        /// <summary>
        /// Data flow reference.
        /// </summary>
        public readonly Outputs.DataFlowReferenceResponse DataFlow;
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponse> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponse? IntegrationRuntime;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Status result of the activity when the state is set to Inactive. This is an optional property and if not provided when the activity is inactive, the status will be Succeeded by default.
        /// </summary>
        public readonly string? OnInactiveMarkAs;
        /// <summary>
        /// Activity policy.
        /// </summary>
        public readonly Outputs.ActivityPolicyResponse? Policy;
        /// <summary>
        /// List of mapping for Power Query mashup query to sink dataset(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.PowerQuerySinkMappingResponse> Queries;
        /// <summary>
        /// Concurrent run setting used for data flow execution. Allows sinks with the same save order to be processed concurrently. Type: boolean (or Expression with resultType boolean)
        /// </summary>
        public readonly object? RunConcurrently;
        /// <summary>
        /// (Deprecated. Please use Queries). List of Power Query activity sinks mapped to a queryName.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.PowerQuerySinkResponse>? Sinks;
        /// <summary>
        /// Specify number of parallel staging for sources applicable to the sink. Type: integer (or Expression with resultType integer)
        /// </summary>
        public readonly object? SourceStagingConcurrency;
        /// <summary>
        /// Staging info for execute data flow activity.
        /// </summary>
        public readonly Outputs.DataFlowStagingInfoResponse? Staging;
        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Trace level setting used for data flow monitoring output. Supported values are: 'coarse', 'fine', and 'none'. Type: string (or Expression with resultType string)
        /// </summary>
        public readonly object? TraceLevel;
        /// <summary>
        /// Type of activity.
        /// Expected value is 'ExecuteWranglingDataflow'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponse> UserProperties;

        [OutputConstructor]
        private ExecuteWranglingDataflowActivityResponse(
            Outputs.ExecuteDataFlowActivityTypePropertiesResponseCompute? compute,

            Outputs.ContinuationSettingsReferenceResponse? continuationSettings,

            object? continueOnError,

            Outputs.DataFlowReferenceResponse dataFlow,

            ImmutableArray<Outputs.ActivityDependencyResponse> dependsOn,

            string? description,

            Outputs.IntegrationRuntimeReferenceResponse? integrationRuntime,

            string name,

            string? onInactiveMarkAs,

            Outputs.ActivityPolicyResponse? policy,

            ImmutableArray<Outputs.PowerQuerySinkMappingResponse> queries,

            object? runConcurrently,

            ImmutableDictionary<string, Outputs.PowerQuerySinkResponse>? sinks,

            object? sourceStagingConcurrency,

            Outputs.DataFlowStagingInfoResponse? staging,

            string? state,

            object? traceLevel,

            string type,

            ImmutableArray<Outputs.UserPropertyResponse> userProperties)
        {
            Compute = compute;
            ContinuationSettings = continuationSettings;
            ContinueOnError = continueOnError;
            DataFlow = dataFlow;
            DependsOn = dependsOn;
            Description = description;
            IntegrationRuntime = integrationRuntime;
            Name = name;
            OnInactiveMarkAs = onInactiveMarkAs;
            Policy = policy;
            Queries = queries;
            RunConcurrently = runConcurrently;
            Sinks = sinks;
            SourceStagingConcurrency = sourceStagingConcurrency;
            Staging = staging;
            State = state;
            TraceLevel = traceLevel;
            Type = type;
            UserProperties = userProperties;
        }
    }
}
