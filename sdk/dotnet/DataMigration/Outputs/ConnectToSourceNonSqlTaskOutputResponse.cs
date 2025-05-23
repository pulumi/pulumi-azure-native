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
    /// Output for connect to MySQL type source
    /// </summary>
    [OutputType]
    public sealed class ConnectToSourceNonSqlTaskOutputResponse
    {
        /// <summary>
        /// List of databases on the server
        /// </summary>
        public readonly ImmutableArray<string> Databases;
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Server properties
        /// </summary>
        public readonly Outputs.ServerPropertiesResponse ServerProperties;
        /// <summary>
        /// Server brand version
        /// </summary>
        public readonly string SourceServerBrandVersion;
        /// <summary>
        /// Validation errors associated with the task
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponse> ValidationErrors;

        [OutputConstructor]
        private ConnectToSourceNonSqlTaskOutputResponse(
            ImmutableArray<string> databases,

            string id,

            Outputs.ServerPropertiesResponse serverProperties,

            string sourceServerBrandVersion,

            ImmutableArray<Outputs.ReportableExceptionResponse> validationErrors)
        {
            Databases = databases;
            Id = id;
            ServerProperties = serverProperties;
            SourceServerBrandVersion = sourceServerBrandVersion;
            ValidationErrors = validationErrors;
        }
    }
}
