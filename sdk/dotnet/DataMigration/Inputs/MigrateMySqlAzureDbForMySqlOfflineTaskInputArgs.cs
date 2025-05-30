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
    /// Input for the task that migrates MySQL databases to Azure Database for MySQL for offline migrations
    /// </summary>
    public sealed class MigrateMySqlAzureDbForMySqlOfflineTaskInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// encrypted key for secure fields
        /// </summary>
        [Input("encryptedKeyForSecureFields")]
        public Input<string>? EncryptedKeyForSecureFields { get; set; }

        /// <summary>
        /// Setting to set the source server read only
        /// </summary>
        [Input("makeSourceServerReadOnly")]
        public Input<bool>? MakeSourceServerReadOnly { get; set; }

        [Input("optionalAgentSettings")]
        private InputMap<string>? _optionalAgentSettings;

        /// <summary>
        /// Optional parameters for fine tuning the data transfer rate during migration
        /// </summary>
        public InputMap<string> OptionalAgentSettings
        {
            get => _optionalAgentSettings ?? (_optionalAgentSettings = new InputMap<string>());
            set => _optionalAgentSettings = value;
        }

        [Input("selectedDatabases", required: true)]
        private InputList<Inputs.MigrateMySqlAzureDbForMySqlOfflineDatabaseInputArgs>? _selectedDatabases;

        /// <summary>
        /// Databases to migrate
        /// </summary>
        public InputList<Inputs.MigrateMySqlAzureDbForMySqlOfflineDatabaseInputArgs> SelectedDatabases
        {
            get => _selectedDatabases ?? (_selectedDatabases = new InputList<Inputs.MigrateMySqlAzureDbForMySqlOfflineDatabaseInputArgs>());
            set => _selectedDatabases = value;
        }

        /// <summary>
        /// Connection information for source MySQL
        /// </summary>
        [Input("sourceConnectionInfo", required: true)]
        public Input<Inputs.MySqlConnectionInfoArgs> SourceConnectionInfo { get; set; } = null!;

        /// <summary>
        /// Parameter to specify when the migration started
        /// </summary>
        [Input("startedOn")]
        public Input<string>? StartedOn { get; set; }

        /// <summary>
        /// Connection information for target Azure Database for MySQL
        /// </summary>
        [Input("targetConnectionInfo", required: true)]
        public Input<Inputs.MySqlConnectionInfoArgs> TargetConnectionInfo { get; set; } = null!;

        public MigrateMySqlAzureDbForMySqlOfflineTaskInputArgs()
        {
            MakeSourceServerReadOnly = false;
        }
        public static new MigrateMySqlAzureDbForMySqlOfflineTaskInputArgs Empty => new MigrateMySqlAzureDbForMySqlOfflineTaskInputArgs();
    }
}
