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
    /// Properties for the task that collects user tables for the given list of Oracle schemas
    /// </summary>
    [OutputType]
    public sealed class GetUserTablesOracleTaskPropertiesResponse
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
        public readonly Outputs.GetUserTablesOracleTaskInputResponse? Input;
        /// <summary>
        /// Task output. This is ignored if submitted.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserTablesOracleTaskOutputResponse> Output;
        /// <summary>
        /// The state of the task. This is ignored if submitted.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Task type.
        /// Expected value is 'GetUserTablesOracle'.
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private GetUserTablesOracleTaskPropertiesResponse(
            ImmutableDictionary<string, string>? clientData,

            ImmutableArray<Union<Outputs.MigrateMISyncCompleteCommandPropertiesResponse, Outputs.MigrateSyncCompleteCommandPropertiesResponse>> commands,

            ImmutableArray<Outputs.ODataErrorResponse> errors,

            Outputs.GetUserTablesOracleTaskInputResponse? input,

            ImmutableArray<Outputs.GetUserTablesOracleTaskOutputResponse> output,

            string state,

            string taskType)
        {
            ClientData = clientData;
            Commands = commands;
            Errors = errors;
            Input = input;
            Output = output;
            State = state;
            TaskType = taskType;
        }
    }
}
