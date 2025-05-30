// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Inputs
{

    /// <summary>
    /// Input for the task that collects user tables for the given list of databases
    /// </summary>
    public sealed class GetUserTablesSqlSyncTaskInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("selectedSourceDatabases", required: true)]
        private InputList<string>? _selectedSourceDatabases;

        /// <summary>
        /// List of source database names to collect tables for
        /// </summary>
        public InputList<string> SelectedSourceDatabases
        {
            get => _selectedSourceDatabases ?? (_selectedSourceDatabases = new InputList<string>());
            set => _selectedSourceDatabases = value;
        }

        [Input("selectedTargetDatabases", required: true)]
        private InputList<string>? _selectedTargetDatabases;

        /// <summary>
        /// List of target database names to collect tables for
        /// </summary>
        public InputList<string> SelectedTargetDatabases
        {
            get => _selectedTargetDatabases ?? (_selectedTargetDatabases = new InputList<string>());
            set => _selectedTargetDatabases = value;
        }

        /// <summary>
        /// Connection information for SQL Server
        /// </summary>
        [Input("sourceConnectionInfo", required: true)]
        public Input<Inputs.SqlConnectionInfoArgs> SourceConnectionInfo { get; set; } = null!;

        /// <summary>
        /// Connection information for SQL DB
        /// </summary>
        [Input("targetConnectionInfo", required: true)]
        public Input<Inputs.SqlConnectionInfoArgs> TargetConnectionInfo { get; set; } = null!;

        public GetUserTablesSqlSyncTaskInputArgs()
        {
        }
        public static new GetUserTablesSqlSyncTaskInputArgs Empty => new GetUserTablesSqlSyncTaskInputArgs();
    }
}
