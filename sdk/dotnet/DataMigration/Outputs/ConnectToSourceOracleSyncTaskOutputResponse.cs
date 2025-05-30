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
    /// Output for the task that validates Oracle database connection
    /// </summary>
    [OutputType]
    public sealed class ConnectToSourceOracleSyncTaskOutputResponse
    {
        /// <summary>
        /// List of schemas on source server
        /// </summary>
        public readonly ImmutableArray<string> Databases;
        /// <summary>
        /// Source server brand version
        /// </summary>
        public readonly string SourceServerBrandVersion;
        /// <summary>
        /// Version of the source server
        /// </summary>
        public readonly string SourceServerVersion;
        /// <summary>
        /// Validation errors associated with the task
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponse> ValidationErrors;

        [OutputConstructor]
        private ConnectToSourceOracleSyncTaskOutputResponse(
            ImmutableArray<string> databases,

            string sourceServerBrandVersion,

            string sourceServerVersion,

            ImmutableArray<Outputs.ReportableExceptionResponse> validationErrors)
        {
            Databases = databases;
            SourceServerBrandVersion = sourceServerBrandVersion;
            SourceServerVersion = sourceServerVersion;
            ValidationErrors = validationErrors;
        }
    }
}
