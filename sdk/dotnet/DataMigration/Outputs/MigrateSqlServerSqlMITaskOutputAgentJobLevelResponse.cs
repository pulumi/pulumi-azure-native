// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    [OutputType]
    public sealed class MigrateSqlServerSqlMITaskOutputAgentJobLevelResponse
    {
        /// <summary>
        /// Migration end time
        /// </summary>
        public readonly string EndedOn;
        /// <summary>
        /// Migration errors and warnings per job
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponse> ExceptionsAndWarnings;
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The state of the original Agent Job.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// Migration progress message
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Agent Job name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Result type
        /// Expected value is 'AgentJobLevelOutput'.
        /// </summary>
        public readonly string ResultType;
        /// <summary>
        /// Migration start time
        /// </summary>
        public readonly string StartedOn;
        /// <summary>
        /// Current state of migration
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private MigrateSqlServerSqlMITaskOutputAgentJobLevelResponse(
            string endedOn,

            ImmutableArray<Outputs.ReportableExceptionResponse> exceptionsAndWarnings,

            string id,

            bool isEnabled,

            string message,

            string name,

            string resultType,

            string startedOn,

            string state)
        {
            EndedOn = endedOn;
            ExceptionsAndWarnings = exceptionsAndWarnings;
            Id = id;
            IsEnabled = isEnabled;
            Message = message;
            Name = name;
            ResultType = resultType;
            StartedOn = startedOn;
            State = state;
        }
    }
}
