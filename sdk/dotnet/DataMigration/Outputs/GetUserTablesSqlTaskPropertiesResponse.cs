// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    /// <summary>
    /// Properties for the task that collects user tables for the given list of databases
    /// </summary>
    [OutputType]
    public sealed class GetUserTablesSqlTaskPropertiesResponse
    {
        /// <summary>
        /// Key value pairs of client data to attach meta data information to task
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ClientData;
        /// <summary>
        /// Array of command properties.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.MigrateMISyncCompleteCommandPropertiesResponse, Outputs.MigrateSyncCompleteCommandPropertiesResponse>> Commands;
        /// <summary>
        /// Array of errors. This is ignored if submitted.
        /// </summary>
        public readonly ImmutableArray<Outputs.ODataErrorResponse> Errors;
        /// <summary>
        /// Task input
        /// </summary>
        public readonly Outputs.GetUserTablesSqlTaskInputResponse? Input;
        /// <summary>
        /// Task output. This is ignored if submitted.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserTablesSqlTaskOutputResponse> Output;
        /// <summary>
        /// The state of the task. This is ignored if submitted.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Task id 
        /// </summary>
        public readonly string? TaskId;
        /// <summary>
        /// Task type.
        /// Expected value is 'GetUserTables.Sql'.
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private GetUserTablesSqlTaskPropertiesResponse(
            ImmutableDictionary<string, string>? clientData,

            ImmutableArray<Union<Outputs.MigrateMISyncCompleteCommandPropertiesResponse, Outputs.MigrateSyncCompleteCommandPropertiesResponse>> commands,

            ImmutableArray<Outputs.ODataErrorResponse> errors,

            Outputs.GetUserTablesSqlTaskInputResponse? input,

            ImmutableArray<Outputs.GetUserTablesSqlTaskOutputResponse> output,

            string state,

            string? taskId,

            string taskType)
        {
            ClientData = clientData;
            Commands = commands;
            Errors = errors;
            Input = input;
            Output = output;
            State = state;
            TaskId = taskId;
            TaskType = taskType;
        }
    }
}
