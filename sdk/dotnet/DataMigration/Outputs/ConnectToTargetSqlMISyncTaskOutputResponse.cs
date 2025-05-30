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
    /// Output for the task that validates connection to Azure SQL Database Managed Instance.
    /// </summary>
    [OutputType]
    public sealed class ConnectToTargetSqlMISyncTaskOutputResponse
    {
        /// <summary>
        /// Target server brand version
        /// </summary>
        public readonly string TargetServerBrandVersion;
        /// <summary>
        /// Target server version
        /// </summary>
        public readonly string TargetServerVersion;
        /// <summary>
        /// Validation errors
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponse> ValidationErrors;

        [OutputConstructor]
        private ConnectToTargetSqlMISyncTaskOutputResponse(
            string targetServerBrandVersion,

            string targetServerVersion,

            ImmutableArray<Outputs.ReportableExceptionResponse> validationErrors)
        {
            TargetServerBrandVersion = targetServerBrandVersion;
            TargetServerVersion = targetServerVersion;
            ValidationErrors = validationErrors;
        }
    }
}
