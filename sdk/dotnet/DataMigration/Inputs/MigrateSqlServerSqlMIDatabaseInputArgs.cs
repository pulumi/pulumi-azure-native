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
    /// Database specific information for SQL to Azure SQL DB Managed Instance migration task inputs
    /// </summary>
    public sealed class MigrateSqlServerSqlMIDatabaseInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupFilePaths")]
        private InputList<string>? _backupFilePaths;

        /// <summary>
        /// The list of backup files to be used in case of existing backups.
        /// </summary>
        public InputList<string> BackupFilePaths
        {
            get => _backupFilePaths ?? (_backupFilePaths = new InputList<string>());
            set => _backupFilePaths = value;
        }

        /// <summary>
        /// Backup file share information for backing up this database.
        /// </summary>
        [Input("backupFileShare")]
        public Input<Inputs.FileShareArgs>? BackupFileShare { get; set; }

        /// <summary>
        /// id of the database
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the database
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the database at destination
        /// </summary>
        [Input("restoreDatabaseName", required: true)]
        public Input<string> RestoreDatabaseName { get; set; } = null!;

        public MigrateSqlServerSqlMIDatabaseInputArgs()
        {
        }
        public static new MigrateSqlServerSqlMIDatabaseInputArgs Empty => new MigrateSqlServerSqlMIDatabaseInputArgs();
    }
}
