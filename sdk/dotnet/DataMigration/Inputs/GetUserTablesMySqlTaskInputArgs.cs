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
    public sealed class GetUserTablesMySqlTaskInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection information for SQL Server
        /// </summary>
        [Input("connectionInfo", required: true)]
        public Input<Inputs.MySqlConnectionInfoArgs> ConnectionInfo { get; set; } = null!;

        [Input("selectedDatabases", required: true)]
        private InputList<string>? _selectedDatabases;

        /// <summary>
        /// List of database names to collect tables for
        /// </summary>
        public InputList<string> SelectedDatabases
        {
            get => _selectedDatabases ?? (_selectedDatabases = new InputList<string>());
            set => _selectedDatabases = value;
        }

        public GetUserTablesMySqlTaskInputArgs()
        {
        }
        public static new GetUserTablesMySqlTaskInputArgs Empty => new GetUserTablesMySqlTaskInputArgs();
    }
}
